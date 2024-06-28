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

// ExadataInsightSummary Summary of an Exadata insight resource.
type ExadataInsightSummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight resource.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The Exadata system name. If the Exadata systems managed by Enterprise Manager, the name is unique amongst the Exadata systems managed by the same Enterprise Manager.
	GetExadataName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Indicates the status of an Exadata insight in Operations Insights
	GetStatus() ResourceStatusEnum

	// The time the the Exadata insight was first enabled. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The current state of the Exadata insight.
	GetLifecycleState() ExadataInsightLifecycleStateEnum

	// The user-friendly name for the Exadata system. The name does not have to be unique.
	GetExadataDisplayName() *string

	// Operations Insights internal representation of the the Exadata system type.
	GetExadataType() ExadataTypeEnum

	// Operations Insights internal representation of the the Exadata system rack type.
	GetExadataRackType() ExadataRackTypeEnum

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the Exadata insight was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string
}

type exadatainsightsummary struct {
	JsonData           []byte
	ExadataDisplayName *string                           `mandatory:"false" json:"exadataDisplayName"`
	ExadataType        ExadataTypeEnum                   `mandatory:"false" json:"exadataType,omitempty"`
	ExadataRackType    ExadataRackTypeEnum               `mandatory:"false" json:"exadataRackType,omitempty"`
	SystemTags         map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	TimeUpdated        *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails   *string                           `mandatory:"false" json:"lifecycleDetails"`
	Id                 *string                           `mandatory:"true" json:"id"`
	CompartmentId      *string                           `mandatory:"true" json:"compartmentId"`
	ExadataName        *string                           `mandatory:"true" json:"exadataName"`
	FreeformTags       map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	Status             ResourceStatusEnum                `mandatory:"true" json:"status"`
	TimeCreated        *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState     ExadataInsightLifecycleStateEnum  `mandatory:"true" json:"lifecycleState"`
	EntitySource       string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *exadatainsightsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexadatainsightsummary exadatainsightsummary
	s := struct {
		Model Unmarshalerexadatainsightsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.ExadataName = s.Model.ExadataName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Status = s.Model.Status
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.ExadataDisplayName = s.Model.ExadataDisplayName
	m.ExadataType = s.Model.ExadataType
	m.ExadataRackType = s.Model.ExadataRackType
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *exadatainsightsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_EXADATA":
		mm := EmManagedExternalExadataInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PE_COMANAGED_EXADATA":
		mm := PeComanagedExadataInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ExadataInsightSummary: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetExadataDisplayName returns ExadataDisplayName
func (m exadatainsightsummary) GetExadataDisplayName() *string {
	return m.ExadataDisplayName
}

// GetExadataType returns ExadataType
func (m exadatainsightsummary) GetExadataType() ExadataTypeEnum {
	return m.ExadataType
}

// GetExadataRackType returns ExadataRackType
func (m exadatainsightsummary) GetExadataRackType() ExadataRackTypeEnum {
	return m.ExadataRackType
}

// GetSystemTags returns SystemTags
func (m exadatainsightsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeUpdated returns TimeUpdated
func (m exadatainsightsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m exadatainsightsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m exadatainsightsummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m exadatainsightsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetExadataName returns ExadataName
func (m exadatainsightsummary) GetExadataName() *string {
	return m.ExadataName
}

// GetFreeformTags returns FreeformTags
func (m exadatainsightsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m exadatainsightsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetStatus returns Status
func (m exadatainsightsummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m exadatainsightsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m exadatainsightsummary) GetLifecycleState() ExadataInsightLifecycleStateEnum {
	return m.LifecycleState
}

func (m exadatainsightsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m exadatainsightsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetResourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInsightLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadataInsightLifecycleStateEnumStringValues(), ",")))
	}

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
