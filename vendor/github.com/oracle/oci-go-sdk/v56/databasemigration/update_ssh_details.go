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

// UpdateSshDetails Details of the SSH key that will be used.
type UpdateSshDetails struct {

	// Name of the host the SSH key is valid for.
	Host *string `mandatory:"false" json:"host"`

	// Private SSH key string.
	Sshkey *string `mandatory:"false" json:"sshkey"`

	// SSH user
	User *string `mandatory:"false" json:"user"`

	// Sudo location
	SudoLocation *string `mandatory:"false" json:"sudoLocation"`
}

func (m UpdateSshDetails) String() string {
	return common.PointerString(m)
}
