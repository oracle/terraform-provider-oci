// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateModelDeploymentDetails Details for updating a model deployment. You can update `modelDeploymentConfigurationDetails` and change `instanceShapeName` and `modelId` when the model deployment is in
// the ACTIVE lifecycle state. The `bandwidthMbps` or `instanceCount` can only be updated while the model deployment is in the `INACTIVE` state. Changes to the `bandwidthMbps`
// or `instanceCount` will take effect the next time the `ActivateModelDeployment` action is invoked on the model deployment resource.
type UpdateModelDeploymentDetails struct {

	// A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information.
	// Example: `My ModelDeployment`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the model deployment.
	Description *string `mandatory:"false" json:"description"`

	ModelDeploymentConfigurationDetails UpdateModelDeploymentConfigurationDetails `mandatory:"false" json:"modelDeploymentConfigurationDetails"`

	CategoryLogDetails *UpdateCategoryLogDetails `mandatory:"false" json:"categoryLogDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateModelDeploymentDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateModelDeploymentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                         *string                                   `json:"displayName"`
		Description                         *string                                   `json:"description"`
		ModelDeploymentConfigurationDetails updatemodeldeploymentconfigurationdetails `json:"modelDeploymentConfigurationDetails"`
		CategoryLogDetails                  *UpdateCategoryLogDetails                 `json:"categoryLogDetails"`
		FreeformTags                        map[string]string                         `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{}         `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.ModelDeploymentConfigurationDetails.UnmarshalPolymorphicJSON(model.ModelDeploymentConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelDeploymentConfigurationDetails = nn.(UpdateModelDeploymentConfigurationDetails)
	} else {
		m.ModelDeploymentConfigurationDetails = nil
	}

	m.CategoryLogDetails = model.CategoryLogDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
