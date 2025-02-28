package types

import "encoding/json"

func GetJsonAsMap(jsonValue string) (map[string]interface{}, error) {
	var mapData map[string]interface{}
	err := json.Unmarshal([]byte(jsonValue), &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

func DumpMapAsJson(mapData map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
