// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOpsiConfigurationDetails Information about OPSI configuration to be created.
type CreateOpsiConfigurationDetails interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// User-friendly display name for the OPSI configuration. The name does not have to be unique.
	GetDisplayName() *string

	// Description of OPSI configuration.
	GetDescription() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// Array of configuration items with custom values. All and only configuration items requiring custom values should be part of this array.
	GetConfigItems() []CreateConfigurationItemDetails
}

type createopsiconfigurationdetails struct {
	JsonData       []byte
	CompartmentId  *string                           `mandatory:"false" json:"compartmentId"`
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	Description    *string                           `mandatory:"false" json:"description"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags     map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	ConfigItems    json.RawMessage                   `mandatory:"false" json:"configItems"`
	OpsiConfigType string                            `json:"opsiConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *createopsiconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateopsiconfigurationdetails createopsiconfigurationdetails
	s := struct {
		Model Unmarshalercreateopsiconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ConfigItems = s.Model.ConfigItems
	m.OpsiConfigType = s.Model.OpsiConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createopsiconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OpsiConfigType {
	case "UX_CONFIGURATION":
		mm := CreateOpsiUxConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateOpsiConfigurationDetails: %s.", m.OpsiConfigType)
		return *m, nil
	}
}

// GetCompartmentId returns CompartmentId
func (m createopsiconfigurationdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m createopsiconfigurationdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m createopsiconfigurationdetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m createopsiconfigurationdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createopsiconfigurationdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m createopsiconfigurationdetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetConfigItems returns ConfigItems
func (m createopsiconfigurationdetails) GetConfigItems() json.RawMessage {
	return m.ConfigItems
}

func (m createopsiconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createopsiconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
