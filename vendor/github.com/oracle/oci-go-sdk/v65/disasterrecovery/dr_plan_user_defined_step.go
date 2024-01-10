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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanUserDefinedStep The details for a user-defined step in a DR plan.
type DrPlanUserDefinedStep interface {
}

type drplanuserdefinedstep struct {
	JsonData []byte
	StepType string `json:"stepType"`
}

// UnmarshalJSON unmarshals json
func (m *drplanuserdefinedstep) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdrplanuserdefinedstep drplanuserdefinedstep
	s := struct {
		Model Unmarshalerdrplanuserdefinedstep
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StepType = s.Model.StepType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *drplanuserdefinedstep) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StepType {
	case "INVOKE_FUNCTION":
		mm := InvokeFunctionStep{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION_PRECHECK":
		mm := InvokeFunctionPrecheckStep{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_LOCAL_SCRIPT":
		mm := RunLocalScriptUserDefinedStep{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_LOCAL_SCRIPT_PRECHECK":
		mm := LocalScriptPrecheckStep{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OBJECTSTORE_SCRIPT_PRECHECK":
		mm := ObjectStoreScriptPrecheckStep{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OBJECTSTORE_SCRIPT":
		mm := RunObjectStoreScriptUserDefinedStep{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DrPlanUserDefinedStep: %s.", m.StepType)
		return *m, nil
	}
}

func (m drplanuserdefinedstep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m drplanuserdefinedstep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrPlanUserDefinedStepStepTypeEnum Enum with underlying type: string
type DrPlanUserDefinedStepStepTypeEnum string

// Set of constants representing the allowable values for DrPlanUserDefinedStepStepTypeEnum
const (
	DrPlanUserDefinedStepStepTypeRunObjectstoreScriptPrecheck DrPlanUserDefinedStepStepTypeEnum = "RUN_OBJECTSTORE_SCRIPT_PRECHECK"
	DrPlanUserDefinedStepStepTypeRunLocalScriptPrecheck       DrPlanUserDefinedStepStepTypeEnum = "RUN_LOCAL_SCRIPT_PRECHECK"
	DrPlanUserDefinedStepStepTypeInvokeFunctionPrecheck       DrPlanUserDefinedStepStepTypeEnum = "INVOKE_FUNCTION_PRECHECK"
	DrPlanUserDefinedStepStepTypeRunObjectstoreScript         DrPlanUserDefinedStepStepTypeEnum = "RUN_OBJECTSTORE_SCRIPT"
	DrPlanUserDefinedStepStepTypeRunLocalScript               DrPlanUserDefinedStepStepTypeEnum = "RUN_LOCAL_SCRIPT"
	DrPlanUserDefinedStepStepTypeInvokeFunction               DrPlanUserDefinedStepStepTypeEnum = "INVOKE_FUNCTION"
)

var mappingDrPlanUserDefinedStepStepTypeEnum = map[string]DrPlanUserDefinedStepStepTypeEnum{
	"RUN_OBJECTSTORE_SCRIPT_PRECHECK": DrPlanUserDefinedStepStepTypeRunObjectstoreScriptPrecheck,
	"RUN_LOCAL_SCRIPT_PRECHECK":       DrPlanUserDefinedStepStepTypeRunLocalScriptPrecheck,
	"INVOKE_FUNCTION_PRECHECK":        DrPlanUserDefinedStepStepTypeInvokeFunctionPrecheck,
	"RUN_OBJECTSTORE_SCRIPT":          DrPlanUserDefinedStepStepTypeRunObjectstoreScript,
	"RUN_LOCAL_SCRIPT":                DrPlanUserDefinedStepStepTypeRunLocalScript,
	"INVOKE_FUNCTION":                 DrPlanUserDefinedStepStepTypeInvokeFunction,
}

var mappingDrPlanUserDefinedStepStepTypeEnumLowerCase = map[string]DrPlanUserDefinedStepStepTypeEnum{
	"run_objectstore_script_precheck": DrPlanUserDefinedStepStepTypeRunObjectstoreScriptPrecheck,
	"run_local_script_precheck":       DrPlanUserDefinedStepStepTypeRunLocalScriptPrecheck,
	"invoke_function_precheck":        DrPlanUserDefinedStepStepTypeInvokeFunctionPrecheck,
	"run_objectstore_script":          DrPlanUserDefinedStepStepTypeRunObjectstoreScript,
	"run_local_script":                DrPlanUserDefinedStepStepTypeRunLocalScript,
	"invoke_function":                 DrPlanUserDefinedStepStepTypeInvokeFunction,
}

// GetDrPlanUserDefinedStepStepTypeEnumValues Enumerates the set of values for DrPlanUserDefinedStepStepTypeEnum
func GetDrPlanUserDefinedStepStepTypeEnumValues() []DrPlanUserDefinedStepStepTypeEnum {
	values := make([]DrPlanUserDefinedStepStepTypeEnum, 0)
	for _, v := range mappingDrPlanUserDefinedStepStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanUserDefinedStepStepTypeEnumStringValues Enumerates the set of values in String for DrPlanUserDefinedStepStepTypeEnum
func GetDrPlanUserDefinedStepStepTypeEnumStringValues() []string {
	return []string{
		"RUN_OBJECTSTORE_SCRIPT_PRECHECK",
		"RUN_LOCAL_SCRIPT_PRECHECK",
		"INVOKE_FUNCTION_PRECHECK",
		"RUN_OBJECTSTORE_SCRIPT",
		"RUN_LOCAL_SCRIPT",
		"INVOKE_FUNCTION",
	}
}

// GetMappingDrPlanUserDefinedStepStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanUserDefinedStepStepTypeEnum(val string) (DrPlanUserDefinedStepStepTypeEnum, bool) {
	enum, ok := mappingDrPlanUserDefinedStepStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
