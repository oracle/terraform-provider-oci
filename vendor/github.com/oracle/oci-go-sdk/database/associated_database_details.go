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

// AssociatedDatabaseDetails Databases associated with a backup destination
type AssociatedDatabaseDetails struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"false" json:"id"`

	// The display name of the database that is associated with the backup destination.
	DbName *string `mandatory:"false" json:"dbName"`
}

func (m AssociatedDatabaseDetails) String() string {
	return common.PointerString(m)
}
