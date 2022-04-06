// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDbSystemSourceFromBackupDetails Use the backupId to specify from which backup the new DB System will be created.
type CreateDbSystemSourceFromBackupDetails struct {

	// The OCID of the backup to be used as the source for the new DB System.
	BackupId *string `mandatory:"true" json:"backupId"`
}

func (m CreateDbSystemSourceFromBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbSystemSourceFromBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDbSystemSourceFromBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbSystemSourceFromBackupDetails CreateDbSystemSourceFromBackupDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateDbSystemSourceFromBackupDetails
	}{
		"BACKUP",
		(MarshalTypeCreateDbSystemSourceFromBackupDetails)(m),
	}

	return json.Marshal(&s)
}
