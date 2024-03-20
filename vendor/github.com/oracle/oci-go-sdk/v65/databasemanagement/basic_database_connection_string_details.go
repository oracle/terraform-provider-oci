// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BasicDatabaseConnectionStringDetails The details of the Oracle Database basic connection string.
type BasicDatabaseConnectionStringDetails struct {

	// The service name of the database.
	Service *string `mandatory:"true" json:"service"`

	// The port number used to connect to the database.
	Port *int `mandatory:"true" json:"port"`

	// The protocol used to connect to the database.
	Protocol BasicDatabaseConnectionStringDetailsProtocolEnum `mandatory:"true" json:"protocol"`
}

func (m BasicDatabaseConnectionStringDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BasicDatabaseConnectionStringDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBasicDatabaseConnectionStringDetailsProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetBasicDatabaseConnectionStringDetailsProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BasicDatabaseConnectionStringDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBasicDatabaseConnectionStringDetails BasicDatabaseConnectionStringDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeBasicDatabaseConnectionStringDetails
	}{
		"BASIC",
		(MarshalTypeBasicDatabaseConnectionStringDetails)(m),
	}

	return json.Marshal(&s)
}

// BasicDatabaseConnectionStringDetailsProtocolEnum Enum with underlying type: string
type BasicDatabaseConnectionStringDetailsProtocolEnum string

// Set of constants representing the allowable values for BasicDatabaseConnectionStringDetailsProtocolEnum
const (
	BasicDatabaseConnectionStringDetailsProtocolTcp  BasicDatabaseConnectionStringDetailsProtocolEnum = "TCP"
	BasicDatabaseConnectionStringDetailsProtocolTcps BasicDatabaseConnectionStringDetailsProtocolEnum = "TCPS"
)

var mappingBasicDatabaseConnectionStringDetailsProtocolEnum = map[string]BasicDatabaseConnectionStringDetailsProtocolEnum{
	"TCP":  BasicDatabaseConnectionStringDetailsProtocolTcp,
	"TCPS": BasicDatabaseConnectionStringDetailsProtocolTcps,
}

var mappingBasicDatabaseConnectionStringDetailsProtocolEnumLowerCase = map[string]BasicDatabaseConnectionStringDetailsProtocolEnum{
	"tcp":  BasicDatabaseConnectionStringDetailsProtocolTcp,
	"tcps": BasicDatabaseConnectionStringDetailsProtocolTcps,
}

// GetBasicDatabaseConnectionStringDetailsProtocolEnumValues Enumerates the set of values for BasicDatabaseConnectionStringDetailsProtocolEnum
func GetBasicDatabaseConnectionStringDetailsProtocolEnumValues() []BasicDatabaseConnectionStringDetailsProtocolEnum {
	values := make([]BasicDatabaseConnectionStringDetailsProtocolEnum, 0)
	for _, v := range mappingBasicDatabaseConnectionStringDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetBasicDatabaseConnectionStringDetailsProtocolEnumStringValues Enumerates the set of values in String for BasicDatabaseConnectionStringDetailsProtocolEnum
func GetBasicDatabaseConnectionStringDetailsProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingBasicDatabaseConnectionStringDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBasicDatabaseConnectionStringDetailsProtocolEnum(val string) (BasicDatabaseConnectionStringDetailsProtocolEnum, bool) {
	enum, ok := mappingBasicDatabaseConnectionStringDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
