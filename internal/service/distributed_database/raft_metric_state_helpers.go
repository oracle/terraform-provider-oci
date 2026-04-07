// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"encoding/json"
	"fmt"
)

// TF_CODE_GEN: TERSI-4920-TOP-28 raft metric specs define additionalProperties
// objects, but the generated datasource schema flattens them as
// TypeMap(TypeString). Marshal nested values to JSON strings so Terraform can
// persist the dynamic payload instead of dropping state.
func flattenRaftMetricStateMap(values map[string]interface{}) (map[string]string, error) {
	result := map[string]string{}

	for key, value := range values {
		if value == nil {
			result[key] = "{}"
			continue
		}

		switch typedValue := value.(type) {
		case string:
			result[key] = typedValue
		default:
			marshaledValue, err := json.Marshal(typedValue)
			if err != nil {
				return nil, fmt.Errorf("marshal raft metric value for key %q: %w", key, err)
			}
			result[key] = string(marshaledValue)
		}
	}

	return result, nil
}
