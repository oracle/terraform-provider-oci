// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccCustomerGroup Details of the customer group resource.
type OccCustomerGroup struct {

	// The OCID of the customer group.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the customer group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the customer group.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// To determine whether the customer group is enabled/disabled.
	Status OccCustomerGroupStatusEnum `mandatory:"true" json:"status"`

	// The current lifecycle state of the resource.
	LifecycleState OccCustomerGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A list containing all the customers that belong to this customer group
	CustomersList []OccCustomer `mandatory:"true" json:"customersList"`

	// The description about the customer group.
	Description *string `mandatory:"false" json:"description"`

	// The time when the customer group was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time when the customer group was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccCustomerGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccCustomerGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccCustomerGroupStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOccCustomerGroupStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccCustomerGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccCustomerGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccCustomerGroupStatusEnum Enum with underlying type: string
type OccCustomerGroupStatusEnum string

// Set of constants representing the allowable values for OccCustomerGroupStatusEnum
const (
	OccCustomerGroupStatusEnabled  OccCustomerGroupStatusEnum = "ENABLED"
	OccCustomerGroupStatusDisabled OccCustomerGroupStatusEnum = "DISABLED"
)

var mappingOccCustomerGroupStatusEnum = map[string]OccCustomerGroupStatusEnum{
	"ENABLED":  OccCustomerGroupStatusEnabled,
	"DISABLED": OccCustomerGroupStatusDisabled,
}

var mappingOccCustomerGroupStatusEnumLowerCase = map[string]OccCustomerGroupStatusEnum{
	"enabled":  OccCustomerGroupStatusEnabled,
	"disabled": OccCustomerGroupStatusDisabled,
}

// GetOccCustomerGroupStatusEnumValues Enumerates the set of values for OccCustomerGroupStatusEnum
func GetOccCustomerGroupStatusEnumValues() []OccCustomerGroupStatusEnum {
	values := make([]OccCustomerGroupStatusEnum, 0)
	for _, v := range mappingOccCustomerGroupStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOccCustomerGroupStatusEnumStringValues Enumerates the set of values in String for OccCustomerGroupStatusEnum
func GetOccCustomerGroupStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingOccCustomerGroupStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccCustomerGroupStatusEnum(val string) (OccCustomerGroupStatusEnum, bool) {
	enum, ok := mappingOccCustomerGroupStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccCustomerGroupLifecycleStateEnum Enum with underlying type: string
type OccCustomerGroupLifecycleStateEnum string

// Set of constants representing the allowable values for OccCustomerGroupLifecycleStateEnum
const (
	OccCustomerGroupLifecycleStateCreating OccCustomerGroupLifecycleStateEnum = "CREATING"
	OccCustomerGroupLifecycleStateUpdating OccCustomerGroupLifecycleStateEnum = "UPDATING"
	OccCustomerGroupLifecycleStateActive   OccCustomerGroupLifecycleStateEnum = "ACTIVE"
	OccCustomerGroupLifecycleStateDeleting OccCustomerGroupLifecycleStateEnum = "DELETING"
	OccCustomerGroupLifecycleStateDeleted  OccCustomerGroupLifecycleStateEnum = "DELETED"
	OccCustomerGroupLifecycleStateFailed   OccCustomerGroupLifecycleStateEnum = "FAILED"
)

var mappingOccCustomerGroupLifecycleStateEnum = map[string]OccCustomerGroupLifecycleStateEnum{
	"CREATING": OccCustomerGroupLifecycleStateCreating,
	"UPDATING": OccCustomerGroupLifecycleStateUpdating,
	"ACTIVE":   OccCustomerGroupLifecycleStateActive,
	"DELETING": OccCustomerGroupLifecycleStateDeleting,
	"DELETED":  OccCustomerGroupLifecycleStateDeleted,
	"FAILED":   OccCustomerGroupLifecycleStateFailed,
}

var mappingOccCustomerGroupLifecycleStateEnumLowerCase = map[string]OccCustomerGroupLifecycleStateEnum{
	"creating": OccCustomerGroupLifecycleStateCreating,
	"updating": OccCustomerGroupLifecycleStateUpdating,
	"active":   OccCustomerGroupLifecycleStateActive,
	"deleting": OccCustomerGroupLifecycleStateDeleting,
	"deleted":  OccCustomerGroupLifecycleStateDeleted,
	"failed":   OccCustomerGroupLifecycleStateFailed,
}

// GetOccCustomerGroupLifecycleStateEnumValues Enumerates the set of values for OccCustomerGroupLifecycleStateEnum
func GetOccCustomerGroupLifecycleStateEnumValues() []OccCustomerGroupLifecycleStateEnum {
	values := make([]OccCustomerGroupLifecycleStateEnum, 0)
	for _, v := range mappingOccCustomerGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccCustomerGroupLifecycleStateEnumStringValues Enumerates the set of values in String for OccCustomerGroupLifecycleStateEnum
func GetOccCustomerGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOccCustomerGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccCustomerGroupLifecycleStateEnum(val string) (OccCustomerGroupLifecycleStateEnum, bool) {
	enum, ok := mappingOccCustomerGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
