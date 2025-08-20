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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FsuHomeSummary Exadata Fleet Update Home resource.
type FsuHomeSummary struct {

	// OCID identifier for the Exadata Fleet Update Home.
	Id *string `mandatory:"true" json:"id"`

	// Type of Exadata Fleet Update Home to create.
	Kind FsuHomeTypesEnum `mandatory:"true" json:"kind"`

	// Exadata Fleet Update Home display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the Exadata Fleet Update Home was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Home.
	LifecycleState HomeLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The time the Exadata Fleet Update Home was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
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

func (m FsuHomeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FsuHomeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFsuHomeTypesEnum(string(m.Kind)); !ok && m.Kind != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Kind: %s. Supported values are: %s.", m.Kind, strings.Join(GetFsuHomeTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHomeLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHomeLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
