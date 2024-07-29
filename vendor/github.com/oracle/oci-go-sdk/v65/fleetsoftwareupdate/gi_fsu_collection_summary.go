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

// GiFsuCollectionSummary 'GI' type Exadata Fleet Update Collection summary.
type GiFsuCollectionSummary struct {

	// OCID identifier for the Exadata Fleet Update Collection.
	Id *string `mandatory:"true" json:"id"`

	// Exadata Fleet Update Collection resource display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the Exadata Fleet Update Collection was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	ActiveFsuCycle *ActiveCycleDetails `mandatory:"false" json:"activeFsuCycle"`

	// Number of targets that are members of this Collection.
	TargetCount *int `mandatory:"false" json:"targetCount"`

	// The time the Exadata Fleet Update Collection was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

	// Grid Infrastructure Major Version of targets to be included in the Exadata Fleet Update Collection.
	// Only GI targets that match the version specified in this value would be added to the Exadata Fleet Update Collection.
	SourceMajorVersion GiSourceMajorVersionsEnum `mandatory:"true" json:"sourceMajorVersion"`
}

// GetId returns Id
func (m GiFsuCollectionSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m GiFsuCollectionSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetServiceType returns ServiceType
func (m GiFsuCollectionSummary) GetServiceType() CollectionServiceTypesEnum {
	return m.ServiceType
}

// GetCompartmentId returns CompartmentId
func (m GiFsuCollectionSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetActiveFsuCycle returns ActiveFsuCycle
func (m GiFsuCollectionSummary) GetActiveFsuCycle() *ActiveCycleDetails {
	return m.ActiveFsuCycle
}

// GetTargetCount returns TargetCount
func (m GiFsuCollectionSummary) GetTargetCount() *int {
	return m.TargetCount
}

// GetTimeCreated returns TimeCreated
func (m GiFsuCollectionSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m GiFsuCollectionSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m GiFsuCollectionSummary) GetLifecycleState() CollectionLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m GiFsuCollectionSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m GiFsuCollectionSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m GiFsuCollectionSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m GiFsuCollectionSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m GiFsuCollectionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiFsuCollectionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCollectionServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetCollectionServiceTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCollectionLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCollectionLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGiSourceMajorVersionsEnum(string(m.SourceMajorVersion)); !ok && m.SourceMajorVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceMajorVersion: %s. Supported values are: %s.", m.SourceMajorVersion, strings.Join(GetGiSourceMajorVersionsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GiFsuCollectionSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGiFsuCollectionSummary GiFsuCollectionSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGiFsuCollectionSummary
	}{
		"GI",
		(MarshalTypeGiFsuCollectionSummary)(m),
	}

	return json.Marshal(&s)
}
