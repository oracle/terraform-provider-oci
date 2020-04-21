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

// CreateConsoleConnectionDetails The details for creating a Db node console connection.
// The Db node console connection is created in the same compartment as the dbNode.
type CreateConsoleConnectionDetails struct {

	// The SSH public key used to authenticate the console connection.
	PublicKey *string `mandatory:"true" json:"publicKey"`
}

func (m CreateConsoleConnectionDetails) String() string {
	return common.PointerString(m)
}
