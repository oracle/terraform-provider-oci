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

// DbIormConfig IORM Config setting response for this database
type DbIormConfig struct {

	// Database Name. For default DbPlan, the dbName will always be `default`
	DbName *string `mandatory:"false" json:"dbName"`

	// Relative priority of a database
	Share *int `mandatory:"false" json:"share"`

	// Flash Cache limit, internally configured based on shares
	FlashCacheLimit *string `mandatory:"false" json:"flashCacheLimit"`
}

func (m DbIormConfig) String() string {
	return common.PointerString(m)
}
