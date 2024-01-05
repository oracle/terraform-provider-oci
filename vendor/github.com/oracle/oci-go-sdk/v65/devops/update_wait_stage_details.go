// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateWaitStageDetails Specifies the Wait stage. You can specify variable wait times or an absolute duration.
type UpdateWaitStageDetails struct {

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the build stage.
	Description *string `mandatory:"false" json:"description"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	WaitCriteria UpdateWaitCriteriaDetails `mandatory:"false" json:"waitCriteria"`
}

// GetDisplayName returns DisplayName
func (m UpdateWaitStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateWaitStageDetails) GetDescription() *string {
	return m.Description
}

// GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m UpdateWaitStageDetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m UpdateWaitStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateWaitStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateWaitStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateWaitStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateWaitStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateWaitStageDetails UpdateWaitStageDetails
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeUpdateWaitStageDetails
	}{
		"WAIT",
		(MarshalTypeUpdateWaitStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateWaitStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                             *string                                  `json:"displayName"`
		Description                             *string                                  `json:"description"`
		BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `json:"buildPipelineStagePredecessorCollection"`
		FreeformTags                            map[string]string                        `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}        `json:"definedTags"`
		WaitCriteria                            updatewaitcriteriadetails                `json:"waitCriteria"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.BuildPipelineStagePredecessorCollection = model.BuildPipelineStagePredecessorCollection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.WaitCriteria.UnmarshalPolymorphicJSON(model.WaitCriteria.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.WaitCriteria = nn.(UpdateWaitCriteriaDetails)
	} else {
		m.WaitCriteria = nil
	}

	return
}
