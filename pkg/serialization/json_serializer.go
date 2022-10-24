package serialization

import (
	"encoding/json"
	"fmt"
	"log"
)

func ToJson(object interface{}) (*string, *string) {
	var result, errorMessage string
	jsonData, err := json.Marshal(object)

	if err != nil {
		errorMessage = fmt.Sprintf("Error Json Serializing Data. Error: %s", err.Error())
		log.Panicf(errorMessage)
		return nil, &errorMessage
	}

	result = string(jsonData)
	return &result, nil
}
