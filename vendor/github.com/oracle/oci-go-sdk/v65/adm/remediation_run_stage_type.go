// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"strings"
)

// RemediationRunStageTypeEnum Enum with underlying type: string
type RemediationRunStageTypeEnum string

// Set of constants representing the allowable values for RemediationRunStageTypeEnum
const (
	RemediationRunStageTypeDetect    RemediationRunStageTypeEnum = "DETECT"
	RemediationRunStageTypeRecommend RemediationRunStageTypeEnum = "RECOMMEND"
	RemediationRunStageTypeVerify    RemediationRunStageTypeEnum = "VERIFY"
	RemediationRunStageTypeApply     RemediationRunStageTypeEnum = "APPLY"
)

var mappingRemediationRunStageTypeEnum = map[string]RemediationRunStageTypeEnum{
	"DETECT":    RemediationRunStageTypeDetect,
	"RECOMMEND": RemediationRunStageTypeRecommend,
	"VERIFY":    RemediationRunStageTypeVerify,
	"APPLY":     RemediationRunStageTypeApply,
}

var mappingRemediationRunStageTypeEnumLowerCase = map[string]RemediationRunStageTypeEnum{
	"detect":    RemediationRunStageTypeDetect,
	"recommend": RemediationRunStageTypeRecommend,
	"verify":    RemediationRunStageTypeVerify,
	"apply":     RemediationRunStageTypeApply,
}

// GetRemediationRunStageTypeEnumValues Enumerates the set of values for RemediationRunStageTypeEnum
func GetRemediationRunStageTypeEnumValues() []RemediationRunStageTypeEnum {
	values := make([]RemediationRunStageTypeEnum, 0)
	for _, v := range mappingRemediationRunStageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemediationRunStageTypeEnumStringValues Enumerates the set of values in String for RemediationRunStageTypeEnum
func GetRemediationRunStageTypeEnumStringValues() []string {
	return []string{
		"DETECT",
		"RECOMMEND",
		"VERIFY",
		"APPLY",
	}
}

// GetMappingRemediationRunStageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemediationRunStageTypeEnum(val string) (RemediationRunStageTypeEnum, bool) {
	enum, ok := mappingRemediationRunStageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
