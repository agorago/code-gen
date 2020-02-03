package util

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

// Servicedetail - the full structure of the interface definition extracted from the AST
type Servicedetail struct {
	InterfaceFile  string            // the path to the .go file that will be parsed for interfaces
	Name           string            // the Base Name of the file (without .go suffix)
	CamelCase      string            // the base name of the file in Camel Case
	CamelCaseLower string            // the base name in camel case with the first letter in lower case
	URL            string            // the URL for the repo as passed in the argument
	BaseErrorCode  string            // the starting error code or 100,000 if not specified
	Operations     []Operationdetail // Details of the operations
}

// Operationdetail - the details of the operation
type Operationdetail struct {
	Operation      string // the name of the Operation
	Params         []Fielddetail
	Results        []Fielddetail
	RequestPayload string // the type of the request payload - short cut rather than parsing
	// through all the Params above
	ResponsePayload string // the type of the response payload - short cut rather than parsing
	// through all the Results above
	RequestPayloadLower         string // the request payload with the first letter in lower case
	ResponsePayloadLower        string // the response payload with the first letter in lower case
	RequestPayloadDefaultValue  string // short cut for default value of the request payload
	ResponsePayloadDefaultValue string // short cut for default value of the response payload
	URL                         string // the operation name with eiphens
	Method                      string // the type of method. GET if no request param type is known
	// POST otherwise
}

// Fielddetail - the details of either the params or the return values
type Fielddetail struct {
	Name         string // the name of the argument (for params) or "" for return values
	Type         string // the type (either the primitive type or the struct)
	Kind         string // the kind (the precise type as defined in reflect package)
	Origin       string // the origin as expected by the ParamOrigin of Param descriptor in B Plus
	DefaultValue string // the default value of this param that can be passed to a function
	PointerType  bool   // Is this a pointer or a normal struct
}

// ParseService - read the interface file passed from command line and extract the
// Servicedetail
func ParseService() Servicedetail {
	sdet := Servicedetail{}

	iFile := os.Args[1]
	sdet.InterfaceFile = iFile
	sdet.Name = trimInterfaceName(iFile)
	sdet.URL = os.Args[2]
	sdet.CamelCase = strcase.ToCamel(sdet.Name)
	sdet.CamelCaseLower = strcase.ToLowerCamel(sdet.Name)
	s := "100000"
	if len(os.Args) > 4 {
		s = os.Args[4]
	}
	sdet.BaseErrorCode = s
	parseFile(iFile, &sdet)

	return sdet
}

func trimInterfaceName(s string) string {
	arr := strings.Split(s, "/")
	if n := len(arr); n > 0 {
		s = arr[n-1]
	}
	return strings.TrimSuffix(s, ".go")
}

func parseFile(iFile string, sdet *Servicedetail) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, iFile, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("File %s is not readable\n", iFile)
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Find Interface definitions
		inter, ok := n.(*ast.InterfaceType)
		if ok {
			err := visitInterface(sdet, inter)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error occurred while parsing the file. Error = %s\n", err.Error())
				return false
			}
			return true
		}
		return true
	})
}

func visitInterface(sdet *Servicedetail, t *ast.InterfaceType) error {
	m := t.Methods

	for _, m1 := range m.List {
		op := m1.Names[0]
		ft, ok := m1.Type.(*ast.FuncType)
		opdata, err := getOpData(sdet, op.Name, ft)
		if err != nil {
			return fmt.Errorf("Error with interface. Err = %s", err.Error())
		}
		if ok {
			sdet.Operations = append(sdet.Operations, opdata)
		}
	}
	return nil
}

func getOpData(sdet *Servicedetail, op string, ft *ast.FuncType) (Operationdetail, error) {
	paramDetails, requestPayloadType, requestPayloadDefaultValue := parseFields(op, ft.Params)
	if len(paramDetails) == 0 || paramDetails[0].Name != "ctx" ||
		paramDetails[0].Type != "context.Context" {
		return Operationdetail{}, fmt.Errorf("First parameter of function %s must be ctx context.Context", op)
	}

	respDetails, responsePayloadType, responsePayloadDefaultValue := parseFields(op, ft.Results)
	method := "POST"
	if requestPayloadType == "" {
		method = "GET"
	}

	return Operationdetail{
		Operation:                   op,
		Params:                      paramDetails,
		Results:                     respDetails,
		RequestPayload:              requestPayloadType,
		ResponsePayload:             responsePayloadType,
		RequestPayloadLower:         strcase.ToLowerCamel(requestPayloadType),
		ResponsePayloadLower:        strcase.ToLowerCamel(responsePayloadType),
		RequestPayloadDefaultValue:  requestPayloadDefaultValue,
		ResponsePayloadDefaultValue: responsePayloadDefaultValue,
		URL:                         strcase.ToDelimited(op, '-'),
		Method:                      method,
	}, nil
}

// returns the details of all fields as well as the type for the
// field that is of kind payload
func parseFields(op string, fl *ast.FieldList) ([]Fielddetail, string, string) {
	var pd []Fielddetail
	var payloadType = ""
	var payloadDefaultValue = ""
	var pointerType = false
	for _, m1 := range fl.List {
		var name = ""
		if len(m1.Names) == 1 {
			name = m1.Names[0].Name
		}
		varType := ""
		switch v := m1.Type.(type) {
		case *ast.Ident:
			varType = v.Name
			//fmt.Fprintf(os.Stderr,"Ident:%#v\n", v)
		case *ast.SelectorExpr:
			varType = v.Sel.Name
			//fmt.Fprintf(os.Stderr,"Sel:%#v X: %#v\n", v.Sel, v.X)
		case *ast.StarExpr:
			v1 := v.X.(*ast.Ident)
			varType = v1.Name
			pointerType = true
			// fmt.Fprintf(os.Stderr,"StarExpr:%#v \n", v.X)
		default:
			fmt.Fprintf(os.Stderr, "Unknown Param or return type - talk to the team that maintains this program :%#v \n", v)
		}

		origin := getOrigin(varType)
		varType = correctPayloadType(varType, origin, pointerType)
		kind, defaultValue := getKindOfDefaultValue(op, name, varType, pointerType)
		pd = append(pd, Fielddetail{name, varType, kind, origin, defaultValue, pointerType})
		if origin == "bplus.PAYLOAD" {
			payloadType = varType
			payloadDefaultValue = defaultValue
		}
	}

	return pd, payloadType, payloadDefaultValue
}

// correct the payload type to reflect the correct value.
// the AST does not seem to give us the fully qualified type name
func correctPayloadType(typ string, origin string, pointerType bool) string {
	switch origin {
	case "bplus.CONTEXT":
		return "context.Context"
	case "bplus.PAYLOAD":
		if pointerType {
			return "*api." + typ
		}
		return "api." + typ
	default:
		return typ
	}
}

func getKindOfDefaultValue(op string, paramname string, paramtype string, pointerType bool) (string, string) {
	switch paramtype {
	case "int", "int8", "int16", "int32", "int64":
		return "reflect.Int", "0"
	case "string":
		return "reflect.String", `""`
	case "float32":
		return "reflect.Float32", "0.0"
	case "float64":
		return "reflect.Float64", "0.0"
	case "Context":
		return "", "context.TODO()"
	case "error":
		return "", "nil"
	default:
		if pointerType {
			return "", "&" + paramtype[1:] + "{}"
		}

		return "", paramtype + "{}"

	}
}

func getOrigin(s string) string {
	switch s {
	case "int", "int8", "string", "int16", "int32", "int64", "float32",
		"float64":
		return "bplus.HEADER"
	case "Context":
		return "bplus.CONTEXT"
	case "error":
		return "error"
	default:
		return "bplus.PAYLOAD"
	}
}
