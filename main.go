package main

import (
	"strings"
)

func selectFieldsRecursively(attributesMap map[string]interface{}, object interface{}) interface{} {
	if len(attributesMap) == 0 {
		return object
	}
	switch o := object.(type) {
	case map[string]interface{}:
		resultMap := make(map[string]interface{})
		for field := range attributesMap {
			newAttributesMap := attributesMap[field].(map[string]interface{})
			if value, ok := o[field]; ok {
				returnedValue := selectFieldsRecursively(newAttributesMap, value)
				if returnedValue != nil {
					resultMap[field] = returnedValue
				}
			}
		}
		if len(resultMap) > 0 {
			return resultMap
		}
		return nil
	case []interface{}:
		var returnedValues []interface{}
		for _, item := range o {
			value := selectFieldsRecursively(attributesMap, item)
			if value != nil {
				returnedValues = append(returnedValues, value)
			}
		}
		if len(returnedValues) > 0 {
			return returnedValues
		}
		return nil
	default:
		//fmt.Println(o)
	}
	return nil
}

func createAttributesMap(attributesList []string) map[string]interface{} {
	attributesMap := make(map[string]interface{})
	for _, attributes := range attributesList {
		currentLayerMap := &attributesMap
		for _, attribute := range strings.Split(attributes, "/") {
			if _, ok := (*currentLayerMap)[attribute]; !ok {
				(*currentLayerMap)[attribute] = make(map[string]interface{})
			}
			object, _ := (*currentLayerMap)[attribute].(map[string]interface{})
			currentLayerMap = &object
		}
	}
	return attributesMap
}

func SelectFields(attributesList []string, data interface{}) interface{} {
	attributesMap := createAttributesMap(attributesList)
	if len(attributesMap) > 0 && data != nil {
		modifiedData := selectFieldsRecursively(attributesMap, data)
		return modifiedData
	}
	return data
}
