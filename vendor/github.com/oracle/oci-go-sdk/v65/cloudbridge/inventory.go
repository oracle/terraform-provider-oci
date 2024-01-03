// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Inventory Description of inventory.
type Inventory struct {

	// Inventory OCID.
	Id *string `mandatory:"true" json:"id"`

	// Inventory display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the inventory.
	LifecycleState InventoryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the tenantId.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when the inventory was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the inventory was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Inventory) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Inventory) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInventoryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInventoryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InventoryLifecycleStateEnum Enum with underlying type: string
type InventoryLifecycleStateEnum string

// Set of constants representing the allowable values for InventoryLifecycleStateEnum
const (
	InventoryLifecycleStateActive   InventoryLifecycleStateEnum = "ACTIVE"
	InventoryLifecycleStateDeleted  InventoryLifecycleStateEnum = "DELETED"
	InventoryLifecycleStateDeleting InventoryLifecycleStateEnum = "DELETING"
	InventoryLifecycleStateCreating InventoryLifecycleStateEnum = "CREATING"
	InventoryLifecycleStateFailed   InventoryLifecycleStateEnum = "FAILED"
)

var mappingInventoryLifecycleStateEnum = map[string]InventoryLifecycleStateEnum{
	"ACTIVE":   InventoryLifecycleStateActive,
	"DELETED":  InventoryLifecycleStateDeleted,
	"DELETING": InventoryLifecycleStateDeleting,
	"CREATING": InventoryLifecycleStateCreating,
	"FAILED":   InventoryLifecycleStateFailed,
}

var mappingInventoryLifecycleStateEnumLowerCase = map[string]InventoryLifecycleStateEnum{
	"active":   InventoryLifecycleStateActive,
	"deleted":  InventoryLifecycleStateDeleted,
	"deleting": InventoryLifecycleStateDeleting,
	"creating": InventoryLifecycleStateCreating,
	"failed":   InventoryLifecycleStateFailed,
}

// GetInventoryLifecycleStateEnumValues Enumerates the set of values for InventoryLifecycleStateEnum
func GetInventoryLifecycleStateEnumValues() []InventoryLifecycleStateEnum {
	values := make([]InventoryLifecycleStateEnum, 0)
	for _, v := range mappingInventoryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInventoryLifecycleStateEnumStringValues Enumerates the set of values in String for InventoryLifecycleStateEnum
func GetInventoryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"DELETING",
		"CREATING",
		"FAILED",
	}
}

// GetMappingInventoryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInventoryLifecycleStateEnum(val string) (InventoryLifecycleStateEnum, bool) {
	enum, ok := mappingInventoryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
