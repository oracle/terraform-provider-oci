// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// UpdateChannelEnum Enum with underlying type: string
type UpdateChannelEnum string

// Set of constants representing the allowable values for UpdateChannelEnum
const (
	UpdateChannelRegular UpdateChannelEnum = "REGULAR"
	UpdateChannelEarly   UpdateChannelEnum = "EARLY"
	UpdateChannelPhase2  UpdateChannelEnum = "PHASE_2"
	UpdateChannelPhase1  UpdateChannelEnum = "PHASE_1"
)

var mappingUpdateChannelEnum = map[string]UpdateChannelEnum{
	"REGULAR": UpdateChannelRegular,
	"EARLY":   UpdateChannelEarly,
	"PHASE_2": UpdateChannelPhase2,
	"PHASE_1": UpdateChannelPhase1,
}

var mappingUpdateChannelEnumLowerCase = map[string]UpdateChannelEnum{
	"regular": UpdateChannelRegular,
	"early":   UpdateChannelEarly,
	"phase_2": UpdateChannelPhase2,
	"phase_1": UpdateChannelPhase1,
}

// GetUpdateChannelEnumValues Enumerates the set of values for UpdateChannelEnum
func GetUpdateChannelEnumValues() []UpdateChannelEnum {
	values := make([]UpdateChannelEnum, 0)
	for _, v := range mappingUpdateChannelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateChannelEnumStringValues Enumerates the set of values in String for UpdateChannelEnum
func GetUpdateChannelEnumStringValues() []string {
	return []string{
		"REGULAR",
		"EARLY",
		"PHASE_2",
		"PHASE_1",
	}
}

// GetMappingUpdateChannelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateChannelEnum(val string) (UpdateChannelEnum, bool) {
	enum, ok := mappingUpdateChannelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
