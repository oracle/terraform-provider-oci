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

// ModelArtifactReferenceDetails Reference to the location where the model artifact is stored.
type ModelArtifactReferenceDetails interface {
}

type modelartifactreferencedetails struct {
	JsonData                   []byte
	ModelArtifactReferenceType string `json:"modelArtifactReferenceType"`
}

// UnmarshalJSON unmarshals json
func (m *modelartifactreferencedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodelartifactreferencedetails modelartifactreferencedetails
	s := struct {
		Model Unmarshalermodelartifactreferencedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelArtifactReferenceType = s.Model.ModelArtifactReferenceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modelartifactreferencedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelArtifactReferenceType {
	case "OSS":
		mm := OssModelArtifactReferenceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ModelArtifactReferenceDetails: %s.", m.ModelArtifactReferenceType)
		return *m, nil
	}
}

func (m modelartifactreferencedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modelartifactreferencedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum Enum with underlying type: string
type ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum string

// Set of constants representing the allowable values for ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum
const (
	ModelArtifactReferenceDetailsModelArtifactReferenceTypeOss ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum = "OSS"
)

var mappingModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum = map[string]ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum{
	"OSS": ModelArtifactReferenceDetailsModelArtifactReferenceTypeOss,
}

var mappingModelArtifactReferenceDetailsModelArtifactReferenceTypeEnumLowerCase = map[string]ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum{
	"oss": ModelArtifactReferenceDetailsModelArtifactReferenceTypeOss,
}

// GetModelArtifactReferenceDetailsModelArtifactReferenceTypeEnumValues Enumerates the set of values for ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum
func GetModelArtifactReferenceDetailsModelArtifactReferenceTypeEnumValues() []ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum {
	values := make([]ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum, 0)
	for _, v := range mappingModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelArtifactReferenceDetailsModelArtifactReferenceTypeEnumStringValues Enumerates the set of values in String for ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum
func GetModelArtifactReferenceDetailsModelArtifactReferenceTypeEnumStringValues() []string {
	return []string{
		"OSS",
	}
}

// GetMappingModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum(val string) (ModelArtifactReferenceDetailsModelArtifactReferenceTypeEnum, bool) {
	enum, ok := mappingModelArtifactReferenceDetailsModelArtifactReferenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
