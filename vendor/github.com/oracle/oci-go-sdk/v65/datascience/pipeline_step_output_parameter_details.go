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

// PipelineStepOutputParameterDetails Pipeline output parameter details
type PipelineStepOutputParameterDetails interface {
}

type pipelinestepoutputparameterdetails struct {
	JsonData            []byte
	OutputParameterType string `json:"outputParameterType"`
}

// UnmarshalJSON unmarshals json
func (m *pipelinestepoutputparameterdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipelinestepoutputparameterdetails pipelinestepoutputparameterdetails
	s := struct {
		Model Unmarshalerpipelinestepoutputparameterdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OutputParameterType = s.Model.OutputParameterType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pipelinestepoutputparameterdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputParameterType {
	case "JSON":
		mm := PipelineJsonStepOutputParameterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PipelineStepOutputParameterDetails: %s.", m.OutputParameterType)
		return *m, nil
	}
}

func (m pipelinestepoutputparameterdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipelinestepoutputparameterdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PipelineStepOutputParameterDetailsOutputParameterTypeEnum Enum with underlying type: string
type PipelineStepOutputParameterDetailsOutputParameterTypeEnum string

// Set of constants representing the allowable values for PipelineStepOutputParameterDetailsOutputParameterTypeEnum
const (
	PipelineStepOutputParameterDetailsOutputParameterTypeJson PipelineStepOutputParameterDetailsOutputParameterTypeEnum = "JSON"
)

var mappingPipelineStepOutputParameterDetailsOutputParameterTypeEnum = map[string]PipelineStepOutputParameterDetailsOutputParameterTypeEnum{
	"JSON": PipelineStepOutputParameterDetailsOutputParameterTypeJson,
}

var mappingPipelineStepOutputParameterDetailsOutputParameterTypeEnumLowerCase = map[string]PipelineStepOutputParameterDetailsOutputParameterTypeEnum{
	"json": PipelineStepOutputParameterDetailsOutputParameterTypeJson,
}

// GetPipelineStepOutputParameterDetailsOutputParameterTypeEnumValues Enumerates the set of values for PipelineStepOutputParameterDetailsOutputParameterTypeEnum
func GetPipelineStepOutputParameterDetailsOutputParameterTypeEnumValues() []PipelineStepOutputParameterDetailsOutputParameterTypeEnum {
	values := make([]PipelineStepOutputParameterDetailsOutputParameterTypeEnum, 0)
	for _, v := range mappingPipelineStepOutputParameterDetailsOutputParameterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineStepOutputParameterDetailsOutputParameterTypeEnumStringValues Enumerates the set of values in String for PipelineStepOutputParameterDetailsOutputParameterTypeEnum
func GetPipelineStepOutputParameterDetailsOutputParameterTypeEnumStringValues() []string {
	return []string{
		"JSON",
	}
}

// GetMappingPipelineStepOutputParameterDetailsOutputParameterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineStepOutputParameterDetailsOutputParameterTypeEnum(val string) (PipelineStepOutputParameterDetailsOutputParameterTypeEnum, bool) {
	enum, ok := mappingPipelineStepOutputParameterDetailsOutputParameterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
