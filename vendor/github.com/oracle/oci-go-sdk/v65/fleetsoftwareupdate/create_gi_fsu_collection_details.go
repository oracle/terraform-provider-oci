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

// CreateGiFsuCollectionDetails Details to create a 'GI' type Exadata Fleet Update Collection.
type CreateGiFsuCollectionDetails struct {

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Exadata Fleet Update Collection Identifier.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	FleetDiscovery GiFleetDiscoveryDetails `mandatory:"false" json:"fleetDiscovery"`

	// Exadata service type for the target resource members.
	ServiceType CollectionServiceTypesEnum `mandatory:"true" json:"serviceType"`

	// Grid Infrastructure Major Version of targets to be included in the Exadata Fleet Update Collection.
	// Only GI targets that match the version specified in this value would be added to the Exadata Fleet Update Collection.
	SourceMajorVersion GiSourceMajorVersionsEnum `mandatory:"true" json:"sourceMajorVersion"`
}

// GetDisplayName returns DisplayName
func (m CreateGiFsuCollectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetServiceType returns ServiceType
func (m CreateGiFsuCollectionDetails) GetServiceType() CollectionServiceTypesEnum {
	return m.ServiceType
}

// GetCompartmentId returns CompartmentId
func (m CreateGiFsuCollectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateGiFsuCollectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateGiFsuCollectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateGiFsuCollectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGiFsuCollectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCollectionServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetCollectionServiceTypesEnumStringValues(), ",")))
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
func (m CreateGiFsuCollectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateGiFsuCollectionDetails CreateGiFsuCollectionDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateGiFsuCollectionDetails
	}{
		"GI",
		(MarshalTypeCreateGiFsuCollectionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateGiFsuCollectionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                           `json:"displayName"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FleetDiscovery     gifleetdiscoverydetails           `json:"fleetDiscovery"`
		ServiceType        CollectionServiceTypesEnum        `json:"serviceType"`
		CompartmentId      *string                           `json:"compartmentId"`
		SourceMajorVersion GiSourceMajorVersionsEnum         `json:"sourceMajorVersion"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.FleetDiscovery.UnmarshalPolymorphicJSON(model.FleetDiscovery.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FleetDiscovery = nn.(GiFleetDiscoveryDetails)
	} else {
		m.FleetDiscovery = nil
	}

	m.ServiceType = model.ServiceType

	m.CompartmentId = model.CompartmentId

	m.SourceMajorVersion = model.SourceMajorVersion

	return
}
