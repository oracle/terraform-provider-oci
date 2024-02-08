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

// Metrics Backup performance and storage utilization metrics for the protected database.
type Metrics struct {

	// Backup storage space, in gigabytes, utilized by the protected database. Oracle charges for the total storage used.
	BackupSpaceUsedInGBs *float32 `mandatory:"false" json:"backupSpaceUsedInGBs"`

	// The estimated backup storage space, in gigabytes, required to meet the recovery window goal, including foot print and backups for the protected database.
	BackupSpaceEstimateInGBs *float32 `mandatory:"false" json:"backupSpaceEstimateInGBs"`

	// This is the time window when there is data loss exposure. The point after which recovery is impossible unless additional redo is available.
	// This is the time we received the last backup or last redo-log shipped.
	UnprotectedWindowInSeconds *float32 `mandatory:"false" json:"unprotectedWindowInSeconds"`

	// The estimated space, in gigabytes, consumed by the protected database. The database size is based on the size of the data files in the catalog, and does not include archive logs.
	DbSizeInGBs *float32 `mandatory:"false" json:"dbSizeInGBs"`

	// The value TRUE indicates that the protected database is configured to use Real-time data protection, and redo-data is sent from the protected database to Recovery Service.
	// Real-time data protection substantially reduces the window of potential data loss that exists between successive archived redo log backups.
	IsRedoLogsEnabled *bool `mandatory:"false" json:"isRedoLogsEnabled"`

	// The maximum number of days to retain backups for a protected database.
	RetentionPeriodInDays *float32 `mandatory:"false" json:"retentionPeriodInDays"`

	// Number of seconds backups are currently retained for this database.
	CurrentRetentionPeriodInSeconds *float32 `mandatory:"false" json:"currentRetentionPeriodInSeconds"`

	// Number of days of redo/archive to be applied to recover database.
	MinimumRecoveryNeededInDays *float32 `mandatory:"false" json:"minimumRecoveryNeededInDays"`
}

func (m Metrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Metrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
