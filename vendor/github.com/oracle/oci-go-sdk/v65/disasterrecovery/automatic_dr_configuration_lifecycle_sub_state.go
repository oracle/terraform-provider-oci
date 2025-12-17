// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AutomaticDrConfigurationLifecycleSubStateEnum Enum with underlying type: string
type AutomaticDrConfigurationLifecycleSubStateEnum string

// Set of constants representing the allowable values for AutomaticDrConfigurationLifecycleSubStateEnum
const (
	AutomaticDrConfigurationLifecycleSubStateResourcePrincipalValidationFailed AutomaticDrConfigurationLifecycleSubStateEnum = "RESOURCE_PRINCIPAL_VALIDATION_FAILED"
	AutomaticDrConfigurationLifecycleSubStateDrPlanNeedsAttention              AutomaticDrConfigurationLifecycleSubStateEnum = "DR_PLAN_NEEDS_ATTENTION"
	AutomaticDrConfigurationLifecycleSubStateDrPlanInactive                    AutomaticDrConfigurationLifecycleSubStateEnum = "DR_PLAN_INACTIVE"
)

var mappingAutomaticDrConfigurationLifecycleSubStateEnum = map[string]AutomaticDrConfigurationLifecycleSubStateEnum{
	"RESOURCE_PRINCIPAL_VALIDATION_FAILED": AutomaticDrConfigurationLifecycleSubStateResourcePrincipalValidationFailed,
	"DR_PLAN_NEEDS_ATTENTION":              AutomaticDrConfigurationLifecycleSubStateDrPlanNeedsAttention,
	"DR_PLAN_INACTIVE":                     AutomaticDrConfigurationLifecycleSubStateDrPlanInactive,
}

var mappingAutomaticDrConfigurationLifecycleSubStateEnumLowerCase = map[string]AutomaticDrConfigurationLifecycleSubStateEnum{
	"resource_principal_validation_failed": AutomaticDrConfigurationLifecycleSubStateResourcePrincipalValidationFailed,
	"dr_plan_needs_attention":              AutomaticDrConfigurationLifecycleSubStateDrPlanNeedsAttention,
	"dr_plan_inactive":                     AutomaticDrConfigurationLifecycleSubStateDrPlanInactive,
}

// GetAutomaticDrConfigurationLifecycleSubStateEnumValues Enumerates the set of values for AutomaticDrConfigurationLifecycleSubStateEnum
func GetAutomaticDrConfigurationLifecycleSubStateEnumValues() []AutomaticDrConfigurationLifecycleSubStateEnum {
	values := make([]AutomaticDrConfigurationLifecycleSubStateEnum, 0)
	for _, v := range mappingAutomaticDrConfigurationLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutomaticDrConfigurationLifecycleSubStateEnumStringValues Enumerates the set of values in String for AutomaticDrConfigurationLifecycleSubStateEnum
func GetAutomaticDrConfigurationLifecycleSubStateEnumStringValues() []string {
	return []string{
		"RESOURCE_PRINCIPAL_VALIDATION_FAILED",
		"DR_PLAN_NEEDS_ATTENTION",
		"DR_PLAN_INACTIVE",
	}
}

// GetMappingAutomaticDrConfigurationLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutomaticDrConfigurationLifecycleSubStateEnum(val string) (AutomaticDrConfigurationLifecycleSubStateEnum, bool) {
	enum, ok := mappingAutomaticDrConfigurationLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
