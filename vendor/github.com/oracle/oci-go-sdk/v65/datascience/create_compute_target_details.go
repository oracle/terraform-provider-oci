// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateComputeTargetDetails Parameters needed to create a new compute target.
type CreateComputeTargetDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the compute target.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ComputeConfigurationDetails ComputeConfigurationDetails `mandatory:"true" json:"computeConfigurationDetails"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the compute target.
	Description *string `mandatory:"false" json:"description"`

	// Metadata for the compute target.
	// The size of metadata must be less than 2048 bytes.
	// Key should be under 32 characters.
	// Key should contain only letters, digits and underscore (_)
	// Key should start with a letter.
	// Key should have at least 2 characters.
	// Key should not end with underscore eg. `TEST_`
	// Key if added cannot be empty. Value can be empty.
	// No specific size limits on individual Values. But overall metadata is limited to 2048 bytes.
	// Key can't be reserved Compute Target metadata.
	Metadata map[string]string `mandatory:"false" json:"metadata"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateComputeTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateComputeTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateComputeTargetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                 *string                           `json:"displayName"`
		Description                 *string                           `json:"description"`
		Metadata                    map[string]string                 `json:"metadata"`
		FreeformTags                map[string]string                 `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId               *string                           `json:"compartmentId"`
		ComputeConfigurationDetails computeconfigurationdetails       `json:"computeConfigurationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.Metadata = model.Metadata

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	nn, e = model.ComputeConfigurationDetails.UnmarshalPolymorphicJSON(model.ComputeConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ComputeConfigurationDetails = nn.(ComputeConfigurationDetails)
	} else {
		m.ComputeConfigurationDetails = nil
	}

	return
}
