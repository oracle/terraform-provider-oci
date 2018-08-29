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

// AutonomousDatabaseConnectionStrings Connection strings to connect to an Oracle Autonomous Database.
type AutonomousDatabaseConnectionStrings struct {

	// The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	High *string `mandatory:"false" json:"high"`

	// The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	Low *string `mandatory:"false" json:"low"`

	// The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
	Medium *string `mandatory:"false" json:"medium"`
}

func (m AutonomousDatabaseConnectionStrings) String() string {
	return common.PointerString(m)
}
