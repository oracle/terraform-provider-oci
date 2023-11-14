// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// A description of the PGSQL Control Plane API
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackupSourceDetails Restoring to a new DbSystem from the backup. The DbSystem details that are part of the CreateDbSystem request are not required, but if present will override the backup's DbSystem details snapshot.
type BackupSourceDetails struct {

	// DbSystem backup identifier.
	BackupId *string `mandatory:"true" json:"backupId"`

	// Restore the DB config overrides from backup. Default is false
	IsHavingRestoreConfigOverrides *bool `mandatory:"false" json:"isHavingRestoreConfigOverrides"`
}

func (m BackupSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BackupSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBackupSourceDetails BackupSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeBackupSourceDetails
	}{
		"BACKUP",
		(MarshalTypeBackupSourceDetails)(m),
	}

	return json.Marshal(&s)
}
