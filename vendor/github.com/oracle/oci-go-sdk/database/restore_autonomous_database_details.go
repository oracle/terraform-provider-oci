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

// RestoreAutonomousDatabaseDetails Details to restore an Oracle Autonomous Database.
type RestoreAutonomousDatabaseDetails struct {

	// The time to restore the database to.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m RestoreAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}
