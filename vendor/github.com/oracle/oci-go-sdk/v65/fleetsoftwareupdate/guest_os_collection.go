// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GuestOsCollection Details of a 'GUEST_OS' type Exadata Fleet Update Collection.
type GuestOsCollection struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Collection.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the Exadata Fleet Update Collection.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the Exadata Fleet Update Collection was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	FleetDiscovery GuestOsFleetDiscoveryDetails `mandatory:"true" json:"fleetDiscovery"`

	ActiveFsuCycle *ActiveCycleDetails `mandatory:"false" json:"activeFsuCycle"`

	// Number of targets that are members of this Collection.
	TargetCount *int `mandatory:"false" json:"targetCount"`

	// The time the Exadata Fleet Update Collection was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of last completed FSU Cycle.
	LastCompletedFsuCycleId *string `mandatory:"false" json:"lastCompletedFsuCycleId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Exadata service type for the target resource members.
	ServiceType CollectionServiceTypesEnum `mandatory:"true" json:"serviceType"`

	// The current state of the Exadata Fleet Update Collection.
	LifecycleState CollectionLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Major version of Exadata Image (Guest OS) release for Exadata VM Cluster targets to be included in the Exadata Fleet Update Collection.
	// Only Exadata VM Clusters whose 'systemVersion' is related to the major version will be added to the Exadata Fleet Update Collection.
	// For more details, refer to Oracle document 2075007.1 (https://support.oracle.com/knowledge/Oracle%20Database%20Products/2075007_1.html)
	SourceMajorVersion GuestOsSourceMajorVersionsEnum `mandatory:"true" json:"sourceMajorVersion"`
}

// GetId returns Id
func (m GuestOsCollection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m GuestOsCollection) GetDisplayName() *string {
	return m.DisplayName
}

// GetServiceType returns ServiceType
func (m GuestOsCollection) GetServiceType() CollectionServiceTypesEnum {
	return m.ServiceType
}

// GetCompartmentId returns CompartmentId
func (m GuestOsCollection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetActiveFsuCycle returns ActiveFsuCycle
func (m GuestOsCollection) GetActiveFsuCycle() *ActiveCycleDetails {
	return m.ActiveFsuCycle
}

// GetTargetCount returns TargetCount
func (m GuestOsCollection) GetTargetCount() *int {
	return m.TargetCount
}

// GetTimeCreated returns TimeCreated
func (m GuestOsCollection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m GuestOsCollection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m GuestOsCollection) GetLifecycleState() CollectionLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m GuestOsCollection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetLastCompletedFsuCycleId returns LastCompletedFsuCycleId
func (m GuestOsCollection) GetLastCompletedFsuCycleId() *string {
	return m.LastCompletedFsuCycleId
}

// GetFreeformTags returns FreeformTags
func (m GuestOsCollection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m GuestOsCollection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m GuestOsCollection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m GuestOsCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GuestOsCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCollectionServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetCollectionServiceTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCollectionLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCollectionLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGuestOsSourceMajorVersionsEnum(string(m.SourceMajorVersion)); !ok && m.SourceMajorVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceMajorVersion: %s. Supported values are: %s.", m.SourceMajorVersion, strings.Join(GetGuestOsSourceMajorVersionsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GuestOsCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGuestOsCollection GuestOsCollection
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGuestOsCollection
	}{
		"GUEST_OS",
		(MarshalTypeGuestOsCollection)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GuestOsCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ActiveFsuCycle          *ActiveCycleDetails               `json:"activeFsuCycle"`
		TargetCount             *int                              `json:"targetCount"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails        *string                           `json:"lifecycleDetails"`
		LastCompletedFsuCycleId *string                           `json:"lastCompletedFsuCycleId"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		SystemTags              map[string]map[string]interface{} `json:"systemTags"`
		Id                      *string                           `json:"id"`
		DisplayName             *string                           `json:"displayName"`
		ServiceType             CollectionServiceTypesEnum        `json:"serviceType"`
		CompartmentId           *string                           `json:"compartmentId"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		LifecycleState          CollectionLifecycleStatesEnum     `json:"lifecycleState"`
		SourceMajorVersion      GuestOsSourceMajorVersionsEnum    `json:"sourceMajorVersion"`
		FleetDiscovery          guestosfleetdiscoverydetails      `json:"fleetDiscovery"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ActiveFsuCycle = model.ActiveFsuCycle

	m.TargetCount = model.TargetCount

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.LastCompletedFsuCycleId = model.LastCompletedFsuCycleId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.ServiceType = model.ServiceType

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.SourceMajorVersion = model.SourceMajorVersion

	nn, e = model.FleetDiscovery.UnmarshalPolymorphicJSON(model.FleetDiscovery.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FleetDiscovery = nn.(GuestOsFleetDiscoveryDetails)
	} else {
		m.FleetDiscovery = nil
	}

	return
}
