package register
{{ with $service := .}}
import (
	bplus "github.com/MenaEnergyVentures/bplus"
)
func init(){
	var sd =  bplus.ServiceDescriptor{
		ServiceToInvoke: service.Make{{$service.CamelCase}}Service(),
		Name:            "{{$service.Name}}",
		Operations:      OperationDescriptors(),
	}
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
func make{{$value.Operation}}Request()({{$value.RequestPayload}}){
	return {{$value.RequestPayload}}{}
}
{{- end}}

{{if $value.ResponsePayload -}}
func make{{$value.Operation}}Response()({{$value.ResponsePayload}}){
	return {{$value.ResponsePayload}}{}
}
{{- end}}
{{- end}}

{{- end}} 
	