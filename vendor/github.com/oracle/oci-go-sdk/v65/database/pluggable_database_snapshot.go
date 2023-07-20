// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PluggableDatabaseSnapshot Details of the Pluggable Database Snapshot.
type PluggableDatabaseSnapshot struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata Pluggable Database Snapshot.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the Database Snapshot. The name should be unique.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata Pluggable Database.
	PluggableDatabaseId *string `mandatory:"true" json:"pluggableDatabaseId"`

	// The current state of the Exadata Pluggable Database Snapshot.
	LifecycleState PluggableDatabaseSnapshotLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time that the Exadata Pluggable Database Snapshot was created, as expressed in RFC 3339 format. For example, 2023-06-27T21:10:29Z
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state of the Exadata Pluggable Database Snapshot.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata VM cluster on Exascale Infrastructure.
	ExadbVmClusterId *string `mandatory:"false" json:"exadbVmClusterId"`
}

func (m PluggableDatabaseSnapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PluggableDatabaseSnapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPluggableDatabaseSnapshotLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPluggableDatabaseSnapshotLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PluggableDatabaseSnapshotLifecycleStateEnum Enum with underlying type: string
type PluggableDatabaseSnapshotLifecycleStateEnum string

// Set of constants representing the allowable values for PluggableDatabaseSnapshotLifecycleStateEnum
const (
	PluggableDatabaseSnapshotLifecycleStateCreating    PluggableDatabaseSnapshotLifecycleStateEnum = "CREATING"
	PluggableDatabaseSnapshotLifecycleStateAvailable   PluggableDatabaseSnapshotLifecycleStateEnum = "AVAILABLE"
	PluggableDatabaseSnapshotLifecycleStateTerminating PluggableDatabaseSnapshotLifecycleStateEnum = "TERMINATING"
	PluggableDatabaseSnapshotLifecycleStateTerminated  PluggableDatabaseSnapshotLifecycleStateEnum = "TERMINATED"
	PluggableDatabaseSnapshotLifecycleStateFailed      PluggableDatabaseSnapshotLifecycleStateEnum = "FAILED"
)

var mappingPluggableDatabaseSnapshotLifecycleStateEnum = map[string]PluggableDatabaseSnapshotLifecycleStateEnum{
	"CREATING":    PluggableDatabaseSnapshotLifecycleStateCreating,
	"AVAILABLE":   PluggableDatabaseSnapshotLifecycleStateAvailable,
	"TERMINATING": PluggableDatabaseSnapshotLifecycleStateTerminating,
	"TERMINATED":  PluggableDatabaseSnapshotLifecycleStateTerminated,
	"FAILED":      PluggableDatabaseSnapshotLifecycleStateFailed,
}

var mappingPluggableDatabaseSnapshotLifecycleStateEnumLowerCase = map[string]PluggableDatabaseSnapshotLifecycleStateEnum{
	"creating":    PluggableDatabaseSnapshotLifecycleStateCreating,
	"available":   PluggableDatabaseSnapshotLifecycleStateAvailable,
	"terminating": PluggableDatabaseSnapshotLifecycleStateTerminating,
	"terminated":  PluggableDatabaseSnapshotLifecycleStateTerminated,
	"failed":      PluggableDatabaseSnapshotLifecycleStateFailed,
}

// GetPluggableDatabaseSnapshotLifecycleStateEnumValues Enumerates the set of values for PluggableDatabaseSnapshotLifecycleStateEnum
func GetPluggableDatabaseSnapshotLifecycleStateEnumValues() []PluggableDatabaseSnapshotLifecycleStateEnum {
	values := make([]PluggableDatabaseSnapshotLifecycleStateEnum, 0)
	for _, v := range mappingPluggableDatabaseSnapshotLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPluggableDatabaseSnapshotLifecycleStateEnumStringValues Enumerates the set of values in String for PluggableDatabaseSnapshotLifecycleStateEnum
func GetPluggableDatabaseSnapshotLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingPluggableDatabaseSnapshotLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluggableDatabaseSnapshotLifecycleStateEnum(val string) (PluggableDatabaseSnapshotLifecycleStateEnum, bool) {
	enum, ok := mappingPluggableDatabaseSnapshotLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
