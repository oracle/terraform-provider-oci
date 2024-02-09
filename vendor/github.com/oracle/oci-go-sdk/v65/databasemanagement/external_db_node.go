// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalDbNode The details of an external database node.
type ExternalDbNode struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB node.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the external DB node. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the external DB node.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the DB node is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The current lifecycle state of the external DB node.
	LifecycleState ExternalDbNodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external DB node was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external DB node was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external connector.
	ExternalConnectorId *string `mandatory:"false" json:"externalConnectorId"`

	// The host name for the DB node.
	HostName *string `mandatory:"false" json:"hostName"`

	// The number of CPU cores available on the DB node.
	CpuCoreCount *float32 `mandatory:"false" json:"cpuCoreCount"`

	// The total memory in gigabytes (GB) on the DB node.
	MemorySizeInGBs *float32 `mandatory:"false" json:"memorySizeInGBs"`

	// The additional details of the external DB node defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Name of the domain.
	DomainName *string `mandatory:"false" json:"domainName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExternalDbNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDbNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDbNodeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDbNodeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDbNodeLifecycleStateEnum Enum with underlying type: string
type ExternalDbNodeLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalDbNodeLifecycleStateEnum
const (
	ExternalDbNodeLifecycleStateCreating     ExternalDbNodeLifecycleStateEnum = "CREATING"
	ExternalDbNodeLifecycleStateNotConnected ExternalDbNodeLifecycleStateEnum = "NOT_CONNECTED"
	ExternalDbNodeLifecycleStateActive       ExternalDbNodeLifecycleStateEnum = "ACTIVE"
	ExternalDbNodeLifecycleStateInactive     ExternalDbNodeLifecycleStateEnum = "INACTIVE"
	ExternalDbNodeLifecycleStateUpdating     ExternalDbNodeLifecycleStateEnum = "UPDATING"
	ExternalDbNodeLifecycleStateDeleting     ExternalDbNodeLifecycleStateEnum = "DELETING"
	ExternalDbNodeLifecycleStateDeleted      ExternalDbNodeLifecycleStateEnum = "DELETED"
	ExternalDbNodeLifecycleStateFailed       ExternalDbNodeLifecycleStateEnum = "FAILED"
)

var mappingExternalDbNodeLifecycleStateEnum = map[string]ExternalDbNodeLifecycleStateEnum{
	"CREATING":      ExternalDbNodeLifecycleStateCreating,
	"NOT_CONNECTED": ExternalDbNodeLifecycleStateNotConnected,
	"ACTIVE":        ExternalDbNodeLifecycleStateActive,
	"INACTIVE":      ExternalDbNodeLifecycleStateInactive,
	"UPDATING":      ExternalDbNodeLifecycleStateUpdating,
	"DELETING":      ExternalDbNodeLifecycleStateDeleting,
	"DELETED":       ExternalDbNodeLifecycleStateDeleted,
	"FAILED":        ExternalDbNodeLifecycleStateFailed,
}

var mappingExternalDbNodeLifecycleStateEnumLowerCase = map[string]ExternalDbNodeLifecycleStateEnum{
	"creating":      ExternalDbNodeLifecycleStateCreating,
	"not_connected": ExternalDbNodeLifecycleStateNotConnected,
	"active":        ExternalDbNodeLifecycleStateActive,
	"inactive":      ExternalDbNodeLifecycleStateInactive,
	"updating":      ExternalDbNodeLifecycleStateUpdating,
	"deleting":      ExternalDbNodeLifecycleStateDeleting,
	"deleted":       ExternalDbNodeLifecycleStateDeleted,
	"failed":        ExternalDbNodeLifecycleStateFailed,
}

// GetExternalDbNodeLifecycleStateEnumValues Enumerates the set of values for ExternalDbNodeLifecycleStateEnum
func GetExternalDbNodeLifecycleStateEnumValues() []ExternalDbNodeLifecycleStateEnum {
	values := make([]ExternalDbNodeLifecycleStateEnum, 0)
	for _, v := range mappingExternalDbNodeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbNodeLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalDbNodeLifecycleStateEnum
func GetExternalDbNodeLifecycleStateEnumStringValues() []string {
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

// GetMappingExternalDbNodeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbNodeLifecycleStateEnum(val string) (ExternalDbNodeLifecycleStateEnum, bool) {
	enum, ok := mappingExternalDbNodeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
