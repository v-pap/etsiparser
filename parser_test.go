package etsiparser

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

func TestNonExistingAttribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"cores"}
	expectedOutput := `null`
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

func TestPartialMatchingAttribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"parts/cores"}
	expectedOutput := `null`
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

func TestBothExistingAndNonExistingAttribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `
	[
	{"parts":[{"id":1}, {"id":2}]},
	{"parts":[{"id":3}, {"id":4}]}
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

func TestNullInput(t *testing.T) {
	input := `null`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `null`
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

func TestEmptyMapInput(t *testing.T) {
	input := `{}`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `null`
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

func TestEmptyListInput(t *testing.T) {
	input := `[]`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `null`
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

func TestIntegerInput(t *testing.T) {
	input := `1`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `null`
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

func TestStringInput(t *testing.T) {
	input := `"hello world"`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `null`
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

func TestNestedListInput(t *testing.T) {
	input := `[[[{"id":1, "color":"red"}, {"id":2, "color":"green"}]]]`
	attributes := []string{"id"}
	expectedOutput := `[[[{"id":1}, {"id":2}]]]`
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

func TestEmptyAttributes(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{}
	expectedOutput := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
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
