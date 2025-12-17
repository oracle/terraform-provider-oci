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

// ModelDeploymentSystemData Model deployment system data.
type ModelDeploymentSystemData interface {

	// The type of the deployed model.
	GetModelType() ModelDeploymentModelTypeEnum
}

type modeldeploymentsystemdata struct {
	JsonData        []byte
	ModelType       ModelDeploymentModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`
	SystemInfraType string                       `json:"systemInfraType"`
}

// UnmarshalJSON unmarshals json
func (m *modeldeploymentsystemdata) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodeldeploymentsystemdata modeldeploymentsystemdata
	s := struct {
		Model Unmarshalermodeldeploymentsystemdata
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType
	m.SystemInfraType = s.Model.SystemInfraType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modeldeploymentsystemdata) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SystemInfraType {
	case "INSTANCE_POOL":
		mm := InstancePoolModelDeploymentSystemData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ModelDeploymentSystemData: %s.", m.SystemInfraType)
		return *m, nil
	}
}

// GetModelType returns ModelType
func (m modeldeploymentsystemdata) GetModelType() ModelDeploymentModelTypeEnum {
	return m.ModelType
}

func (m modeldeploymentsystemdata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modeldeploymentsystemdata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingModelDeploymentModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetModelDeploymentModelTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelDeploymentSystemDataSystemInfraTypeEnum Enum with underlying type: string
type ModelDeploymentSystemDataSystemInfraTypeEnum string

// Set of constants representing the allowable values for ModelDeploymentSystemDataSystemInfraTypeEnum
const (
	ModelDeploymentSystemDataSystemInfraTypeInstancePool ModelDeploymentSystemDataSystemInfraTypeEnum = "INSTANCE_POOL"
)

var mappingModelDeploymentSystemDataSystemInfraTypeEnum = map[string]ModelDeploymentSystemDataSystemInfraTypeEnum{
	"INSTANCE_POOL": ModelDeploymentSystemDataSystemInfraTypeInstancePool,
}

var mappingModelDeploymentSystemDataSystemInfraTypeEnumLowerCase = map[string]ModelDeploymentSystemDataSystemInfraTypeEnum{
	"instance_pool": ModelDeploymentSystemDataSystemInfraTypeInstancePool,
}

// GetModelDeploymentSystemDataSystemInfraTypeEnumValues Enumerates the set of values for ModelDeploymentSystemDataSystemInfraTypeEnum
func GetModelDeploymentSystemDataSystemInfraTypeEnumValues() []ModelDeploymentSystemDataSystemInfraTypeEnum {
	values := make([]ModelDeploymentSystemDataSystemInfraTypeEnum, 0)
	for _, v := range mappingModelDeploymentSystemDataSystemInfraTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDeploymentSystemDataSystemInfraTypeEnumStringValues Enumerates the set of values in String for ModelDeploymentSystemDataSystemInfraTypeEnum
func GetModelDeploymentSystemDataSystemInfraTypeEnumStringValues() []string {
	return []string{
		"INSTANCE_POOL",
	}
}

// GetMappingModelDeploymentSystemDataSystemInfraTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDeploymentSystemDataSystemInfraTypeEnum(val string) (ModelDeploymentSystemDataSystemInfraTypeEnum, bool) {
	enum, ok := mappingModelDeploymentSystemDataSystemInfraTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
