package register
{{ with $service := .}}
import (
	"context"
	api "{{$service.URL}}/api"
	service "{{$service.URL}}/internal/service"
	bplus "github.com/MenaEnergyVentures/bplus/fw"
)
func init(){
	var sd =  bplus.ServiceDescriptor{
		ServiceToInvoke: service.Make{{$service.CamelCase}}Service(),
		Name:            "{{$service.Name}}",
		Operations:      OperationDescriptors(),
	}
	bplus.RegisterService("{{$service.Name}}", sd)
}

func OperationDescriptors()([]bplus.OperationDescriptor){
	return []bplus.OperationDescriptor{
		{{range $index,$elem := $service.Operations}}
		bplus.OperationDescriptor{
			Name:        "{{$elem.Operation}}",
			URL:             "/{{$elem.URL}}",
			HTTPMethod:      "{{$elem.Method}}",
			OpRequestMaker:  make{{$elem.Operation}}Request,
			OpResponseMaker: make{{$elem.Operation}}Response,
			Params:          {{$elem.Operation}}PD(),
		},
		{{end}}
	}
}

{{range $index,$elem := $service.Operations}}
func {{$elem.Operation}}PD()([]bplus.ParamDescriptor){
	
	return []bplus.ParamDescriptor{
		{{range $index,$p := $elem.Params}}
		bplus.ParamDescriptor{
			Name:        "{{$p.Name}}",
			ParamOrigin: {{$p.Origin}},
			{{if $p.Kind -}} ParamKind: {{$p.Kind}}, {{- end}}
		},
		{{end}}
	}
}
{{end}}

{{range $index,$value := $service.Operations}}
{{if $value.RequestPayload -}}
func make{{$value.Operation}}Request(ctx context.Context)(interface{},error){
	return {{$value.RequestPayload}}{},nil
}
{{- end}}

{{if $value.ResponsePayload -}}
func make{{$value.Operation}}Response(ctx context.Context)(interface{},error){
	return {{$value.ResponsePayload}}{},nil
}
{{- end}}
{{- end}}

{{- end}} 
	