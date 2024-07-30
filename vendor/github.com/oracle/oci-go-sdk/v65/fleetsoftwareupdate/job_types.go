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

// JobTypesEnum Enum with underlying type: string
type JobTypesEnum string

// Set of constants representing the allowable values for JobTypesEnum
const (
	JobTypesStage                   JobTypesEnum = "STAGE"
	JobTypesPrecheck                JobTypesEnum = "PRECHECK"
	JobTypesApply                   JobTypesEnum = "APPLY"
	JobTypesRollbackAndRemoveTarget JobTypesEnum = "ROLLBACK_AND_REMOVE_TARGET"
	JobTypesCleanup                 JobTypesEnum = "CLEANUP"
)

var mappingJobTypesEnum = map[string]JobTypesEnum{
	"STAGE":                      JobTypesStage,
	"PRECHECK":                   JobTypesPrecheck,
	"APPLY":                      JobTypesApply,
	"ROLLBACK_AND_REMOVE_TARGET": JobTypesRollbackAndRemoveTarget,
	"CLEANUP":                    JobTypesCleanup,
}

var mappingJobTypesEnumLowerCase = map[string]JobTypesEnum{
	"stage":                      JobTypesStage,
	"precheck":                   JobTypesPrecheck,
	"apply":                      JobTypesApply,
	"rollback_and_remove_target": JobTypesRollbackAndRemoveTarget,
	"cleanup":                    JobTypesCleanup,
}

// GetJobTypesEnumValues Enumerates the set of values for JobTypesEnum
func GetJobTypesEnumValues() []JobTypesEnum {
	values := make([]JobTypesEnum, 0)
	for _, v := range mappingJobTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetJobTypesEnumStringValues Enumerates the set of values in String for JobTypesEnum
func GetJobTypesEnumStringValues() []string {
	return []string{
		"STAGE",
		"PRECHECK",
		"APPLY",
		"ROLLBACK_AND_REMOVE_TARGET",
		"CLEANUP",
	}
}

// GetMappingJobTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobTypesEnum(val string) (JobTypesEnum, bool) {
	enum, ok := mappingJobTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
