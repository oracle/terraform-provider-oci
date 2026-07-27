// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOracleOnlinePhysicalMigrationSettings Create settings for an online physical Oracle migration.
type CreateOracleOnlinePhysicalMigrationSettings struct {
	ObjectStorageBucket *CreateObjectStoreBucket `mandatory:"true" json:"objectStorageBucket"`

	// The OCID of the resource being referenced.
	SourceDatabaseHostOracleUserPasswordSecretId *string `mandatory:"true" json:"sourceDatabaseHostOracleUserPasswordSecretId"`

	// Destination location to use for the target database DB_CREATE_FILE_DEST parameter during
	// ONLINE_PHYSICAL migrations. Oracle uses this location for Oracle-managed target database
	// data files, temp files, and control files.
	TargetDatabaseCreateFileDestination *string `mandatory:"true" json:"targetDatabaseCreateFileDestination"`

	// Destination location to use for the target database DB_CREATE_ONLINE_LOG_DEST_1 parameter
	// during ONLINE_PHYSICAL migrations. Oracle uses this location for all Oracle-managed target
	// database online redo log files.
	TargetDatabaseCreateOnlineLogDestination *string `mandatory:"true" json:"targetDatabaseCreateOnlineLogDestination"`
}

func (m CreateOracleOnlinePhysicalMigrationSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOracleOnlinePhysicalMigrationSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOracleOnlinePhysicalMigrationSettings) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOracleOnlinePhysicalMigrationSettings CreateOracleOnlinePhysicalMigrationSettings
	s := struct {
		DiscriminatorParam string `json:"migrationMethod"`
		MarshalTypeCreateOracleOnlinePhysicalMigrationSettings
	}{
		"ONLINE_PHYSICAL",
		(MarshalTypeCreateOracleOnlinePhysicalMigrationSettings)(m),
	}

	return json.Marshal(&s)
}
