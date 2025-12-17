// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package batch

import (
	"fmt"

	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
)

func EnvironmentVariableToMap(obj oci_batch.EnvironmentVariable) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func FleetAssignmentPolicyToMap(obj oci_batch.FleetAssignmentPolicy) (map[string]interface{}, error) {
	if obj == nil {
		return nil, fmt.Errorf("FleetAssignmentPolicy cannot be nil")
	}

	result := map[string]interface{}{}
	switch v := obj.(type) {
	case oci_batch.BestFitFleetAssignmentPolicy:
		result["type"] = "BEST_FIT"
	case oci_batch.FlexFitFleetAssignmentPolicy:
		result["type"] = "FLEX_FIT"

		if v.Threshold != nil {
			result["threshold"] = float32(*v.Threshold)
		}
	default:
		return nil, fmt.Errorf("unknown FleetAssignmentPolicy type: %T", obj)
	}

	return result, nil
}
