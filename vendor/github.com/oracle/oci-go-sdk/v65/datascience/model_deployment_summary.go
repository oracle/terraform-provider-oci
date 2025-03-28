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

// ModelDeploymentSummary Summary information for a model deployment.
type ModelDeploymentSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2019-08-25T21:10:29.41Z
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information.
	// Example: `My ModelDeployment`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model deployment.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model deployment.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The URL to interact with the model deployment.
	ModelDeploymentUrl *string `mandatory:"true" json:"modelDeploymentUrl"`

	// The state of the model deployment.
	LifecycleState ModelDeploymentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A short description of the model deployment.
	Description *string `mandatory:"false" json:"description"`

	ModelDeploymentConfigurationDetails ModelDeploymentConfigurationDetails `mandatory:"false" json:"modelDeploymentConfigurationDetails"`

	CategoryLogDetails *CategoryLogDetails `mandatory:"false" json:"categoryLogDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ModelDeploymentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelDeploymentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelDeploymentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetModelDeploymentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ModelDeploymentSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                         *string                             `json:"description"`
		ModelDeploymentConfigurationDetails modeldeploymentconfigurationdetails `json:"modelDeploymentConfigurationDetails"`
		CategoryLogDetails                  *CategoryLogDetails                 `json:"categoryLogDetails"`
		FreeformTags                        map[string]string                   `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{}   `json:"definedTags"`
		Id                                  *string                             `json:"id"`
		TimeCreated                         *common.SDKTime                     `json:"timeCreated"`
		DisplayName                         *string                             `json:"displayName"`
		ProjectId                           *string                             `json:"projectId"`
		CreatedBy                           *string                             `json:"createdBy"`
		CompartmentId                       *string                             `json:"compartmentId"`
		ModelDeploymentUrl                  *string                             `json:"modelDeploymentUrl"`
		LifecycleState                      ModelDeploymentLifecycleStateEnum   `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	nn, e = model.ModelDeploymentConfigurationDetails.UnmarshalPolymorphicJSON(model.ModelDeploymentConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelDeploymentConfigurationDetails = nn.(ModelDeploymentConfigurationDetails)
	} else {
		m.ModelDeploymentConfigurationDetails = nil
	}

	m.CategoryLogDetails = model.CategoryLogDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.TimeCreated = model.TimeCreated

	m.DisplayName = model.DisplayName

	m.ProjectId = model.ProjectId

	m.CreatedBy = model.CreatedBy

	m.CompartmentId = model.CompartmentId

	m.ModelDeploymentUrl = model.ModelDeploymentUrl

	m.LifecycleState = model.LifecycleState

	return
}
