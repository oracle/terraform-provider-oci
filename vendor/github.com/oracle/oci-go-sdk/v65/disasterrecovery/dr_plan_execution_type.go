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

// DrPlanExecutionTypeEnum Enum with underlying type: string
type DrPlanExecutionTypeEnum string

// Set of constants representing the allowable values for DrPlanExecutionTypeEnum
const (
	DrPlanExecutionTypeSwitchover         DrPlanExecutionTypeEnum = "SWITCHOVER"
	DrPlanExecutionTypeSwitchoverPrecheck DrPlanExecutionTypeEnum = "SWITCHOVER_PRECHECK"
	DrPlanExecutionTypeFailover           DrPlanExecutionTypeEnum = "FAILOVER"
	DrPlanExecutionTypeFailoverPrecheck   DrPlanExecutionTypeEnum = "FAILOVER_PRECHECK"
	DrPlanExecutionTypeStartDrill         DrPlanExecutionTypeEnum = "START_DRILL"
	DrPlanExecutionTypeStartDrillPrecheck DrPlanExecutionTypeEnum = "START_DRILL_PRECHECK"
	DrPlanExecutionTypeStopDrill          DrPlanExecutionTypeEnum = "STOP_DRILL"
	DrPlanExecutionTypeStopDrillPrecheck  DrPlanExecutionTypeEnum = "STOP_DRILL_PRECHECK"
)

var mappingDrPlanExecutionTypeEnum = map[string]DrPlanExecutionTypeEnum{
	"SWITCHOVER":           DrPlanExecutionTypeSwitchover,
	"SWITCHOVER_PRECHECK":  DrPlanExecutionTypeSwitchoverPrecheck,
	"FAILOVER":             DrPlanExecutionTypeFailover,
	"FAILOVER_PRECHECK":    DrPlanExecutionTypeFailoverPrecheck,
	"START_DRILL":          DrPlanExecutionTypeStartDrill,
	"START_DRILL_PRECHECK": DrPlanExecutionTypeStartDrillPrecheck,
	"STOP_DRILL":           DrPlanExecutionTypeStopDrill,
	"STOP_DRILL_PRECHECK":  DrPlanExecutionTypeStopDrillPrecheck,
}

var mappingDrPlanExecutionTypeEnumLowerCase = map[string]DrPlanExecutionTypeEnum{
	"switchover":           DrPlanExecutionTypeSwitchover,
	"switchover_precheck":  DrPlanExecutionTypeSwitchoverPrecheck,
	"failover":             DrPlanExecutionTypeFailover,
	"failover_precheck":    DrPlanExecutionTypeFailoverPrecheck,
	"start_drill":          DrPlanExecutionTypeStartDrill,
	"start_drill_precheck": DrPlanExecutionTypeStartDrillPrecheck,
	"stop_drill":           DrPlanExecutionTypeStopDrill,
	"stop_drill_precheck":  DrPlanExecutionTypeStopDrillPrecheck,
}

// GetDrPlanExecutionTypeEnumValues Enumerates the set of values for DrPlanExecutionTypeEnum
func GetDrPlanExecutionTypeEnumValues() []DrPlanExecutionTypeEnum {
	values := make([]DrPlanExecutionTypeEnum, 0)
	for _, v := range mappingDrPlanExecutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanExecutionTypeEnumStringValues Enumerates the set of values in String for DrPlanExecutionTypeEnum
func GetDrPlanExecutionTypeEnumStringValues() []string {
	return []string{
		"SWITCHOVER",
		"SWITCHOVER_PRECHECK",
		"FAILOVER",
		"FAILOVER_PRECHECK",
		"START_DRILL",
		"START_DRILL_PRECHECK",
		"STOP_DRILL",
		"STOP_DRILL_PRECHECK",
	}
}

// GetMappingDrPlanExecutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanExecutionTypeEnum(val string) (DrPlanExecutionTypeEnum, bool) {
	enum, ok := mappingDrPlanExecutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
