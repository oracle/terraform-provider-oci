// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ClonePipelineDetails Details for a pipeline clone.
type ClonePipelineDetails interface {

	// An object's Display Name.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline. This option applies when retrieving a pipeline.
	GetSourcePipelineId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	GetCompartmentId() *string

	// Metadata about this specific object.
	GetDescription() *string

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	GetSystemTags() map[string]map[string]interface{}
}

type clonepipelinedetails struct {
	JsonData         []byte
	Description      *string                           `mandatory:"false" json:"description"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	SourcePipelineId *string                           `mandatory:"true" json:"sourcePipelineId"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	Type             string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *clonepipelinedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerclonepipelinedetails clonepipelinedetails
	s := struct {
		Model Unmarshalerclonepipelinedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.SourcePipelineId = s.Model.SourcePipelineId
	m.CompartmentId = s.Model.CompartmentId
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *clonepipelinedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DefaultClonePipelineDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ClonePipelineDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m clonepipelinedetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m clonepipelinedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m clonepipelinedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m clonepipelinedetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetDisplayName returns DisplayName
func (m clonepipelinedetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetSourcePipelineId returns SourcePipelineId
func (m clonepipelinedetails) GetSourcePipelineId() *string {
	return m.SourcePipelineId
}

// GetCompartmentId returns CompartmentId
func (m clonepipelinedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m clonepipelinedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m clonepipelinedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
