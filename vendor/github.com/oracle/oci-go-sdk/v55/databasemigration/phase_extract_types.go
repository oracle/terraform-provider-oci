// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// PhaseExtractTypesEnum Enum with underlying type: string
type PhaseExtractTypesEnum string

// Set of constants representing the allowable values for PhaseExtractTypesEnum
const (
	PhaseExtractTypesError PhaseExtractTypesEnum = "ERROR"
)

var mappingPhaseExtractTypes = map[string]PhaseExtractTypesEnum{
	"ERROR": PhaseExtractTypesError,
}

// GetPhaseExtractTypesEnumValues Enumerates the set of values for PhaseExtractTypesEnum
func GetPhaseExtractTypesEnumValues() []PhaseExtractTypesEnum {
	values := make([]PhaseExtractTypesEnum, 0)
	for _, v := range mappingPhaseExtractTypes {
		values = append(values, v)
	}
	return values
}
