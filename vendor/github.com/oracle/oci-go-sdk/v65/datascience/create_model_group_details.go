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

// CreateModelGroupDetails Parameters needed to create a model group. Model Group is a group of models.
type CreateModelGroupDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the modelGroup in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the modelGroup.
	ProjectId *string `mandatory:"true" json:"projectId"`

	ModelGroupDetails ModelGroupDetails `mandatory:"true" json:"modelGroupDetails"`

	MemberModelEntries *MemberModelEntries `mandatory:"true" json:"memberModelEntries"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	// Example: `My ModelGroup`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the modelGroup.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model group version history to which the modelGroup is associated.
	ModelGroupVersionHistoryId *string `mandatory:"false" json:"modelGroupVersionHistoryId"`

	// An additional description of the lifecycle state of the model group.
	VersionLabel *string `mandatory:"false" json:"versionLabel"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetCompartmentId returns CompartmentId
func (m CreateModelGroupDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetProjectId returns ProjectId
func (m CreateModelGroupDetails) GetProjectId() *string {
	return m.ProjectId
}

func (m CreateModelGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateModelGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateModelGroupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateModelGroupDetails CreateModelGroupDetails
	s := struct {
		DiscriminatorParam string `json:"createType"`
		MarshalTypeCreateModelGroupDetails
	}{
		"CREATE",
		(MarshalTypeCreateModelGroupDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateModelGroupDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                *string                           `json:"displayName"`
		Description                *string                           `json:"description"`
		ModelGroupVersionHistoryId *string                           `json:"modelGroupVersionHistoryId"`
		VersionLabel               *string                           `json:"versionLabel"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId              *string                           `json:"compartmentId"`
		ProjectId                  *string                           `json:"projectId"`
		ModelGroupDetails          modelgroupdetails                 `json:"modelGroupDetails"`
		MemberModelEntries         *MemberModelEntries               `json:"memberModelEntries"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.ModelGroupVersionHistoryId = model.ModelGroupVersionHistoryId

	m.VersionLabel = model.VersionLabel

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	nn, e = model.ModelGroupDetails.UnmarshalPolymorphicJSON(model.ModelGroupDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelGroupDetails = nn.(ModelGroupDetails)
	} else {
		m.ModelGroupDetails = nil
	}

	m.MemberModelEntries = model.MemberModelEntries

	return
}
