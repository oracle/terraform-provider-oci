// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CompleteExternalBackupJobDetails The representation of CompleteExternalBackupJobDetails
type CompleteExternalBackupJobDetails struct {

	// The handle of the control file backup.
	CfBackupHandle *string `mandatory:"false" json:"cfBackupHandle"`

	// The size of the data in the database, in megabytes.
	DataSize *int64 `mandatory:"false" json:"dataSize"`

	// The size of the redo in the database, in megabytes.
	RedoSize *int64 `mandatory:"false" json:"redoSize"`

	// The handle of the spfile backup.
	SpfBackupHandle *string `mandatory:"false" json:"spfBackupHandle"`

	// The list of SQL patches that need to be applied to the backup during the restore.
	SqlPatches []string `mandatory:"false" json:"sqlPatches"`

	// If the database being backed up is TDE enabled, this will be the path to the associated TDE wallet in Object Storage.
	TdeWalletPath *string `mandatory:"false" json:"tdeWalletPath"`
}

func (m CompleteExternalBackupJobDetails) String() string {
	return common.PointerString(m)
}
