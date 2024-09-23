// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"strings"
)

// SelectionEnum Enum with underlying type: string
type SelectionEnum string

// Set of constants representing the allowable values for SelectionEnum
const (
	SelectionSingleChoice SelectionEnum = "SINGLE_CHOICE"
	SelectionMultiChoice  SelectionEnum = "MULTI_CHOICE"
	SelectionDefaultText  SelectionEnum = "DEFAULT_TEXT"
)

var mappingSelectionEnum = map[string]SelectionEnum{
	"SINGLE_CHOICE": SelectionSingleChoice,
	"MULTI_CHOICE":  SelectionMultiChoice,
	"DEFAULT_TEXT":  SelectionDefaultText,
}

var mappingSelectionEnumLowerCase = map[string]SelectionEnum{
	"single_choice": SelectionSingleChoice,
	"multi_choice":  SelectionMultiChoice,
	"default_text":  SelectionDefaultText,
}

// GetSelectionEnumValues Enumerates the set of values for SelectionEnum
func GetSelectionEnumValues() []SelectionEnum {
	values := make([]SelectionEnum, 0)
	for _, v := range mappingSelectionEnum {
		values = append(values, v)
	}
	return values
}

// GetSelectionEnumStringValues Enumerates the set of values in String for SelectionEnum
func GetSelectionEnumStringValues() []string {
	return []string{
		"SINGLE_CHOICE",
		"MULTI_CHOICE",
		"DEFAULT_TEXT",
	}
}

// GetMappingSelectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSelectionEnum(val string) (SelectionEnum, bool) {
	enum, ok := mappingSelectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
