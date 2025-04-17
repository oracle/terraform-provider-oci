// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"strings"
)

// OccmDemandSignalItemRequestTypeEnum Enum with underlying type: string
type OccmDemandSignalItemRequestTypeEnum string

// Set of constants representing the allowable values for OccmDemandSignalItemRequestTypeEnum
const (
	OccmDemandSignalItemRequestTypeDemand OccmDemandSignalItemRequestTypeEnum = "DEMAND"
)

var mappingOccmDemandSignalItemRequestTypeEnum = map[string]OccmDemandSignalItemRequestTypeEnum{
	"DEMAND": OccmDemandSignalItemRequestTypeDemand,
}

var mappingOccmDemandSignalItemRequestTypeEnumLowerCase = map[string]OccmDemandSignalItemRequestTypeEnum{
	"demand": OccmDemandSignalItemRequestTypeDemand,
}

// GetOccmDemandSignalItemRequestTypeEnumValues Enumerates the set of values for OccmDemandSignalItemRequestTypeEnum
func GetOccmDemandSignalItemRequestTypeEnumValues() []OccmDemandSignalItemRequestTypeEnum {
	values := make([]OccmDemandSignalItemRequestTypeEnum, 0)
	for _, v := range mappingOccmDemandSignalItemRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOccmDemandSignalItemRequestTypeEnumStringValues Enumerates the set of values in String for OccmDemandSignalItemRequestTypeEnum
func GetOccmDemandSignalItemRequestTypeEnumStringValues() []string {
	return []string{
		"DEMAND",
	}
}

// GetMappingOccmDemandSignalItemRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccmDemandSignalItemRequestTypeEnum(val string) (OccmDemandSignalItemRequestTypeEnum, bool) {
	enum, ok := mappingOccmDemandSignalItemRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
