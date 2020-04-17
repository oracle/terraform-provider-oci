// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RestoreAutonomousDataWarehouseDetails **Deprecated.** See RestoreAutonomousDatabaseDetails for reference information about restoring an Autonomous Data Warehouse.
type RestoreAutonomousDataWarehouseDetails struct {

	// The time to restore the database to.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m RestoreAutonomousDataWarehouseDetails) String() string {
	return common.PointerString(m)
}
