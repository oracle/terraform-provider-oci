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

// CloudDbNode The details of a cloud database node.
type CloudDbNode struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the cloud DB node. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the cloud DB node.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the DB node is a part of.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The current lifecycle state of the cloud DB node.
	LifecycleState CloudDbNodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the cloud DB node was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the cloud DB node was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node in DBaas service.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
	CloudConnectorId *string `mandatory:"false" json:"cloudConnectorId"`

	// The host name for the DB node.
	HostName *string `mandatory:"false" json:"hostName"`

	// The number of CPU cores available on the DB node.
	CpuCoreCount *float32 `mandatory:"false" json:"cpuCoreCount"`

	// The total memory in gigabytes (GB) on the DB node.
	MemorySizeInGBs *float32 `mandatory:"false" json:"memorySizeInGBs"`

	// The additional details of the cloud DB node defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Name of the domain.
	DomainName *string `mandatory:"false" json:"domainName"`

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

func (m CloudDbNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudDbNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudDbNodeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudDbNodeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudDbNodeLifecycleStateEnum Enum with underlying type: string
type CloudDbNodeLifecycleStateEnum string

// Set of constants representing the allowable values for CloudDbNodeLifecycleStateEnum
const (
	CloudDbNodeLifecycleStateCreating     CloudDbNodeLifecycleStateEnum = "CREATING"
	CloudDbNodeLifecycleStateNotConnected CloudDbNodeLifecycleStateEnum = "NOT_CONNECTED"
	CloudDbNodeLifecycleStateActive       CloudDbNodeLifecycleStateEnum = "ACTIVE"
	CloudDbNodeLifecycleStateInactive     CloudDbNodeLifecycleStateEnum = "INACTIVE"
	CloudDbNodeLifecycleStateUpdating     CloudDbNodeLifecycleStateEnum = "UPDATING"
	CloudDbNodeLifecycleStateDeleting     CloudDbNodeLifecycleStateEnum = "DELETING"
	CloudDbNodeLifecycleStateDeleted      CloudDbNodeLifecycleStateEnum = "DELETED"
	CloudDbNodeLifecycleStateFailed       CloudDbNodeLifecycleStateEnum = "FAILED"
)

var mappingCloudDbNodeLifecycleStateEnum = map[string]CloudDbNodeLifecycleStateEnum{
	"CREATING":      CloudDbNodeLifecycleStateCreating,
	"NOT_CONNECTED": CloudDbNodeLifecycleStateNotConnected,
	"ACTIVE":        CloudDbNodeLifecycleStateActive,
	"INACTIVE":      CloudDbNodeLifecycleStateInactive,
	"UPDATING":      CloudDbNodeLifecycleStateUpdating,
	"DELETING":      CloudDbNodeLifecycleStateDeleting,
	"DELETED":       CloudDbNodeLifecycleStateDeleted,
	"FAILED":        CloudDbNodeLifecycleStateFailed,
}

var mappingCloudDbNodeLifecycleStateEnumLowerCase = map[string]CloudDbNodeLifecycleStateEnum{
	"creating":      CloudDbNodeLifecycleStateCreating,
	"not_connected": CloudDbNodeLifecycleStateNotConnected,
	"active":        CloudDbNodeLifecycleStateActive,
	"inactive":      CloudDbNodeLifecycleStateInactive,
	"updating":      CloudDbNodeLifecycleStateUpdating,
	"deleting":      CloudDbNodeLifecycleStateDeleting,
	"deleted":       CloudDbNodeLifecycleStateDeleted,
	"failed":        CloudDbNodeLifecycleStateFailed,
}

// GetCloudDbNodeLifecycleStateEnumValues Enumerates the set of values for CloudDbNodeLifecycleStateEnum
func GetCloudDbNodeLifecycleStateEnumValues() []CloudDbNodeLifecycleStateEnum {
	values := make([]CloudDbNodeLifecycleStateEnum, 0)
	for _, v := range mappingCloudDbNodeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbNodeLifecycleStateEnumStringValues Enumerates the set of values in String for CloudDbNodeLifecycleStateEnum
func GetCloudDbNodeLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NOT_CONNECTED",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCloudDbNodeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbNodeLifecycleStateEnum(val string) (CloudDbNodeLifecycleStateEnum, bool) {
	enum, ok := mappingCloudDbNodeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
