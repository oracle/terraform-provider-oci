// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// TriggerAction The trigger action to be performed.
type TriggerAction interface {
	GetFilter() Filter
}

type triggeraction struct {
	JsonData []byte
	Filter   Filter `mandatory:"false" json:"filter"`
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *triggeraction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertriggeraction triggeraction
	s := struct {
		Model Unmarshalertriggeraction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Filter = s.Model.Filter
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *triggeraction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TRIGGER_BUILD_PIPELINE":
		mm := TriggerBuildPipelineAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetFilter returns Filter
func (m triggeraction) GetFilter() Filter {
	return m.Filter
}

func (m triggeraction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m triggeraction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TriggerActionTypeEnum Enum with underlying type: string
type TriggerActionTypeEnum string

// Set of constants representing the allowable values for TriggerActionTypeEnum
const (
	TriggerActionTypeTriggerBuildPipeline TriggerActionTypeEnum = "TRIGGER_BUILD_PIPELINE"
)

var mappingTriggerActionTypeEnum = map[string]TriggerActionTypeEnum{
	"TRIGGER_BUILD_PIPELINE": TriggerActionTypeTriggerBuildPipeline,
}

// GetTriggerActionTypeEnumValues Enumerates the set of values for TriggerActionTypeEnum
func GetTriggerActionTypeEnumValues() []TriggerActionTypeEnum {
	values := make([]TriggerActionTypeEnum, 0)
	for _, v := range mappingTriggerActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTriggerActionTypeEnumStringValues Enumerates the set of values in String for TriggerActionTypeEnum
func GetTriggerActionTypeEnumStringValues() []string {
	return []string{
		"TRIGGER_BUILD_PIPELINE",
	}
}

// GetMappingTriggerActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTriggerActionTypeEnum(val string) (TriggerActionTypeEnum, bool) {
	mappingTriggerActionTypeEnumIgnoreCase := make(map[string]TriggerActionTypeEnum)
	for k, v := range mappingTriggerActionTypeEnum {
		mappingTriggerActionTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTriggerActionTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
