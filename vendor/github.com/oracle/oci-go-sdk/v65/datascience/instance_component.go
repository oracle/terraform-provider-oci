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

// InstanceComponent Rereference to instance component
type InstanceComponent interface {
}

type instancecomponent struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *instancecomponent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstancecomponent instancecomponent
	s := struct {
		Model Unmarshalerinstancecomponent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instancecomponent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DATA_SCIENCE_MODEL_DEPLOYMENT":
		mm := DataScienceModelDeploymentInstanceComponent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ML_APPLICATION_INSTANCE_INTERNAL_TRIGGER":
		mm := MlApplicationInstanceInternalTrigger{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_BUCKET":
		mm := ObjectStorageBucketInstanceComponent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for InstanceComponent: %s.", m.Type)
		return *m, nil
	}
}

func (m instancecomponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m instancecomponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceComponentTypeEnum Enum with underlying type: string
type InstanceComponentTypeEnum string

// Set of constants representing the allowable values for InstanceComponentTypeEnum
const (
	InstanceComponentTypeDataScienceModelDeployment InstanceComponentTypeEnum = "DATA_SCIENCE_MODEL_DEPLOYMENT"
	InstanceComponentTypeObjectStorageBucket        InstanceComponentTypeEnum = "OBJECT_STORAGE_BUCKET"
)

var mappingInstanceComponentTypeEnum = map[string]InstanceComponentTypeEnum{
	"DATA_SCIENCE_MODEL_DEPLOYMENT": InstanceComponentTypeDataScienceModelDeployment,
	"OBJECT_STORAGE_BUCKET":         InstanceComponentTypeObjectStorageBucket,
}

var mappingInstanceComponentTypeEnumLowerCase = map[string]InstanceComponentTypeEnum{
	"data_science_model_deployment": InstanceComponentTypeDataScienceModelDeployment,
	"object_storage_bucket":         InstanceComponentTypeObjectStorageBucket,
}

// GetInstanceComponentTypeEnumValues Enumerates the set of values for InstanceComponentTypeEnum
func GetInstanceComponentTypeEnumValues() []InstanceComponentTypeEnum {
	values := make([]InstanceComponentTypeEnum, 0)
	for _, v := range mappingInstanceComponentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceComponentTypeEnumStringValues Enumerates the set of values in String for InstanceComponentTypeEnum
func GetInstanceComponentTypeEnumStringValues() []string {
	return []string{
		"DATA_SCIENCE_MODEL_DEPLOYMENT",
		"OBJECT_STORAGE_BUCKET",
	}
}

// GetMappingInstanceComponentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceComponentTypeEnum(val string) (InstanceComponentTypeEnum, bool) {
	enum, ok := mappingInstanceComponentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
