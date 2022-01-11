// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ResponderExecutionModesEnum Enum with underlying type: string
type ResponderExecutionModesEnum string

// Set of constants representing the allowable values for ResponderExecutionModesEnum
const (
	ResponderExecutionModesManual    ResponderExecutionModesEnum = "MANUAL"
	ResponderExecutionModesAutomated ResponderExecutionModesEnum = "AUTOMATED"
	ResponderExecutionModesAll       ResponderExecutionModesEnum = "ALL"
)

var mappingResponderExecutionModes = map[string]ResponderExecutionModesEnum{
	"MANUAL":    ResponderExecutionModesManual,
	"AUTOMATED": ResponderExecutionModesAutomated,
	"ALL":       ResponderExecutionModesAll,
}

// GetResponderExecutionModesEnumValues Enumerates the set of values for ResponderExecutionModesEnum
func GetResponderExecutionModesEnumValues() []ResponderExecutionModesEnum {
	values := make([]ResponderExecutionModesEnum, 0)
	for _, v := range mappingResponderExecutionModes {
		values = append(values, v)
	}
	return values
}
