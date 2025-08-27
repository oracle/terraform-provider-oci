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

// CloneModelGroupDetails Parameters needed to clone a model group.
type CloneModelGroupDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the modelGroup in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the modelGroup.
	ProjectId *string `mandatory:"true" json:"projectId"`

	ModelGroupCloneSourceDetails ModelGroupCloneSourceDetails `mandatory:"true" json:"modelGroupCloneSourceDetails"`
}

// GetCompartmentId returns CompartmentId
func (m CloneModelGroupDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetProjectId returns ProjectId
func (m CloneModelGroupDetails) GetProjectId() *string {
	return m.ProjectId
}

func (m CloneModelGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloneModelGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloneModelGroupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloneModelGroupDetails CloneModelGroupDetails
	s := struct {
		DiscriminatorParam string `json:"createType"`
		MarshalTypeCloneModelGroupDetails
	}{
		"CLONE",
		(MarshalTypeCloneModelGroupDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CloneModelGroupDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId                *string                      `json:"compartmentId"`
		ProjectId                    *string                      `json:"projectId"`
		ModelGroupCloneSourceDetails modelgroupclonesourcedetails `json:"modelGroupCloneSourceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	nn, e = model.ModelGroupCloneSourceDetails.UnmarshalPolymorphicJSON(model.ModelGroupCloneSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelGroupCloneSourceDetails = nn.(ModelGroupCloneSourceDetails)
	} else {
		m.ModelGroupCloneSourceDetails = nil
	}

	return
}
