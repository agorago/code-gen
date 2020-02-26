package test
	
	import (
		"context"
	
		"github.com/DATA-DOG/godog"
		_ "{{.URL}}" // init module
		api "{{.URL}}/api"
		"{{.URL}}/proxy"
		e "{{.URL}}/internal/err"
	)
	type {{.CamelCase}} struct {
		{{ range $index,$op := .Operations}}
		{{$op.ResponsePayloadLower}}  {{$op.ResponsePayload}}
		{{end}}
	}
	var {{.CamelCaseLower}} = {{.CamelCase}}{}
	var {{.CamelCaseLower}}Proxy = proxy.{{.CamelCase}}{}
	
	func FeatureContext(s *godog.Suite) {
		{{with $service := .}}
		{{ range $index,$op := .Operations}}
		s.Step("^I invoke {{$op.Operation}} with ...<fill args here>$", {{$service.CamelCaseLower}}.iInvoke{{$op.Operation}}With)
		{{end}}
		
	}

	{{ range $index,$op := .Operations}}
	func ({{$service.CamelCaseLower}} {{$service.CamelCase}})iInvoke{{$op.Operation}}With()error{
		// Construct the request with the arguments passed
 
		resp, err := {{$service.CamelCaseLower}}Proxy.{{$op.Operation}}(context.TODO() {{range $index,$val := .Params}}{{if $index}},{{$val.DefaultValue}}{{end}}{{end}})
		if err != nil {
			return e.MakeBplusError(context.TODO(), e.CannotInvokeOperation,map[string]in terface{}{
				"Service": "{{$service.CamelCase}}", 
				"Operation":"{{.Operation}}", 
				"Error":err.Error(),}
			)
		}
		{{$service.CamelCaseLower}}.{{$op.ResponsePayloadLower}} = resp
		return nil
	}
	{{end}}
	{{end}}
