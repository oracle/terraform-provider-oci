// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
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

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// **Deprecated.** The local directory path on each VM cluster node where the NFS server location is mounted. The local directory path and the NFS server location must each be the same across all of the VM cluster nodes. Ensure that the NFS mount is maintained continuously on all of the VM cluster nodes.
	// This field is deprecated. Use the mountTypeDetails field instead to specify the mount type for NFS.
	LocalMountPointPath *string `mandatory:"false" json:"localMountPointPath"`

	MountTypeDetails MountTypeDetails `mandatory:"false" json:"mountTypeDetails"`
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

// UnmarshalJSON unmarshals from json
func (m *CreateNfsBackupDestinationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags        map[string]string                 `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		LocalMountPointPath *string                           `json:"localMountPointPath"`
		MountTypeDetails    mounttypedetails                  `json:"mountTypeDetails"`
		DisplayName         *string                           `json:"displayName"`
		CompartmentId       *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.LocalMountPointPath = model.LocalMountPointPath

	nn, e = model.MountTypeDetails.UnmarshalPolymorphicJSON(model.MountTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.MountTypeDetails = nn.(MountTypeDetails)
	} else {
		m.MountTypeDetails = nil
	}

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	return
}
