package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

// Write a program to Map Flattening.
// input data
const inputJson = `{
	"name":"Piyush Singariya",
	"email":"xxxx@avesha.io",
	"password":"xxxxx",
	"phone_number": "+91907xxxxxx",
	"test": {
		"id":"1234",
		"tel": {
			"tags":["hello", "world"]
		}
	}
}`

// Output data
// {
// 	"name": "Piyush Singariya",
// 	"email": "xxxx@avesha.io",
// 	"password": "xxxxx",
// 	"phone_number": "+91907xxxxxx",
// 	"test_password":"1234",
// 	"test_tel_tags_0": "hello",
//  "test_tel_tags_1": "world"
// }

func Test_flatten(t *testing.T) {
	inputData := make(map[string]interface{})
	err := json.Unmarshal([]byte(inputJson), &inputData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inputData)

	outputData := make(map[string]interface{})
	flatten("", inputData, outputData)
	fmt.Println(outputData)
}

func flatten(prefix string, data map[string]interface{}, output map[string]interface{}) {
	if prefix != "" {
		prefix = prefix + "_"
	}

	for key, value := range data {
		switch value.(type) {
		case map[string]interface{}: // object
			flatten(prefix+key, value.(map[string]interface{}), output)
		case []interface{}: // array
			for i, v := range value.([]interface{}) {
				output[prefix+key+"_"+strconv.Itoa(i)] = v
			}
		default:
			output[prefix+key] = value
		}
	}
}
