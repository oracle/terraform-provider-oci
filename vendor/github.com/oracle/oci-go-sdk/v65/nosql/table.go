// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Table Complete metadata about a table.
type Table struct {

	// Unique identifier that is immutable.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Human-friendly table name, immutable.
	Name *string `mandatory:"false" json:"name"`

	// The time the the table was created. An RFC3339 formatted
	// datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the the table's metadata was last updated. An
	// RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	TableLimits *TableLimits `mandatory:"false" json:"tableLimits"`

	// The state of a table.
	LifecycleState TableLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// True if this table can be reclaimed after an idle period.
	IsAutoReclaimable *bool `mandatory:"false" json:"isAutoReclaimable"`

	// If lifecycleState is INACTIVE, indicates when
	// this table will be automatically removed.
	// An RFC3339 formatted datetime string.
	TimeOfExpiration *common.SDKTime `mandatory:"false" json:"timeOfExpiration"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	Schema *Schema `mandatory:"false" json:"schema"`

	// A DDL statement representing the schema.
	DdlStatement *string `mandatory:"false" json:"ddlStatement"`

	// The current state of this table's schema. Available states are
	// MUTABLE - The schema can be changed. The table is not eligible for replication.
	// FROZEN - The schema is immutable. The table is eligible for replication.
	SchemaState TableSchemaStateEnum `mandatory:"false" json:"schemaState,omitempty"`

	// True if this table is currently a member of a replication set.
	IsMultiRegion *bool `mandatory:"false" json:"isMultiRegion"`

	// If this table is in a replication set, this value represents
	// the progress of the initialization of the replica's data.  A
	// value of 100 indicates that initialization has completed.
	LocalReplicaInitializationInPercent *int `mandatory:"false" json:"localReplicaInitializationInPercent"`

	// An array of Replica listing this table's replicas, if any
	Replicas []Replica `mandatory:"false" json:"replicas"`

	// Simple key-value pair that is applied without any predefined
	// name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and
	// scoped to a namespace.  Example: `{"foo-namespace":
	// {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Read-only system tag. These predefined keys are scoped to
	// namespaces.  At present the only supported namespace is
	// `"orcl-cloud"`; and the only key in that namespace is
	// `"free-tier-retained"`.
	// Example: `{"orcl-cloud"": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Table) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Table) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTableLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTableLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTableSchemaStateEnum(string(m.SchemaState)); !ok && m.SchemaState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SchemaState: %s. Supported values are: %s.", m.SchemaState, strings.Join(GetTableSchemaStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TableLifecycleStateEnum Enum with underlying type: string
type TableLifecycleStateEnum string

// Set of constants representing the allowable values for TableLifecycleStateEnum
const (
	TableLifecycleStateCreating TableLifecycleStateEnum = "CREATING"
	TableLifecycleStateUpdating TableLifecycleStateEnum = "UPDATING"
	TableLifecycleStateActive   TableLifecycleStateEnum = "ACTIVE"
	TableLifecycleStateDeleting TableLifecycleStateEnum = "DELETING"
	TableLifecycleStateDeleted  TableLifecycleStateEnum = "DELETED"
	TableLifecycleStateFailed   TableLifecycleStateEnum = "FAILED"
	TableLifecycleStateInactive TableLifecycleStateEnum = "INACTIVE"
)

var mappingTableLifecycleStateEnum = map[string]TableLifecycleStateEnum{
	"CREATING": TableLifecycleStateCreating,
	"UPDATING": TableLifecycleStateUpdating,
	"ACTIVE":   TableLifecycleStateActive,
	"DELETING": TableLifecycleStateDeleting,
	"DELETED":  TableLifecycleStateDeleted,
	"FAILED":   TableLifecycleStateFailed,
	"INACTIVE": TableLifecycleStateInactive,
}

var mappingTableLifecycleStateEnumLowerCase = map[string]TableLifecycleStateEnum{
	"creating": TableLifecycleStateCreating,
	"updating": TableLifecycleStateUpdating,
	"active":   TableLifecycleStateActive,
	"deleting": TableLifecycleStateDeleting,
	"deleted":  TableLifecycleStateDeleted,
	"failed":   TableLifecycleStateFailed,
	"inactive": TableLifecycleStateInactive,
}

// GetTableLifecycleStateEnumValues Enumerates the set of values for TableLifecycleStateEnum
func GetTableLifecycleStateEnumValues() []TableLifecycleStateEnum {
	values := make([]TableLifecycleStateEnum, 0)
	for _, v := range mappingTableLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTableLifecycleStateEnumStringValues Enumerates the set of values in String for TableLifecycleStateEnum
func GetTableLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingTableLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTableLifecycleStateEnum(val string) (TableLifecycleStateEnum, bool) {
	enum, ok := mappingTableLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TableSchemaStateEnum Enum with underlying type: string
type TableSchemaStateEnum string

// Set of constants representing the allowable values for TableSchemaStateEnum
const (
	TableSchemaStateMutable TableSchemaStateEnum = "MUTABLE"
	TableSchemaStateFrozen  TableSchemaStateEnum = "FROZEN"
)

var mappingTableSchemaStateEnum = map[string]TableSchemaStateEnum{
	"MUTABLE": TableSchemaStateMutable,
	"FROZEN":  TableSchemaStateFrozen,
}

var mappingTableSchemaStateEnumLowerCase = map[string]TableSchemaStateEnum{
	"mutable": TableSchemaStateMutable,
	"frozen":  TableSchemaStateFrozen,
}

// GetTableSchemaStateEnumValues Enumerates the set of values for TableSchemaStateEnum
func GetTableSchemaStateEnumValues() []TableSchemaStateEnum {
	values := make([]TableSchemaStateEnum, 0)
	for _, v := range mappingTableSchemaStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTableSchemaStateEnumStringValues Enumerates the set of values in String for TableSchemaStateEnum
func GetTableSchemaStateEnumStringValues() []string {
	return []string{
		"MUTABLE",
		"FROZEN",
	}
}

// GetMappingTableSchemaStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTableSchemaStateEnum(val string) (TableSchemaStateEnum, bool) {
	enum, ok := mappingTableSchemaStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
