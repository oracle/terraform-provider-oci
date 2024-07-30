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

// DetailedActionTypesEnum Enum with underlying type: string
type DetailedActionTypesEnum string

// Set of constants representing the allowable values for DetailedActionTypesEnum
const (
	DetailedActionTypesStage                   DetailedActionTypesEnum = "STAGE"
	DetailedActionTypesPrecheckStage           DetailedActionTypesEnum = "PRECHECK_STAGE"
	DetailedActionTypesPrecheckApply           DetailedActionTypesEnum = "PRECHECK_APPLY"
	DetailedActionTypesApply                   DetailedActionTypesEnum = "APPLY"
	DetailedActionTypesRollbackAndRemoveTarget DetailedActionTypesEnum = "ROLLBACK_AND_REMOVE_TARGET"
	DetailedActionTypesCleanup                 DetailedActionTypesEnum = "CLEANUP"
)

var mappingDetailedActionTypesEnum = map[string]DetailedActionTypesEnum{
	"STAGE":                      DetailedActionTypesStage,
	"PRECHECK_STAGE":             DetailedActionTypesPrecheckStage,
	"PRECHECK_APPLY":             DetailedActionTypesPrecheckApply,
	"APPLY":                      DetailedActionTypesApply,
	"ROLLBACK_AND_REMOVE_TARGET": DetailedActionTypesRollbackAndRemoveTarget,
	"CLEANUP":                    DetailedActionTypesCleanup,
}

var mappingDetailedActionTypesEnumLowerCase = map[string]DetailedActionTypesEnum{
	"stage":                      DetailedActionTypesStage,
	"precheck_stage":             DetailedActionTypesPrecheckStage,
	"precheck_apply":             DetailedActionTypesPrecheckApply,
	"apply":                      DetailedActionTypesApply,
	"rollback_and_remove_target": DetailedActionTypesRollbackAndRemoveTarget,
	"cleanup":                    DetailedActionTypesCleanup,
}

// GetDetailedActionTypesEnumValues Enumerates the set of values for DetailedActionTypesEnum
func GetDetailedActionTypesEnumValues() []DetailedActionTypesEnum {
	values := make([]DetailedActionTypesEnum, 0)
	for _, v := range mappingDetailedActionTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDetailedActionTypesEnumStringValues Enumerates the set of values in String for DetailedActionTypesEnum
func GetDetailedActionTypesEnumStringValues() []string {
	return []string{
		"STAGE",
		"PRECHECK_STAGE",
		"PRECHECK_APPLY",
		"APPLY",
		"ROLLBACK_AND_REMOVE_TARGET",
		"CLEANUP",
	}
}

// GetMappingDetailedActionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetailedActionTypesEnum(val string) (DetailedActionTypesEnum, bool) {
	enum, ok := mappingDetailedActionTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
