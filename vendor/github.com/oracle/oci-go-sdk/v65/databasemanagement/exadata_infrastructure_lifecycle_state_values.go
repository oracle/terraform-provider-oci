// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadataInfrastructureLifecycleStateValues the lifecycle state values for the Exadata infrastructure.
type ExadataInfrastructureLifecycleStateValues struct {

	// The current lifecycle state of the database resource.
	LifecycleState ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ExadataInfrastructureLifecycleStateValues) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInfrastructureLifecycleStateValues) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataInfrastructureLifecycleStateValuesLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadataInfrastructureLifecycleStateValuesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum Enum with underlying type: string
type ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum string

// Set of constants representing the allowable values for ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum
const (
	ExadataInfrastructureLifecycleStateValuesLifecycleStateCreating ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum = "CREATING"
	ExadataInfrastructureLifecycleStateValuesLifecycleStateActive   ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum = "ACTIVE"
	ExadataInfrastructureLifecycleStateValuesLifecycleStateInactive ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum = "INACTIVE"
	ExadataInfrastructureLifecycleStateValuesLifecycleStateUpdating ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum = "UPDATING"
	ExadataInfrastructureLifecycleStateValuesLifecycleStateDeleting ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum = "DELETING"
	ExadataInfrastructureLifecycleStateValuesLifecycleStateDeleted  ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum = "DELETED"
	ExadataInfrastructureLifecycleStateValuesLifecycleStateFailed   ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum = "FAILED"
)

var mappingExadataInfrastructureLifecycleStateValuesLifecycleStateEnum = map[string]ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum{
	"CREATING": ExadataInfrastructureLifecycleStateValuesLifecycleStateCreating,
	"ACTIVE":   ExadataInfrastructureLifecycleStateValuesLifecycleStateActive,
	"INACTIVE": ExadataInfrastructureLifecycleStateValuesLifecycleStateInactive,
	"UPDATING": ExadataInfrastructureLifecycleStateValuesLifecycleStateUpdating,
	"DELETING": ExadataInfrastructureLifecycleStateValuesLifecycleStateDeleting,
	"DELETED":  ExadataInfrastructureLifecycleStateValuesLifecycleStateDeleted,
	"FAILED":   ExadataInfrastructureLifecycleStateValuesLifecycleStateFailed,
}

var mappingExadataInfrastructureLifecycleStateValuesLifecycleStateEnumLowerCase = map[string]ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum{
	"creating": ExadataInfrastructureLifecycleStateValuesLifecycleStateCreating,
	"active":   ExadataInfrastructureLifecycleStateValuesLifecycleStateActive,
	"inactive": ExadataInfrastructureLifecycleStateValuesLifecycleStateInactive,
	"updating": ExadataInfrastructureLifecycleStateValuesLifecycleStateUpdating,
	"deleting": ExadataInfrastructureLifecycleStateValuesLifecycleStateDeleting,
	"deleted":  ExadataInfrastructureLifecycleStateValuesLifecycleStateDeleted,
	"failed":   ExadataInfrastructureLifecycleStateValuesLifecycleStateFailed,
}

// GetExadataInfrastructureLifecycleStateValuesLifecycleStateEnumValues Enumerates the set of values for ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum
func GetExadataInfrastructureLifecycleStateValuesLifecycleStateEnumValues() []ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum {
	values := make([]ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum, 0)
	for _, v := range mappingExadataInfrastructureLifecycleStateValuesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureLifecycleStateValuesLifecycleStateEnumStringValues Enumerates the set of values in String for ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum
func GetExadataInfrastructureLifecycleStateValuesLifecycleStateEnumStringValues() []string {
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

// GetMappingExadataInfrastructureLifecycleStateValuesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureLifecycleStateValuesLifecycleStateEnum(val string) (ExadataInfrastructureLifecycleStateValuesLifecycleStateEnum, bool) {
	enum, ok := mappingExadataInfrastructureLifecycleStateValuesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
