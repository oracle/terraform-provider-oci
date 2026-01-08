// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// NotebookSessionEnvironmentTypeEnum Enum with underlying type: string
type NotebookSessionEnvironmentTypeEnum string

// Set of constants representing the allowable values for NotebookSessionEnvironmentTypeEnum
const (
	NotebookSessionEnvironmentTypeOcirContainer NotebookSessionEnvironmentTypeEnum = "OCIR_CONTAINER"
)

var mappingNotebookSessionEnvironmentTypeEnum = map[string]NotebookSessionEnvironmentTypeEnum{
	"OCIR_CONTAINER": NotebookSessionEnvironmentTypeOcirContainer,
}

var mappingNotebookSessionEnvironmentTypeEnumLowerCase = map[string]NotebookSessionEnvironmentTypeEnum{
	"ocir_container": NotebookSessionEnvironmentTypeOcirContainer,
}

// GetNotebookSessionEnvironmentTypeEnumValues Enumerates the set of values for NotebookSessionEnvironmentTypeEnum
func GetNotebookSessionEnvironmentTypeEnumValues() []NotebookSessionEnvironmentTypeEnum {
	values := make([]NotebookSessionEnvironmentTypeEnum, 0)
	for _, v := range mappingNotebookSessionEnvironmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNotebookSessionEnvironmentTypeEnumStringValues Enumerates the set of values in String for NotebookSessionEnvironmentTypeEnum
func GetNotebookSessionEnvironmentTypeEnumStringValues() []string {
	return []string{
		"OCIR_CONTAINER",
	}
}

// GetMappingNotebookSessionEnvironmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotebookSessionEnvironmentTypeEnum(val string) (NotebookSessionEnvironmentTypeEnum, bool) {
	enum, ok := mappingNotebookSessionEnvironmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
