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

// AutonomousDatabaseConsoleTokenDetails The token that allows the OCI Console to access the Autonomous Transaction Processing Service Console.
type AutonomousDatabaseConsoleTokenDetails struct {

	// The token that allows the OCI Console to access the Autonomous Transaction Processing Service Console.
	Token *string `mandatory:"false" json:"token"`

	// The login URL that allows the OCI Console to access the Autonomous Transaction Processing Service Console.
	LoginUrl *string `mandatory:"false" json:"loginUrl"`
}

func (m AutonomousDatabaseConsoleTokenDetails) String() string {
	return common.PointerString(m)
}
