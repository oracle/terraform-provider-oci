// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DbNode A DB node of a shared-storage DB cluster. A DB node comprises of a single MySQL instance.
type DbNode struct {

	// The OCID of the DB node.
	Id *string `mandatory:"true" json:"id"`

	// The name of the availability domain the DB node is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// Name for the DB node. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the MySQL instance that forms a DB node of the shared-storage DB cluster
	LifecycleState DbNodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Intended MySQL version used by the DB node.
	// This version is directly controlled by the consumer via the create and update shared-storage DB cluster API calls.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// Preference of a DB node as a potential failover target.
	// This is an integer that defines the relative weight/priority
	// of a particular DB node versus another, where lower the number,
	// the higher the preference for that DB node to become a new primary on a failover.
	PromotionTier *int `mandatory:"true" json:"promotionTier"`

	// Role of a DB node within a shared-storage DB cluster.
	// The PRIMARY DB node can handle writes and reads.
	// A SECONDARY DB node can serve read traffic and can be promoted to a PRIMARY.
	Role DbNodeRoleEnum `mandatory:"true" json:"role"`

	// The date and time the DB node was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DB node was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Description of the DB node.
	Description *string `mandatory:"false" json:"description"`

	// Actual version of MySQL version used by the DB node.
	// This version is controlled by the service and could be different from the intended MySQL version (mysqlVersion)
	// for the DB node as a side effect of service maintenance events.
	CurrentMysqlVersion *string `mandatory:"false" json:"currentMysqlVersion"`

	ReadEndpoint *DbNodeReadEndpoint `mandatory:"false" json:"readEndpoint"`
}

func (m DbNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbNodeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbNodeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbNodeRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDbNodeRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbNodeLifecycleStateEnum Enum with underlying type: string
type DbNodeLifecycleStateEnum string

// Set of constants representing the allowable values for DbNodeLifecycleStateEnum
const (
	DbNodeLifecycleStateCreating DbNodeLifecycleStateEnum = "CREATING"
	DbNodeLifecycleStateActive   DbNodeLifecycleStateEnum = "ACTIVE"
	DbNodeLifecycleStateInactive DbNodeLifecycleStateEnum = "INACTIVE"
	DbNodeLifecycleStateUpdating DbNodeLifecycleStateEnum = "UPDATING"
	DbNodeLifecycleStateDeleting DbNodeLifecycleStateEnum = "DELETING"
	DbNodeLifecycleStateDeleted  DbNodeLifecycleStateEnum = "DELETED"
	DbNodeLifecycleStateFailed   DbNodeLifecycleStateEnum = "FAILED"
)

var mappingDbNodeLifecycleStateEnum = map[string]DbNodeLifecycleStateEnum{
	"CREATING": DbNodeLifecycleStateCreating,
	"ACTIVE":   DbNodeLifecycleStateActive,
	"INACTIVE": DbNodeLifecycleStateInactive,
	"UPDATING": DbNodeLifecycleStateUpdating,
	"DELETING": DbNodeLifecycleStateDeleting,
	"DELETED":  DbNodeLifecycleStateDeleted,
	"FAILED":   DbNodeLifecycleStateFailed,
}

var mappingDbNodeLifecycleStateEnumLowerCase = map[string]DbNodeLifecycleStateEnum{
	"creating": DbNodeLifecycleStateCreating,
	"active":   DbNodeLifecycleStateActive,
	"inactive": DbNodeLifecycleStateInactive,
	"updating": DbNodeLifecycleStateUpdating,
	"deleting": DbNodeLifecycleStateDeleting,
	"deleted":  DbNodeLifecycleStateDeleted,
	"failed":   DbNodeLifecycleStateFailed,
}

// GetDbNodeLifecycleStateEnumValues Enumerates the set of values for DbNodeLifecycleStateEnum
func GetDbNodeLifecycleStateEnumValues() []DbNodeLifecycleStateEnum {
	values := make([]DbNodeLifecycleStateEnum, 0)
	for _, v := range mappingDbNodeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbNodeLifecycleStateEnumStringValues Enumerates the set of values in String for DbNodeLifecycleStateEnum
func GetDbNodeLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDbNodeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbNodeLifecycleStateEnum(val string) (DbNodeLifecycleStateEnum, bool) {
	enum, ok := mappingDbNodeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbNodeRoleEnum Enum with underlying type: string
type DbNodeRoleEnum string

// Set of constants representing the allowable values for DbNodeRoleEnum
const (
	DbNodeRolePrimary   DbNodeRoleEnum = "PRIMARY"
	DbNodeRoleSecondary DbNodeRoleEnum = "SECONDARY"
)

var mappingDbNodeRoleEnum = map[string]DbNodeRoleEnum{
	"PRIMARY":   DbNodeRolePrimary,
	"SECONDARY": DbNodeRoleSecondary,
}

var mappingDbNodeRoleEnumLowerCase = map[string]DbNodeRoleEnum{
	"primary":   DbNodeRolePrimary,
	"secondary": DbNodeRoleSecondary,
}

// GetDbNodeRoleEnumValues Enumerates the set of values for DbNodeRoleEnum
func GetDbNodeRoleEnumValues() []DbNodeRoleEnum {
	values := make([]DbNodeRoleEnum, 0)
	for _, v := range mappingDbNodeRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDbNodeRoleEnumStringValues Enumerates the set of values in String for DbNodeRoleEnum
func GetDbNodeRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"SECONDARY",
	}
}

// GetMappingDbNodeRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbNodeRoleEnum(val string) (DbNodeRoleEnum, bool) {
	enum, ok := mappingDbNodeRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
