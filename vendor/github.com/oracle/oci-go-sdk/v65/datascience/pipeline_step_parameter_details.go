// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PipelineStepParameterDetails Pipeline step parameter details
type PipelineStepParameterDetails interface {
}

type pipelinestepparameterdetails struct {
	JsonData      []byte
	ParameterType string `json:"parameterType"`
}

// UnmarshalJSON unmarshals json
func (m *pipelinestepparameterdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipelinestepparameterdetails pipelinestepparameterdetails
	s := struct {
		Model Unmarshalerpipelinestepparameterdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ParameterType = s.Model.ParameterType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pipelinestepparameterdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ParameterType {
	case "DEFAULT":
		mm := PipelineDefaultStepParameterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PipelineStepParameterDetails: %s.", m.ParameterType)
		return *m, nil
	}
}

func (m pipelinestepparameterdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipelinestepparameterdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PipelineStepParameterDetailsParameterTypeEnum Enum with underlying type: string
type PipelineStepParameterDetailsParameterTypeEnum string

// Set of constants representing the allowable values for PipelineStepParameterDetailsParameterTypeEnum
const (
	PipelineStepParameterDetailsParameterTypeDefault PipelineStepParameterDetailsParameterTypeEnum = "DEFAULT"
)

var mappingPipelineStepParameterDetailsParameterTypeEnum = map[string]PipelineStepParameterDetailsParameterTypeEnum{
	"DEFAULT": PipelineStepParameterDetailsParameterTypeDefault,
}

var mappingPipelineStepParameterDetailsParameterTypeEnumLowerCase = map[string]PipelineStepParameterDetailsParameterTypeEnum{
	"default": PipelineStepParameterDetailsParameterTypeDefault,
}

// GetPipelineStepParameterDetailsParameterTypeEnumValues Enumerates the set of values for PipelineStepParameterDetailsParameterTypeEnum
func GetPipelineStepParameterDetailsParameterTypeEnumValues() []PipelineStepParameterDetailsParameterTypeEnum {
	values := make([]PipelineStepParameterDetailsParameterTypeEnum, 0)
	for _, v := range mappingPipelineStepParameterDetailsParameterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineStepParameterDetailsParameterTypeEnumStringValues Enumerates the set of values in String for PipelineStepParameterDetailsParameterTypeEnum
func GetPipelineStepParameterDetailsParameterTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingPipelineStepParameterDetailsParameterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineStepParameterDetailsParameterTypeEnum(val string) (PipelineStepParameterDetailsParameterTypeEnum, bool) {
	enum, ok := mappingPipelineStepParameterDetailsParameterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
