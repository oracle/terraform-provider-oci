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

// ManagementPolicyDetails Posgresql DB system management policy update details
type ManagementPolicyDetails struct {

	// The start of the maintenance window.
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
