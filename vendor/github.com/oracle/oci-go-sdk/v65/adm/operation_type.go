// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateRemediationRecipe     OperationTypeEnum = "CREATE_REMEDIATION_RECIPE"
	OperationTypeDeleteRemediationRecipe     OperationTypeEnum = "DELETE_REMEDIATION_RECIPE"
	OperationTypeMoveRemediationRecipe       OperationTypeEnum = "MOVE_REMEDIATION_RECIPE"
	OperationTypeUpdateRemediationRecipe     OperationTypeEnum = "UPDATE_REMEDIATION_RECIPE"
	OperationTypeActivateRemediationRecipe   OperationTypeEnum = "ACTIVATE_REMEDIATION_RECIPE"
	OperationTypeDeactivateRemediationRecipe OperationTypeEnum = "DEACTIVATE_REMEDIATION_RECIPE"
	OperationTypeCreateKnowledgeBase         OperationTypeEnum = "CREATE_KNOWLEDGE_BASE"
	OperationTypeDeleteKnowledgeBase         OperationTypeEnum = "DELETE_KNOWLEDGE_BASE"
	OperationTypeMoveKnowledgeBase           OperationTypeEnum = "MOVE_KNOWLEDGE_BASE"
	OperationTypeUpdateKnowledgeBase         OperationTypeEnum = "UPDATE_KNOWLEDGE_BASE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_REMEDIATION_RECIPE":     OperationTypeCreateRemediationRecipe,
	"DELETE_REMEDIATION_RECIPE":     OperationTypeDeleteRemediationRecipe,
	"MOVE_REMEDIATION_RECIPE":       OperationTypeMoveRemediationRecipe,
	"UPDATE_REMEDIATION_RECIPE":     OperationTypeUpdateRemediationRecipe,
	"ACTIVATE_REMEDIATION_RECIPE":   OperationTypeActivateRemediationRecipe,
	"DEACTIVATE_REMEDIATION_RECIPE": OperationTypeDeactivateRemediationRecipe,
	"CREATE_KNOWLEDGE_BASE":         OperationTypeCreateKnowledgeBase,
	"DELETE_KNOWLEDGE_BASE":         OperationTypeDeleteKnowledgeBase,
	"MOVE_KNOWLEDGE_BASE":           OperationTypeMoveKnowledgeBase,
	"UPDATE_KNOWLEDGE_BASE":         OperationTypeUpdateKnowledgeBase,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_remediation_recipe":     OperationTypeCreateRemediationRecipe,
	"delete_remediation_recipe":     OperationTypeDeleteRemediationRecipe,
	"move_remediation_recipe":       OperationTypeMoveRemediationRecipe,
	"update_remediation_recipe":     OperationTypeUpdateRemediationRecipe,
	"activate_remediation_recipe":   OperationTypeActivateRemediationRecipe,
	"deactivate_remediation_recipe": OperationTypeDeactivateRemediationRecipe,
	"create_knowledge_base":         OperationTypeCreateKnowledgeBase,
	"delete_knowledge_base":         OperationTypeDeleteKnowledgeBase,
	"move_knowledge_base":           OperationTypeMoveKnowledgeBase,
	"update_knowledge_base":         OperationTypeUpdateKnowledgeBase,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_REMEDIATION_RECIPE",
		"DELETE_REMEDIATION_RECIPE",
		"MOVE_REMEDIATION_RECIPE",
		"UPDATE_REMEDIATION_RECIPE",
		"ACTIVATE_REMEDIATION_RECIPE",
		"DEACTIVATE_REMEDIATION_RECIPE",
		"CREATE_KNOWLEDGE_BASE",
		"DELETE_KNOWLEDGE_BASE",
		"MOVE_KNOWLEDGE_BASE",
		"UPDATE_KNOWLEDGE_BASE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
