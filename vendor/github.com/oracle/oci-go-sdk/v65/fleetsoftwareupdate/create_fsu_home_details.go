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

// CreateFsuHomeDetails The information about the new Exadata Fleet Update Home.
type CreateFsuHomeDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Image for the Home.
	GetFsuImageId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// resource.
	GetCompartmentId() *string

	// Unique name of the home.
	// It cannot be the same as an existing Exadata Fleet Update Home resource.
	GetHomeName() *string

	// ORACLE_BASE path for provisioning Oracle database home or Oracle Grid Infrastructure home
	GetOracleBase() *string

	// Oracle groups to be configured for the Exadata Fleet Update Home.
	GetGroups() []CreateFsuHomeGroup

	// Name of the user for whom the software home is being provisioned.
	GetUser() *string

	// Exadata Fleet Update Home display name.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createfsuhomedetails struct {
	JsonData      []byte
	OracleBase    *string                           `mandatory:"false" json:"oracleBase"`
	Groups        []CreateFsuHomeGroup              `mandatory:"false" json:"groups"`
	User          *string                           `mandatory:"false" json:"user"`
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FsuImageId    *string                           `mandatory:"true" json:"fsuImageId"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	HomeName      *string                           `mandatory:"true" json:"homeName"`
	Kind          string                            `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *createfsuhomedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatefsuhomedetails createfsuhomedetails
	s := struct {
		Model Unmarshalercreatefsuhomedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FsuImageId = s.Model.FsuImageId
	m.CompartmentId = s.Model.CompartmentId
	m.HomeName = s.Model.HomeName
	m.OracleBase = s.Model.OracleBase
	m.Groups = s.Model.Groups
	m.User = s.Model.User
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createfsuhomedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "DBHOME":
		mm := CreateDbHomeFsuHomeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VMCLUSTER":
		mm := CreateVmClusterFsuHomeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateFsuHomeDetails: %s.", m.Kind)
		return *m, nil
	}
}

// GetOracleBase returns OracleBase
func (m createfsuhomedetails) GetOracleBase() *string {
	return m.OracleBase
}

// GetGroups returns Groups
func (m createfsuhomedetails) GetGroups() []CreateFsuHomeGroup {
	return m.Groups
}

// GetUser returns User
func (m createfsuhomedetails) GetUser() *string {
	return m.User
}

// GetDisplayName returns DisplayName
func (m createfsuhomedetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m createfsuhomedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createfsuhomedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFsuImageId returns FsuImageId
func (m createfsuhomedetails) GetFsuImageId() *string {
	return m.FsuImageId
}

// GetCompartmentId returns CompartmentId
func (m createfsuhomedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHomeName returns HomeName
func (m createfsuhomedetails) GetHomeName() *string {
	return m.HomeName
}

func (m createfsuhomedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createfsuhomedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
