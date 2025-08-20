// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceWindowSchedule MaintenanceWindow Information for Schedule.
type MaintenanceWindowSchedule struct {

	// Start Date for the schedule. An RFC3339 formatted datetime string
	ExecutionStartdate *common.SDKTime `mandatory:"true" json:"executionStartdate"`

	// Provide MaintenanceWindowId
	MaintenanceWindowId *string `mandatory:"true" json:"maintenanceWindowId"`
}

// GetExecutionStartdate returns ExecutionStartdate
func (m MaintenanceWindowSchedule) GetExecutionStartdate() *common.SDKTime {
	return m.ExecutionStartdate
}

func (m MaintenanceWindowSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceWindowSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MaintenanceWindowSchedule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMaintenanceWindowSchedule MaintenanceWindowSchedule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeMaintenanceWindowSchedule
	}{
		"MAINTENANCE_WINDOW",
		(MarshalTypeMaintenanceWindowSchedule)(m),
	}

	return json.Marshal(&s)
}
