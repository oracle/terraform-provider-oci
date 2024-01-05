// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// OdmsPhaseActionsEnum Enum with underlying type: string
type OdmsPhaseActionsEnum string

// Set of constants representing the allowable values for OdmsPhaseActionsEnum
const (
	OdmsPhaseActionsWait OdmsPhaseActionsEnum = "WAIT"
)

var mappingOdmsPhaseActionsEnum = map[string]OdmsPhaseActionsEnum{
	"WAIT": OdmsPhaseActionsWait,
}

var mappingOdmsPhaseActionsEnumLowerCase = map[string]OdmsPhaseActionsEnum{
	"wait": OdmsPhaseActionsWait,
}

// GetOdmsPhaseActionsEnumValues Enumerates the set of values for OdmsPhaseActionsEnum
func GetOdmsPhaseActionsEnumValues() []OdmsPhaseActionsEnum {
	values := make([]OdmsPhaseActionsEnum, 0)
	for _, v := range mappingOdmsPhaseActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetOdmsPhaseActionsEnumStringValues Enumerates the set of values in String for OdmsPhaseActionsEnum
func GetOdmsPhaseActionsEnumStringValues() []string {
	return []string{
		"WAIT",
	}
}

// GetMappingOdmsPhaseActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdmsPhaseActionsEnum(val string) (OdmsPhaseActionsEnum, bool) {
	enum, ok := mappingOdmsPhaseActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
