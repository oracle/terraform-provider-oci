// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ResponderActivityTypeEnum Enum with underlying type: string
type ResponderActivityTypeEnum string

// Set of constants representing the allowable values for ResponderActivityTypeEnum
const (
	ResponderActivityTypeStarted   ResponderActivityTypeEnum = "STARTED"
	ResponderActivityTypeCompleted ResponderActivityTypeEnum = "COMPLETED"
)

var mappingResponderActivityType = map[string]ResponderActivityTypeEnum{
	"STARTED":   ResponderActivityTypeStarted,
	"COMPLETED": ResponderActivityTypeCompleted,
}

// GetResponderActivityTypeEnumValues Enumerates the set of values for ResponderActivityTypeEnum
func GetResponderActivityTypeEnumValues() []ResponderActivityTypeEnum {
	values := make([]ResponderActivityTypeEnum, 0)
	for _, v := range mappingResponderActivityType {
		values = append(values, v)
	}
	return values
}
