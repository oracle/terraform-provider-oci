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

// UpdateModelGroupDeploymentConfigurationDetails The model group type deployment for update.
type UpdateModelGroupDeploymentConfigurationDetails struct {
	ModelGroupConfigurationDetails *UpdateModelGroupConfigurationDetails `mandatory:"false" json:"modelGroupConfigurationDetails"`

	InfrastructureConfigurationDetails UpdateInfrastructureConfigurationDetails `mandatory:"false" json:"infrastructureConfigurationDetails"`

	EnvironmentConfigurationDetails UpdateModelDeploymentEnvironmentConfigurationDetails `mandatory:"false" json:"environmentConfigurationDetails"`

	// The type of update operation.
	UpdateType UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum `mandatory:"false" json:"updateType,omitempty"`
}

func (m UpdateModelGroupDeploymentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateModelGroupDeploymentConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateModelGroupDeploymentConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateModelGroupDeploymentConfigurationDetails UpdateModelGroupDeploymentConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"deploymentType"`
		MarshalTypeUpdateModelGroupDeploymentConfigurationDetails
	}{
		"MODEL_GROUP",
		(MarshalTypeUpdateModelGroupDeploymentConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateModelGroupDeploymentConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		UpdateType                         UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum `json:"updateType"`
		ModelGroupConfigurationDetails     *UpdateModelGroupConfigurationDetails                        `json:"modelGroupConfigurationDetails"`
		InfrastructureConfigurationDetails updateinfrastructureconfigurationdetails                     `json:"infrastructureConfigurationDetails"`
		EnvironmentConfigurationDetails    updatemodeldeploymentenvironmentconfigurationdetails         `json:"environmentConfigurationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.UpdateType = model.UpdateType

	m.ModelGroupConfigurationDetails = model.ModelGroupConfigurationDetails

	nn, e = model.InfrastructureConfigurationDetails.UnmarshalPolymorphicJSON(model.InfrastructureConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InfrastructureConfigurationDetails = nn.(UpdateInfrastructureConfigurationDetails)
	} else {
		m.InfrastructureConfigurationDetails = nil
	}

	nn, e = model.EnvironmentConfigurationDetails.UnmarshalPolymorphicJSON(model.EnvironmentConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EnvironmentConfigurationDetails = nn.(UpdateModelDeploymentEnvironmentConfigurationDetails)
	} else {
		m.EnvironmentConfigurationDetails = nil
	}

	return
}

// UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum Enum with underlying type: string
type UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum
const (
	UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeZdt  UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum = "ZDT"
	UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeLive UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum = "LIVE"
)

var mappingUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum = map[string]UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum{
	"ZDT":  UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeZdt,
	"LIVE": UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeLive,
}

var mappingUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnumLowerCase = map[string]UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum{
	"zdt":  UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeZdt,
	"live": UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeLive,
}

// GetUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnumValues Enumerates the set of values for UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum
func GetUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnumValues() []UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum {
	values := make([]UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum, 0)
	for _, v := range mappingUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnumStringValues Enumerates the set of values in String for UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum
func GetUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnumStringValues() []string {
	return []string{
		"ZDT",
		"LIVE",
	}
}

// GetMappingUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum(val string) (UpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnum, bool) {
	enum, ok := mappingUpdateModelGroupDeploymentConfigurationDetailsUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
