// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// ReleaseTypeEnum Enum with underlying type: string
type ReleaseTypeEnum string

// Set of constants representing the allowable values for ReleaseTypeEnum
const (
	ReleaseTypeCpu          ReleaseTypeEnum = "CPU"
	ReleaseTypeFeature      ReleaseTypeEnum = "FEATURE"
	ReleaseTypeBpr          ReleaseTypeEnum = "BPR"
	ReleaseTypePatchRelease ReleaseTypeEnum = "PATCH_RELEASE"
)

var mappingReleaseTypeEnum = map[string]ReleaseTypeEnum{
	"CPU":           ReleaseTypeCpu,
	"FEATURE":       ReleaseTypeFeature,
	"BPR":           ReleaseTypeBpr,
	"PATCH_RELEASE": ReleaseTypePatchRelease,
}

var mappingReleaseTypeEnumLowerCase = map[string]ReleaseTypeEnum{
	"cpu":           ReleaseTypeCpu,
	"feature":       ReleaseTypeFeature,
	"bpr":           ReleaseTypeBpr,
	"patch_release": ReleaseTypePatchRelease,
}

// GetReleaseTypeEnumValues Enumerates the set of values for ReleaseTypeEnum
func GetReleaseTypeEnumValues() []ReleaseTypeEnum {
	values := make([]ReleaseTypeEnum, 0)
	for _, v := range mappingReleaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReleaseTypeEnumStringValues Enumerates the set of values in String for ReleaseTypeEnum
func GetReleaseTypeEnumStringValues() []string {
	return []string{
		"CPU",
		"FEATURE",
		"BPR",
		"PATCH_RELEASE",
	}
}

// GetMappingReleaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReleaseTypeEnum(val string) (ReleaseTypeEnum, bool) {
	enum, ok := mappingReleaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
