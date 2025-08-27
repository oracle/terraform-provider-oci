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

// CloudDbHome The details of a cloud database home.
type CloudDbHome struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB home.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the cloud DB home. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the DB home is a part of.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The current lifecycle state of the cloud DB home.
	LifecycleState CloudDbHomeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the cloud DB home was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the cloud DB home was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB home in DBaas service.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// The name of the cloud DB home.
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

func (m CloudDbHome) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudDbHome) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudDbHomeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudDbHomeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudDbHomeLifecycleStateEnum Enum with underlying type: string
type CloudDbHomeLifecycleStateEnum string

// Set of constants representing the allowable values for CloudDbHomeLifecycleStateEnum
const (
	CloudDbHomeLifecycleStateCreating CloudDbHomeLifecycleStateEnum = "CREATING"
	CloudDbHomeLifecycleStateActive   CloudDbHomeLifecycleStateEnum = "ACTIVE"
	CloudDbHomeLifecycleStateInactive CloudDbHomeLifecycleStateEnum = "INACTIVE"
	CloudDbHomeLifecycleStateUpdating CloudDbHomeLifecycleStateEnum = "UPDATING"
	CloudDbHomeLifecycleStateDeleting CloudDbHomeLifecycleStateEnum = "DELETING"
	CloudDbHomeLifecycleStateDeleted  CloudDbHomeLifecycleStateEnum = "DELETED"
	CloudDbHomeLifecycleStateFailed   CloudDbHomeLifecycleStateEnum = "FAILED"
)

var mappingCloudDbHomeLifecycleStateEnum = map[string]CloudDbHomeLifecycleStateEnum{
	"CREATING": CloudDbHomeLifecycleStateCreating,
	"ACTIVE":   CloudDbHomeLifecycleStateActive,
	"INACTIVE": CloudDbHomeLifecycleStateInactive,
	"UPDATING": CloudDbHomeLifecycleStateUpdating,
	"DELETING": CloudDbHomeLifecycleStateDeleting,
	"DELETED":  CloudDbHomeLifecycleStateDeleted,
	"FAILED":   CloudDbHomeLifecycleStateFailed,
}

var mappingCloudDbHomeLifecycleStateEnumLowerCase = map[string]CloudDbHomeLifecycleStateEnum{
	"creating": CloudDbHomeLifecycleStateCreating,
	"active":   CloudDbHomeLifecycleStateActive,
	"inactive": CloudDbHomeLifecycleStateInactive,
	"updating": CloudDbHomeLifecycleStateUpdating,
	"deleting": CloudDbHomeLifecycleStateDeleting,
	"deleted":  CloudDbHomeLifecycleStateDeleted,
	"failed":   CloudDbHomeLifecycleStateFailed,
}

// GetCloudDbHomeLifecycleStateEnumValues Enumerates the set of values for CloudDbHomeLifecycleStateEnum
func GetCloudDbHomeLifecycleStateEnumValues() []CloudDbHomeLifecycleStateEnum {
	values := make([]CloudDbHomeLifecycleStateEnum, 0)
	for _, v := range mappingCloudDbHomeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbHomeLifecycleStateEnumStringValues Enumerates the set of values in String for CloudDbHomeLifecycleStateEnum
func GetCloudDbHomeLifecycleStateEnumStringValues() []string {
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

// GetMappingCloudDbHomeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbHomeLifecycleStateEnum(val string) (CloudDbHomeLifecycleStateEnum, bool) {
	enum, ok := mappingCloudDbHomeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
