// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// UpdateTypesEnum Enum with underlying type: string
type UpdateTypesEnum string

// Set of constants representing the allowable values for UpdateTypesEnum
const (
	UpdateTypesSecurity         UpdateTypesEnum = "SECURITY"
	UpdateTypesBugfix           UpdateTypesEnum = "BUGFIX"
	UpdateTypesEnhancement      UpdateTypesEnum = "ENHANCEMENT"
	UpdateTypesOther            UpdateTypesEnum = "OTHER"
	UpdateTypesKspliceKernel    UpdateTypesEnum = "KSPLICE_KERNEL"
	UpdateTypesKspliceUserspace UpdateTypesEnum = "KSPLICE_USERSPACE"
	UpdateTypesAll              UpdateTypesEnum = "ALL"
)

var mappingUpdateTypesEnum = map[string]UpdateTypesEnum{
	"SECURITY":          UpdateTypesSecurity,
	"BUGFIX":            UpdateTypesBugfix,
	"ENHANCEMENT":       UpdateTypesEnhancement,
	"OTHER":             UpdateTypesOther,
	"KSPLICE_KERNEL":    UpdateTypesKspliceKernel,
	"KSPLICE_USERSPACE": UpdateTypesKspliceUserspace,
	"ALL":               UpdateTypesAll,
}

var mappingUpdateTypesEnumLowerCase = map[string]UpdateTypesEnum{
	"security":          UpdateTypesSecurity,
	"bugfix":            UpdateTypesBugfix,
	"enhancement":       UpdateTypesEnhancement,
	"other":             UpdateTypesOther,
	"ksplice_kernel":    UpdateTypesKspliceKernel,
	"ksplice_userspace": UpdateTypesKspliceUserspace,
	"all":               UpdateTypesAll,
}

// GetUpdateTypesEnumValues Enumerates the set of values for UpdateTypesEnum
func GetUpdateTypesEnumValues() []UpdateTypesEnum {
	values := make([]UpdateTypesEnum, 0)
	for _, v := range mappingUpdateTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTypesEnumStringValues Enumerates the set of values in String for UpdateTypesEnum
func GetUpdateTypesEnumStringValues() []string {
	return []string{
		"SECURITY",
		"BUGFIX",
		"ENHANCEMENT",
		"OTHER",
		"KSPLICE_KERNEL",
		"KSPLICE_USERSPACE",
		"ALL",
	}
}

// GetMappingUpdateTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTypesEnum(val string) (UpdateTypesEnum, bool) {
	enum, ok := mappingUpdateTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
