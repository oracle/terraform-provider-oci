// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

// EventStatusEnum Enum with underlying type: string
type EventStatusEnum string

// Set of constants representing the allowable values for EventStatusEnum
const (
	EventStatusReopen  EventStatusEnum = "REOPEN"
	EventStatusOpen    EventStatusEnum = "OPEN"
	EventStatusUpdate  EventStatusEnum = "UPDATE"
	EventStatusResolve EventStatusEnum = "RESOLVE"
	EventStatusDismiss EventStatusEnum = "DISMISS"
	EventStatusDelete  EventStatusEnum = "DELETE"
)

var mappingEventStatusEnum = map[string]EventStatusEnum{
	"REOPEN":  EventStatusReopen,
	"OPEN":    EventStatusOpen,
	"UPDATE":  EventStatusUpdate,
	"RESOLVE": EventStatusResolve,
	"DISMISS": EventStatusDismiss,
	"DELETE":  EventStatusDelete,
}

// GetEventStatusEnumValues Enumerates the set of values for EventStatusEnum
func GetEventStatusEnumValues() []EventStatusEnum {
	values := make([]EventStatusEnum, 0)
	for _, v := range mappingEventStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetEventStatusEnumStringValues Enumerates the set of values in String for EventStatusEnum
func GetEventStatusEnumStringValues() []string {
	return []string{
		"REOPEN",
		"OPEN",
		"UPDATE",
		"RESOLVE",
		"DISMISS",
		"DELETE",
	}
}
