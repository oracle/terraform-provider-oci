// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// DatabaseConnectionString The Oracle Database connection string.
type DatabaseConnectionString struct {

	// The host name of the database.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The port used to connect to the database.
	Port *int `mandatory:"true" json:"port"`

	// The name of the service alias used to connect to the database.
	Service *string `mandatory:"true" json:"service"`

	// The protocol used to connect to the database.
	Protocol DatabaseConnectionStringProtocolEnum `mandatory:"true" json:"protocol"`
}

func (m DatabaseConnectionString) String() string {
	return common.PointerString(m)
}

// DatabaseConnectionStringProtocolEnum Enum with underlying type: string
type DatabaseConnectionStringProtocolEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProtocolEnum
const (
	DatabaseConnectionStringProtocolTcp DatabaseConnectionStringProtocolEnum = "TCP"
)

var mappingDatabaseConnectionStringProtocol = map[string]DatabaseConnectionStringProtocolEnum{
	"TCP": DatabaseConnectionStringProtocolTcp,
}

// GetDatabaseConnectionStringProtocolEnumValues Enumerates the set of values for DatabaseConnectionStringProtocolEnum
func GetDatabaseConnectionStringProtocolEnumValues() []DatabaseConnectionStringProtocolEnum {
	values := make([]DatabaseConnectionStringProtocolEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProtocol {
		values = append(values, v)
	}
	return values
}
