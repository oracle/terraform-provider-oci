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

// FsuDiscovery Exadata Fleet Update Discovery resource details.
type FsuDiscovery struct {

	// OCID identifier for the Exadata Fleet Update Discovery.
	Id *string `mandatory:"true" json:"id"`

	// Exadata Fleet Update Discovery display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Details DiscoveryDetails `mandatory:"true" json:"details"`

	// The date and time the Exadata Fleet Update Discovery was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Exadata Fleet Update Discovery.
	LifecycleState DiscoveryLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Exadata Fleet Update Discovery was updated,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the Exadata Fleet Update Discovery was finished,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

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
}

func (m FsuDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FsuDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDiscoveryLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FsuDiscovery) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		TimeFinished     *common.SDKTime                   `json:"timeFinished"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		DisplayName      *string                           `json:"displayName"`
		CompartmentId    *string                           `json:"compartmentId"`
		Details          discoverydetails                  `json:"details"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		LifecycleState   DiscoveryLifecycleStatesEnum      `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeUpdated = model.TimeUpdated

	m.TimeFinished = model.TimeFinished

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	nn, e = model.Details.UnmarshalPolymorphicJSON(model.Details.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Details = nn.(DiscoveryDetails)
	} else {
		m.Details = nil
	}

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
