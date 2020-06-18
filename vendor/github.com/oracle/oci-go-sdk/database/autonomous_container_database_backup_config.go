// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousContainerDatabaseBackupConfig Backup options for the Autonomous Container Database.
type AutonomousContainerDatabaseBackupConfig struct {

	// Backup destination details.
	BackupDestinationDetails []BackupDestinationDetails `mandatory:"false" json:"backupDestinationDetails"`

	// Number of days between the current and the earliest point of recoverability covered by automatic backups.
	// This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window.
	// When the value is updated, it is applied to all existing automatic backups.
	RecoveryWindowInDays *int `mandatory:"false" json:"recoveryWindowInDays"`
}

func (m AutonomousContainerDatabaseBackupConfig) String() string {
	return common.PointerString(m)
}
