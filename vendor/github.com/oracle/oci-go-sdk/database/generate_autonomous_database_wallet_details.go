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

// GenerateAutonomousDatabaseWalletDetails Details to create and download an Oracle Autonomous Transaction Processing database wallet.
type GenerateAutonomousDatabaseWalletDetails struct {

	// The password to encrypt the keys inside the wallet.
	Password *string `mandatory:"true" json:"password"`
}

func (m GenerateAutonomousDatabaseWalletDetails) String() string {
	return common.PointerString(m)
}
