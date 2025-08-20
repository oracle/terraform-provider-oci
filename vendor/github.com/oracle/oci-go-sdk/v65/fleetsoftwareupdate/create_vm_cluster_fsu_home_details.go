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

// CreateVmClusterFsuHomeDetails The information about the new Exadata Fleet Update Home using a VmCluster or CloudVmCluster resource as source.
type CreateVmClusterFsuHomeDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Image for the Home.
	FsuImageId *string `mandatory:"true" json:"fsuImageId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique name of the home.
	// It cannot be the same as an existing Exadata Fleet Update Home resource.
	HomeName *string `mandatory:"true" json:"homeName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster.
	// Only CloudVmCluster or VmCluster resources are allowed.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// Absolute path location of the software home to be imported (For
	// database images, this will be the ORACLE_HOME).
	Path *string `mandatory:"true" json:"path"`

	// ORACLE_BASE path for provisioning Oracle database home or Oracle Grid Infrastructure home
	OracleBase *string `mandatory:"false" json:"oracleBase"`

	// Oracle groups to be configured for the Exadata Fleet Update Home.
	Groups []CreateFsuHomeGroup `mandatory:"false" json:"groups"`

	// Name of the user for whom the software home is being provisioned.
	User *string `mandatory:"false" json:"user"`

	// Exadata Fleet Update Home display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetFsuImageId returns FsuImageId
func (m CreateVmClusterFsuHomeDetails) GetFsuImageId() *string {
	return m.FsuImageId
}

// GetCompartmentId returns CompartmentId
func (m CreateVmClusterFsuHomeDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHomeName returns HomeName
func (m CreateVmClusterFsuHomeDetails) GetHomeName() *string {
	return m.HomeName
}

// GetOracleBase returns OracleBase
func (m CreateVmClusterFsuHomeDetails) GetOracleBase() *string {
	return m.OracleBase
}

// GetGroups returns Groups
func (m CreateVmClusterFsuHomeDetails) GetGroups() []CreateFsuHomeGroup {
	return m.Groups
}

// GetUser returns User
func (m CreateVmClusterFsuHomeDetails) GetUser() *string {
	return m.User
}

// GetDisplayName returns DisplayName
func (m CreateVmClusterFsuHomeDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m CreateVmClusterFsuHomeDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateVmClusterFsuHomeDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateVmClusterFsuHomeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVmClusterFsuHomeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateVmClusterFsuHomeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateVmClusterFsuHomeDetails CreateVmClusterFsuHomeDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeCreateVmClusterFsuHomeDetails
	}{
		"VMCLUSTER",
		(MarshalTypeCreateVmClusterFsuHomeDetails)(m),
	}

	return json.Marshal(&s)
}
