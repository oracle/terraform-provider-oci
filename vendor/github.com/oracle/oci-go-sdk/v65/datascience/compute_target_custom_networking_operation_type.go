// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ComputeTargetCustomNetworkingOperationTypeEnum Enum with underlying type: string
type ComputeTargetCustomNetworkingOperationTypeEnum string

// Set of constants representing the allowable values for ComputeTargetCustomNetworkingOperationTypeEnum
const (
	ComputeTargetCustomNetworkingOperationTypeAttachSecondaryVnic ComputeTargetCustomNetworkingOperationTypeEnum = "ATTACH_SECONDARY_VNIC"
	ComputeTargetCustomNetworkingOperationTypeDetachSecondaryVnic ComputeTargetCustomNetworkingOperationTypeEnum = "DETACH_SECONDARY_VNIC"
)

var mappingComputeTargetCustomNetworkingOperationTypeEnum = map[string]ComputeTargetCustomNetworkingOperationTypeEnum{
	"ATTACH_SECONDARY_VNIC": ComputeTargetCustomNetworkingOperationTypeAttachSecondaryVnic,
	"DETACH_SECONDARY_VNIC": ComputeTargetCustomNetworkingOperationTypeDetachSecondaryVnic,
}

var mappingComputeTargetCustomNetworkingOperationTypeEnumLowerCase = map[string]ComputeTargetCustomNetworkingOperationTypeEnum{
	"attach_secondary_vnic": ComputeTargetCustomNetworkingOperationTypeAttachSecondaryVnic,
	"detach_secondary_vnic": ComputeTargetCustomNetworkingOperationTypeDetachSecondaryVnic,
}

// GetComputeTargetCustomNetworkingOperationTypeEnumValues Enumerates the set of values for ComputeTargetCustomNetworkingOperationTypeEnum
func GetComputeTargetCustomNetworkingOperationTypeEnumValues() []ComputeTargetCustomNetworkingOperationTypeEnum {
	values := make([]ComputeTargetCustomNetworkingOperationTypeEnum, 0)
	for _, v := range mappingComputeTargetCustomNetworkingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeTargetCustomNetworkingOperationTypeEnumStringValues Enumerates the set of values in String for ComputeTargetCustomNetworkingOperationTypeEnum
func GetComputeTargetCustomNetworkingOperationTypeEnumStringValues() []string {
	return []string{
		"ATTACH_SECONDARY_VNIC",
		"DETACH_SECONDARY_VNIC",
	}
}

// GetMappingComputeTargetCustomNetworkingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeTargetCustomNetworkingOperationTypeEnum(val string) (ComputeTargetCustomNetworkingOperationTypeEnum, bool) {
	enum, ok := mappingComputeTargetCustomNetworkingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
