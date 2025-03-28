// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateAnalyticsInstance                WorkRequestOperationTypeEnum = "CREATE_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeDeleteAnalyticsInstance                WorkRequestOperationTypeEnum = "DELETE_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeStartAnalyticsInstance                 WorkRequestOperationTypeEnum = "START_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeStopAnalyticsInstance                  WorkRequestOperationTypeEnum = "STOP_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeScaleAnalyticsInstance                 WorkRequestOperationTypeEnum = "SCALE_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeChangeAnalyticsInstanceCompartment     WorkRequestOperationTypeEnum = "CHANGE_ANALYTICS_INSTANCE_COMPARTMENT"
	WorkRequestOperationTypeChangeAnalyticsInstanceNetworkEndpoint WorkRequestOperationTypeEnum = "CHANGE_ANALYTICS_INSTANCE_NETWORK_ENDPOINT"
	WorkRequestOperationTypeCreateVanityUrl                        WorkRequestOperationTypeEnum = "CREATE_VANITY_URL"
	WorkRequestOperationTypeUpdateVanityUrl                        WorkRequestOperationTypeEnum = "UPDATE_VANITY_URL"
	WorkRequestOperationTypeDeleteVanityUrl                        WorkRequestOperationTypeEnum = "DELETE_VANITY_URL"
	WorkRequestOperationTypeCreatePrivateAccessChannel             WorkRequestOperationTypeEnum = "CREATE_PRIVATE_ACCESS_CHANNEL"
	WorkRequestOperationTypeUpdatePrivateAccessChannel             WorkRequestOperationTypeEnum = "UPDATE_PRIVATE_ACCESS_CHANNEL"
	WorkRequestOperationTypeDeletePrivateAccessChannel             WorkRequestOperationTypeEnum = "DELETE_PRIVATE_ACCESS_CHANNEL"
	WorkRequestOperationTypeUpdateInstanceEncryptionKey            WorkRequestOperationTypeEnum = "UPDATE_INSTANCE_ENCRYPTION_KEY"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"CREATE_ANALYTICS_INSTANCE":                  WorkRequestOperationTypeCreateAnalyticsInstance,
	"DELETE_ANALYTICS_INSTANCE":                  WorkRequestOperationTypeDeleteAnalyticsInstance,
	"START_ANALYTICS_INSTANCE":                   WorkRequestOperationTypeStartAnalyticsInstance,
	"STOP_ANALYTICS_INSTANCE":                    WorkRequestOperationTypeStopAnalyticsInstance,
	"SCALE_ANALYTICS_INSTANCE":                   WorkRequestOperationTypeScaleAnalyticsInstance,
	"CHANGE_ANALYTICS_INSTANCE_COMPARTMENT":      WorkRequestOperationTypeChangeAnalyticsInstanceCompartment,
	"CHANGE_ANALYTICS_INSTANCE_NETWORK_ENDPOINT": WorkRequestOperationTypeChangeAnalyticsInstanceNetworkEndpoint,
	"CREATE_VANITY_URL":                          WorkRequestOperationTypeCreateVanityUrl,
	"UPDATE_VANITY_URL":                          WorkRequestOperationTypeUpdateVanityUrl,
	"DELETE_VANITY_URL":                          WorkRequestOperationTypeDeleteVanityUrl,
	"CREATE_PRIVATE_ACCESS_CHANNEL":              WorkRequestOperationTypeCreatePrivateAccessChannel,
	"UPDATE_PRIVATE_ACCESS_CHANNEL":              WorkRequestOperationTypeUpdatePrivateAccessChannel,
	"DELETE_PRIVATE_ACCESS_CHANNEL":              WorkRequestOperationTypeDeletePrivateAccessChannel,
	"UPDATE_INSTANCE_ENCRYPTION_KEY":             WorkRequestOperationTypeUpdateInstanceEncryptionKey,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_analytics_instance":                  WorkRequestOperationTypeCreateAnalyticsInstance,
	"delete_analytics_instance":                  WorkRequestOperationTypeDeleteAnalyticsInstance,
	"start_analytics_instance":                   WorkRequestOperationTypeStartAnalyticsInstance,
	"stop_analytics_instance":                    WorkRequestOperationTypeStopAnalyticsInstance,
	"scale_analytics_instance":                   WorkRequestOperationTypeScaleAnalyticsInstance,
	"change_analytics_instance_compartment":      WorkRequestOperationTypeChangeAnalyticsInstanceCompartment,
	"change_analytics_instance_network_endpoint": WorkRequestOperationTypeChangeAnalyticsInstanceNetworkEndpoint,
	"create_vanity_url":                          WorkRequestOperationTypeCreateVanityUrl,
	"update_vanity_url":                          WorkRequestOperationTypeUpdateVanityUrl,
	"delete_vanity_url":                          WorkRequestOperationTypeDeleteVanityUrl,
	"create_private_access_channel":              WorkRequestOperationTypeCreatePrivateAccessChannel,
	"update_private_access_channel":              WorkRequestOperationTypeUpdatePrivateAccessChannel,
	"delete_private_access_channel":              WorkRequestOperationTypeDeletePrivateAccessChannel,
	"update_instance_encryption_key":             WorkRequestOperationTypeUpdateInstanceEncryptionKey,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_ANALYTICS_INSTANCE",
		"DELETE_ANALYTICS_INSTANCE",
		"START_ANALYTICS_INSTANCE",
		"STOP_ANALYTICS_INSTANCE",
		"SCALE_ANALYTICS_INSTANCE",
		"CHANGE_ANALYTICS_INSTANCE_COMPARTMENT",
		"CHANGE_ANALYTICS_INSTANCE_NETWORK_ENDPOINT",
		"CREATE_VANITY_URL",
		"UPDATE_VANITY_URL",
		"DELETE_VANITY_URL",
		"CREATE_PRIVATE_ACCESS_CHANNEL",
		"UPDATE_PRIVATE_ACCESS_CHANNEL",
		"DELETE_PRIVATE_ACCESS_CHANNEL",
		"UPDATE_INSTANCE_ENCRYPTION_KEY",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
