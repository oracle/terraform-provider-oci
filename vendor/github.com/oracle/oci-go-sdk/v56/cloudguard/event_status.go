// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
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

var mappingEventStatus = map[string]EventStatusEnum{
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
	for _, v := range mappingEventStatus {
		values = append(values, v)
	}
	return values
}
