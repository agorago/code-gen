package util_test

import (
	"github.com/magiconair/properties/assert"
	"github.com/agorago/wego-gen/util"
	"testing"
)

func TestParseService(t *testing.T) {
	sdet := util.ParseService("stringdemo_test.go","xxx","10000")
	// _,_ := json.Marshal(sdet)
	//t.Errorf("%s\n",s)
	assert.Equal(t,sdet.InterfaceName,"StringDemoService")
	assert.Equal(t,sdet.Description," StringDemoService - the interface that is going to be implemented by the string demo service This has methods to illustrate features of the BPlus framework")
	assert.Equal(t,sdet.Name,"stringdemo_test")
	assert.Equal(t,sdet.CamelCase,"StringdemoTest")
	assert.Equal(t,sdet.CamelCaseLower,"stringdemoTest")
	assert.Equal(t,sdet.DoesServiceHaveGetOperations,true)
	assert.Equal(t,sdet.URL,"xxx")
	for index,op := range sdet.Operations {
		assertOp(t,index,op)
	}
	t.Fail()
}

func assertOp(t *testing.T, index int, op *util.Operationdetail) {
	if index == 0 {
		assert.Equal(t, op.Operation, "Uppercase")
		assert.Equal(t, op.Description, " Uppercase - Converts the input string into upper case")
		assert.Equal(t, op.ResponseDescription, " UpperCaseResponse - the  Uppercase service response")
		assert.Equal(t, op.RequestDescription, " UpperCaseRequest - the payload for Uppercase service")
		assert.Equal(t, op.UnqualifiedResponsePayload, "UpperCaseResponse")
		assert.Equal(t, op.UnqualifiedRequestPayload, "UpperCaseRequest")
		assert.Equal(t, op.RequestPayload, "*api.UpperCaseRequest")
		assert.Equal(t, op.ResponsePayload, "api.UpperCaseResponse")
		assert.Equal(t, op.Method, "POST")
		assert.Equal(t, op.URL, "uppercase")
		assert.Equal(t, op.RequestPayloadDefaultValue, "&api.UpperCaseRequest{}")
		assert.Equal(t, op.ResponsePayloadDefaultValue, "api.UpperCaseResponse{}")
		assert.Equal(t, op.RequestPayloadLower, "upperCaseRequest")
		assert.Equal(t, op.ResponsePayloadLower, "upperCaseResponse")
		for ind,param := range op.Params {
			assertParam(t,ind,param)
		}

		for ind,param := range op.Results {
			assertResult(t,ind,param)
		}
	}
}

func assertParam(t *testing.T,index int, param util.Fielddetail){
	if index == 0 {
		assert.Equal(t,param.Name,"ctx")
		assert.Equal(t,param.Description,"")
		assert.Equal(t,param.Origin,"bplus.CONTEXT")
		assert.Equal(t,param.Type,"context.Context")
		assert.Equal(t,param.DefaultValue,"context.Context{}")
		assert.Equal(t,param.Kind,"")
		assert.Equal(t,param.UnqualifiedType,"Context")
		assert.Equal(t,param.PointerType,false)
	}else if index == 1 {
		assert.Equal(t,param.Name,"ucr")
		assert.Equal(t,param.Description,"")
		assert.Equal(t,param.Origin,"bplus.PAYLOAD")
		assert.Equal(t,param.Type,"*api.UpperCaseRequest")
		assert.Equal(t,param.DefaultValue,"&api.UpperCaseRequest{}")
		assert.Equal(t,param.Kind,"")
		assert.Equal(t,param.UnqualifiedType,"UpperCaseRequest")
		assert.Equal(t,param.PointerType,true)
	}
}

func assertResult(t *testing.T,index int, param util.Fielddetail){
	if index == 0 {
		assert.Equal(t,param.Name,"")
		assert.Equal(t,param.Description,"")
		assert.Equal(t,param.Origin,"bplus.PAYLOAD")
		assert.Equal(t,param.Type,"api.UpperCaseResponse")
		assert.Equal(t,param.DefaultValue,"api.UpperCaseResponse{}")
		assert.Equal(t,param.Kind,"")
		assert.Equal(t,param.UnqualifiedType,"UpperCaseResponse")
		assert.Equal(t,param.PointerType,false)
	}else if index == 1 {
		assert.Equal(t,param.Name,"")
		assert.Equal(t,param.Description,"")
		assert.Equal(t,param.Origin,"error")
		assert.Equal(t,param.Type,"error")
		assert.Equal(t,param.DefaultValue,"nil")
		assert.Equal(t,param.Kind,"")
		assert.Equal(t,param.UnqualifiedType,"error")
		assert.Equal(t,param.PointerType,false)
	}
}