// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"encoding/json"
	"fmt"
)

func JsonStringToMap(jsonString string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("invalid json string %s", jsonString)
	}
	return jsonMap, nil
}

func MapToJsonString(jsonMap map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(jsonMap)
	if err != nil {
		return "", fmt.Errorf("error during json conversion")
	}
	return string(bytes), nil
}
