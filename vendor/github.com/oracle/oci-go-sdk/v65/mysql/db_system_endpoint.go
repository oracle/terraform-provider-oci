// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbSystemEndpoint A particular functional endpoint for access to a DB System, and the properties that apply to it.
type DbSystemEndpoint struct {

	// The IP address the DB System is configured to listen on.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The port the MySQL instance listens on.
	Port *int `mandatory:"true" json:"port"`

	// The network port where to connect to use this endpoint using the X protocol.
	PortX *int `mandatory:"true" json:"portX"`

	// The network address of the DB System.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The access modes from the client that this endpoint supports.
	Modes []DbSystemEndpointModesEnum `mandatory:"false" json:"modes,omitempty"`

	// The state of the endpoints, as far as it can seen from the DB System.
	// There may be some inconsistency with the actual state of the MySQL service.
	Status DbSystemEndpointStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Additional information about the current endpoint status.
	StatusDetails *string `mandatory:"false" json:"statusDetails"`
}

func (m DbSystemEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Modes {
		if _, ok := GetMappingDbSystemEndpointModesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Modes: %s. Supported values are: %s.", val, strings.Join(GetDbSystemEndpointModesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingDbSystemEndpointStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDbSystemEndpointStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemEndpointModesEnum Enum with underlying type: string
type DbSystemEndpointModesEnum string

// Set of constants representing the allowable values for DbSystemEndpointModesEnum
const (
	DbSystemEndpointModesRead  DbSystemEndpointModesEnum = "READ"
	DbSystemEndpointModesWrite DbSystemEndpointModesEnum = "WRITE"
)

var mappingDbSystemEndpointModesEnum = map[string]DbSystemEndpointModesEnum{
	"READ":  DbSystemEndpointModesRead,
	"WRITE": DbSystemEndpointModesWrite,
}

var mappingDbSystemEndpointModesEnumLowerCase = map[string]DbSystemEndpointModesEnum{
	"read":  DbSystemEndpointModesRead,
	"write": DbSystemEndpointModesWrite,
}

// GetDbSystemEndpointModesEnumValues Enumerates the set of values for DbSystemEndpointModesEnum
func GetDbSystemEndpointModesEnumValues() []DbSystemEndpointModesEnum {
	values := make([]DbSystemEndpointModesEnum, 0)
	for _, v := range mappingDbSystemEndpointModesEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemEndpointModesEnumStringValues Enumerates the set of values in String for DbSystemEndpointModesEnum
func GetDbSystemEndpointModesEnumStringValues() []string {
	return []string{
		"READ",
		"WRITE",
	}
}

// GetMappingDbSystemEndpointModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemEndpointModesEnum(val string) (DbSystemEndpointModesEnum, bool) {
	enum, ok := mappingDbSystemEndpointModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbSystemEndpointStatusEnum Enum with underlying type: string
type DbSystemEndpointStatusEnum string

// Set of constants representing the allowable values for DbSystemEndpointStatusEnum
const (
	DbSystemEndpointStatusActive   DbSystemEndpointStatusEnum = "ACTIVE"
	DbSystemEndpointStatusInactive DbSystemEndpointStatusEnum = "INACTIVE"
	DbSystemEndpointStatusUpdating DbSystemEndpointStatusEnum = "UPDATING"
)

var mappingDbSystemEndpointStatusEnum = map[string]DbSystemEndpointStatusEnum{
	"ACTIVE":   DbSystemEndpointStatusActive,
	"INACTIVE": DbSystemEndpointStatusInactive,
	"UPDATING": DbSystemEndpointStatusUpdating,
}

var mappingDbSystemEndpointStatusEnumLowerCase = map[string]DbSystemEndpointStatusEnum{
	"active":   DbSystemEndpointStatusActive,
	"inactive": DbSystemEndpointStatusInactive,
	"updating": DbSystemEndpointStatusUpdating,
}

// GetDbSystemEndpointStatusEnumValues Enumerates the set of values for DbSystemEndpointStatusEnum
func GetDbSystemEndpointStatusEnumValues() []DbSystemEndpointStatusEnum {
	values := make([]DbSystemEndpointStatusEnum, 0)
	for _, v := range mappingDbSystemEndpointStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemEndpointStatusEnumStringValues Enumerates the set of values in String for DbSystemEndpointStatusEnum
func GetDbSystemEndpointStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
	}
}

// GetMappingDbSystemEndpointStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemEndpointStatusEnum(val string) (DbSystemEndpointStatusEnum, bool) {
	enum, ok := mappingDbSystemEndpointStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
