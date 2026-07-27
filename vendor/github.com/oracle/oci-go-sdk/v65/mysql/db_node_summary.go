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

// DbNodeSummary A summary of a DB node in a shared-storage DB cluster.
type DbNodeSummary struct {

	// The OCID of the DB node.
	Id *string `mandatory:"true" json:"id"`

	// Name of the DB node. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the MySQL instance that forms a DB node of the shared-storage DB cluster.
	LifecycleState DbNodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the DB node was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the DB node was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the availability domain for the DB node to be located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Description of the DB node.
	Description *string `mandatory:"false" json:"description"`

	// Preference of a DB node as a potential failover target.
	// This is an integer  that defines the relative weight/priority
	// of a particular DB node versus another, where the lower the number,
	// the higher the preference for that DB node to become a new primary on a failover.
	PromotionTier *int `mandatory:"false" json:"promotionTier"`

	ReadEndpoint *DbNodeReadEndpoint `mandatory:"false" json:"readEndpoint"`

	// Role of a DB node within a shared-storage DB cluster.
	// The PRIMARY DB node handles writes and reads.
	// A SECONDARY DB node serves the read traffic and can be promoted to a PRIMARY.
	Role DbNodeSummaryRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m DbNodeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbNodeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbNodeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbNodeLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDbNodeSummaryRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDbNodeSummaryRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbNodeSummaryRoleEnum Enum with underlying type: string
type DbNodeSummaryRoleEnum string

// Set of constants representing the allowable values for DbNodeSummaryRoleEnum
const (
	DbNodeSummaryRolePrimary   DbNodeSummaryRoleEnum = "PRIMARY"
	DbNodeSummaryRoleSecondary DbNodeSummaryRoleEnum = "SECONDARY"
)

var mappingDbNodeSummaryRoleEnum = map[string]DbNodeSummaryRoleEnum{
	"PRIMARY":   DbNodeSummaryRolePrimary,
	"SECONDARY": DbNodeSummaryRoleSecondary,
}

var mappingDbNodeSummaryRoleEnumLowerCase = map[string]DbNodeSummaryRoleEnum{
	"primary":   DbNodeSummaryRolePrimary,
	"secondary": DbNodeSummaryRoleSecondary,
}

// GetDbNodeSummaryRoleEnumValues Enumerates the set of values for DbNodeSummaryRoleEnum
func GetDbNodeSummaryRoleEnumValues() []DbNodeSummaryRoleEnum {
	values := make([]DbNodeSummaryRoleEnum, 0)
	for _, v := range mappingDbNodeSummaryRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDbNodeSummaryRoleEnumStringValues Enumerates the set of values in String for DbNodeSummaryRoleEnum
func GetDbNodeSummaryRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"SECONDARY",
	}
}

// GetMappingDbNodeSummaryRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbNodeSummaryRoleEnum(val string) (DbNodeSummaryRoleEnum, bool) {
	enum, ok := mappingDbNodeSummaryRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
