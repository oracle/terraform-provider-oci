// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// TableSummary Summary of the table.
type TableSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Human-friendly table name, also immutable.
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

	// True if this table can be reclaimed after an idle period.
	IsAutoReclaimable *bool `mandatory:"false" json:"isAutoReclaimable"`

	// If lifecycleState is INACTIVE, indicates when
	// this table will be automatically removed.
	// An RFC3339 formatted datetime string.
	TimeOfExpiration *common.SDKTime `mandatory:"false" json:"timeOfExpiration"`

	// The current state of this table's schema. Available states are
	// MUTABLE - The schema can be changed. The table is not eligible for replication.
	// FROZEN - The schema is immutable. The table is eligible for replication.
	SchemaState TableSummarySchemaStateEnum `mandatory:"false" json:"schemaState,omitempty"`

	// True if this table is currently a member of a replication set.
	IsMultiRegion *bool `mandatory:"false" json:"isMultiRegion"`

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

func (m TableSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TableSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTableLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTableLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTableSummarySchemaStateEnum(string(m.SchemaState)); !ok && m.SchemaState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SchemaState: %s. Supported values are: %s.", m.SchemaState, strings.Join(GetTableSummarySchemaStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TableSummarySchemaStateEnum Enum with underlying type: string
type TableSummarySchemaStateEnum string

// Set of constants representing the allowable values for TableSummarySchemaStateEnum
const (
	TableSummarySchemaStateMutable TableSummarySchemaStateEnum = "MUTABLE"
	TableSummarySchemaStateFrozen  TableSummarySchemaStateEnum = "FROZEN"
)

var mappingTableSummarySchemaStateEnum = map[string]TableSummarySchemaStateEnum{
	"MUTABLE": TableSummarySchemaStateMutable,
	"FROZEN":  TableSummarySchemaStateFrozen,
}

var mappingTableSummarySchemaStateEnumLowerCase = map[string]TableSummarySchemaStateEnum{
	"mutable": TableSummarySchemaStateMutable,
	"frozen":  TableSummarySchemaStateFrozen,
}

// GetTableSummarySchemaStateEnumValues Enumerates the set of values for TableSummarySchemaStateEnum
func GetTableSummarySchemaStateEnumValues() []TableSummarySchemaStateEnum {
	values := make([]TableSummarySchemaStateEnum, 0)
	for _, v := range mappingTableSummarySchemaStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTableSummarySchemaStateEnumStringValues Enumerates the set of values in String for TableSummarySchemaStateEnum
func GetTableSummarySchemaStateEnumStringValues() []string {
	return []string{
		"MUTABLE",
		"FROZEN",
	}
}

// GetMappingTableSummarySchemaStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTableSummarySchemaStateEnum(val string) (TableSummarySchemaStateEnum, bool) {
	enum, ok := mappingTableSummarySchemaStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
