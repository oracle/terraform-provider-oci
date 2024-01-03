// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// DetectorRuleResourceTypeEnum Enum with underlying type: string
type DetectorRuleResourceTypeEnum string

// Set of constants representing the allowable values for DetectorRuleResourceTypeEnum
const (
	DetectorRuleResourceTypeUser             DetectorRuleResourceTypeEnum = "USER"
	DetectorRuleResourceTypeComputeInstance  DetectorRuleResourceTypeEnum = "COMPUTE_INSTANCE"
	DetectorRuleResourceTypeOtherOciResource DetectorRuleResourceTypeEnum = "OTHER_OCI_RESOURCE"
	DetectorRuleResourceTypeExternalIp       DetectorRuleResourceTypeEnum = "EXTERNAL_IP"
	DetectorRuleResourceTypeInternalIp       DetectorRuleResourceTypeEnum = "INTERNAL_IP"
	DetectorRuleResourceTypeCustom           DetectorRuleResourceTypeEnum = "CUSTOM"
)

var mappingDetectorRuleResourceTypeEnum = map[string]DetectorRuleResourceTypeEnum{
	"USER":               DetectorRuleResourceTypeUser,
	"COMPUTE_INSTANCE":   DetectorRuleResourceTypeComputeInstance,
	"OTHER_OCI_RESOURCE": DetectorRuleResourceTypeOtherOciResource,
	"EXTERNAL_IP":        DetectorRuleResourceTypeExternalIp,
	"INTERNAL_IP":        DetectorRuleResourceTypeInternalIp,
	"CUSTOM":             DetectorRuleResourceTypeCustom,
}

var mappingDetectorRuleResourceTypeEnumLowerCase = map[string]DetectorRuleResourceTypeEnum{
	"user":               DetectorRuleResourceTypeUser,
	"compute_instance":   DetectorRuleResourceTypeComputeInstance,
	"other_oci_resource": DetectorRuleResourceTypeOtherOciResource,
	"external_ip":        DetectorRuleResourceTypeExternalIp,
	"internal_ip":        DetectorRuleResourceTypeInternalIp,
	"custom":             DetectorRuleResourceTypeCustom,
}

// GetDetectorRuleResourceTypeEnumValues Enumerates the set of values for DetectorRuleResourceTypeEnum
func GetDetectorRuleResourceTypeEnumValues() []DetectorRuleResourceTypeEnum {
	values := make([]DetectorRuleResourceTypeEnum, 0)
	for _, v := range mappingDetectorRuleResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectorRuleResourceTypeEnumStringValues Enumerates the set of values in String for DetectorRuleResourceTypeEnum
func GetDetectorRuleResourceTypeEnumStringValues() []string {
	return []string{
		"USER",
		"COMPUTE_INSTANCE",
		"OTHER_OCI_RESOURCE",
		"EXTERNAL_IP",
		"INTERNAL_IP",
		"CUSTOM",
	}
}

// GetMappingDetectorRuleResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectorRuleResourceTypeEnum(val string) (DetectorRuleResourceTypeEnum, bool) {
	enum, ok := mappingDetectorRuleResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
