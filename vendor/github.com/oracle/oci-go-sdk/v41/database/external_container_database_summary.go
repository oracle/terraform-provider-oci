// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// ExternalContainerDatabaseSummary An Oracle Cloud Infrastructure resource that allows you to manage an external Oracle container database.
type ExternalContainerDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Oracle Cloud Infrastructure external database resource.
	LifecycleState ExternalContainerDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The `DB_UNIQUE_NAME` of the external database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The Oracle Database ID, which identifies an Oracle Database located outside of Oracle Cloud.
	DbId *string `mandatory:"false" json:"dbId"`

	// The Oracle Database version.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// The Oracle Database edition.
	DatabaseEdition ExternalContainerDatabaseSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The time zone of the external database.
	// It is a time zone offset (a character type in the format '[+|-]TZH:TZM') or a time zone region name,
	// depending on how the time zone value was specified when the database was created / last altered.
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The character set of the external database.
	CharacterSet *string `mandatory:"false" json:"characterSet"`

	// The national character of the external database.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet"`

	// The database packs licensed for the external Oracle Database.
	DbPacks *string `mandatory:"false" json:"dbPacks"`

	// The Oracle Database configuration
	DatabaseConfiguration ExternalContainerDatabaseSummaryDatabaseConfigurationEnum `mandatory:"false" json:"databaseConfiguration,omitempty"`

	DatabaseManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`
}

func (m ExternalContainerDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ExternalContainerDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type ExternalContainerDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseSummaryLifecycleStateEnum
const (
	ExternalContainerDatabaseSummaryLifecycleStateProvisioning ExternalContainerDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	ExternalContainerDatabaseSummaryLifecycleStateNotConnected ExternalContainerDatabaseSummaryLifecycleStateEnum = "NOT_CONNECTED"
	ExternalContainerDatabaseSummaryLifecycleStateAvailable    ExternalContainerDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	ExternalContainerDatabaseSummaryLifecycleStateUpdating     ExternalContainerDatabaseSummaryLifecycleStateEnum = "UPDATING"
	ExternalContainerDatabaseSummaryLifecycleStateTerminating  ExternalContainerDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	ExternalContainerDatabaseSummaryLifecycleStateTerminated   ExternalContainerDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	ExternalContainerDatabaseSummaryLifecycleStateFailed       ExternalContainerDatabaseSummaryLifecycleStateEnum = "FAILED"
)

var mappingExternalContainerDatabaseSummaryLifecycleState = map[string]ExternalContainerDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":  ExternalContainerDatabaseSummaryLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalContainerDatabaseSummaryLifecycleStateNotConnected,
	"AVAILABLE":     ExternalContainerDatabaseSummaryLifecycleStateAvailable,
	"UPDATING":      ExternalContainerDatabaseSummaryLifecycleStateUpdating,
	"TERMINATING":   ExternalContainerDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":    ExternalContainerDatabaseSummaryLifecycleStateTerminated,
	"FAILED":        ExternalContainerDatabaseSummaryLifecycleStateFailed,
}

// GetExternalContainerDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for ExternalContainerDatabaseSummaryLifecycleStateEnum
func GetExternalContainerDatabaseSummaryLifecycleStateEnumValues() []ExternalContainerDatabaseSummaryLifecycleStateEnum {
	values := make([]ExternalContainerDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExternalContainerDatabaseSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// ExternalContainerDatabaseSummaryDatabaseEditionEnum Enum with underlying type: string
type ExternalContainerDatabaseSummaryDatabaseEditionEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseSummaryDatabaseEditionEnum
const (
	ExternalContainerDatabaseSummaryDatabaseEditionStandardEdition                     ExternalContainerDatabaseSummaryDatabaseEditionEnum = "STANDARD_EDITION"
	ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEdition                   ExternalContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance    ExternalContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance ExternalContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingExternalContainerDatabaseSummaryDatabaseEdition = map[string]ExternalContainerDatabaseSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalContainerDatabaseSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalContainerDatabaseSummaryDatabaseEditionEnumValues Enumerates the set of values for ExternalContainerDatabaseSummaryDatabaseEditionEnum
func GetExternalContainerDatabaseSummaryDatabaseEditionEnumValues() []ExternalContainerDatabaseSummaryDatabaseEditionEnum {
	values := make([]ExternalContainerDatabaseSummaryDatabaseEditionEnum, 0)
	for _, v := range mappingExternalContainerDatabaseSummaryDatabaseEdition {
		values = append(values, v)
	}
	return values
}

// ExternalContainerDatabaseSummaryDatabaseConfigurationEnum Enum with underlying type: string
type ExternalContainerDatabaseSummaryDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseSummaryDatabaseConfigurationEnum
const (
	ExternalContainerDatabaseSummaryDatabaseConfigurationRac            ExternalContainerDatabaseSummaryDatabaseConfigurationEnum = "RAC"
	ExternalContainerDatabaseSummaryDatabaseConfigurationSingleInstance ExternalContainerDatabaseSummaryDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalContainerDatabaseSummaryDatabaseConfiguration = map[string]ExternalContainerDatabaseSummaryDatabaseConfigurationEnum{
	"RAC":             ExternalContainerDatabaseSummaryDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalContainerDatabaseSummaryDatabaseConfigurationSingleInstance,
}

// GetExternalContainerDatabaseSummaryDatabaseConfigurationEnumValues Enumerates the set of values for ExternalContainerDatabaseSummaryDatabaseConfigurationEnum
func GetExternalContainerDatabaseSummaryDatabaseConfigurationEnumValues() []ExternalContainerDatabaseSummaryDatabaseConfigurationEnum {
	values := make([]ExternalContainerDatabaseSummaryDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalContainerDatabaseSummaryDatabaseConfiguration {
		values = append(values, v)
	}
	return values
}
