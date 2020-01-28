package util

import (
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
	RequestPayloadLower  string // the request payload with the first letter in lower case
	ResponsePayloadLower string // the response payload with the first letter in lower case
	URL                  string // the operation name with eiphens
	Method               string // the type of method. GET if no request param type is known
	// POST otherwise
}

// Fielddetail - the details of either the params or the return values
type Fielddetail struct {
	Name   string // the name of the argument (for params) or "" for return values
	Type   string // the type (either the primitive type or the struct)
	Kind   string // the kind (the precise type as defined in reflect package)
	Origin string // the origin as expected by the ParamOrigin of Param descriptor in B Plus
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
			visitInterface(sdet, inter)
			return true
		}
		return true
	})
}

func visitInterface(sdet *Servicedetail, t *ast.InterfaceType) {
	m := t.Methods

	for _, m1 := range m.List {
		op := m1.Names[0]
		ft, ok := m1.Type.(*ast.FuncType)
		if ok {
			sdet.Operations = append(sdet.Operations, getOpData(sdet, op.Name, ft))
		}
	}
}

func getOpData(sdet *Servicedetail, op string, ft *ast.FuncType) Operationdetail {
	paramDetails, requestPayloadType := parseFields(op, ft.Params)
	respDetails, responsePayloadType := parseFields(op, ft.Results)
	method := "POST"
	if requestPayloadType == "" {
		method = "GET"
	}

	return Operationdetail{
		Operation:            op,
		Params:               paramDetails,
		Results:              respDetails,
		RequestPayload:       requestPayloadType,
		ResponsePayload:      responsePayloadType,
		RequestPayloadLower:  strcase.ToLowerCamel(requestPayloadType),
		ResponsePayloadLower: strcase.ToLowerCamel(responsePayloadType),
		URL:                  strcase.ToDelimited(op, '-'),
		Method:               method,
	}
}

// returns the details of all fields as well as the type for the
// field that is of kind payload
func parseFields(op string, fl *ast.FieldList) ([]Fielddetail, string) {
	var pd []Fielddetail
	var payloadType = ""
	for _, m1 := range fl.List {
		var name = ""
		if len(m1.Names) == 1 {
			name = m1.Names[0].Name
		}
		varType := ""
		switch v := m1.Type.(type) {
		case *ast.Ident:
			varType = v.Name
		case *ast.SelectorExpr:
			varType = v.Sel.Name
		}
		kind := getKindOf(op, name, varType)
		origin := getOrigin(varType)
		pd = append(pd, Fielddetail{name, varType, kind, origin})
		if origin == "bplus.PAYLOAD" {
			payloadType = varType
		}
	}

	return pd, payloadType
}

func getKindOf(op string, paramname string, paramtype string) string {
	switch paramtype {
	case "int", "int8", "int16", "int32", "int64":
		return "reflect.Int"
	case "string":
		return "reflect.String"
	case "float32":
		return "reflect.Float32"
	case "float64":
		return "reflect.Float64"
	default:
		return ""
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
