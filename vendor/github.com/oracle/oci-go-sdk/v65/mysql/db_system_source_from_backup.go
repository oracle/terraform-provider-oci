// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DbSystemSourceFromBackup From which backup this DB System was created.
type DbSystemSourceFromBackup struct {

	// The OCID of the backup to be used as the source for the new DB System.
	BackupId *string `mandatory:"true" json:"backupId"`
}

func (m DbSystemSourceFromBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemSourceFromBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbSystemSourceFromBackup) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbSystemSourceFromBackup DbSystemSourceFromBackup
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbSystemSourceFromBackup
	}{
		"BACKUP",
		(MarshalTypeDbSystemSourceFromBackup)(m),
	}

	return json.Marshal(&s)
}
