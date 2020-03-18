// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
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
