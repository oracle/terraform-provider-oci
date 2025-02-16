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

// CreateMlApplicationInstanceDetails The information about new MlApplicationInstance.
type CreateMlApplicationInstanceDetails struct {

	// The OCID of ML Application. This resource is an instance of ML Application referenced by this OCID.
	MlApplicationId *string `mandatory:"true" json:"mlApplicationId"`

	// The OCID of ML Application Implementation selected as a certain solution for a given ML problem (ML Application)
	MlApplicationImplementationId *string `mandatory:"true" json:"mlApplicationImplementationId"`

	// The OCID of the compartment where the MlApplicationInstance is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of MlApplicationInstance. System will generate displayName when not provided.
	DisplayName *string `mandatory:"false" json:"displayName"`

	AuthConfiguration CreateAuthConfigurationDetails `mandatory:"false" json:"authConfiguration"`

	// Data that are used for provisioning of the given MlApplicationInstance. These are validated against configurationSchema defined in referenced MlApplicationImplementation.
	Configuration []ConfigurationProperty `mandatory:"false" json:"configuration"`

	// Defines whether the MlApplicationInstance will be created in ACTIVE (true value) or INACTIVE (false value) lifecycle state.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMlApplicationInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMlApplicationInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateMlApplicationInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                   *string                           `json:"displayName"`
		AuthConfiguration             createauthconfigurationdetails    `json:"authConfiguration"`
		Configuration                 []ConfigurationProperty           `json:"configuration"`
		IsEnabled                     *bool                             `json:"isEnabled"`
		FreeformTags                  map[string]string                 `json:"freeformTags"`
		DefinedTags                   map[string]map[string]interface{} `json:"definedTags"`
		MlApplicationId               *string                           `json:"mlApplicationId"`
		MlApplicationImplementationId *string                           `json:"mlApplicationImplementationId"`
		CompartmentId                 *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.AuthConfiguration.UnmarshalPolymorphicJSON(model.AuthConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthConfiguration = nn.(CreateAuthConfigurationDetails)
	} else {
		m.AuthConfiguration = nil
	}

	m.Configuration = make([]ConfigurationProperty, len(model.Configuration))
	copy(m.Configuration, model.Configuration)
	m.IsEnabled = model.IsEnabled

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.MlApplicationId = model.MlApplicationId

	m.MlApplicationImplementationId = model.MlApplicationImplementationId

	m.CompartmentId = model.CompartmentId

	return
}
