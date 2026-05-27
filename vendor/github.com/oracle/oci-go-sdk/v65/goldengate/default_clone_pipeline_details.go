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

// DefaultClonePipelineDetails Attribute details for a default pipeline clone.
type DefaultClonePipelineDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline. This option applies when retrieving a pipeline.
	SourcePipelineId *string `mandatory:"true" json:"sourcePipelineId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

// GetDisplayName returns DisplayName
func (m DefaultClonePipelineDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetSourcePipelineId returns SourcePipelineId
func (m DefaultClonePipelineDetails) GetSourcePipelineId() *string {
	return m.SourcePipelineId
}

// GetCompartmentId returns CompartmentId
func (m DefaultClonePipelineDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDescription returns Description
func (m DefaultClonePipelineDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m DefaultClonePipelineDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m DefaultClonePipelineDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m DefaultClonePipelineDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DefaultClonePipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultClonePipelineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DefaultClonePipelineDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultClonePipelineDetails DefaultClonePipelineDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDefaultClonePipelineDetails
	}{
		"DEFAULT",
		(MarshalTypeDefaultClonePipelineDetails)(m),
	}

	return json.Marshal(&s)
}
