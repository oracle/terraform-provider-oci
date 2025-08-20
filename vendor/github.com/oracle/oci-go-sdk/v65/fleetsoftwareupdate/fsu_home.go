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

// FsuHome Exadata Fleet Update Home resource.
type FsuHome interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Image.
	GetFsuImageId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// resource.
	GetCompartmentId() *string

	// Unique name of the home.
	GetHomeName() *string

	// OCID identifier for the Exadata Fleet Update Home.
	GetId() *string

	// Exadata Fleet Update Home display name.
	GetDisplayName() *string

	// The time the Exadata Fleet Update Home was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the Exadata Fleet Update Home.
	GetLifecycleState() HomeLifecycleStatesEnum

	// ORACLE_BASE path for provisioning Oracle database home or Oracle Grid
	// Infrastructure home
	GetOracleBase() *string

	// Oracle groups configured for the Exadata Fleet Update Home.
	GetGroups() []FsuHomeGroup

	// User who provisioned the Exadata Fleet Update Home.
	GetUser() *string

	// The time the Exadata Fleet Update Home was updated. An RFC3339 formatted datetime string.
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

type fsuhome struct {
	JsonData         []byte
	OracleBase       *string                           `mandatory:"false" json:"oracleBase"`
	Groups           []FsuHomeGroup                    `mandatory:"false" json:"groups"`
	User             *string                           `mandatory:"false" json:"user"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	FsuImageId       *string                           `mandatory:"true" json:"fsuImageId"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	HomeName         *string                           `mandatory:"true" json:"homeName"`
	Id               *string                           `mandatory:"true" json:"id"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState   HomeLifecycleStatesEnum           `mandatory:"true" json:"lifecycleState"`
	Kind             string                            `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *fsuhome) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfsuhome fsuhome
	s := struct {
		Model Unmarshalerfsuhome
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FsuImageId = s.Model.FsuImageId
	m.CompartmentId = s.Model.CompartmentId
	m.HomeName = s.Model.HomeName
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.OracleBase = s.Model.OracleBase
	m.Groups = s.Model.Groups
	m.User = s.Model.User
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fsuhome) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "VMCLUSTER":
		mm := VmClusterFsuHome{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DBHOME":
		mm := DbHomeFsuHome{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for FsuHome: %s.", m.Kind)
		return *m, nil
	}
}

// GetOracleBase returns OracleBase
func (m fsuhome) GetOracleBase() *string {
	return m.OracleBase
}

// GetGroups returns Groups
func (m fsuhome) GetGroups() []FsuHomeGroup {
	return m.Groups
}

// GetUser returns User
func (m fsuhome) GetUser() *string {
	return m.User
}

// GetTimeUpdated returns TimeUpdated
func (m fsuhome) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m fsuhome) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m fsuhome) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m fsuhome) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m fsuhome) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetFsuImageId returns FsuImageId
func (m fsuhome) GetFsuImageId() *string {
	return m.FsuImageId
}

// GetCompartmentId returns CompartmentId
func (m fsuhome) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHomeName returns HomeName
func (m fsuhome) GetHomeName() *string {
	return m.HomeName
}

// GetId returns Id
func (m fsuhome) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m fsuhome) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m fsuhome) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m fsuhome) GetLifecycleState() HomeLifecycleStatesEnum {
	return m.LifecycleState
}

func (m fsuhome) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fsuhome) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHomeLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHomeLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
