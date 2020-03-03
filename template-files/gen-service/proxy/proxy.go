package proxy
	{{ with $service := .}}
	import (
		"context"
	
		bplus "gitlab.intelligentb.com/devops/bplus/http"
		api "{{$service.URL}}/api"
		e "{{$service.URL}}/internal/err"
	)

	type {{$service.CamelCase}} struct {}

{{range $opindex,$op := $service.Operations}}
// {{$op.Operation}} - proxies the {{$op.Operation}} and calls the server using HTTP
func ({{$service.CamelCase}}) {{$op.Operation}}({{range $index,$val := $op.Params}}{{if $index}},{{end}}{{$val.Name}} {{$val.Type}}{{end}})({{range $index,$val := $op.Results}}{{if $index}},{{end}}{{$val.Type}}{{end}}){
	resp, err := bplus.ProxyRequest(ctx, "{{$service.Name}}", "{{$op.Operation}}" {{range $index,$val := $op.Params}}{{if $index}},{{$val.Name}}{{end}}{{end}})
	if err != nil {
		return {{$op.ResponsePayload}}{}, e.MakeBplusError(ctx, e.CannotInvokeOperation,map[string]interface{}{
			"Service":"{{$service.Name}}", 
			"Operation":"{{$op.Operation}}", 
			"Error":err.Error(),})
	}
	r, ok := resp.(*{{$op.ResponsePayload}})
	if ok {
		return *r,nil
	}

	return {{$op.ResponsePayload}}{}, e.MakeBplusError(ctx, e.CannotInvokeOperation,map[string]interface{}{
		"Service":"{{$service.Name}}", "Operation":"{{$op.Operation}}", 
		"Error": err.Error(),})

}
{{end}}
{{end}}