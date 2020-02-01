package service
{{ with $service := .}}
import (
	"context"
	api "{{$service.URL}}/api"
)
type {{$service.CamelCaseLower}} struct {}

func Make{{$service.CamelCase}}Service(){{$service.CamelCaseLower}}{
	return {{$service.CamelCaseLower}}{}
}
{{range $indexop,$op := $service.Operations}}
func ({{$service.CamelCaseLower}}){{$op.Operation}}({{range $index,$val := $op.Params}}{{if $index}},{{end}}{{$val.Name}} {{$val.Type}}{{end}})({{$op.ResponsePayload}},error){
	return {{$op.ResponsePayload}}{},nil
}
{{end}}
{{end}}