// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementPolicyDetails PostgreSQL database system management policy update details.
type ManagementPolicyDetails struct {

	// The start of the maintenance window in UTC.
	// This string is of the format: "{day-of-week} {time-of-day}".
	// "{day-of-week}" is a case-insensitive string like "mon", "tue", &c.
	// "{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.
	MaintenanceWindowStart *string `mandatory:"false" json:"maintenanceWindowStart"`

	BackupPolicy BackupPolicy `mandatory:"false" json:"backupPolicy"`
}

func (m ManagementPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ManagementPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		MaintenanceWindowStart *string      `json:"maintenanceWindowStart"`
		BackupPolicy           backuppolicy `json:"backupPolicy"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.MaintenanceWindowStart = model.MaintenanceWindowStart

	nn, e = model.BackupPolicy.UnmarshalPolymorphicJSON(model.BackupPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BackupPolicy = nn.(BackupPolicy)
	} else {
		m.BackupPolicy = nil
	}

	return
}
