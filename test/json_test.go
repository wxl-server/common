package test

import (
	"log"
	"testing"

	"github.com/wxl-server/common/json"
)

func TestJson(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	testStruct := TestStruct{
		Name: "test",
		Age:  18,
	}
	jsonStr := json.MarshalWithoutError[string](testStruct)
	log.Print(jsonStr)
	testStruct2 := json.UnmarshalWithoutError[TestStruct](jsonStr)
	log.Print(json.MarshalWithoutError[string](testStruct2))
	testStruct3 := json.UnmarshalWithoutError[*TestStruct](jsonStr)
	log.Print(json.MarshalWithoutError[string](testStruct3))

	type TestStruct2 struct {
		Named string `json:"named"`
		Ages  int    `json:"ages"`
	}
	testStruct4 := json.UnmarshalWithoutError[TestStruct2](jsonStr)
	log.Print(json.MarshalWithoutError[string](testStruct4))

	var nilSlice []string
	jsonStr = json.MarshalWithoutError[string](nilSlice)
	log.Print(jsonStr)

	var nilMap map[string]string
	jsonStr = json.MarshalWithoutError[string](nilMap)
	log.Print(jsonStr)

	var nilStruct *TestStruct
	jsonStr = json.MarshalWithoutError[string](nilStruct)
	log.Print(jsonStr)
}
