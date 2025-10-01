// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateResourceAnalyticsInstance OperationTypeEnum = "CREATE_RESOURCE_ANALYTICS_INSTANCE"
	OperationTypeUpdateResourceAnalyticsInstance OperationTypeEnum = "UPDATE_RESOURCE_ANALYTICS_INSTANCE"
	OperationTypeDeleteResourceAnalyticsInstance OperationTypeEnum = "DELETE_RESOURCE_ANALYTICS_INSTANCE"
	OperationTypeMoveResourceAnalyticsInstance   OperationTypeEnum = "MOVE_RESOURCE_ANALYTICS_INSTANCE"
	OperationTypeCreateTenancyAttachment         OperationTypeEnum = "CREATE_TENANCY_ATTACHMENT"
	OperationTypeUpdateTenancyAttachment         OperationTypeEnum = "UPDATE_TENANCY_ATTACHMENT"
	OperationTypeDeleteTenancyAttachment         OperationTypeEnum = "DELETE_TENANCY_ATTACHMENT"
	OperationTypeCreateMonitoredRegion           OperationTypeEnum = "CREATE_MONITORED_REGION"
	OperationTypeDeleteMonitoredRegion           OperationTypeEnum = "DELETE_MONITORED_REGION"
	OperationTypeEnableOac                       OperationTypeEnum = "ENABLE_OAC"
	OperationTypeDisableOac                      OperationTypeEnum = "DISABLE_OAC"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_RESOURCE_ANALYTICS_INSTANCE": OperationTypeCreateResourceAnalyticsInstance,
	"UPDATE_RESOURCE_ANALYTICS_INSTANCE": OperationTypeUpdateResourceAnalyticsInstance,
	"DELETE_RESOURCE_ANALYTICS_INSTANCE": OperationTypeDeleteResourceAnalyticsInstance,
	"MOVE_RESOURCE_ANALYTICS_INSTANCE":   OperationTypeMoveResourceAnalyticsInstance,
	"CREATE_TENANCY_ATTACHMENT":          OperationTypeCreateTenancyAttachment,
	"UPDATE_TENANCY_ATTACHMENT":          OperationTypeUpdateTenancyAttachment,
	"DELETE_TENANCY_ATTACHMENT":          OperationTypeDeleteTenancyAttachment,
	"CREATE_MONITORED_REGION":            OperationTypeCreateMonitoredRegion,
	"DELETE_MONITORED_REGION":            OperationTypeDeleteMonitoredRegion,
	"ENABLE_OAC":                         OperationTypeEnableOac,
	"DISABLE_OAC":                        OperationTypeDisableOac,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_resource_analytics_instance": OperationTypeCreateResourceAnalyticsInstance,
	"update_resource_analytics_instance": OperationTypeUpdateResourceAnalyticsInstance,
	"delete_resource_analytics_instance": OperationTypeDeleteResourceAnalyticsInstance,
	"move_resource_analytics_instance":   OperationTypeMoveResourceAnalyticsInstance,
	"create_tenancy_attachment":          OperationTypeCreateTenancyAttachment,
	"update_tenancy_attachment":          OperationTypeUpdateTenancyAttachment,
	"delete_tenancy_attachment":          OperationTypeDeleteTenancyAttachment,
	"create_monitored_region":            OperationTypeCreateMonitoredRegion,
	"delete_monitored_region":            OperationTypeDeleteMonitoredRegion,
	"enable_oac":                         OperationTypeEnableOac,
	"disable_oac":                        OperationTypeDisableOac,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_RESOURCE_ANALYTICS_INSTANCE",
		"UPDATE_RESOURCE_ANALYTICS_INSTANCE",
		"DELETE_RESOURCE_ANALYTICS_INSTANCE",
		"MOVE_RESOURCE_ANALYTICS_INSTANCE",
		"CREATE_TENANCY_ATTACHMENT",
		"UPDATE_TENANCY_ATTACHMENT",
		"DELETE_TENANCY_ATTACHMENT",
		"CREATE_MONITORED_REGION",
		"DELETE_MONITORED_REGION",
		"ENABLE_OAC",
		"DISABLE_OAC",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
