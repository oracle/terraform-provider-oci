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

// DrPlanExecutionOptionDetails The options for a plan execution.
type DrPlanExecutionOptionDetails interface {
}

type drplanexecutionoptiondetails struct {
	JsonData          []byte
	PlanExecutionType string `json:"planExecutionType"`
}

// UnmarshalJSON unmarshals json
func (m *drplanexecutionoptiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdrplanexecutionoptiondetails drplanexecutionoptiondetails
	s := struct {
		Model Unmarshalerdrplanexecutionoptiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PlanExecutionType = s.Model.PlanExecutionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *drplanexecutionoptiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PlanExecutionType {
	case "STOP_DRILL_PRECHECK":
		mm := StopDrillPrecheckExecutionOptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SWITCHOVER_PRECHECK":
		mm := SwitchoverPrecheckExecutionOptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STOP_DRILL":
		mm := StopDrillExecutionOptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAILOVER_PRECHECK":
		mm := FailoverPrecheckExecutionOptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "START_DRILL":
		mm := StartDrillExecutionOptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "START_DRILL_PRECHECK":
		mm := StartDrillPrecheckExecutionOptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SWITCHOVER":
		mm := SwitchoverExecutionOptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAILOVER":
		mm := FailoverExecutionOptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DrPlanExecutionOptionDetails: %s.", m.PlanExecutionType)
		return *m, nil
	}
}

func (m drplanexecutionoptiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m drplanexecutionoptiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrPlanExecutionOptionDetailsPlanExecutionTypeEnum Enum with underlying type: string
type DrPlanExecutionOptionDetailsPlanExecutionTypeEnum string

// Set of constants representing the allowable values for DrPlanExecutionOptionDetailsPlanExecutionTypeEnum
const (
	DrPlanExecutionOptionDetailsPlanExecutionTypeSwitchover         DrPlanExecutionOptionDetailsPlanExecutionTypeEnum = "SWITCHOVER"
	DrPlanExecutionOptionDetailsPlanExecutionTypeSwitchoverPrecheck DrPlanExecutionOptionDetailsPlanExecutionTypeEnum = "SWITCHOVER_PRECHECK"
	DrPlanExecutionOptionDetailsPlanExecutionTypeFailover           DrPlanExecutionOptionDetailsPlanExecutionTypeEnum = "FAILOVER"
	DrPlanExecutionOptionDetailsPlanExecutionTypeFailoverPrecheck   DrPlanExecutionOptionDetailsPlanExecutionTypeEnum = "FAILOVER_PRECHECK"
	DrPlanExecutionOptionDetailsPlanExecutionTypeStartDrill         DrPlanExecutionOptionDetailsPlanExecutionTypeEnum = "START_DRILL"
	DrPlanExecutionOptionDetailsPlanExecutionTypeStartDrillPrecheck DrPlanExecutionOptionDetailsPlanExecutionTypeEnum = "START_DRILL_PRECHECK"
	DrPlanExecutionOptionDetailsPlanExecutionTypeStopDrill          DrPlanExecutionOptionDetailsPlanExecutionTypeEnum = "STOP_DRILL"
	DrPlanExecutionOptionDetailsPlanExecutionTypeStopDrillPrecheck  DrPlanExecutionOptionDetailsPlanExecutionTypeEnum = "STOP_DRILL_PRECHECK"
)

var mappingDrPlanExecutionOptionDetailsPlanExecutionTypeEnum = map[string]DrPlanExecutionOptionDetailsPlanExecutionTypeEnum{
	"SWITCHOVER":           DrPlanExecutionOptionDetailsPlanExecutionTypeSwitchover,
	"SWITCHOVER_PRECHECK":  DrPlanExecutionOptionDetailsPlanExecutionTypeSwitchoverPrecheck,
	"FAILOVER":             DrPlanExecutionOptionDetailsPlanExecutionTypeFailover,
	"FAILOVER_PRECHECK":    DrPlanExecutionOptionDetailsPlanExecutionTypeFailoverPrecheck,
	"START_DRILL":          DrPlanExecutionOptionDetailsPlanExecutionTypeStartDrill,
	"START_DRILL_PRECHECK": DrPlanExecutionOptionDetailsPlanExecutionTypeStartDrillPrecheck,
	"STOP_DRILL":           DrPlanExecutionOptionDetailsPlanExecutionTypeStopDrill,
	"STOP_DRILL_PRECHECK":  DrPlanExecutionOptionDetailsPlanExecutionTypeStopDrillPrecheck,
}

var mappingDrPlanExecutionOptionDetailsPlanExecutionTypeEnumLowerCase = map[string]DrPlanExecutionOptionDetailsPlanExecutionTypeEnum{
	"switchover":           DrPlanExecutionOptionDetailsPlanExecutionTypeSwitchover,
	"switchover_precheck":  DrPlanExecutionOptionDetailsPlanExecutionTypeSwitchoverPrecheck,
	"failover":             DrPlanExecutionOptionDetailsPlanExecutionTypeFailover,
	"failover_precheck":    DrPlanExecutionOptionDetailsPlanExecutionTypeFailoverPrecheck,
	"start_drill":          DrPlanExecutionOptionDetailsPlanExecutionTypeStartDrill,
	"start_drill_precheck": DrPlanExecutionOptionDetailsPlanExecutionTypeStartDrillPrecheck,
	"stop_drill":           DrPlanExecutionOptionDetailsPlanExecutionTypeStopDrill,
	"stop_drill_precheck":  DrPlanExecutionOptionDetailsPlanExecutionTypeStopDrillPrecheck,
}

// GetDrPlanExecutionOptionDetailsPlanExecutionTypeEnumValues Enumerates the set of values for DrPlanExecutionOptionDetailsPlanExecutionTypeEnum
func GetDrPlanExecutionOptionDetailsPlanExecutionTypeEnumValues() []DrPlanExecutionOptionDetailsPlanExecutionTypeEnum {
	values := make([]DrPlanExecutionOptionDetailsPlanExecutionTypeEnum, 0)
	for _, v := range mappingDrPlanExecutionOptionDetailsPlanExecutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanExecutionOptionDetailsPlanExecutionTypeEnumStringValues Enumerates the set of values in String for DrPlanExecutionOptionDetailsPlanExecutionTypeEnum
func GetDrPlanExecutionOptionDetailsPlanExecutionTypeEnumStringValues() []string {
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

// GetMappingDrPlanExecutionOptionDetailsPlanExecutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanExecutionOptionDetailsPlanExecutionTypeEnum(val string) (DrPlanExecutionOptionDetailsPlanExecutionTypeEnum, bool) {
	enum, ok := mappingDrPlanExecutionOptionDetailsPlanExecutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
