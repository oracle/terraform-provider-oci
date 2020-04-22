// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ndcs-control-plane API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
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

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	Schema *Schema `mandatory:"false" json:"schema"`

	// A DDL statement representing the schema.
	DdlStatement *string `mandatory:"false" json:"ddlStatement"`

	// Simple key-value pair that is applied without any predefined
	// name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and
	// scoped to a namespace.  Example: `{"foo-namespace":
	// {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Table) String() string {
	return common.PointerString(m)
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
)

var mappingTableLifecycleState = map[string]TableLifecycleStateEnum{
	"CREATING": TableLifecycleStateCreating,
	"UPDATING": TableLifecycleStateUpdating,
	"ACTIVE":   TableLifecycleStateActive,
	"DELETING": TableLifecycleStateDeleting,
	"DELETED":  TableLifecycleStateDeleted,
	"FAILED":   TableLifecycleStateFailed,
}

// GetTableLifecycleStateEnumValues Enumerates the set of values for TableLifecycleStateEnum
func GetTableLifecycleStateEnumValues() []TableLifecycleStateEnum {
	values := make([]TableLifecycleStateEnum, 0)
	for _, v := range mappingTableLifecycleState {
		values = append(values, v)
	}
	return values
}
