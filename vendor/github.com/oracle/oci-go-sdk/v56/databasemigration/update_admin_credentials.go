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

// UpdateAdminCredentials Database Administrator Credentials details. An empty object would result in the removal of the stored details.
type UpdateAdminCredentials struct {

	// Administrator username
	Username *string `mandatory:"false" json:"username"`

	// Administrator password
	Password *string `mandatory:"false" json:"password"`
}

func (m UpdateAdminCredentials) String() string {
	return common.PointerString(m)
}
