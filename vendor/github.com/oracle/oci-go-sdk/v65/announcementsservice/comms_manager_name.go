// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"strings"
)

// CommsManagerNameEnum Enum with underlying type: string
type CommsManagerNameEnum string

// Set of constants representing the allowable values for CommsManagerNameEnum
const (
	CommsManagerNameCn     CommsManagerNameEnum = "CN"
	CommsManagerNameFusion CommsManagerNameEnum = "FUSION"
	CommsManagerNameAs     CommsManagerNameEnum = "AS"
	CommsManagerNameErf    CommsManagerNameEnum = "ERF"
)

var mappingCommsManagerNameEnum = map[string]CommsManagerNameEnum{
	"CN":     CommsManagerNameCn,
	"FUSION": CommsManagerNameFusion,
	"AS":     CommsManagerNameAs,
	"ERF":    CommsManagerNameErf,
}

var mappingCommsManagerNameEnumLowerCase = map[string]CommsManagerNameEnum{
	"cn":     CommsManagerNameCn,
	"fusion": CommsManagerNameFusion,
	"as":     CommsManagerNameAs,
	"erf":    CommsManagerNameErf,
}

// GetCommsManagerNameEnumValues Enumerates the set of values for CommsManagerNameEnum
func GetCommsManagerNameEnumValues() []CommsManagerNameEnum {
	values := make([]CommsManagerNameEnum, 0)
	for _, v := range mappingCommsManagerNameEnum {
		values = append(values, v)
	}
	return values
}

// GetCommsManagerNameEnumStringValues Enumerates the set of values in String for CommsManagerNameEnum
func GetCommsManagerNameEnumStringValues() []string {
	return []string{
		"CN",
		"FUSION",
		"AS",
		"ERF",
	}
}

// GetMappingCommsManagerNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCommsManagerNameEnum(val string) (CommsManagerNameEnum, bool) {
	enum, ok := mappingCommsManagerNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
