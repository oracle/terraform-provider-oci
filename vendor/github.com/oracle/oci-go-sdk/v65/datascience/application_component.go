// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ApplicationComponent Reference to an application component
type ApplicationComponent interface {

	// Name of application component
	GetName() *string
}

type applicationcomponent struct {
	JsonData []byte
	Name     *string `mandatory:"true" json:"name"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *applicationcomponent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerapplicationcomponent applicationcomponent
	s := struct {
		Model Unmarshalerapplicationcomponent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *applicationcomponent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DATA_SCIENCE_JOB":
		mm := DataScienceJobApplicationComponent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_INTEGRATION_PIPELINE":
		mm := DataIntegrationPipelineApplicationComponent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_SCIENCE_PIPELINE":
		mm := DataSciencePipelineApplicationComponent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ApplicationComponent: %s.", m.Type)
		return *m, nil
	}
}

// GetName returns Name
func (m applicationcomponent) GetName() *string {
	return m.Name
}

func (m applicationcomponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m applicationcomponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApplicationComponentTypeEnum Enum with underlying type: string
type ApplicationComponentTypeEnum string

// Set of constants representing the allowable values for ApplicationComponentTypeEnum
const (
	ApplicationComponentTypeSciencePipeline     ApplicationComponentTypeEnum = "DATA_SCIENCE_PIPELINE"
	ApplicationComponentTypeScienceJob          ApplicationComponentTypeEnum = "DATA_SCIENCE_JOB"
	ApplicationComponentTypeIntegrationPipeline ApplicationComponentTypeEnum = "DATA_INTEGRATION_PIPELINE"
)

var mappingApplicationComponentTypeEnum = map[string]ApplicationComponentTypeEnum{
	"DATA_SCIENCE_PIPELINE":     ApplicationComponentTypeSciencePipeline,
	"DATA_SCIENCE_JOB":          ApplicationComponentTypeScienceJob,
	"DATA_INTEGRATION_PIPELINE": ApplicationComponentTypeIntegrationPipeline,
}

var mappingApplicationComponentTypeEnumLowerCase = map[string]ApplicationComponentTypeEnum{
	"data_science_pipeline":     ApplicationComponentTypeSciencePipeline,
	"data_science_job":          ApplicationComponentTypeScienceJob,
	"data_integration_pipeline": ApplicationComponentTypeIntegrationPipeline,
}

// GetApplicationComponentTypeEnumValues Enumerates the set of values for ApplicationComponentTypeEnum
func GetApplicationComponentTypeEnumValues() []ApplicationComponentTypeEnum {
	values := make([]ApplicationComponentTypeEnum, 0)
	for _, v := range mappingApplicationComponentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationComponentTypeEnumStringValues Enumerates the set of values in String for ApplicationComponentTypeEnum
func GetApplicationComponentTypeEnumStringValues() []string {
	return []string{
		"DATA_SCIENCE_PIPELINE",
		"DATA_SCIENCE_JOB",
		"DATA_INTEGRATION_PIPELINE",
	}
}

// GetMappingApplicationComponentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationComponentTypeEnum(val string) (ApplicationComponentTypeEnum, bool) {
	enum, ok := mappingApplicationComponentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
