// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"strings"
)

// WorkflowActionEnum Enum with underlying type: string
type WorkflowActionEnum string

// Set of constants representing the allowable values for WorkflowActionEnum
const (
	WorkflowActionCreate WorkflowActionEnum = "CREATE"
	WorkflowActionUpdate WorkflowActionEnum = "UPDATE"
)

var mappingWorkflowActionEnum = map[string]WorkflowActionEnum{
	"CREATE": WorkflowActionCreate,
	"UPDATE": WorkflowActionUpdate,
}

var mappingWorkflowActionEnumLowerCase = map[string]WorkflowActionEnum{
	"create": WorkflowActionCreate,
	"update": WorkflowActionUpdate,
}

// GetWorkflowActionEnumValues Enumerates the set of values for WorkflowActionEnum
func GetWorkflowActionEnumValues() []WorkflowActionEnum {
	values := make([]WorkflowActionEnum, 0)
	for _, v := range mappingWorkflowActionEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkflowActionEnumStringValues Enumerates the set of values in String for WorkflowActionEnum
func GetWorkflowActionEnumStringValues() []string {
	return []string{
		"CREATE",
		"UPDATE",
	}
}

// GetMappingWorkflowActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkflowActionEnum(val string) (WorkflowActionEnum, bool) {
	enum, ok := mappingWorkflowActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
