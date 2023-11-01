// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanExecutionOptions The options for a plan execution.
type DrPlanExecutionOptions interface {
}

type drplanexecutionoptions struct {
	JsonData          []byte
	PlanExecutionType string `json:"planExecutionType"`
}

// UnmarshalJSON unmarshals json
func (m *drplanexecutionoptions) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdrplanexecutionoptions drplanexecutionoptions
	s := struct {
		Model Unmarshalerdrplanexecutionoptions
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PlanExecutionType = s.Model.PlanExecutionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *drplanexecutionoptions) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PlanExecutionType {
	case "STOP_DRILL":
		mm := StopDrillExecutionOptions{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SWITCHOVER":
		mm := SwitchoverExecutionOptions{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAILOVER_PRECHECK":
		mm := FailoverPrecheckExecutionOptions{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "START_DRILL":
		mm := StartDrillExecutionOptions{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SWITCHOVER_PRECHECK":
		mm := SwitchoverPrecheckExecutionOptions{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "START_DRILL_PRECHECK":
		mm := StartDrillPrecheckExecutionOptions{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STOP_DRILL_PRECHECK":
		mm := StopDrillPrecheckExecutionOptions{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAILOVER":
		mm := FailoverExecutionOptions{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DrPlanExecutionOptions: %s.", m.PlanExecutionType)
		return *m, nil
	}
}

func (m drplanexecutionoptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m drplanexecutionoptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrPlanExecutionOptionsPlanExecutionTypeEnum Enum with underlying type: string
type DrPlanExecutionOptionsPlanExecutionTypeEnum string

// Set of constants representing the allowable values for DrPlanExecutionOptionsPlanExecutionTypeEnum
const (
	DrPlanExecutionOptionsPlanExecutionTypeSwitchover         DrPlanExecutionOptionsPlanExecutionTypeEnum = "SWITCHOVER"
	DrPlanExecutionOptionsPlanExecutionTypeSwitchoverPrecheck DrPlanExecutionOptionsPlanExecutionTypeEnum = "SWITCHOVER_PRECHECK"
	DrPlanExecutionOptionsPlanExecutionTypeFailover           DrPlanExecutionOptionsPlanExecutionTypeEnum = "FAILOVER"
	DrPlanExecutionOptionsPlanExecutionTypeFailoverPrecheck   DrPlanExecutionOptionsPlanExecutionTypeEnum = "FAILOVER_PRECHECK"
	DrPlanExecutionOptionsPlanExecutionTypeStartDrillPrecheck DrPlanExecutionOptionsPlanExecutionTypeEnum = "START_DRILL_PRECHECK"
	DrPlanExecutionOptionsPlanExecutionTypeStartDrill         DrPlanExecutionOptionsPlanExecutionTypeEnum = "START_DRILL"
	DrPlanExecutionOptionsPlanExecutionTypeStopDrillPrecheck  DrPlanExecutionOptionsPlanExecutionTypeEnum = "STOP_DRILL_PRECHECK"
	DrPlanExecutionOptionsPlanExecutionTypeStopDrill          DrPlanExecutionOptionsPlanExecutionTypeEnum = "STOP_DRILL"
)

var mappingDrPlanExecutionOptionsPlanExecutionTypeEnum = map[string]DrPlanExecutionOptionsPlanExecutionTypeEnum{
	"SWITCHOVER":           DrPlanExecutionOptionsPlanExecutionTypeSwitchover,
	"SWITCHOVER_PRECHECK":  DrPlanExecutionOptionsPlanExecutionTypeSwitchoverPrecheck,
	"FAILOVER":             DrPlanExecutionOptionsPlanExecutionTypeFailover,
	"FAILOVER_PRECHECK":    DrPlanExecutionOptionsPlanExecutionTypeFailoverPrecheck,
	"START_DRILL_PRECHECK": DrPlanExecutionOptionsPlanExecutionTypeStartDrillPrecheck,
	"START_DRILL":          DrPlanExecutionOptionsPlanExecutionTypeStartDrill,
	"STOP_DRILL_PRECHECK":  DrPlanExecutionOptionsPlanExecutionTypeStopDrillPrecheck,
	"STOP_DRILL":           DrPlanExecutionOptionsPlanExecutionTypeStopDrill,
}

var mappingDrPlanExecutionOptionsPlanExecutionTypeEnumLowerCase = map[string]DrPlanExecutionOptionsPlanExecutionTypeEnum{
	"switchover":           DrPlanExecutionOptionsPlanExecutionTypeSwitchover,
	"switchover_precheck":  DrPlanExecutionOptionsPlanExecutionTypeSwitchoverPrecheck,
	"failover":             DrPlanExecutionOptionsPlanExecutionTypeFailover,
	"failover_precheck":    DrPlanExecutionOptionsPlanExecutionTypeFailoverPrecheck,
	"start_drill_precheck": DrPlanExecutionOptionsPlanExecutionTypeStartDrillPrecheck,
	"start_drill":          DrPlanExecutionOptionsPlanExecutionTypeStartDrill,
	"stop_drill_precheck":  DrPlanExecutionOptionsPlanExecutionTypeStopDrillPrecheck,
	"stop_drill":           DrPlanExecutionOptionsPlanExecutionTypeStopDrill,
}

// GetDrPlanExecutionOptionsPlanExecutionTypeEnumValues Enumerates the set of values for DrPlanExecutionOptionsPlanExecutionTypeEnum
func GetDrPlanExecutionOptionsPlanExecutionTypeEnumValues() []DrPlanExecutionOptionsPlanExecutionTypeEnum {
	values := make([]DrPlanExecutionOptionsPlanExecutionTypeEnum, 0)
	for _, v := range mappingDrPlanExecutionOptionsPlanExecutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanExecutionOptionsPlanExecutionTypeEnumStringValues Enumerates the set of values in String for DrPlanExecutionOptionsPlanExecutionTypeEnum
func GetDrPlanExecutionOptionsPlanExecutionTypeEnumStringValues() []string {
	return []string{
		"SWITCHOVER",
		"SWITCHOVER_PRECHECK",
		"FAILOVER",
		"FAILOVER_PRECHECK",
		"START_DRILL_PRECHECK",
		"START_DRILL",
		"STOP_DRILL_PRECHECK",
		"STOP_DRILL",
	}
}

// GetMappingDrPlanExecutionOptionsPlanExecutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanExecutionOptionsPlanExecutionTypeEnum(val string) (DrPlanExecutionOptionsPlanExecutionTypeEnum, bool) {
	enum, ok := mappingDrPlanExecutionOptionsPlanExecutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
