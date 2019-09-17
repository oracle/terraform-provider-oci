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

// CreateDbHomeWithVmClusterIdDetails Note that a valid `vmClusterId` value must be supplied for the `CreateDbHomeWithVmClusterId` API operation to successfully complete.
type CreateDbHomeWithVmClusterIdDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// A valid Oracle Database version. To get a list of supported versions, use the ListDbVersions operation.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	Database *CreateDatabaseDetails `mandatory:"true" json:"database"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithVmClusterIdDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateDbHomeWithVmClusterIdDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithVmClusterIdDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithVmClusterIdDetails CreateDbHomeWithVmClusterIdDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithVmClusterIdDetails
	}{
		"VM_CLUSTER_NEW",
		(MarshalTypeCreateDbHomeWithVmClusterIdDetails)(m),
	}

	return json.Marshal(&s)
}
