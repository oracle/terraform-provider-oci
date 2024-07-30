// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FsuCollectionSummary Exadata Fleet Update Collection Resource.
type FsuCollectionSummary interface {

	// OCID identifier for the Exadata Fleet Update Collection.
	GetId() *string

	// Exadata Fleet Update Collection resource display name.
	GetDisplayName() *string

	// Exadata service type for the target resource members.
	GetServiceType() CollectionServiceTypesEnum

	// Compartment Identifier
	GetCompartmentId() *string

	// The time the Exadata Fleet Update Collection was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the Exadata Fleet Update Collection.
	GetLifecycleState() CollectionLifecycleStatesEnum

	GetActiveFsuCycle() *ActiveCycleDetails

	// Number of targets that are members of this Collection.
	GetTargetCount() *int

	// The time the Exadata Fleet Update Collection was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type fsucollectionsummary struct {
	JsonData         []byte
	ActiveFsuCycle   *ActiveCycleDetails               `mandatory:"false" json:"activeFsuCycle"`
	TargetCount      *int                              `mandatory:"false" json:"targetCount"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id               *string                           `mandatory:"true" json:"id"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	ServiceType      CollectionServiceTypesEnum        `mandatory:"true" json:"serviceType"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState   CollectionLifecycleStatesEnum     `mandatory:"true" json:"lifecycleState"`
	Type             string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fsucollectionsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfsucollectionsummary fsucollectionsummary
	s := struct {
		Model Unmarshalerfsucollectionsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.ServiceType = s.Model.ServiceType
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.ActiveFsuCycle = s.Model.ActiveFsuCycle
	m.TargetCount = s.Model.TargetCount
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fsucollectionsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "GI":
		mm := GiFsuCollectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB":
		mm := DbFsuCollectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FsuCollectionSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetActiveFsuCycle returns ActiveFsuCycle
func (m fsucollectionsummary) GetActiveFsuCycle() *ActiveCycleDetails {
	return m.ActiveFsuCycle
}

// GetTargetCount returns TargetCount
func (m fsucollectionsummary) GetTargetCount() *int {
	return m.TargetCount
}

// GetTimeUpdated returns TimeUpdated
func (m fsucollectionsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m fsucollectionsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m fsucollectionsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m fsucollectionsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m fsucollectionsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m fsucollectionsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m fsucollectionsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetServiceType returns ServiceType
func (m fsucollectionsummary) GetServiceType() CollectionServiceTypesEnum {
	return m.ServiceType
}

// GetCompartmentId returns CompartmentId
func (m fsucollectionsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m fsucollectionsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m fsucollectionsummary) GetLifecycleState() CollectionLifecycleStatesEnum {
	return m.LifecycleState
}

func (m fsucollectionsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fsucollectionsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCollectionServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetCollectionServiceTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCollectionLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCollectionLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
