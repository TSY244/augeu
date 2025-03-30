package utils

import "encoding/json"

func JsonToStruct(jsonStr []byte, structPtr any) error {
	return json.Unmarshal(jsonStr, structPtr)
}

func StructToJson(structPtr any) (string, error) {
	jsonBytes, err := json.Marshal(structPtr)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
