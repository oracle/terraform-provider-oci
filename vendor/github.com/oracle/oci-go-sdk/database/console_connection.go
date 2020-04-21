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

// ConsoleConnection The representation of ConsoleConnection
type ConsoleConnection struct {

	// The OCID of the console connection.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment to contain the console connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the database node.
	DbNodeId *string `mandatory:"true" json:"dbNodeId"`

	// The SSH connection string for the console connection.
	ConnectionString *string `mandatory:"true" json:"connectionString"`

	// The SSH public key fingerprint for the console connection.
	Fingerprint *string `mandatory:"true" json:"fingerprint"`

	// The current state of the console connection.
	LifecycleState ConsoleConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m ConsoleConnection) String() string {
	return common.PointerString(m)
}

// ConsoleConnectionLifecycleStateEnum Enum with underlying type: string
type ConsoleConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for ConsoleConnectionLifecycleStateEnum
const (
	ConsoleConnectionLifecycleStateActive   ConsoleConnectionLifecycleStateEnum = "ACTIVE"
	ConsoleConnectionLifecycleStateCreating ConsoleConnectionLifecycleStateEnum = "CREATING"
	ConsoleConnectionLifecycleStateDeleted  ConsoleConnectionLifecycleStateEnum = "DELETED"
	ConsoleConnectionLifecycleStateDeleting ConsoleConnectionLifecycleStateEnum = "DELETING"
	ConsoleConnectionLifecycleStateFailed   ConsoleConnectionLifecycleStateEnum = "FAILED"
)

var mappingConsoleConnectionLifecycleState = map[string]ConsoleConnectionLifecycleStateEnum{
	"ACTIVE":   ConsoleConnectionLifecycleStateActive,
	"CREATING": ConsoleConnectionLifecycleStateCreating,
	"DELETED":  ConsoleConnectionLifecycleStateDeleted,
	"DELETING": ConsoleConnectionLifecycleStateDeleting,
	"FAILED":   ConsoleConnectionLifecycleStateFailed,
}

// GetConsoleConnectionLifecycleStateEnumValues Enumerates the set of values for ConsoleConnectionLifecycleStateEnum
func GetConsoleConnectionLifecycleStateEnumValues() []ConsoleConnectionLifecycleStateEnum {
	values := make([]ConsoleConnectionLifecycleStateEnum, 0)
	for _, v := range mappingConsoleConnectionLifecycleState {
		values = append(values, v)
	}
	return values
}
