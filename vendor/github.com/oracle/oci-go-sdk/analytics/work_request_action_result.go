// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

// WorkRequestActionResultEnum Enum with underlying type: string
type WorkRequestActionResultEnum string

// Set of constants representing the allowable values for WorkRequestActionResultEnum
const (
	WorkRequestActionResultCompartmentChanged WorkRequestActionResultEnum = "COMPARTMENT_CHANGED"
	WorkRequestActionResultCreated            WorkRequestActionResultEnum = "CREATED"
	WorkRequestActionResultDeleted            WorkRequestActionResultEnum = "DELETED"
	WorkRequestActionResultStarted            WorkRequestActionResultEnum = "STARTED"
	WorkRequestActionResultStopped            WorkRequestActionResultEnum = "STOPPED"
	WorkRequestActionResultScaled             WorkRequestActionResultEnum = "SCALED"
	WorkRequestActionResultNone               WorkRequestActionResultEnum = "NONE"
)

var mappingWorkRequestActionResult = map[string]WorkRequestActionResultEnum{
	"COMPARTMENT_CHANGED": WorkRequestActionResultCompartmentChanged,
	"CREATED":             WorkRequestActionResultCreated,
	"DELETED":             WorkRequestActionResultDeleted,
	"STARTED":             WorkRequestActionResultStarted,
	"STOPPED":             WorkRequestActionResultStopped,
	"SCALED":              WorkRequestActionResultScaled,
	"NONE":                WorkRequestActionResultNone,
}

// GetWorkRequestActionResultEnumValues Enumerates the set of values for WorkRequestActionResultEnum
func GetWorkRequestActionResultEnumValues() []WorkRequestActionResultEnum {
	values := make([]WorkRequestActionResultEnum, 0)
	for _, v := range mappingWorkRequestActionResult {
		values = append(values, v)
	}
	return values
}
