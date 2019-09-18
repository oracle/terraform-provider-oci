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

// CreateDbHomeWithDbSystemIdFromBackupDetails Note that a valid `dbSystemId` value must be supplied for the `CreateDbHomeWithDbSystemIdFromBackup` API operation to successfully complete.
type CreateDbHomeWithDbSystemIdFromBackupDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	Database *CreateDatabaseFromBackupDetails `mandatory:"true" json:"database"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateDbHomeWithDbSystemIdFromBackupDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails CreateDbHomeWithDbSystemIdFromBackupDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails
	}{
		"DB_BACKUP",
		(MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails)(m),
	}

	return json.Marshal(&s)
}
