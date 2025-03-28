// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOpsiConfigurationDetails Information to be updated in OPSI configuration resource.
type UpdateOpsiConfigurationDetails interface {

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
	// This array overwrites the existing custom configuration items array for this resource.
	GetConfigItems() []UpdateConfigurationItemDetails
}

type updateopsiconfigurationdetails struct {
	JsonData       []byte
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	Description    *string                           `mandatory:"false" json:"description"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags     map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	ConfigItems    json.RawMessage                   `mandatory:"false" json:"configItems"`
	OpsiConfigType string                            `json:"opsiConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *updateopsiconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateopsiconfigurationdetails updateopsiconfigurationdetails
	s := struct {
		Model Unmarshalerupdateopsiconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
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
func (m *updateopsiconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OpsiConfigType {
	case "UX_CONFIGURATION":
		mm := UpdateOpsiUxConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateOpsiConfigurationDetails: %s.", m.OpsiConfigType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updateopsiconfigurationdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m updateopsiconfigurationdetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m updateopsiconfigurationdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updateopsiconfigurationdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m updateopsiconfigurationdetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetConfigItems returns ConfigItems
func (m updateopsiconfigurationdetails) GetConfigItems() json.RawMessage {
	return m.ConfigItems
}

func (m updateopsiconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateopsiconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
