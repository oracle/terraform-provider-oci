// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompleteExternalBackupJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
