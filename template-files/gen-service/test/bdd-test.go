package test
	
	import (
		"context"
		"fmt"
	
		"github.com/DATA-DOG/godog"
		_ "{{.URL}}" // init module
		"{{.URL}}/api"
		"{{.URL}}/proxy"
		e "{{.URL}}/internal/err"
	)
	type {{.CamelCase}} struct {
		{{ range $index,$op := .Operations}}
		{{$op.ResponsePayloadLower}}  api.{{$op.ResponsePayload}}
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
	func ({{$service.CamelCaseLower}} {{$service.CamelCase}})iInvoke{{$op.Operation}}With(){
		// Construct the request with the arguments passed
 
		resp, err := {{$service.CamelCaseLower}}Proxy.{{$op.Operation}}(context.TODO() {{range $index,$val := .Params}}{{if $index}},{{$val.Name}}{{end}}{{end}})
		if err != nil {
			return e.MakeBplusError(ctx, e.CannotInvokeService,"{{$service.CamelCase}}", "{{.Operation}}", err.Error())
		}
		{{$service.CamelCaseLower}}.{{$op.ResponsePayloadLower}} = resp
		return nil
	}
	{{end}}
	{{end}}
