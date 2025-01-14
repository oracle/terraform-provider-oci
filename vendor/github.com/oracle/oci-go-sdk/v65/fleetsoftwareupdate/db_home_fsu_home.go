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

// DbHomeFsuHome Exadata Fleet Update Home using a VmCluster or CloudVmCluster resource as source.
type DbHomeFsuHome struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Image.
	FsuImageId *string `mandatory:"true" json:"fsuImageId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique name of the home.
	HomeName *string `mandatory:"true" json:"homeName"`

	// OCID identifier for the Exadata Fleet Update Home.
	Id *string `mandatory:"true" json:"id"`

	// Exadata Fleet Update Home display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the Exadata Fleet Update Home was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"true" json:"dbHomeId"`

	// ORACLE_BASE path for provisioning Oracle database home or Oracle Grid
	// Infrastructure home
	OracleBase *string `mandatory:"false" json:"oracleBase"`

	// Oracle groups configured for the Exadata Fleet Update Home.
	Groups []FsuHomeGroup `mandatory:"false" json:"groups"`

	// User who provisioned the Exadata Fleet Update Home.
	User *string `mandatory:"false" json:"user"`

	// The time the Exadata Fleet Update Home was updated. An RFC3339 formatted datetime string.
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

	// Absolute path location of the software home to be imported (For
	// database images, this will be the ORACLE_HOME).
	Path *string `mandatory:"false" json:"path"`

	// The current state of the Exadata Fleet Update Home.
	LifecycleState HomeLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`
}

// GetFsuImageId returns FsuImageId
func (m DbHomeFsuHome) GetFsuImageId() *string {
	return m.FsuImageId
}

// GetCompartmentId returns CompartmentId
func (m DbHomeFsuHome) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHomeName returns HomeName
func (m DbHomeFsuHome) GetHomeName() *string {
	return m.HomeName
}

// GetOracleBase returns OracleBase
func (m DbHomeFsuHome) GetOracleBase() *string {
	return m.OracleBase
}

// GetGroups returns Groups
func (m DbHomeFsuHome) GetGroups() []FsuHomeGroup {
	return m.Groups
}

// GetUser returns User
func (m DbHomeFsuHome) GetUser() *string {
	return m.User
}

// GetId returns Id
func (m DbHomeFsuHome) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m DbHomeFsuHome) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m DbHomeFsuHome) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DbHomeFsuHome) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m DbHomeFsuHome) GetLifecycleState() HomeLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m DbHomeFsuHome) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m DbHomeFsuHome) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m DbHomeFsuHome) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m DbHomeFsuHome) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DbHomeFsuHome) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbHomeFsuHome) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHomeLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHomeLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbHomeFsuHome) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbHomeFsuHome DbHomeFsuHome
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeDbHomeFsuHome
	}{
		"DBHOME",
		(MarshalTypeDbHomeFsuHome)(m),
	}

	return json.Marshal(&s)
}
