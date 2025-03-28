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

// UpdateJobDetails Details for updating a job.
type UpdateJobDetails struct {

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	JobInfrastructureConfigurationDetails JobInfrastructureConfigurationDetails `mandatory:"false" json:"jobInfrastructureConfigurationDetails"`

	// Collection of JobStorageMountConfigurationDetails.
	JobStorageMountConfigurationDetailsList []StorageMountConfigurationDetails `mandatory:"false" json:"jobStorageMountConfigurationDetailsList"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                             *string                               `json:"displayName"`
		Description                             *string                               `json:"description"`
		JobInfrastructureConfigurationDetails   jobinfrastructureconfigurationdetails `json:"jobInfrastructureConfigurationDetails"`
		JobStorageMountConfigurationDetailsList []storagemountconfigurationdetails    `json:"jobStorageMountConfigurationDetailsList"`
		FreeformTags                            map[string]string                     `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}     `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.JobInfrastructureConfigurationDetails.UnmarshalPolymorphicJSON(model.JobInfrastructureConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobInfrastructureConfigurationDetails = nn.(JobInfrastructureConfigurationDetails)
	} else {
		m.JobInfrastructureConfigurationDetails = nil
	}

	m.JobStorageMountConfigurationDetailsList = make([]StorageMountConfigurationDetails, len(model.JobStorageMountConfigurationDetailsList))
	for i, n := range model.JobStorageMountConfigurationDetailsList {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.JobStorageMountConfigurationDetailsList[i] = nn.(StorageMountConfigurationDetails)
		} else {
			m.JobStorageMountConfigurationDetailsList[i] = nil
		}
	}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
