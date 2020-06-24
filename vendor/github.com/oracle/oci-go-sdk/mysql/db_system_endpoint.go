// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
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

// DbSystemEndpointModesEnum Enum with underlying type: string
type DbSystemEndpointModesEnum string

// Set of constants representing the allowable values for DbSystemEndpointModesEnum
const (
	DbSystemEndpointModesRead  DbSystemEndpointModesEnum = "READ"
	DbSystemEndpointModesWrite DbSystemEndpointModesEnum = "WRITE"
)

var mappingDbSystemEndpointModes = map[string]DbSystemEndpointModesEnum{
	"READ":  DbSystemEndpointModesRead,
	"WRITE": DbSystemEndpointModesWrite,
}

// GetDbSystemEndpointModesEnumValues Enumerates the set of values for DbSystemEndpointModesEnum
func GetDbSystemEndpointModesEnumValues() []DbSystemEndpointModesEnum {
	values := make([]DbSystemEndpointModesEnum, 0)
	for _, v := range mappingDbSystemEndpointModes {
		values = append(values, v)
	}
	return values
}

// DbSystemEndpointStatusEnum Enum with underlying type: string
type DbSystemEndpointStatusEnum string

// Set of constants representing the allowable values for DbSystemEndpointStatusEnum
const (
	DbSystemEndpointStatusActive   DbSystemEndpointStatusEnum = "ACTIVE"
	DbSystemEndpointStatusInactive DbSystemEndpointStatusEnum = "INACTIVE"
	DbSystemEndpointStatusUpdating DbSystemEndpointStatusEnum = "UPDATING"
)

var mappingDbSystemEndpointStatus = map[string]DbSystemEndpointStatusEnum{
	"ACTIVE":   DbSystemEndpointStatusActive,
	"INACTIVE": DbSystemEndpointStatusInactive,
	"UPDATING": DbSystemEndpointStatusUpdating,
}

// GetDbSystemEndpointStatusEnumValues Enumerates the set of values for DbSystemEndpointStatusEnum
func GetDbSystemEndpointStatusEnumValues() []DbSystemEndpointStatusEnum {
	values := make([]DbSystemEndpointStatusEnum, 0)
	for _, v := range mappingDbSystemEndpointStatus {
		values = append(values, v)
	}
	return values
}
