// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbInstance Information about a database instance node.
type DbInstance struct {

	// A unique identifier for the database instance node. Immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The availability domain in which the database instance node is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The current state of the database instance node.
	LifecycleState DbInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time that the database instance node was created, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly display name for the database instance node. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the database instance node.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time that the database instance node was updated, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m DbInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbInstanceLifecycleStateEnum Enum with underlying type: string
type DbInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for DbInstanceLifecycleStateEnum
const (
	DbInstanceLifecycleStateCreating DbInstanceLifecycleStateEnum = "CREATING"
	DbInstanceLifecycleStateUpdating DbInstanceLifecycleStateEnum = "UPDATING"
	DbInstanceLifecycleStateActive   DbInstanceLifecycleStateEnum = "ACTIVE"
	DbInstanceLifecycleStateInactive DbInstanceLifecycleStateEnum = "INACTIVE"
	DbInstanceLifecycleStateDeleting DbInstanceLifecycleStateEnum = "DELETING"
	DbInstanceLifecycleStateDeleted  DbInstanceLifecycleStateEnum = "DELETED"
	DbInstanceLifecycleStateFailed   DbInstanceLifecycleStateEnum = "FAILED"
)

var mappingDbInstanceLifecycleStateEnum = map[string]DbInstanceLifecycleStateEnum{
	"CREATING": DbInstanceLifecycleStateCreating,
	"UPDATING": DbInstanceLifecycleStateUpdating,
	"ACTIVE":   DbInstanceLifecycleStateActive,
	"INACTIVE": DbInstanceLifecycleStateInactive,
	"DELETING": DbInstanceLifecycleStateDeleting,
	"DELETED":  DbInstanceLifecycleStateDeleted,
	"FAILED":   DbInstanceLifecycleStateFailed,
}

var mappingDbInstanceLifecycleStateEnumLowerCase = map[string]DbInstanceLifecycleStateEnum{
	"creating": DbInstanceLifecycleStateCreating,
	"updating": DbInstanceLifecycleStateUpdating,
	"active":   DbInstanceLifecycleStateActive,
	"inactive": DbInstanceLifecycleStateInactive,
	"deleting": DbInstanceLifecycleStateDeleting,
	"deleted":  DbInstanceLifecycleStateDeleted,
	"failed":   DbInstanceLifecycleStateFailed,
}

// GetDbInstanceLifecycleStateEnumValues Enumerates the set of values for DbInstanceLifecycleStateEnum
func GetDbInstanceLifecycleStateEnumValues() []DbInstanceLifecycleStateEnum {
	values := make([]DbInstanceLifecycleStateEnum, 0)
	for _, v := range mappingDbInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for DbInstanceLifecycleStateEnum
func GetDbInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDbInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbInstanceLifecycleStateEnum(val string) (DbInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingDbInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
