// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InventoryRecord Information about an inventory target.
type InventoryRecord struct {

	// The OCID of the Inventory target.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Name of the target
	TargetName *string `mandatory:"true" json:"targetName"`

	// Version of the product on the target
	Version *string `mandatory:"true" json:"version"`

	// OCID of the resource associated with the target
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the product installed at the target path
	TargetProductName *string `mandatory:"true" json:"targetProductName"`

	// OS installed on the resource associated with the target
	OsType *string `mandatory:"true" json:"osType"`

	// Architecture of the resource associated with the target
	Architecture *string `mandatory:"true" json:"architecture"`

	// List of target properties
	Properties []InventoryRecordProperty `mandatory:"true" json:"properties"`

	// List of target components
	Components []InventoryRecordComponent `mandatory:"true" json:"components"`

	// List of details on the patches currently installed on the target
	InstalledPatches []InventoryRecordPatchDetails `mandatory:"true" json:"installedPatches"`

	// Name of the resource associated with the target
	TargetResourceName *string `mandatory:"false" json:"targetResourceName"`

	// OCID of the product installed at the target path
	TargetProductId *string `mandatory:"false" json:"targetProductId"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Inventory target.
	LifecycleState InventoryRecordLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m InventoryRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InventoryRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInventoryRecordLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInventoryRecordLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InventoryRecordLifecycleStateEnum Enum with underlying type: string
type InventoryRecordLifecycleStateEnum string

// Set of constants representing the allowable values for InventoryRecordLifecycleStateEnum
const (
	InventoryRecordLifecycleStateActive InventoryRecordLifecycleStateEnum = "ACTIVE"
)

var mappingInventoryRecordLifecycleStateEnum = map[string]InventoryRecordLifecycleStateEnum{
	"ACTIVE": InventoryRecordLifecycleStateActive,
}

var mappingInventoryRecordLifecycleStateEnumLowerCase = map[string]InventoryRecordLifecycleStateEnum{
	"active": InventoryRecordLifecycleStateActive,
}

// GetInventoryRecordLifecycleStateEnumValues Enumerates the set of values for InventoryRecordLifecycleStateEnum
func GetInventoryRecordLifecycleStateEnumValues() []InventoryRecordLifecycleStateEnum {
	values := make([]InventoryRecordLifecycleStateEnum, 0)
	for _, v := range mappingInventoryRecordLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInventoryRecordLifecycleStateEnumStringValues Enumerates the set of values in String for InventoryRecordLifecycleStateEnum
func GetInventoryRecordLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingInventoryRecordLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInventoryRecordLifecycleStateEnum(val string) (InventoryRecordLifecycleStateEnum, bool) {
	enum, ok := mappingInventoryRecordLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
