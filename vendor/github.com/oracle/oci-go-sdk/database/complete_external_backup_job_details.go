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

// CompleteExternalBackupJobDetails The representation of CompleteExternalBackupJobDetails
type CompleteExternalBackupJobDetails struct {

	// If the database being backed up is TDE enabled, this will be the path to the associated TDE wallet in Object Storage.
	TdeWalletPath *string `mandatory:"false" json:"tdeWalletPath"`

	// The handle of the control file backup.
	CfBackupHandle *string `mandatory:"false" json:"cfBackupHandle"`

	// The handle of the spfile backup.
	SpfBackupHandle *string `mandatory:"false" json:"spfBackupHandle"`

	// The list of SQL patches that need to be applied to the backup during the restore.
	SqlPatches []string `mandatory:"false" json:"sqlPatches"`

	// The size of the data in the database, in megabytes.
	DataSize *int64 `mandatory:"false" json:"dataSize"`

	// The size of the redo in the database, in megabytes.
	RedoSize *int64 `mandatory:"false" json:"redoSize"`
}

func (m CompleteExternalBackupJobDetails) String() string {
	return common.PointerString(m)
}
