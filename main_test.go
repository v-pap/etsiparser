package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1Attribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"parts/color"}
	expectedOutput := `
	[
	{"parts":[{"color":"red"}, {"color":"green"}]},
	{"parts":[{"color":"green"}, {"color":"blue"}]}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := SelectFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}
func Test2Attributes(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"parts/color", "parts/id"}
	expectedOutput := `
	[
	{"parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := SelectFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestWithoutNestedLists(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":{"id":1, "color":"red"}},
	{"id":456, "weight":500, "parts":{"id":3, "color":"green"}}
	]
	`
	attributes := []string{"parts/color"}
	expectedOutput := `
	[
	{"parts":{"color":"red"}},
	{"parts":{"color":"green"}}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := SelectFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestWithListOfDifferentObjects(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":{"id":3, "color":"green"}}
	]
	`
	attributes := []string{"parts/color", "parts/id"}
	expectedOutput := `
	[
	{"parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"parts":{"id":3, "color":"green"}}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := SelectFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestWithSingleObjects(t *testing.T) {
	input := `
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]}
	`
	attributes := []string{"parts/color", "parts/id"}
	expectedOutput := `
	{"parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]}
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := SelectFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestShallowAttribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"parts"}
	expectedOutput := `
	[
	{"parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := SelectFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}
