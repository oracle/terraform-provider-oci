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

// UpdateDrPlanUserDefinedStepDetails The details for updating a user-defined step in a DR plan.
type UpdateDrPlanUserDefinedStepDetails interface {
}

type updatedrplanuserdefinedstepdetails struct {
	JsonData []byte
	StepType string `json:"stepType"`
}

// UnmarshalJSON unmarshals json
func (m *updatedrplanuserdefinedstepdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedrplanuserdefinedstepdetails updatedrplanuserdefinedstepdetails
	s := struct {
		Model Unmarshalerupdatedrplanuserdefinedstepdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StepType = s.Model.StepType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedrplanuserdefinedstepdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StepType {
	case "RUN_LOCAL_SCRIPT_PRECHECK":
		mm := UpdateLocalScriptPrecheckStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION_PRECHECK":
		mm := UpdateInvokeFunctionPrecheckStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := UpdateInvokeFunctionUserDefinedStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OBJECTSTORE_SCRIPT":
		mm := UpdateRunObjectStoreScriptUserDefinedStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OBJECTSTORE_SCRIPT_PRECHECK":
		mm := UpdateObjectStoreScriptPrecheckStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_LOCAL_SCRIPT":
		mm := UpdateRunLocalScriptUserDefinedStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateDrPlanUserDefinedStepDetails: %s.", m.StepType)
		return *m, nil
	}
}

func (m updatedrplanuserdefinedstepdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedrplanuserdefinedstepdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDrPlanUserDefinedStepDetailsStepTypeEnum Enum with underlying type: string
type UpdateDrPlanUserDefinedStepDetailsStepTypeEnum string

// Set of constants representing the allowable values for UpdateDrPlanUserDefinedStepDetailsStepTypeEnum
const (
	UpdateDrPlanUserDefinedStepDetailsStepTypeRunObjectstoreScriptPrecheck UpdateDrPlanUserDefinedStepDetailsStepTypeEnum = "RUN_OBJECTSTORE_SCRIPT_PRECHECK"
	UpdateDrPlanUserDefinedStepDetailsStepTypeRunLocalScriptPrecheck       UpdateDrPlanUserDefinedStepDetailsStepTypeEnum = "RUN_LOCAL_SCRIPT_PRECHECK"
	UpdateDrPlanUserDefinedStepDetailsStepTypeInvokeFunctionPrecheck       UpdateDrPlanUserDefinedStepDetailsStepTypeEnum = "INVOKE_FUNCTION_PRECHECK"
	UpdateDrPlanUserDefinedStepDetailsStepTypeRunObjectstoreScript         UpdateDrPlanUserDefinedStepDetailsStepTypeEnum = "RUN_OBJECTSTORE_SCRIPT"
	UpdateDrPlanUserDefinedStepDetailsStepTypeRunLocalScript               UpdateDrPlanUserDefinedStepDetailsStepTypeEnum = "RUN_LOCAL_SCRIPT"
	UpdateDrPlanUserDefinedStepDetailsStepTypeInvokeFunction               UpdateDrPlanUserDefinedStepDetailsStepTypeEnum = "INVOKE_FUNCTION"
)

var mappingUpdateDrPlanUserDefinedStepDetailsStepTypeEnum = map[string]UpdateDrPlanUserDefinedStepDetailsStepTypeEnum{
	"RUN_OBJECTSTORE_SCRIPT_PRECHECK": UpdateDrPlanUserDefinedStepDetailsStepTypeRunObjectstoreScriptPrecheck,
	"RUN_LOCAL_SCRIPT_PRECHECK":       UpdateDrPlanUserDefinedStepDetailsStepTypeRunLocalScriptPrecheck,
	"INVOKE_FUNCTION_PRECHECK":        UpdateDrPlanUserDefinedStepDetailsStepTypeInvokeFunctionPrecheck,
	"RUN_OBJECTSTORE_SCRIPT":          UpdateDrPlanUserDefinedStepDetailsStepTypeRunObjectstoreScript,
	"RUN_LOCAL_SCRIPT":                UpdateDrPlanUserDefinedStepDetailsStepTypeRunLocalScript,
	"INVOKE_FUNCTION":                 UpdateDrPlanUserDefinedStepDetailsStepTypeInvokeFunction,
}

var mappingUpdateDrPlanUserDefinedStepDetailsStepTypeEnumLowerCase = map[string]UpdateDrPlanUserDefinedStepDetailsStepTypeEnum{
	"run_objectstore_script_precheck": UpdateDrPlanUserDefinedStepDetailsStepTypeRunObjectstoreScriptPrecheck,
	"run_local_script_precheck":       UpdateDrPlanUserDefinedStepDetailsStepTypeRunLocalScriptPrecheck,
	"invoke_function_precheck":        UpdateDrPlanUserDefinedStepDetailsStepTypeInvokeFunctionPrecheck,
	"run_objectstore_script":          UpdateDrPlanUserDefinedStepDetailsStepTypeRunObjectstoreScript,
	"run_local_script":                UpdateDrPlanUserDefinedStepDetailsStepTypeRunLocalScript,
	"invoke_function":                 UpdateDrPlanUserDefinedStepDetailsStepTypeInvokeFunction,
}

// GetUpdateDrPlanUserDefinedStepDetailsStepTypeEnumValues Enumerates the set of values for UpdateDrPlanUserDefinedStepDetailsStepTypeEnum
func GetUpdateDrPlanUserDefinedStepDetailsStepTypeEnumValues() []UpdateDrPlanUserDefinedStepDetailsStepTypeEnum {
	values := make([]UpdateDrPlanUserDefinedStepDetailsStepTypeEnum, 0)
	for _, v := range mappingUpdateDrPlanUserDefinedStepDetailsStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDrPlanUserDefinedStepDetailsStepTypeEnumStringValues Enumerates the set of values in String for UpdateDrPlanUserDefinedStepDetailsStepTypeEnum
func GetUpdateDrPlanUserDefinedStepDetailsStepTypeEnumStringValues() []string {
	return []string{
		"RUN_OBJECTSTORE_SCRIPT_PRECHECK",
		"RUN_LOCAL_SCRIPT_PRECHECK",
		"INVOKE_FUNCTION_PRECHECK",
		"RUN_OBJECTSTORE_SCRIPT",
		"RUN_LOCAL_SCRIPT",
		"INVOKE_FUNCTION",
	}
}

// GetMappingUpdateDrPlanUserDefinedStepDetailsStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDrPlanUserDefinedStepDetailsStepTypeEnum(val string) (UpdateDrPlanUserDefinedStepDetailsStepTypeEnum, bool) {
	enum, ok := mappingUpdateDrPlanUserDefinedStepDetailsStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
