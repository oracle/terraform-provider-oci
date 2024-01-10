// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PipelineConfigurationDetails The configuration details of a pipeline.
type PipelineConfigurationDetails interface {
}

type pipelineconfigurationdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *pipelineconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipelineconfigurationdetails pipelineconfigurationdetails
	s := struct {
		Model Unmarshalerpipelineconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pipelineconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := PipelineDefaultConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PipelineConfigurationDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m pipelineconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipelineconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PipelineConfigurationDetailsTypeEnum Enum with underlying type: string
type PipelineConfigurationDetailsTypeEnum string

// Set of constants representing the allowable values for PipelineConfigurationDetailsTypeEnum
const (
	PipelineConfigurationDetailsTypeDefault PipelineConfigurationDetailsTypeEnum = "DEFAULT"
)

var mappingPipelineConfigurationDetailsTypeEnum = map[string]PipelineConfigurationDetailsTypeEnum{
	"DEFAULT": PipelineConfigurationDetailsTypeDefault,
}

var mappingPipelineConfigurationDetailsTypeEnumLowerCase = map[string]PipelineConfigurationDetailsTypeEnum{
	"default": PipelineConfigurationDetailsTypeDefault,
}

// GetPipelineConfigurationDetailsTypeEnumValues Enumerates the set of values for PipelineConfigurationDetailsTypeEnum
func GetPipelineConfigurationDetailsTypeEnumValues() []PipelineConfigurationDetailsTypeEnum {
	values := make([]PipelineConfigurationDetailsTypeEnum, 0)
	for _, v := range mappingPipelineConfigurationDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineConfigurationDetailsTypeEnumStringValues Enumerates the set of values in String for PipelineConfigurationDetailsTypeEnum
func GetPipelineConfigurationDetailsTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingPipelineConfigurationDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineConfigurationDetailsTypeEnum(val string) (PipelineConfigurationDetailsTypeEnum, bool) {
	enum, ok := mappingPipelineConfigurationDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
