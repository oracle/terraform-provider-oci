// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// ReportTypeEnum Enum with underlying type: string
type ReportTypeEnum string

// Set of constants representing the allowable values for ReportTypeEnum
const (
	ReportTypeErrata ReportTypeEnum = "ERRATA"
	ReportTypeCve    ReportTypeEnum = "CVE"
)

var mappingReportTypeEnum = map[string]ReportTypeEnum{
	"ERRATA": ReportTypeErrata,
	"CVE":    ReportTypeCve,
}

var mappingReportTypeEnumLowerCase = map[string]ReportTypeEnum{
	"errata": ReportTypeErrata,
	"cve":    ReportTypeCve,
}

// GetReportTypeEnumValues Enumerates the set of values for ReportTypeEnum
func GetReportTypeEnumValues() []ReportTypeEnum {
	values := make([]ReportTypeEnum, 0)
	for _, v := range mappingReportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReportTypeEnumStringValues Enumerates the set of values in String for ReportTypeEnum
func GetReportTypeEnumStringValues() []string {
	return []string{
		"ERRATA",
		"CVE",
	}
}

// GetMappingReportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportTypeEnum(val string) (ReportTypeEnum, bool) {
	enum, ok := mappingReportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
