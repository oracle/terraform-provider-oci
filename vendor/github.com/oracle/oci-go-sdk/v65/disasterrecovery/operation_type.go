// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDrProtectionGroup       OperationTypeEnum = "CREATE_DR_PROTECTION_GROUP"
	OperationTypeUpdateDrProtectionGroup       OperationTypeEnum = "UPDATE_DR_PROTECTION_GROUP"
	OperationTypeDeleteDrProtectionGroup       OperationTypeEnum = "DELETE_DR_PROTECTION_GROUP"
	OperationTypeMoveDrProtectionGroup         OperationTypeEnum = "MOVE_DR_PROTECTION_GROUP"
	OperationTypeAssociateDrProtectionGroup    OperationTypeEnum = "ASSOCIATE_DR_PROTECTION_GROUP"
	OperationTypeDisassociateDrProtectionGroup OperationTypeEnum = "DISASSOCIATE_DR_PROTECTION_GROUP"
	OperationTypeUpdateRoleDrProtectionGroup   OperationTypeEnum = "UPDATE_ROLE_DR_PROTECTION_GROUP"
	OperationTypeCreateDrPlan                  OperationTypeEnum = "CREATE_DR_PLAN"
	OperationTypeUpdateDrPlan                  OperationTypeEnum = "UPDATE_DR_PLAN"
	OperationTypeDeleteDrPlan                  OperationTypeEnum = "DELETE_DR_PLAN"
	OperationTypeCreateDrPlanExecution         OperationTypeEnum = "CREATE_DR_PLAN_EXECUTION"
	OperationTypeUpdateDrPlanExecution         OperationTypeEnum = "UPDATE_DR_PLAN_EXECUTION"
	OperationTypeDeleteDrPlanExecution         OperationTypeEnum = "DELETE_DR_PLAN_EXECUTION"
	OperationTypeRetryDrPlanExecution          OperationTypeEnum = "RETRY_DR_PLAN_EXECUTION"
	OperationTypeIgnoreDrPlanExecution         OperationTypeEnum = "IGNORE_DR_PLAN_EXECUTION"
	OperationTypeCancelDrPlanExecution         OperationTypeEnum = "CANCEL_DR_PLAN_EXECUTION"
	OperationTypePauseDrPlanExecution          OperationTypeEnum = "PAUSE_DR_PLAN_EXECUTION"
	OperationTypeResumeDrPlanExecution         OperationTypeEnum = "RESUME_DR_PLAN_EXECUTION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DR_PROTECTION_GROUP":       OperationTypeCreateDrProtectionGroup,
	"UPDATE_DR_PROTECTION_GROUP":       OperationTypeUpdateDrProtectionGroup,
	"DELETE_DR_PROTECTION_GROUP":       OperationTypeDeleteDrProtectionGroup,
	"MOVE_DR_PROTECTION_GROUP":         OperationTypeMoveDrProtectionGroup,
	"ASSOCIATE_DR_PROTECTION_GROUP":    OperationTypeAssociateDrProtectionGroup,
	"DISASSOCIATE_DR_PROTECTION_GROUP": OperationTypeDisassociateDrProtectionGroup,
	"UPDATE_ROLE_DR_PROTECTION_GROUP":  OperationTypeUpdateRoleDrProtectionGroup,
	"CREATE_DR_PLAN":                   OperationTypeCreateDrPlan,
	"UPDATE_DR_PLAN":                   OperationTypeUpdateDrPlan,
	"DELETE_DR_PLAN":                   OperationTypeDeleteDrPlan,
	"CREATE_DR_PLAN_EXECUTION":         OperationTypeCreateDrPlanExecution,
	"UPDATE_DR_PLAN_EXECUTION":         OperationTypeUpdateDrPlanExecution,
	"DELETE_DR_PLAN_EXECUTION":         OperationTypeDeleteDrPlanExecution,
	"RETRY_DR_PLAN_EXECUTION":          OperationTypeRetryDrPlanExecution,
	"IGNORE_DR_PLAN_EXECUTION":         OperationTypeIgnoreDrPlanExecution,
	"CANCEL_DR_PLAN_EXECUTION":         OperationTypeCancelDrPlanExecution,
	"PAUSE_DR_PLAN_EXECUTION":          OperationTypePauseDrPlanExecution,
	"RESUME_DR_PLAN_EXECUTION":         OperationTypeResumeDrPlanExecution,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_dr_protection_group":       OperationTypeCreateDrProtectionGroup,
	"update_dr_protection_group":       OperationTypeUpdateDrProtectionGroup,
	"delete_dr_protection_group":       OperationTypeDeleteDrProtectionGroup,
	"move_dr_protection_group":         OperationTypeMoveDrProtectionGroup,
	"associate_dr_protection_group":    OperationTypeAssociateDrProtectionGroup,
	"disassociate_dr_protection_group": OperationTypeDisassociateDrProtectionGroup,
	"update_role_dr_protection_group":  OperationTypeUpdateRoleDrProtectionGroup,
	"create_dr_plan":                   OperationTypeCreateDrPlan,
	"update_dr_plan":                   OperationTypeUpdateDrPlan,
	"delete_dr_plan":                   OperationTypeDeleteDrPlan,
	"create_dr_plan_execution":         OperationTypeCreateDrPlanExecution,
	"update_dr_plan_execution":         OperationTypeUpdateDrPlanExecution,
	"delete_dr_plan_execution":         OperationTypeDeleteDrPlanExecution,
	"retry_dr_plan_execution":          OperationTypeRetryDrPlanExecution,
	"ignore_dr_plan_execution":         OperationTypeIgnoreDrPlanExecution,
	"cancel_dr_plan_execution":         OperationTypeCancelDrPlanExecution,
	"pause_dr_plan_execution":          OperationTypePauseDrPlanExecution,
	"resume_dr_plan_execution":         OperationTypeResumeDrPlanExecution,
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
		"CREATE_DR_PROTECTION_GROUP",
		"UPDATE_DR_PROTECTION_GROUP",
		"DELETE_DR_PROTECTION_GROUP",
		"MOVE_DR_PROTECTION_GROUP",
		"ASSOCIATE_DR_PROTECTION_GROUP",
		"DISASSOCIATE_DR_PROTECTION_GROUP",
		"UPDATE_ROLE_DR_PROTECTION_GROUP",
		"CREATE_DR_PLAN",
		"UPDATE_DR_PLAN",
		"DELETE_DR_PLAN",
		"CREATE_DR_PLAN_EXECUTION",
		"UPDATE_DR_PLAN_EXECUTION",
		"DELETE_DR_PLAN_EXECUTION",
		"RETRY_DR_PLAN_EXECUTION",
		"IGNORE_DR_PLAN_EXECUTION",
		"CANCEL_DR_PLAN_EXECUTION",
		"PAUSE_DR_PLAN_EXECUTION",
		"RESUME_DR_PLAN_EXECUTION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
