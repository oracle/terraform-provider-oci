// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeDblmSubscribe            OperationTypeEnum = "DBLM_SUBSCRIBE"
	OperationTypeDblmEnable               OperationTypeEnum = "DBLM_ENABLE"
	OperationTypeScanVulCve               OperationTypeEnum = "SCAN_VUL_CVE"
	OperationTypeScanVulPatch             OperationTypeEnum = "SCAN_VUL_PATCH"
	OperationTypeScanVulImage             OperationTypeEnum = "SCAN_VUL_IMAGE"
	OperationTypeUpdateSubscription       OperationTypeEnum = "UPDATE_SUBSCRIPTION"
	OperationTypeUpdateSubscribedResource OperationTypeEnum = "UPDATE_SUBSCRIBED_RESOURCE"
	OperationTypeCreateSharedDatastore    OperationTypeEnum = "CREATE_SHARED_DATASTORE"
	OperationTypeDeleteSharedDatastore    OperationTypeEnum = "DELETE_SHARED_DATASTORE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"DBLM_SUBSCRIBE":             OperationTypeDblmSubscribe,
	"DBLM_ENABLE":                OperationTypeDblmEnable,
	"SCAN_VUL_CVE":               OperationTypeScanVulCve,
	"SCAN_VUL_PATCH":             OperationTypeScanVulPatch,
	"SCAN_VUL_IMAGE":             OperationTypeScanVulImage,
	"UPDATE_SUBSCRIPTION":        OperationTypeUpdateSubscription,
	"UPDATE_SUBSCRIBED_RESOURCE": OperationTypeUpdateSubscribedResource,
	"CREATE_SHARED_DATASTORE":    OperationTypeCreateSharedDatastore,
	"DELETE_SHARED_DATASTORE":    OperationTypeDeleteSharedDatastore,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"dblm_subscribe":             OperationTypeDblmSubscribe,
	"dblm_enable":                OperationTypeDblmEnable,
	"scan_vul_cve":               OperationTypeScanVulCve,
	"scan_vul_patch":             OperationTypeScanVulPatch,
	"scan_vul_image":             OperationTypeScanVulImage,
	"update_subscription":        OperationTypeUpdateSubscription,
	"update_subscribed_resource": OperationTypeUpdateSubscribedResource,
	"create_shared_datastore":    OperationTypeCreateSharedDatastore,
	"delete_shared_datastore":    OperationTypeDeleteSharedDatastore,
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
		"DBLM_SUBSCRIBE",
		"DBLM_ENABLE",
		"SCAN_VUL_CVE",
		"SCAN_VUL_PATCH",
		"SCAN_VUL_IMAGE",
		"UPDATE_SUBSCRIPTION",
		"UPDATE_SUBSCRIBED_RESOURCE",
		"CREATE_SHARED_DATASTORE",
		"DELETE_SHARED_DATASTORE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
