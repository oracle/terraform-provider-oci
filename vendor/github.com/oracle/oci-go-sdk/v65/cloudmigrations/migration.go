// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Migration A top-level container to track all aspects of a long-running migration workflow to OCI.
type Migration struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of migration.
	LifecycleState MigrationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the migration project was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Migration Identifier that can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time when the migration project was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Replication schedule identifier
	ReplicationScheduleId *string `mandatory:"false" json:"replicationScheduleId"`

	// Indicates whether migration is marked as completed.
	IsCompleted *bool `mandatory:"false" json:"isCompleted"`

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Migration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Migration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMigrationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MigrationLifecycleStateEnum Enum with underlying type: string
type MigrationLifecycleStateEnum string

// Set of constants representing the allowable values for MigrationLifecycleStateEnum
const (
	MigrationLifecycleStateCreating       MigrationLifecycleStateEnum = "CREATING"
	MigrationLifecycleStateUpdating       MigrationLifecycleStateEnum = "UPDATING"
	MigrationLifecycleStateNeedsAttention MigrationLifecycleStateEnum = "NEEDS_ATTENTION"
	MigrationLifecycleStateActive         MigrationLifecycleStateEnum = "ACTIVE"
	MigrationLifecycleStateDeleting       MigrationLifecycleStateEnum = "DELETING"
	MigrationLifecycleStateDeleted        MigrationLifecycleStateEnum = "DELETED"
	MigrationLifecycleStateFailed         MigrationLifecycleStateEnum = "FAILED"
)

var mappingMigrationLifecycleStateEnum = map[string]MigrationLifecycleStateEnum{
	"CREATING":        MigrationLifecycleStateCreating,
	"UPDATING":        MigrationLifecycleStateUpdating,
	"NEEDS_ATTENTION": MigrationLifecycleStateNeedsAttention,
	"ACTIVE":          MigrationLifecycleStateActive,
	"DELETING":        MigrationLifecycleStateDeleting,
	"DELETED":         MigrationLifecycleStateDeleted,
	"FAILED":          MigrationLifecycleStateFailed,
}

var mappingMigrationLifecycleStateEnumLowerCase = map[string]MigrationLifecycleStateEnum{
	"creating":        MigrationLifecycleStateCreating,
	"updating":        MigrationLifecycleStateUpdating,
	"needs_attention": MigrationLifecycleStateNeedsAttention,
	"active":          MigrationLifecycleStateActive,
	"deleting":        MigrationLifecycleStateDeleting,
	"deleted":         MigrationLifecycleStateDeleted,
	"failed":          MigrationLifecycleStateFailed,
}

// GetMigrationLifecycleStateEnumValues Enumerates the set of values for MigrationLifecycleStateEnum
func GetMigrationLifecycleStateEnumValues() []MigrationLifecycleStateEnum {
	values := make([]MigrationLifecycleStateEnum, 0)
	for _, v := range mappingMigrationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationLifecycleStateEnumStringValues Enumerates the set of values in String for MigrationLifecycleStateEnum
func GetMigrationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"NEEDS_ATTENTION",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMigrationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationLifecycleStateEnum(val string) (MigrationLifecycleStateEnum, bool) {
	enum, ok := mappingMigrationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
