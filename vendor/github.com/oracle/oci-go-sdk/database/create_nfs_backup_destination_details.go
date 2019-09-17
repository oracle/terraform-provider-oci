// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateNfsBackupDestinationDetails Used for creating NFS backup destinations.
type CreateNfsBackupDestinationDetails struct {

	// The user-provided name of the backup destination.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The local directory path on each VM cluster node where the NFS server location is mounted. The local directory path and the NFS server location must each be the same across all of the VM cluster nodes. Ensure that the NFS mount is maintained continuously on all of the VM cluster nodes.
	LocalMountPointPath *string `mandatory:"true" json:"localMountPointPath"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m CreateNfsBackupDestinationDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateNfsBackupDestinationDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m CreateNfsBackupDestinationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateNfsBackupDestinationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateNfsBackupDestinationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateNfsBackupDestinationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateNfsBackupDestinationDetails CreateNfsBackupDestinationDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateNfsBackupDestinationDetails
	}{
		"NFS",
		(MarshalTypeCreateNfsBackupDestinationDetails)(m),
	}

	return json.Marshal(&s)
}
