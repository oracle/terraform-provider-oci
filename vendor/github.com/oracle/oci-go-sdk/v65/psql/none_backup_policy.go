// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NoneBackupPolicy No backup policy.
type NoneBackupPolicy struct {

	// How many days the data should be stored after the database system deletion.
	RetentionDays *int `mandatory:"false" json:"retentionDays"`
}

// GetRetentionDays returns RetentionDays
func (m NoneBackupPolicy) GetRetentionDays() *int {
	return m.RetentionDays
}

func (m NoneBackupPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NoneBackupPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NoneBackupPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNoneBackupPolicy NoneBackupPolicy
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeNoneBackupPolicy
	}{
		"NONE",
		(MarshalTypeNoneBackupPolicy)(m),
	}

	return json.Marshal(&s)
}
