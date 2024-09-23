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

// ExadataExaccConfigurationSummary Configuration summary of a macs managed Exacc exadata machine.
type ExadataExaccConfigurationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
	ExadataInsightId *string `mandatory:"true" json:"exadataInsightId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Exadata system name. If the Exadata systems managed by Enterprise Manager, the name is unique amongst the Exadata systems managed by the same Enterprise Manager.
	ExadataName *string `mandatory:"true" json:"exadataName"`

	// The user-friendly name for the Exadata system. The name does not have to be unique.
	ExadataDisplayName *string `mandatory:"true" json:"exadataDisplayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	ParentId *string `mandatory:"true" json:"parentId"`

	// Array of objects containing VM cluster information.
	VmclusterDetails []VmClusterSummary `mandatory:"false" json:"vmclusterDetails"`

	// Operations Insights internal representation of the the Exadata system type.
	ExadataType ExadataTypeEnum `mandatory:"true" json:"exadataType"`

	// Exadata rack type.
	ExadataRackType ExadataRackTypeEnum `mandatory:"true" json:"exadataRackType"`
}

// GetExadataInsightId returns ExadataInsightId
func (m ExadataExaccConfigurationSummary) GetExadataInsightId() *string {
	return m.ExadataInsightId
}

// GetCompartmentId returns CompartmentId
func (m ExadataExaccConfigurationSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetExadataName returns ExadataName
func (m ExadataExaccConfigurationSummary) GetExadataName() *string {
	return m.ExadataName
}

// GetExadataDisplayName returns ExadataDisplayName
func (m ExadataExaccConfigurationSummary) GetExadataDisplayName() *string {
	return m.ExadataDisplayName
}

// GetExadataType returns ExadataType
func (m ExadataExaccConfigurationSummary) GetExadataType() ExadataTypeEnum {
	return m.ExadataType
}

// GetExadataRackType returns ExadataRackType
func (m ExadataExaccConfigurationSummary) GetExadataRackType() ExadataRackTypeEnum {
	return m.ExadataRackType
}

// GetDefinedTags returns DefinedTags
func (m ExadataExaccConfigurationSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m ExadataExaccConfigurationSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetVmclusterDetails returns VmclusterDetails
func (m ExadataExaccConfigurationSummary) GetVmclusterDetails() []VmClusterSummary {
	return m.VmclusterDetails
}

func (m ExadataExaccConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataExaccConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataTypeEnum(string(m.ExadataType)); !ok && m.ExadataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataType: %s. Supported values are: %s.", m.ExadataType, strings.Join(GetExadataTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataRackTypeEnum(string(m.ExadataRackType)); !ok && m.ExadataRackType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataRackType: %s. Supported values are: %s.", m.ExadataRackType, strings.Join(GetExadataRackTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExadataExaccConfigurationSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadataExaccConfigurationSummary ExadataExaccConfigurationSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeExadataExaccConfigurationSummary
	}{
		"MACS_MANAGED_CLOUD_EXADATA",
		(MarshalTypeExadataExaccConfigurationSummary)(m),
	}

	return json.Marshal(&s)
}
