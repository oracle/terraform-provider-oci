// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseConnectionString The Oracle Database connection string.
type DatabaseConnectionString struct {

	// The host name of the database or the SCAN name in case of a RAC database.
	HostName *string `mandatory:"true" json:"hostName"`

	// The port used to connect to the database.
	Port *int `mandatory:"true" json:"port"`

	// The service name of the database.
	Service *string `mandatory:"true" json:"service"`

	// The protocol used to connect to the database.
	Protocol DatabaseConnectionStringProtocolEnum `mandatory:"true" json:"protocol"`
}

func (m DatabaseConnectionString) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseConnectionString) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseConnectionStringProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetDatabaseConnectionStringProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseConnectionStringProtocolEnum Enum with underlying type: string
type DatabaseConnectionStringProtocolEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProtocolEnum
const (
	DatabaseConnectionStringProtocolTcp  DatabaseConnectionStringProtocolEnum = "TCP"
	DatabaseConnectionStringProtocolTcps DatabaseConnectionStringProtocolEnum = "TCPS"
)

var mappingDatabaseConnectionStringProtocolEnum = map[string]DatabaseConnectionStringProtocolEnum{
	"TCP":  DatabaseConnectionStringProtocolTcp,
	"TCPS": DatabaseConnectionStringProtocolTcps,
}

var mappingDatabaseConnectionStringProtocolEnumLowerCase = map[string]DatabaseConnectionStringProtocolEnum{
	"tcp":  DatabaseConnectionStringProtocolTcp,
	"tcps": DatabaseConnectionStringProtocolTcps,
}

// GetDatabaseConnectionStringProtocolEnumValues Enumerates the set of values for DatabaseConnectionStringProtocolEnum
func GetDatabaseConnectionStringProtocolEnumValues() []DatabaseConnectionStringProtocolEnum {
	values := make([]DatabaseConnectionStringProtocolEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionStringProtocolEnumStringValues Enumerates the set of values in String for DatabaseConnectionStringProtocolEnum
func GetDatabaseConnectionStringProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingDatabaseConnectionStringProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionStringProtocolEnum(val string) (DatabaseConnectionStringProtocolEnum, bool) {
	enum, ok := mappingDatabaseConnectionStringProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
