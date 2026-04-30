// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectionOption Types of connection supported by Data Safe.
type ConnectionOption struct {

	// The connection type used by Data Safe to connect to the database. Allowed values:
	// - PRIVATE_ENDPOINT - Represents connection through private endpoint in Data Safe.
	// - ONPREM_CONNECTOR - Represents connection through on-premises connector in Data Safe.
	ConnectionType ConnectionOptionConnectionTypeEnum `mandatory:"true" json:"connectionType"`

	// List of OCIDs required to establish the connection.
	// - For `PRIVATE_ENDPOINT`, provide the OCID of Data Safe private endpoint.
	// - For `ONPREM_CONNECTOR`, provide the OCID(s) of on-premises connector(s).
	Identifiers []string `mandatory:"true" json:"identifiers"`
}

func (m ConnectionOption) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectionOption) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionOptionConnectionTypeEnum(string(m.ConnectionType)); !ok && m.ConnectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionType: %s. Supported values are: %s.", m.ConnectionType, strings.Join(GetConnectionOptionConnectionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionOptionConnectionTypeEnum Enum with underlying type: string
type ConnectionOptionConnectionTypeEnum string

// Set of constants representing the allowable values for ConnectionOptionConnectionTypeEnum
const (
	ConnectionOptionConnectionTypePrivateEndpoint ConnectionOptionConnectionTypeEnum = "PRIVATE_ENDPOINT"
	ConnectionOptionConnectionTypeOnpremConnector ConnectionOptionConnectionTypeEnum = "ONPREM_CONNECTOR"
)

var mappingConnectionOptionConnectionTypeEnum = map[string]ConnectionOptionConnectionTypeEnum{
	"PRIVATE_ENDPOINT": ConnectionOptionConnectionTypePrivateEndpoint,
	"ONPREM_CONNECTOR": ConnectionOptionConnectionTypeOnpremConnector,
}

var mappingConnectionOptionConnectionTypeEnumLowerCase = map[string]ConnectionOptionConnectionTypeEnum{
	"private_endpoint": ConnectionOptionConnectionTypePrivateEndpoint,
	"onprem_connector": ConnectionOptionConnectionTypeOnpremConnector,
}

// GetConnectionOptionConnectionTypeEnumValues Enumerates the set of values for ConnectionOptionConnectionTypeEnum
func GetConnectionOptionConnectionTypeEnumValues() []ConnectionOptionConnectionTypeEnum {
	values := make([]ConnectionOptionConnectionTypeEnum, 0)
	for _, v := range mappingConnectionOptionConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionOptionConnectionTypeEnumStringValues Enumerates the set of values in String for ConnectionOptionConnectionTypeEnum
func GetConnectionOptionConnectionTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_ENDPOINT",
		"ONPREM_CONNECTOR",
	}
}

// GetMappingConnectionOptionConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionOptionConnectionTypeEnum(val string) (ConnectionOptionConnectionTypeEnum, bool) {
	enum, ok := mappingConnectionOptionConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
