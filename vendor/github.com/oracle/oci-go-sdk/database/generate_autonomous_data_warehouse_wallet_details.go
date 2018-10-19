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

// GenerateAutonomousDataWarehouseWalletDetails Details to create and download a wallet for an Oracle Autonomous Data Warehouse.
type GenerateAutonomousDataWarehouseWalletDetails struct {

	// The password to encrypt the keys inside the wallet.
	Password *string `mandatory:"true" json:"password"`
}

func (m GenerateAutonomousDataWarehouseWalletDetails) String() string {
	return common.PointerString(m)
}
