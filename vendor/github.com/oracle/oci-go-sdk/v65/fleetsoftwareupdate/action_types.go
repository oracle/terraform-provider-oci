// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// ActionTypesEnum Enum with underlying type: string
type ActionTypesEnum string

// Set of constants representing the allowable values for ActionTypesEnum
const (
	ActionTypesStage                   ActionTypesEnum = "STAGE"
	ActionTypesPrecheck                ActionTypesEnum = "PRECHECK"
	ActionTypesApply                   ActionTypesEnum = "APPLY"
	ActionTypesRollbackAndRemoveTarget ActionTypesEnum = "ROLLBACK_AND_REMOVE_TARGET"
	ActionTypesCleanup                 ActionTypesEnum = "CLEANUP"
)

var mappingActionTypesEnum = map[string]ActionTypesEnum{
	"STAGE":                      ActionTypesStage,
	"PRECHECK":                   ActionTypesPrecheck,
	"APPLY":                      ActionTypesApply,
	"ROLLBACK_AND_REMOVE_TARGET": ActionTypesRollbackAndRemoveTarget,
	"CLEANUP":                    ActionTypesCleanup,
}

var mappingActionTypesEnumLowerCase = map[string]ActionTypesEnum{
	"stage":                      ActionTypesStage,
	"precheck":                   ActionTypesPrecheck,
	"apply":                      ActionTypesApply,
	"rollback_and_remove_target": ActionTypesRollbackAndRemoveTarget,
	"cleanup":                    ActionTypesCleanup,
}

// GetActionTypesEnumValues Enumerates the set of values for ActionTypesEnum
func GetActionTypesEnumValues() []ActionTypesEnum {
	values := make([]ActionTypesEnum, 0)
	for _, v := range mappingActionTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetActionTypesEnumStringValues Enumerates the set of values in String for ActionTypesEnum
func GetActionTypesEnumStringValues() []string {
	return []string{
		"STAGE",
		"PRECHECK",
		"APPLY",
		"ROLLBACK_AND_REMOVE_TARGET",
		"CLEANUP",
	}
}

// GetMappingActionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypesEnum(val string) (ActionTypesEnum, bool) {
	enum, ok := mappingActionTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
