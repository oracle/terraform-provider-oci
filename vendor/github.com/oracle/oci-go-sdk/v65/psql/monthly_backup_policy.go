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

// MonthlyBackupPolicy Monthly backup policy
type MonthlyBackupPolicy struct {

	// Hour of the day when backup starts.
	BackupStart *string `mandatory:"true" json:"backupStart"`

	// Days of the month when backup should start.
	// If the day is greater last day of the current month, then it will be triggered on the last day of the current month
	DaysOfTheMonth []int `mandatory:"true" json:"daysOfTheMonth"`

	// How many days the customers data should be stored after the db system deletion.
	RetentionDays *int `mandatory:"false" json:"retentionDays"`
}

// GetRetentionDays returns RetentionDays
func (m MonthlyBackupPolicy) GetRetentionDays() *int {
	return m.RetentionDays
}

func (m MonthlyBackupPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonthlyBackupPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MonthlyBackupPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMonthlyBackupPolicy MonthlyBackupPolicy
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeMonthlyBackupPolicy
	}{
		"MONTHLY",
		(MarshalTypeMonthlyBackupPolicy)(m),
	}

	return json.Marshal(&s)
}
