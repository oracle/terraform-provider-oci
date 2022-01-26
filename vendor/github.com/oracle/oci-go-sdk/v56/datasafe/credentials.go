// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Credentials The database credentials required for Data Safe to connect to the database.
type Credentials struct {

	// The database user name.
	UserName *string `mandatory:"true" json:"userName"`

	// The password of the database user.
	Password *string `mandatory:"true" json:"password"`
}

func (m Credentials) String() string {
	return common.PointerString(m)
}
