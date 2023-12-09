// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// DailyBackupPolicy Daily backup policy.
type DailyBackupPolicy struct {

	// Hour of the day when the backup starts.
	BackupStart *string `mandatory:"true" json:"backupStart"`

	// How many days the data should be stored after the database system deletion.
	RetentionDays *int `mandatory:"false" json:"retentionDays"`
}

// GetRetentionDays returns RetentionDays
func (m DailyBackupPolicy) GetRetentionDays() *int {
	return m.RetentionDays
}

func (m DailyBackupPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DailyBackupPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DailyBackupPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDailyBackupPolicy DailyBackupPolicy
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeDailyBackupPolicy
	}{
		"DAILY",
		(MarshalTypeDailyBackupPolicy)(m),
	}

	return json.Marshal(&s)
}
