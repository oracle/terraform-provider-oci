// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// License Manager API
//
// Use the License Manager API to manage product licenses and license records. For more information, see License Manager Overview (https://docs.cloud.oracle.com/iaas/Content/LicenseManager/Concepts/licensemanageroverview.htm).
//

package licensemanager

import (
	"strings"
)

// StatusEnum Enum with underlying type: string
type StatusEnum string

// Set of constants representing the allowable values for StatusEnum
const (
	StatusIncomplete  StatusEnum = "INCOMPLETE"
	StatusIssuesFound StatusEnum = "ISSUES_FOUND"
	StatusWarning     StatusEnum = "WARNING"
	StatusOk          StatusEnum = "OK"
)

var mappingStatusEnum = map[string]StatusEnum{
	"INCOMPLETE":   StatusIncomplete,
	"ISSUES_FOUND": StatusIssuesFound,
	"WARNING":      StatusWarning,
	"OK":           StatusOk,
}

var mappingStatusEnumLowerCase = map[string]StatusEnum{
	"incomplete":   StatusIncomplete,
	"issues_found": StatusIssuesFound,
	"warning":      StatusWarning,
	"ok":           StatusOk,
}

// GetStatusEnumValues Enumerates the set of values for StatusEnum
func GetStatusEnumValues() []StatusEnum {
	values := make([]StatusEnum, 0)
	for _, v := range mappingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetStatusEnumStringValues Enumerates the set of values in String for StatusEnum
func GetStatusEnumStringValues() []string {
	return []string{
		"INCOMPLETE",
		"ISSUES_FOUND",
		"WARNING",
		"OK",
	}
}

// GetMappingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStatusEnum(val string) (StatusEnum, bool) {
	enum, ok := mappingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
