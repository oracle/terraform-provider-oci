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

// PipelineContainerConfigurationDetails Container Details for a step in pipeline.
type PipelineContainerConfigurationDetails interface {
}

type pipelinecontainerconfigurationdetails struct {
	JsonData      []byte
	ContainerType string `json:"containerType"`
}

// UnmarshalJSON unmarshals json
func (m *pipelinecontainerconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipelinecontainerconfigurationdetails pipelinecontainerconfigurationdetails
	s := struct {
		Model Unmarshalerpipelinecontainerconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ContainerType = s.Model.ContainerType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pipelinecontainerconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ContainerType {
	case "OCIR_CONTAINER":
		mm := PipelineOcirContainerConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PipelineContainerConfigurationDetails: %s.", m.ContainerType)
		return *m, nil
	}
}

func (m pipelinecontainerconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipelinecontainerconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PipelineContainerConfigurationDetailsContainerTypeEnum Enum with underlying type: string
type PipelineContainerConfigurationDetailsContainerTypeEnum string

// Set of constants representing the allowable values for PipelineContainerConfigurationDetailsContainerTypeEnum
const (
	PipelineContainerConfigurationDetailsContainerTypeOcirContainer PipelineContainerConfigurationDetailsContainerTypeEnum = "OCIR_CONTAINER"
)

var mappingPipelineContainerConfigurationDetailsContainerTypeEnum = map[string]PipelineContainerConfigurationDetailsContainerTypeEnum{
	"OCIR_CONTAINER": PipelineContainerConfigurationDetailsContainerTypeOcirContainer,
}

var mappingPipelineContainerConfigurationDetailsContainerTypeEnumLowerCase = map[string]PipelineContainerConfigurationDetailsContainerTypeEnum{
	"ocir_container": PipelineContainerConfigurationDetailsContainerTypeOcirContainer,
}

// GetPipelineContainerConfigurationDetailsContainerTypeEnumValues Enumerates the set of values for PipelineContainerConfigurationDetailsContainerTypeEnum
func GetPipelineContainerConfigurationDetailsContainerTypeEnumValues() []PipelineContainerConfigurationDetailsContainerTypeEnum {
	values := make([]PipelineContainerConfigurationDetailsContainerTypeEnum, 0)
	for _, v := range mappingPipelineContainerConfigurationDetailsContainerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineContainerConfigurationDetailsContainerTypeEnumStringValues Enumerates the set of values in String for PipelineContainerConfigurationDetailsContainerTypeEnum
func GetPipelineContainerConfigurationDetailsContainerTypeEnumStringValues() []string {
	return []string{
		"OCIR_CONTAINER",
	}
}

// GetMappingPipelineContainerConfigurationDetailsContainerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineContainerConfigurationDetailsContainerTypeEnum(val string) (PipelineContainerConfigurationDetailsContainerTypeEnum, bool) {
	enum, ok := mappingPipelineContainerConfigurationDetailsContainerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
