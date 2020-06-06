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
	Description string // the extracted comments for this service interface
	Name           string            // the Base Name of the file (without .go suffix) this can have eiphens
	CamelCase      string            // the base name of the file in Camel Case
	CamelCaseLower string // the base name in camel case with the first letter in lower case
	InterfaceName string // the name of the actual interface that was defined in the file
	WegoURL  string // the URL of the wego library
	ApiURL            string            // the URL for the APUI repo
	ServiceURL string  // the URL for the service Repo
	BaseErrorCode  string            // the starting error code or 100,000 if not specified
	Operations     []*Operationdetail // Details of the operations
	DoesServiceHaveGetOperations bool // true if there is even a single operation that is pure GET (no payload)
}

// Operationdetail - the details of the operation
type Operationdetail struct {
	Operation      string // the name of the Operation
	Description string // the extracted comments for this operation
	Params         []Fielddetail
	Results        []Fielddetail
	UnqualifiedRequestPayload string
	UnqualifiedResponsePayload string
	RequestPayload string // the type of the request payload
	ResponsePayload string // the type of the response payload
	RequestPayloadLower         string // the request payload with the first letter in lower case
	ResponsePayloadLower        string // the response payload with the first letter in lower case
	RequestPayloadDefaultValue  string // Default value of the request payload
	ResponsePayloadDefaultValue string // Default value of the response payload
	URL                         string // constructed url for operation
	Method                      string // the type of method. GET if no request param type is known POST otherwise
	RequestDescription string // the description of the request payload
	ResponseDescription string // the description of the response of this operation
}

// Fielddetail - the details of either the params or the return values
type Fielddetail struct {
	Name         string // the name of the argument (for params) or "" for unnamed return values
	Description string // the extracted comments for this field
	UnqualifiedType string // the type that is returned by the AST. This does not have package name etc.
	Type         string // the type (either the primitive type or the struct)
	Kind         string // the kind (the precise type as defined in reflect package)
	Origin       string // the origin as expected by the ParamOrigin of Param descriptor
	DefaultValue string // the default value of this param that can be passed to a function
	PointerType  bool   // Is this a pointer or a normal struct
}

// ParseService - read the interface file passed from command line and extract the
// Servicedetail along with all other info. Parses the AST using GO reflection
func ParseService(iFile string,apiURL string,serviceURL string, wegoURL string, errorcode string) Servicedetail {
	sdet := Servicedetail{}

	sdet.InterfaceFile = iFile
	sdet.Name = trimInterfaceName(iFile)
	sdet.ApiURL = apiURL
	sdet.WegoURL = wegoURL
	sdet.ServiceURL = serviceURL
	sdet.CamelCase = strcase.ToCamel(sdet.Name)
	sdet.CamelCaseLower = strcase.ToLowerCamel(sdet.Name)
	sdet.DoesServiceHaveGetOperations = false
	sdet.BaseErrorCode = errorcode
	parseFile(iFile, &sdet)

	return sdet
}

// typeInfo - for every type the comments are preserved in this.
// Comments are used for constructing descriptions that are needed against the service and operation
type typeInfo struct {
	Name string
	Comments string
}

var allTypes map[string]typeInfo

func init(){
	allTypes = make(map[string]typeInfo)
}

// trimInterfaceName - extracts the interface from the fully qualified path.
// extracted name may contain eiphens
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
		log.Fatalf("File %s is not readable or parseable as a GO program. Error = %s\n", iFile,err.Error())
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type){
		case *ast.GenDecl:
			err = visitGenDecl(sdet,t)
		default:
			// log.Printf("Encountered this object %#v \n",t)
			return true
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred while parsing the file. Error = %s\n", err.Error())
			return false
		}
		return true
	})
	extractDescriptions(sdet)
}

// gets the description for service and for request and response payloads for each operation
func extractDescriptions(sdet *Servicedetail){
	sdet.Description = getCommentForType(sdet.InterfaceName)
	for _,op := range sdet.Operations{
		op.RequestDescription = getCommentForType(op.UnqualifiedRequestPayload)
		op.ResponseDescription = getCommentForType(op.UnqualifiedResponsePayload)
	}
}

func getCommentForType(t string)string{
	if t == "" {
		return ""
	}
	ti,ok := allTypes[t]
	if !ok {
		return ""
	}
	return ti.Comments
}

//  visitGenDecl - extract each declaration from the GO file.
// Specific behaviour for each specific type of declaration
func visitGenDecl(sdet *Servicedetail,t *ast.GenDecl) error{
	if t.Specs == nil{
		return nil
	}
	for _,spec := range t.Specs{
		ti,err := visitSpec(sdet,spec)
		if err != nil {
			return err
		}
		if ti.Name != "" {
			ti.Comments += extractDocumentation(t.Doc)
			allTypes[ti.Name] = ti
		}
	}
	return nil
}

// visitSpec - process the declaration if it defines a type
func visitSpec(sdet *Servicedetail,spec ast.Spec) (typeInfo,error) {
	switch s := spec.(type){
	case *ast.TypeSpec:
		return visitType(sdet, s)
	default:
		// log.Printf("Encountered a new spec %#v\n",s)
		return typeInfo{}, nil
	}
}

// visitType - for every type extract the documentation
// if it is interface type extract other details and store them in sdet
func visitType(sdet *Servicedetail,t *ast.TypeSpec) (typeInfo,error) {
	switch theType := t.Type.(type){
	case *ast.InterfaceType:
		err := visitInterface(sdet,theType)
		sdet.InterfaceName = t.Name.Name
		if err != nil {
			return typeInfo{},err
		}
	}
	return typeInfo {
		Name: t.Name.Name,
		Comments: extractDocumentation(t.Doc,t.Comment),
	},nil
}

func visitInterface(sdet *Servicedetail, t *ast.InterfaceType) error {
	m := t.Methods

	for _, m1 := range m.List {
		comments := extractDocumentation(m1.Doc,m1.Comment)
		op := m1.Names[0]
		ft, ok := m1.Type.(*ast.FuncType)
		if ok {
			opdata, err := getOpData(sdet, op.Name, ft)
			opdata.Description = comments
			if err != nil {
				return fmt.Errorf("Error with interface. Err = %s", err.Error())
			}
			sdet.Operations = append(sdet.Operations, &opdata)
		}
	}
	return nil
}

func extractDocumentation(comms ...*ast.CommentGroup) string {
	comments := ""
	for _,comm := range comms {
		comments = concatCommentGroup(comments,comm)
	}
	return comments
}

func concatCommentGroup(comments string, group *ast.CommentGroup)string{
	if group == nil {
		return comments
	}
	for _,comment := range group.List{
		s := comment.Text
		if len(s) > 2 {
			s = s[2:] // skip the "//"
		}
		comments += s
	}
	return comments
}

func getOpData(sdet *Servicedetail, op string, ft *ast.FuncType) (Operationdetail, error) {
	paramDetails, unqualifiedRequestPayloadType,requestPayloadType, requestPayloadDefaultValue := parseFields(op, ft.Params)
	if len(paramDetails) == 0 || paramDetails[0].Name != "ctx" ||
		paramDetails[0].Type != "context.Context" {
		return Operationdetail{},
			fmt.Errorf("First parameter of function %s must be ctx context.Context", op)
	}

	respDetails, unqualifiedResponsePayloadType,responsePayloadType, responsePayloadDefaultValue := parseFields(op, ft.Results)
	if len(respDetails) != 2 || respDetails[1].Type != "error" {
		return Operationdetail{},
			fmt.Errorf("function %s must return 2 values and the second one must be of type error", op)
	}

	method := "POST"
	if requestPayloadType == "" {
		method = "GET"
		sdet.DoesServiceHaveGetOperations = true // service has at least one GET operation
	}

	return Operationdetail{
		Operation:                   op,
		Params:                      paramDetails,
		Results:                     respDetails,
		UnqualifiedRequestPayload:   unqualifiedRequestPayloadType,
		UnqualifiedResponsePayload:  unqualifiedResponsePayloadType,
		RequestPayload:              requestPayloadType,
		ResponsePayload:             responsePayloadType,
		RequestPayloadLower:         strcase.ToLowerCamel(unqualifiedRequestPayloadType),
		ResponsePayloadLower:        strcase.ToLowerCamel(unqualifiedResponsePayloadType),
		RequestPayloadDefaultValue:  requestPayloadDefaultValue,
		ResponsePayloadDefaultValue: responsePayloadDefaultValue,
		URL:                         strcase.ToDelimited(op, '-'),
		Method:                      method,
	}, nil
}

func extractStuff(tag *ast.BasicLit){
	if tag == nil {
		return
	}
	log.Printf("extractStuff: the tag is %s",tag.Value)
}

// returns the details of all fields as well as the type for the
// field that is of kind payload
func parseFields(op string, fl *ast.FieldList) ([]Fielddetail, string, string, string) {
	var pd []Fielddetail
	var payloadType = ""
	var payloadDefaultValue = ""
	var unqualifiedPayloadType = ""
	var pointerType = false
	for _, m1 := range fl.List {
		var name = ""
		extractStuff(m1.Tag)
		comments := extractDocumentation(m1.Doc,m1.Comment)
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
		qualifiedType := correctPayloadType(varType, origin, pointerType)
		kind, defaultValue := getKindOfDefaultValue(op, name, qualifiedType, pointerType)
		pd = append(pd, Fielddetail{
			Name: name,
			Description:comments,
			UnqualifiedType: varType,
			Type: qualifiedType,
			Kind: kind,
			Origin: origin,
			DefaultValue: defaultValue,
			PointerType: pointerType,
		})
		if origin == "fw.PAYLOAD" {
			unqualifiedPayloadType = varType
			payloadType = qualifiedType
			payloadDefaultValue = defaultValue
		}
	}

	return pd, unqualifiedPayloadType,payloadType, payloadDefaultValue
}

// correct the payload type to reflect the correct value.
// the AST does not seem to give us the fully qualified type name
func correctPayloadType(typ string, origin string, pointerType bool) string {
	switch origin {
	case "fw.CONTEXT":
		return "context.Context"
	case "fw.PAYLOAD":
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
		return "fw.HEADER"
	case "Context":
		return "fw.CONTEXT"
	case "error":
		return "error"
	default:
		return "fw.PAYLOAD"
	}
}
