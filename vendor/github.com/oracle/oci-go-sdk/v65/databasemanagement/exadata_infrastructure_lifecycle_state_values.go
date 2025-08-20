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

	// The current lifecycle state of the Exadata infrastructure resource.
	State ExadataInfrastructureLifecycleStateValuesStateEnum `mandatory:"false" json:"state,omitempty"`
}

func (m ExadataInfrastructureLifecycleStateValues) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInfrastructureLifecycleStateValues) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataInfrastructureLifecycleStateValuesStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetExadataInfrastructureLifecycleStateValuesStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataInfrastructureLifecycleStateValuesStateEnum Enum with underlying type: string
type ExadataInfrastructureLifecycleStateValuesStateEnum string

// Set of constants representing the allowable values for ExadataInfrastructureLifecycleStateValuesStateEnum
const (
	ExadataInfrastructureLifecycleStateValuesStateCreating ExadataInfrastructureLifecycleStateValuesStateEnum = "CREATING"
	ExadataInfrastructureLifecycleStateValuesStateActive   ExadataInfrastructureLifecycleStateValuesStateEnum = "ACTIVE"
	ExadataInfrastructureLifecycleStateValuesStateInactive ExadataInfrastructureLifecycleStateValuesStateEnum = "INACTIVE"
	ExadataInfrastructureLifecycleStateValuesStateUpdating ExadataInfrastructureLifecycleStateValuesStateEnum = "UPDATING"
	ExadataInfrastructureLifecycleStateValuesStateDeleting ExadataInfrastructureLifecycleStateValuesStateEnum = "DELETING"
	ExadataInfrastructureLifecycleStateValuesStateDeleted  ExadataInfrastructureLifecycleStateValuesStateEnum = "DELETED"
	ExadataInfrastructureLifecycleStateValuesStateFailed   ExadataInfrastructureLifecycleStateValuesStateEnum = "FAILED"
	ExadataInfrastructureLifecycleStateValuesStateUnknown  ExadataInfrastructureLifecycleStateValuesStateEnum = "UNKNOWN"
)

var mappingExadataInfrastructureLifecycleStateValuesStateEnum = map[string]ExadataInfrastructureLifecycleStateValuesStateEnum{
	"CREATING": ExadataInfrastructureLifecycleStateValuesStateCreating,
	"ACTIVE":   ExadataInfrastructureLifecycleStateValuesStateActive,
	"INACTIVE": ExadataInfrastructureLifecycleStateValuesStateInactive,
	"UPDATING": ExadataInfrastructureLifecycleStateValuesStateUpdating,
	"DELETING": ExadataInfrastructureLifecycleStateValuesStateDeleting,
	"DELETED":  ExadataInfrastructureLifecycleStateValuesStateDeleted,
	"FAILED":   ExadataInfrastructureLifecycleStateValuesStateFailed,
	"UNKNOWN":  ExadataInfrastructureLifecycleStateValuesStateUnknown,
}

var mappingExadataInfrastructureLifecycleStateValuesStateEnumLowerCase = map[string]ExadataInfrastructureLifecycleStateValuesStateEnum{
	"creating": ExadataInfrastructureLifecycleStateValuesStateCreating,
	"active":   ExadataInfrastructureLifecycleStateValuesStateActive,
	"inactive": ExadataInfrastructureLifecycleStateValuesStateInactive,
	"updating": ExadataInfrastructureLifecycleStateValuesStateUpdating,
	"deleting": ExadataInfrastructureLifecycleStateValuesStateDeleting,
	"deleted":  ExadataInfrastructureLifecycleStateValuesStateDeleted,
	"failed":   ExadataInfrastructureLifecycleStateValuesStateFailed,
	"unknown":  ExadataInfrastructureLifecycleStateValuesStateUnknown,
}

// GetExadataInfrastructureLifecycleStateValuesStateEnumValues Enumerates the set of values for ExadataInfrastructureLifecycleStateValuesStateEnum
func GetExadataInfrastructureLifecycleStateValuesStateEnumValues() []ExadataInfrastructureLifecycleStateValuesStateEnum {
	values := make([]ExadataInfrastructureLifecycleStateValuesStateEnum, 0)
	for _, v := range mappingExadataInfrastructureLifecycleStateValuesStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureLifecycleStateValuesStateEnumStringValues Enumerates the set of values in String for ExadataInfrastructureLifecycleStateValuesStateEnum
func GetExadataInfrastructureLifecycleStateValuesStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"UNKNOWN",
	}
}

// GetMappingExadataInfrastructureLifecycleStateValuesStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureLifecycleStateValuesStateEnum(val string) (ExadataInfrastructureLifecycleStateValuesStateEnum, bool) {
	enum, ok := mappingExadataInfrastructureLifecycleStateValuesStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
