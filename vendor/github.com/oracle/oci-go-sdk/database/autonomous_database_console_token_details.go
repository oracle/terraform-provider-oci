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

// AutonomousDatabaseConsoleTokenDetails The token that allows the OCI Console to access the Autonomous Database Service Console.
type AutonomousDatabaseConsoleTokenDetails struct {

	// The token that allows the OCI Console to access the Autonomous Transaction Processing Service Console.
	Token *string `mandatory:"false" json:"token"`

	// The login URL that allows the OCI Console to access the Autonomous Transaction Processing Service Console.
	LoginUrl *string `mandatory:"false" json:"loginUrl"`
}

func (m AutonomousDatabaseConsoleTokenDetails) String() string {
	return common.PointerString(m)
}
