package utility

import (
	"encoding/json"
	"log"
	"strings"
)

func QueryLike(value string) string {
	if len(value) > 0 {
		return "%" + strings.ToLower(value) + "%"
	} else {
		return value
	}
}

func ToJsonString(object interface{}) string {
	result := "{}"
	if object != nil {
		jsonByte, err := json.Marshal(object)
		log.Println(err)
		result = string(jsonByte)
	}
	return result
}

func JsonStringToMap(dataString string) map[string]interface{} {
	result := map[string]interface{}{}
	if err := json.Unmarshal([]byte(dataString), &result); err != nil {
		panic(err)
	}
	return result
}
