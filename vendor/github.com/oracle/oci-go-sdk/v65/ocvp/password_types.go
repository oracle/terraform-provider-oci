// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// PasswordTypesEnum Enum with underlying type: string
type PasswordTypesEnum string

// Set of constants representing the allowable values for PasswordTypesEnum
const (
	PasswordTypesVcenter PasswordTypesEnum = "VCENTER"
	PasswordTypesNsx     PasswordTypesEnum = "NSX"
	PasswordTypesHcx     PasswordTypesEnum = "HCX"
)

var mappingPasswordTypesEnum = map[string]PasswordTypesEnum{
	"VCENTER": PasswordTypesVcenter,
	"NSX":     PasswordTypesNsx,
	"HCX":     PasswordTypesHcx,
}

var mappingPasswordTypesEnumLowerCase = map[string]PasswordTypesEnum{
	"vcenter": PasswordTypesVcenter,
	"nsx":     PasswordTypesNsx,
	"hcx":     PasswordTypesHcx,
}

// GetPasswordTypesEnumValues Enumerates the set of values for PasswordTypesEnum
func GetPasswordTypesEnumValues() []PasswordTypesEnum {
	values := make([]PasswordTypesEnum, 0)
	for _, v := range mappingPasswordTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetPasswordTypesEnumStringValues Enumerates the set of values in String for PasswordTypesEnum
func GetPasswordTypesEnumStringValues() []string {
	return []string{
		"VCENTER",
		"NSX",
		"HCX",
	}
}

// GetMappingPasswordTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPasswordTypesEnum(val string) (PasswordTypesEnum, bool) {
	enum, ok := mappingPasswordTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
