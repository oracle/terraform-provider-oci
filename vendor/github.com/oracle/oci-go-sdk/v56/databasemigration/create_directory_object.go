// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateDirectoryObject Directory object details, used to define either import or export directory objects in Data Pump Settings.
// Import directory is required for Non-Autonomous target connections. If specified for an autonomous target, it will show an error.
// Export directory will error if there are database link details specified.
type CreateDirectoryObject struct {

	// Name of directory object in database
	Name *string `mandatory:"true" json:"name"`

	// Absolute path of directory on database server
	Path *string `mandatory:"true" json:"path"`
}

func (m CreateDirectoryObject) String() string {
	return common.PointerString(m)
}
