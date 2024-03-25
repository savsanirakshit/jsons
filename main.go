package jsons

import (
	"encoding/json"
	"github.com/PaesslerAG/jsonpath"
	"reflect"
	"strings"
)

// ConvertToMap converts JSON data byte slice to a map[string]interface{}
func ConvertToMap(data []byte) (map[string]interface{}, error) {
	jsonValue := map[string]interface{}{}
	err := json.Unmarshal(data, &jsonValue)
	return jsonValue, err
}

// Get retrieves the value corresponding to the specified key path from the provided JSON data.
func Get(jsonData map[string]interface{}, key string) (interface{}, error) {
	var value interface{}
	var err error
	if strings.Contains(key, "$") {
		key = strings.ReplaceAll(key, "${", "")
		key = strings.ReplaceAll(key, "}", "")
		if strings.Contains(key, ".") {
			val, err := jsonpath.Get(key, jsonData)
			if err == nil {
				value = val
			} else {
				return nil, err
			}
		} else {
			value = jsonData[key]
		}
	} else {
		value = jsonData[key]
	}
	return value, err
}

// Set updates the value at the specified key path within the provided JSON data.
func Set(jsonMap map[string]interface{}, key string, value interface{}) {
	if strings.Contains(key, "$") {
		key = strings.ReplaceAll(key, "${", "")
		key = strings.ReplaceAll(key, "}", "")
		key = strings.Trim(key, " ")
		if strings.Contains(key, ".") {
			keyList := strings.Split(key, ".")
			setValueToJson(jsonMap, keyList, value)
		} else {
			jsonMap[key] = value
		}
	} else {
		jsonMap[key] = value
	}
}

func setValueToJson(jsonMap map[string]interface{}, path []string, value interface{}) {
	if len(path) == 1 {
		jsonMap[path[0]] = value
		return
	}

	if _, ok := jsonMap[path[0]]; !ok {
		jsonMap[path[0]] = make(map[string]interface{})
	}

	setValueToJson(jsonMap[path[0]].(map[string]interface{}), path[1:], value)
}

// Delete removes the value at the specified key path from the provided JSON data.
func Delete(jsonData map[string]interface{}, key string) (map[string]interface{}, error) {
	var err error
	if strings.Contains(key, "$") {
		key = strings.ReplaceAll(key, "${", "")
		key = strings.ReplaceAll(key, "}", "")
		if strings.Contains(key, ".") {
			keyList := strings.Split(key, ".")
			if len(keyList) == 2 {
				val, err := jsonpath.Get(key, jsonData)
				if err == nil && val != nil {
					val, err := jsonpath.Get(keyList[0], jsonData)
					if err == nil && val != nil && strings.Contains(reflect.TypeOf(val).Name(), "map") {
						tempVal := val.(map[string]interface{})
						delete(tempVal, keyList[1])
						jsonData[keyList[0]] = tempVal
					}
				}
			} else if len(keyList) > 2 {
				val, err := jsonpath.Get(key, jsonData)
				if err == nil && val != nil {
					tempVal := val.(map[string]interface{})
					for i := len(keyList) - 2; i >= 0; i++ {
						val, err := jsonpath.Get(strings.Join(keyList[:i], "."), jsonData)
						if err == nil && val != nil {
							tempMap := val.(map[string]interface{})
							delete(tempMap, keyList[i+1])
							tempVal = tempMap
						}
					}
					jsonData[keyList[0]] = tempVal
				}
			}
		} else {
			delete(jsonData, key)
		}
	} else {
		delete(jsonData, key)
	}
	return jsonData, err
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
