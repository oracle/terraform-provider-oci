// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleProtectedDatabaseDeletionDetails The details for scheduling deletion of the protected database
type ScheduleProtectedDatabaseDeletionDetails struct {

	// Defines a preferred schedule to delete a protected database after you terminate the source database.
	// * The default schedule is DELETE_AFTER_72_HOURS, so that the delete operation can occur 72 hours (3 days) after the source database is terminated.
	// * The alternate schedule is DELETE_AFTER_RETENTION_PERIOD. Specify this option if you want to delete a protected database only after the policy-defined backup retention period expires.
	DeletionSchedule DeletionScheduleEnum `mandatory:"false" json:"deletionSchedule,omitempty"`
}

func (m ScheduleProtectedDatabaseDeletionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleProtectedDatabaseDeletionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeletionScheduleEnum(string(m.DeletionSchedule)); !ok && m.DeletionSchedule != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeletionSchedule: %s. Supported values are: %s.", m.DeletionSchedule, strings.Join(GetDeletionScheduleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
