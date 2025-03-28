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

// ExternalDbHome The details of an external database home.
type ExternalDbHome struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB home.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the external DB home. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the DB home is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The current lifecycle state of the external DB home.
	LifecycleState ExternalDbHomeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external DB home was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external DB home was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the external DB home.
	ComponentName *string `mandatory:"false" json:"componentName"`

	// The location of the DB home.
	HomeDirectory *string `mandatory:"false" json:"homeDirectory"`

	// The additional details of the DB home defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ExternalDbHome) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDbHome) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDbHomeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDbHomeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDbHomeLifecycleStateEnum Enum with underlying type: string
type ExternalDbHomeLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalDbHomeLifecycleStateEnum
const (
	ExternalDbHomeLifecycleStateCreating ExternalDbHomeLifecycleStateEnum = "CREATING"
	ExternalDbHomeLifecycleStateActive   ExternalDbHomeLifecycleStateEnum = "ACTIVE"
	ExternalDbHomeLifecycleStateInactive ExternalDbHomeLifecycleStateEnum = "INACTIVE"
	ExternalDbHomeLifecycleStateUpdating ExternalDbHomeLifecycleStateEnum = "UPDATING"
	ExternalDbHomeLifecycleStateDeleting ExternalDbHomeLifecycleStateEnum = "DELETING"
	ExternalDbHomeLifecycleStateDeleted  ExternalDbHomeLifecycleStateEnum = "DELETED"
	ExternalDbHomeLifecycleStateFailed   ExternalDbHomeLifecycleStateEnum = "FAILED"
)

var mappingExternalDbHomeLifecycleStateEnum = map[string]ExternalDbHomeLifecycleStateEnum{
	"CREATING": ExternalDbHomeLifecycleStateCreating,
	"ACTIVE":   ExternalDbHomeLifecycleStateActive,
	"INACTIVE": ExternalDbHomeLifecycleStateInactive,
	"UPDATING": ExternalDbHomeLifecycleStateUpdating,
	"DELETING": ExternalDbHomeLifecycleStateDeleting,
	"DELETED":  ExternalDbHomeLifecycleStateDeleted,
	"FAILED":   ExternalDbHomeLifecycleStateFailed,
}

var mappingExternalDbHomeLifecycleStateEnumLowerCase = map[string]ExternalDbHomeLifecycleStateEnum{
	"creating": ExternalDbHomeLifecycleStateCreating,
	"active":   ExternalDbHomeLifecycleStateActive,
	"inactive": ExternalDbHomeLifecycleStateInactive,
	"updating": ExternalDbHomeLifecycleStateUpdating,
	"deleting": ExternalDbHomeLifecycleStateDeleting,
	"deleted":  ExternalDbHomeLifecycleStateDeleted,
	"failed":   ExternalDbHomeLifecycleStateFailed,
}

// GetExternalDbHomeLifecycleStateEnumValues Enumerates the set of values for ExternalDbHomeLifecycleStateEnum
func GetExternalDbHomeLifecycleStateEnumValues() []ExternalDbHomeLifecycleStateEnum {
	values := make([]ExternalDbHomeLifecycleStateEnum, 0)
	for _, v := range mappingExternalDbHomeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbHomeLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalDbHomeLifecycleStateEnum
func GetExternalDbHomeLifecycleStateEnumStringValues() []string {
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

// GetMappingExternalDbHomeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbHomeLifecycleStateEnum(val string) (ExternalDbHomeLifecycleStateEnum, bool) {
	enum, ok := mappingExternalDbHomeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
