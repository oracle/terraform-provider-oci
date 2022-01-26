// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ExadataConfigurationSummary Summary of a exadata configuration for a resource.
type ExadataConfigurationSummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
	GetExadataInsightId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The Exadata system name. If the Exadata systems managed by Enterprise Manager, the name is unique amongst the Exadata systems managed by the same Enterprise Manager.
	GetExadataName() *string

	// The user-friendly name for the Exadata system. The name does not have to be unique.
	GetExadataDisplayName() *string

	// Operations Insights internal representation of the the Exadata system type.
	GetExadataType() ExadataTypeEnum

	// Exadata rack type.
	GetExadataRackType() ExadataRackTypeEnum

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string
}

type exadataconfigurationsummary struct {
	JsonData           []byte
	ExadataInsightId   *string                           `mandatory:"true" json:"exadataInsightId"`
	CompartmentId      *string                           `mandatory:"true" json:"compartmentId"`
	ExadataName        *string                           `mandatory:"true" json:"exadataName"`
	ExadataDisplayName *string                           `mandatory:"true" json:"exadataDisplayName"`
	ExadataType        ExadataTypeEnum                   `mandatory:"true" json:"exadataType"`
	ExadataRackType    ExadataRackTypeEnum               `mandatory:"true" json:"exadataRackType"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	FreeformTags       map[string]string                 `mandatory:"true" json:"freeformTags"`
	EntitySource       string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *exadataconfigurationsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexadataconfigurationsummary exadataconfigurationsummary
	s := struct {
		Model Unmarshalerexadataconfigurationsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ExadataInsightId = s.Model.ExadataInsightId
	m.CompartmentId = s.Model.CompartmentId
	m.ExadataName = s.Model.ExadataName
	m.ExadataDisplayName = s.Model.ExadataDisplayName
	m.ExadataType = s.Model.ExadataType
	m.ExadataRackType = s.Model.ExadataRackType
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *exadataconfigurationsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_EXADATA":
		mm := ExadataDatabaseMachineConfigurationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetExadataInsightId returns ExadataInsightId
func (m exadataconfigurationsummary) GetExadataInsightId() *string {
	return m.ExadataInsightId
}

//GetCompartmentId returns CompartmentId
func (m exadataconfigurationsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetExadataName returns ExadataName
func (m exadataconfigurationsummary) GetExadataName() *string {
	return m.ExadataName
}

//GetExadataDisplayName returns ExadataDisplayName
func (m exadataconfigurationsummary) GetExadataDisplayName() *string {
	return m.ExadataDisplayName
}

//GetExadataType returns ExadataType
func (m exadataconfigurationsummary) GetExadataType() ExadataTypeEnum {
	return m.ExadataType
}

//GetExadataRackType returns ExadataRackType
func (m exadataconfigurationsummary) GetExadataRackType() ExadataRackTypeEnum {
	return m.ExadataRackType
}

//GetDefinedTags returns DefinedTags
func (m exadataconfigurationsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m exadataconfigurationsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m exadataconfigurationsummary) String() string {
	return common.PointerString(m)
}
