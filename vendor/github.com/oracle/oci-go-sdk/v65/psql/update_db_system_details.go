// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDbSystemDetails The information to be updated.
type UpdateDbSystemDetails struct {

	// A user-friendly display name for the database system. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-provided description of the database system.
	Description *string `mandatory:"false" json:"description"`

	// The name of the shape for the database system nodes.
	// Example: `VM.Standard.E4.Flex`
	Shape *string `mandatory:"false" json:"shape"`

	// The total number of OCPUs available to each database system node.
	InstanceOcpuCount *int `mandatory:"false" json:"instanceOcpuCount"`

	// The total amount of memory available to each database system node, in gigabytes.
	InstanceMemorySizeInGBs *int `mandatory:"false" json:"instanceMemorySizeInGBs"`

	DbConfigurationParams *UpdateDbConfigParams `mandatory:"false" json:"dbConfigurationParams"`

	ManagementPolicy *ManagementPolicyDetails `mandatory:"false" json:"managementPolicy"`

	StorageDetails *UpdateStorageDetailsParams `mandatory:"false" json:"storageDetails"`

	NetworkDetails *UpdateNetworkDetails `mandatory:"false" json:"networkDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
