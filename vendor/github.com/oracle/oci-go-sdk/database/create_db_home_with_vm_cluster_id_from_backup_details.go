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

// CreateDbHomeWithVmClusterIdFromBackupDetails Note that a valid `vmClusterId` value must be supplied for the `CreateDbHomeWithVmClusterIdFromBackup` API operation to successfully complete.
type CreateDbHomeWithVmClusterIdFromBackupDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	Database *CreateDatabaseFromBackupDetails `mandatory:"true" json:"database"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithVmClusterIdFromBackupDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateDbHomeWithVmClusterIdFromBackupDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithVmClusterIdFromBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithVmClusterIdFromBackupDetails CreateDbHomeWithVmClusterIdFromBackupDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithVmClusterIdFromBackupDetails
	}{
		"VM_CLUSTER_BACKUP",
		(MarshalTypeCreateDbHomeWithVmClusterIdFromBackupDetails)(m),
	}

	return json.Marshal(&s)
}
