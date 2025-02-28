// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDatabaseFromBackupDetails The representation of CreateDatabaseFromBackupDetails
type CreateDatabaseFromBackupDetails struct {

	// The backup OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	BackupId *string `mandatory:"true" json:"backupId"`

	// A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The password to open the TDE wallet.
	BackupTDEPassword *string `mandatory:"false" json:"backupTDEPassword"`

	SourceEncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"sourceEncryptionKeyLocationDetails"`

	// The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	DbName *string `mandatory:"false" json:"dbName"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	SidPrefix *string `mandatory:"false" json:"sidPrefix"`

	// The list of pluggable databases that needs to be restored into new database.
	PluggableDatabases []string `mandatory:"false" json:"pluggableDatabases"`
}

func (m CreateDatabaseFromBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseFromBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDatabaseFromBackupDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupTDEPassword                  *string                      `json:"backupTDEPassword"`
		SourceEncryptionKeyLocationDetails encryptionkeylocationdetails `json:"sourceEncryptionKeyLocationDetails"`
		DbUniqueName                       *string                      `json:"dbUniqueName"`
		DbName                             *string                      `json:"dbName"`
		SidPrefix                          *string                      `json:"sidPrefix"`
		PluggableDatabases                 []string                     `json:"pluggableDatabases"`
		BackupId                           *string                      `json:"backupId"`
		AdminPassword                      *string                      `json:"adminPassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BackupTDEPassword = model.BackupTDEPassword

	nn, e = model.SourceEncryptionKeyLocationDetails.UnmarshalPolymorphicJSON(model.SourceEncryptionKeyLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceEncryptionKeyLocationDetails = nn.(EncryptionKeyLocationDetails)
	} else {
		m.SourceEncryptionKeyLocationDetails = nil
	}

	m.DbUniqueName = model.DbUniqueName

	m.DbName = model.DbName

	m.SidPrefix = model.SidPrefix

	m.PluggableDatabases = make([]string, len(model.PluggableDatabases))
	copy(m.PluggableDatabases, model.PluggableDatabases)
	m.BackupId = model.BackupId

	m.AdminPassword = model.AdminPassword

	return
}
