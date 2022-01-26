// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PluggableDatabase A pluggable database (PDB) is portable collection of schemas, schema objects, and non-schema objects that appears to an Oracle client as a non-container database. To use a PDB, it needs to be plugged into a CDB.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to a tenancy administrator. If you are an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type PluggableDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the pluggable database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the CDB.
	ContainerDatabaseId *string `mandatory:"true" json:"containerDatabaseId"`

	// The name for the pluggable database (PDB). The name is unique in the context of a Database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	PdbName *string `mandatory:"true" json:"pdbName"`

	// The current state of the pluggable database.
	LifecycleState PluggableDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the pluggable database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
	OpenMode PluggableDatabaseOpenModeEnum `mandatory:"true" json:"openMode"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Detailed message for the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	ConnectionStrings *PluggableDatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	// The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode,
	// the user needs both create a session and have restricted session privileges to connect to it.
	IsRestricted *bool `mandatory:"false" json:"isRestricted"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m PluggableDatabase) String() string {
	return common.PointerString(m)
}

// PluggableDatabaseLifecycleStateEnum Enum with underlying type: string
type PluggableDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for PluggableDatabaseLifecycleStateEnum
const (
	PluggableDatabaseLifecycleStateProvisioning PluggableDatabaseLifecycleStateEnum = "PROVISIONING"
	PluggableDatabaseLifecycleStateAvailable    PluggableDatabaseLifecycleStateEnum = "AVAILABLE"
	PluggableDatabaseLifecycleStateTerminating  PluggableDatabaseLifecycleStateEnum = "TERMINATING"
	PluggableDatabaseLifecycleStateTerminated   PluggableDatabaseLifecycleStateEnum = "TERMINATED"
	PluggableDatabaseLifecycleStateUpdating     PluggableDatabaseLifecycleStateEnum = "UPDATING"
	PluggableDatabaseLifecycleStateFailed       PluggableDatabaseLifecycleStateEnum = "FAILED"
)

var mappingPluggableDatabaseLifecycleState = map[string]PluggableDatabaseLifecycleStateEnum{
	"PROVISIONING": PluggableDatabaseLifecycleStateProvisioning,
	"AVAILABLE":    PluggableDatabaseLifecycleStateAvailable,
	"TERMINATING":  PluggableDatabaseLifecycleStateTerminating,
	"TERMINATED":   PluggableDatabaseLifecycleStateTerminated,
	"UPDATING":     PluggableDatabaseLifecycleStateUpdating,
	"FAILED":       PluggableDatabaseLifecycleStateFailed,
}

// GetPluggableDatabaseLifecycleStateEnumValues Enumerates the set of values for PluggableDatabaseLifecycleStateEnum
func GetPluggableDatabaseLifecycleStateEnumValues() []PluggableDatabaseLifecycleStateEnum {
	values := make([]PluggableDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingPluggableDatabaseLifecycleState {
		values = append(values, v)
	}
	return values
}

// PluggableDatabaseOpenModeEnum Enum with underlying type: string
type PluggableDatabaseOpenModeEnum string

// Set of constants representing the allowable values for PluggableDatabaseOpenModeEnum
const (
	PluggableDatabaseOpenModeReadOnly  PluggableDatabaseOpenModeEnum = "READ_ONLY"
	PluggableDatabaseOpenModeReadWrite PluggableDatabaseOpenModeEnum = "READ_WRITE"
	PluggableDatabaseOpenModeMounted   PluggableDatabaseOpenModeEnum = "MOUNTED"
	PluggableDatabaseOpenModeMigrate   PluggableDatabaseOpenModeEnum = "MIGRATE"
)

var mappingPluggableDatabaseOpenMode = map[string]PluggableDatabaseOpenModeEnum{
	"READ_ONLY":  PluggableDatabaseOpenModeReadOnly,
	"READ_WRITE": PluggableDatabaseOpenModeReadWrite,
	"MOUNTED":    PluggableDatabaseOpenModeMounted,
	"MIGRATE":    PluggableDatabaseOpenModeMigrate,
}

// GetPluggableDatabaseOpenModeEnumValues Enumerates the set of values for PluggableDatabaseOpenModeEnum
func GetPluggableDatabaseOpenModeEnumValues() []PluggableDatabaseOpenModeEnum {
	values := make([]PluggableDatabaseOpenModeEnum, 0)
	for _, v := range mappingPluggableDatabaseOpenMode {
		values = append(values, v)
	}
	return values
}
