package etsiparser

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect1Attribute(t *testing.T) {
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

func TestExclude1Attribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"parts/color"}
	expectedOutput := `
	[
	{"id":123, "weight":100, "parts":[{"id":1}, {"id":2}]},
	{"id":456, "weight":500, "parts":[{"id":3}, {"id":4}]}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}
func TestSelect2Attributes(t *testing.T) {
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

func TestExclude2Attributes(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"parts/color", "parts/id"}
	expectedOutput := `
	[
	{"id":123, "weight":100, "parts":[{}, {}]},
	{"id":456, "weight":500, "parts":[{}, {}]}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectWithoutNestedLists(t *testing.T) {
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

func TestExcludeWithoutNestedLists(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":{"id":1, "color":"red"}},
	{"id":456, "weight":500, "parts":{"id":3, "color":"green"}}
	]
	`
	attributes := []string{"parts/color"}
	expectedOutput := `
	[
	{"id":123, "weight":100, "parts":{"id":1}},
	{"id":456, "weight":500, "parts":{"id":3}}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectWithListOfDifferentObjects(t *testing.T) {
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

func TestExcludeWithListOfDifferentObjects(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":{"id":3, "color":"green"}}
	]
	`
	attributes := []string{"parts/color", "parts/id"}
	expectedOutput := `
	[
	{"id":123, "weight":100, "parts":[{}, {}]},
	{"id":456, "weight":500, "parts":{}}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectWithSingleObjects(t *testing.T) {
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

func TestExcludeWithSingleObjects(t *testing.T) {
	input := `
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]}
	`
	attributes := []string{"parts/color", "parts/id"}
	expectedOutput := `
	{"id":123, "weight":100, "parts":[{}, {}]}
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectShallowAttribute(t *testing.T) {
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

func TestExcludeShallowAttribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"parts"}
	expectedOutput := `
	[
	{"id":123, "weight":100},
	{"id":456, "weight":500}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectNonExistingAttribute(t *testing.T) {
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

func TestExcludeNonExistingAttribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"cores"}
	expectedOutput := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectPartialMatchingAttribute(t *testing.T) {
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

func TestExcludePartialMatchingAttribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"parts/cores"}
	expectedOutput := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectBothExistingAndNonExistingAttribute(t *testing.T) {
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

func TestExcludeBothExistingAndNonExistingAttribute(t *testing.T) {
	input := `
	[
	{"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
	{"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
	]
	`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `
	[
	{"id":123, "weight":100, "parts":[{"color":"red"}, {"color":"green"}]},
	{"id":456, "weight":500, "parts":[{"color":"green"}, {"color":"blue"}]}
	]
	`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectNullInput(t *testing.T) {
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

func TestExcludeNullInput(t *testing.T) {
	input := `null`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `null`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectEmptyMapInput(t *testing.T) {
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

func TestExcludeEmptyMapInput(t *testing.T) {
	input := `{}`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `{}`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectEmptyListInput(t *testing.T) {
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

func TestExcludeEmptyListInput(t *testing.T) {
	input := `[]`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `[]`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectIntegerInput(t *testing.T) {
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

func TestExcludeIntegerInput(t *testing.T) {
	input := `1`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `1`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectStringInput(t *testing.T) {
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

func TestExcludeStringInput(t *testing.T) {
	input := `"hello world"`
	attributes := []string{"cores", "parts/id"}
	expectedOutput := `"hello world"`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectNestedListInput(t *testing.T) {
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

func TestExcludeNestedListInput(t *testing.T) {
	input := `[[[{"id":1, "color":"red"}, {"id":2, "color":"green"}]]]`
	attributes := []string{"id"}
	expectedOutput := `[[[{"color":"red"}, {"color":"green"}]]]`
	var payload interface{}
	err := json.Unmarshal([]byte(input), &payload)

	assert.Nil(t, err)

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}

func TestSelectEmptyAttributes(t *testing.T) {
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

func TestExcludeEmptyAttributes(t *testing.T) {
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

	modifiedPayload := ExcludeFields(attributes, payload)

	res, err := json.MarshalIndent(modifiedPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	assert.JSONEq(t, expectedOutput, string(res))
}
