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

// ModelGroupCloneSourceDetails Model Group clone source details.
type ModelGroupCloneSourceDetails interface {
}

type modelgroupclonesourcedetails struct {
	JsonData                  []byte
	ModelGroupCloneSourceType string `json:"modelGroupCloneSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *modelgroupclonesourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodelgroupclonesourcedetails modelgroupclonesourcedetails
	s := struct {
		Model Unmarshalermodelgroupclonesourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelGroupCloneSourceType = s.Model.ModelGroupCloneSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modelgroupclonesourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelGroupCloneSourceType {
	case "MODEL_GROUP_VERSION_HISTORY":
		mm := CloneCreateFromModelGroupVersionHistoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MODEL_GROUP":
		mm := CloneCreateFromModelGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ModelGroupCloneSourceDetails: %s.", m.ModelGroupCloneSourceType)
		return *m, nil
	}
}

func (m modelgroupclonesourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modelgroupclonesourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum Enum with underlying type: string
type ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum string

// Set of constants representing the allowable values for ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum
const (
	ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeGroup               ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum = "MODEL_GROUP"
	ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeGroupVersionHistory ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum = "MODEL_GROUP_VERSION_HISTORY"
)

var mappingModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum = map[string]ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum{
	"MODEL_GROUP":                 ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeGroup,
	"MODEL_GROUP_VERSION_HISTORY": ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeGroupVersionHistory,
}

var mappingModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnumLowerCase = map[string]ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum{
	"model_group":                 ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeGroup,
	"model_group_version_history": ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeGroupVersionHistory,
}

// GetModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnumValues Enumerates the set of values for ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum
func GetModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnumValues() []ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum {
	values := make([]ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum, 0)
	for _, v := range mappingModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnumStringValues Enumerates the set of values in String for ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum
func GetModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnumStringValues() []string {
	return []string{
		"MODEL_GROUP",
		"MODEL_GROUP_VERSION_HISTORY",
	}
}

// GetMappingModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum(val string) (ModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnum, bool) {
	enum, ok := mappingModelGroupCloneSourceDetailsModelGroupCloneSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
