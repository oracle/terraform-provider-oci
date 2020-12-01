// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

// OsFamiliesEnum Enum with underlying type: string
type OsFamiliesEnum string

// Set of constants representing the allowable values for OsFamiliesEnum
const (
	OsFamiliesLinux   OsFamiliesEnum = "LINUX"
	OsFamiliesWindows OsFamiliesEnum = "WINDOWS"
	OsFamiliesAll     OsFamiliesEnum = "ALL"
)

var mappingOsFamilies = map[string]OsFamiliesEnum{
	"LINUX":   OsFamiliesLinux,
	"WINDOWS": OsFamiliesWindows,
	"ALL":     OsFamiliesAll,
}

// GetOsFamiliesEnumValues Enumerates the set of values for OsFamiliesEnum
func GetOsFamiliesEnumValues() []OsFamiliesEnum {
	values := make([]OsFamiliesEnum, 0)
	for _, v := range mappingOsFamilies {
		values = append(values, v)
	}
	return values
}
