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

// DrPlanExecutionControlDetails The details for controlling plan execution.
type DrPlanExecutionControlDetails interface {
}

type drplanexecutioncontroldetails struct {
	JsonData   []byte
	ActionType string `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *drplanexecutioncontroldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdrplanexecutioncontroldetails drplanexecutioncontroldetails
	s := struct {
		Model Unmarshalerdrplanexecutioncontroldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *drplanexecutioncontroldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "PAUSE":
		mm := PauseDrPlanExecutionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CANCEL":
		mm := CancelDrPlanExecutionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESUME":
		mm := ResumeDrPlanExecutionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DrPlanExecutionControlDetails: %s.", m.ActionType)
		return *m, nil
	}
}

func (m drplanexecutioncontroldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m drplanexecutioncontroldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrPlanExecutionControlDetailsActionTypeEnum Enum with underlying type: string
type DrPlanExecutionControlDetailsActionTypeEnum string

// Set of constants representing the allowable values for DrPlanExecutionControlDetailsActionTypeEnum
const (
	DrPlanExecutionControlDetailsActionTypeCancel DrPlanExecutionControlDetailsActionTypeEnum = "CANCEL"
	DrPlanExecutionControlDetailsActionTypePause  DrPlanExecutionControlDetailsActionTypeEnum = "PAUSE"
	DrPlanExecutionControlDetailsActionTypeResume DrPlanExecutionControlDetailsActionTypeEnum = "RESUME"
)

var mappingDrPlanExecutionControlDetailsActionTypeEnum = map[string]DrPlanExecutionControlDetailsActionTypeEnum{
	"CANCEL": DrPlanExecutionControlDetailsActionTypeCancel,
	"PAUSE":  DrPlanExecutionControlDetailsActionTypePause,
	"RESUME": DrPlanExecutionControlDetailsActionTypeResume,
}

var mappingDrPlanExecutionControlDetailsActionTypeEnumLowerCase = map[string]DrPlanExecutionControlDetailsActionTypeEnum{
	"cancel": DrPlanExecutionControlDetailsActionTypeCancel,
	"pause":  DrPlanExecutionControlDetailsActionTypePause,
	"resume": DrPlanExecutionControlDetailsActionTypeResume,
}

// GetDrPlanExecutionControlDetailsActionTypeEnumValues Enumerates the set of values for DrPlanExecutionControlDetailsActionTypeEnum
func GetDrPlanExecutionControlDetailsActionTypeEnumValues() []DrPlanExecutionControlDetailsActionTypeEnum {
	values := make([]DrPlanExecutionControlDetailsActionTypeEnum, 0)
	for _, v := range mappingDrPlanExecutionControlDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanExecutionControlDetailsActionTypeEnumStringValues Enumerates the set of values in String for DrPlanExecutionControlDetailsActionTypeEnum
func GetDrPlanExecutionControlDetailsActionTypeEnumStringValues() []string {
	return []string{
		"CANCEL",
		"PAUSE",
		"RESUME",
	}
}

// GetMappingDrPlanExecutionControlDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanExecutionControlDetailsActionTypeEnum(val string) (DrPlanExecutionControlDetailsActionTypeEnum, bool) {
	enum, ok := mappingDrPlanExecutionControlDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
