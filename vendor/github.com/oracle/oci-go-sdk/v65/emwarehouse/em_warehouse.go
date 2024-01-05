// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// EM Warehouse API
//
// Use the EM Warehouse API to manage EM Warehouse data collection.
//

package emwarehouse

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmWarehouse Description of EmWarehouse.
type EmWarehouse struct {

	// operations Insights Warehouse Identifier
	OperationsInsightsWarehouseId *string `mandatory:"true" json:"operationsInsightsWarehouseId"`

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// EmWarehouse Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the EmWarehouse.
	EmWarehouseType *string `mandatory:"true" json:"emWarehouseType"`

	// The time the the EmWarehouse was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the EmWarehouse.
	LifecycleState EmWarehouseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Data Flow Run Status
	LatestEtlRunStatus *string `mandatory:"false" json:"latestEtlRunStatus"`

	// Data Flow Run Status Message
	LatestEtlRunMessage *string `mandatory:"false" json:"latestEtlRunMessage"`

	// Data Flow Run Total Time
	LatestEtlRunTime *string `mandatory:"false" json:"latestEtlRunTime"`

	// EMBridge Identifier
	EmBridgeId *string `mandatory:"false" json:"emBridgeId"`

	// The time the EmWarehouse was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m EmWarehouse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmWarehouse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmWarehouseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmWarehouseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmWarehouseLifecycleStateEnum Enum with underlying type: string
type EmWarehouseLifecycleStateEnum string

// Set of constants representing the allowable values for EmWarehouseLifecycleStateEnum
const (
	EmWarehouseLifecycleStateCreating EmWarehouseLifecycleStateEnum = "CREATING"
	EmWarehouseLifecycleStateUpdating EmWarehouseLifecycleStateEnum = "UPDATING"
	EmWarehouseLifecycleStateActive   EmWarehouseLifecycleStateEnum = "ACTIVE"
	EmWarehouseLifecycleStateDeleting EmWarehouseLifecycleStateEnum = "DELETING"
	EmWarehouseLifecycleStateDeleted  EmWarehouseLifecycleStateEnum = "DELETED"
	EmWarehouseLifecycleStateFailed   EmWarehouseLifecycleStateEnum = "FAILED"
)

var mappingEmWarehouseLifecycleStateEnum = map[string]EmWarehouseLifecycleStateEnum{
	"CREATING": EmWarehouseLifecycleStateCreating,
	"UPDATING": EmWarehouseLifecycleStateUpdating,
	"ACTIVE":   EmWarehouseLifecycleStateActive,
	"DELETING": EmWarehouseLifecycleStateDeleting,
	"DELETED":  EmWarehouseLifecycleStateDeleted,
	"FAILED":   EmWarehouseLifecycleStateFailed,
}

var mappingEmWarehouseLifecycleStateEnumLowerCase = map[string]EmWarehouseLifecycleStateEnum{
	"creating": EmWarehouseLifecycleStateCreating,
	"updating": EmWarehouseLifecycleStateUpdating,
	"active":   EmWarehouseLifecycleStateActive,
	"deleting": EmWarehouseLifecycleStateDeleting,
	"deleted":  EmWarehouseLifecycleStateDeleted,
	"failed":   EmWarehouseLifecycleStateFailed,
}

// GetEmWarehouseLifecycleStateEnumValues Enumerates the set of values for EmWarehouseLifecycleStateEnum
func GetEmWarehouseLifecycleStateEnumValues() []EmWarehouseLifecycleStateEnum {
	values := make([]EmWarehouseLifecycleStateEnum, 0)
	for _, v := range mappingEmWarehouseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmWarehouseLifecycleStateEnumStringValues Enumerates the set of values in String for EmWarehouseLifecycleStateEnum
func GetEmWarehouseLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEmWarehouseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmWarehouseLifecycleStateEnum(val string) (EmWarehouseLifecycleStateEnum, bool) {
	enum, ok := mappingEmWarehouseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
