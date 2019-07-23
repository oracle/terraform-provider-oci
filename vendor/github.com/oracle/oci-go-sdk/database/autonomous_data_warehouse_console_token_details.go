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

// AutonomousDataWarehouseConsoleTokenDetails **Deprecated.** See AutonomousDatabaseConsoleTokenDetails for reference information about the token that allows the OCI Console to access the Autonomous Data Warehouse Service Console.
type AutonomousDataWarehouseConsoleTokenDetails struct {

	// The token that allows the OCI Console to access the Autonomous Data Warehouse Service Console.
	Token *string `mandatory:"false" json:"token"`

	// The login URL that allows the OCI Console to access the Autonomous Data Warehouse Service Console.
	LoginUrl *string `mandatory:"false" json:"loginUrl"`
}

func (m AutonomousDataWarehouseConsoleTokenDetails) String() string {
	return common.PointerString(m)
}
