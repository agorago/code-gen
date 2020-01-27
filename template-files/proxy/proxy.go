package proxy
	{{ with $service := .}}
	import (
		"context"
		"fmt"
	
		bplus "github.com/MenaEnergyVentures/bplus/http"
		api "{{$service.URL}}/api"
		e "{{$service.URL}}/internal/err"
	)

	type {{$service.CamelCaseLower}} struct {}

{{range $opindex,$op := $service.Operations}}
// {{$op.Operation}} - proxies the {{$op.Operation}} and calls the server using HTTP
func ({{$service.CamelCaseLower}}) {{$op.Operation}}({{range $index,$val := $op.Params}}{{if $index}},{{end}}{{$val.Name}} {{$val.Type}}{{end}})({{range $index,$val := $op.Results}}{{if $index}},{{end}}{{$val.Type}}{{end}}){
	resp, err := bplus.ProxyRequest(ctx, "{{$service.Name}}", "{{$op.Operation}}" {{range $index,$val := $op.Params}}{{if $index}},{{$val.Name}}{{end}}{{end}})
	if err != nil {
		return api.{{$op.ResponsePayload}}{}, e.MakeBplusError(ctx, e.CannotInvokeService,"{{$service.Name}}", "{{$op.Operation}}", err.Error())
	}
	r, ok := resp.(*api.{{$op.ResponsePayload}})
	if ok {
		return r,nil
	}

	return api.{{$op.ResponsePayload}}{}, e.MakeBplusError(ctx, e.CannotInvokeService,"{{$service.Name}}", "{{$op.Operation}}", err.Error())

}
{{end}}
{{end}}