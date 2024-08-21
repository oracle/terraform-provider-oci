// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalDatabaseSummary The summary of an external database.
type ExternalDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the external database resource.
	LifecycleState ExternalDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external DB system was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The `DB_UNIQUE_NAME` of the external database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The type of Oracle Database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"false" json:"databaseType,omitempty"`

	// The subtype of Oracle Database. Indicates whether the database is a Container Database,
	// Pluggable Database, or Non-container Database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent Container Database (CDB)
	// if this is a Pluggable Database (PDB).
	ExternalContainerDatabaseId *string `mandatory:"false" json:"externalContainerDatabaseId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB home.
	ExternalDbHomeId *string `mandatory:"false" json:"externalDbHomeId"`

	DbSystemInfo *ExternalDbSystemBasicInfo `mandatory:"false" json:"dbSystemInfo"`

	DbManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"dbManagementConfig"`

	// The list of database instances if the database is a RAC database.
	InstanceDetails []ExternalDatabaseInstance `mandatory:"false" json:"instanceDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The list of feature configurations
	DbmgmtFeatureConfigs []DatabaseFeatureConfiguration `mandatory:"false" json:"dbmgmtFeatureConfigs"`

	// The Oracle database version.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// The operating system of database.
	DatabasePlatformName *string `mandatory:"false" json:"databasePlatformName"`
}

func (m ExternalDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDatabaseSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ExternalDatabaseSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DbUniqueName                *string                                   `json:"dbUniqueName"`
		DatabaseType                DatabaseTypeEnum                          `json:"databaseType"`
		DatabaseSubType             DatabaseSubTypeEnum                       `json:"databaseSubType"`
		ExternalContainerDatabaseId *string                                   `json:"externalContainerDatabaseId"`
		ExternalDbHomeId            *string                                   `json:"externalDbHomeId"`
		DbSystemInfo                *ExternalDbSystemBasicInfo                `json:"dbSystemInfo"`
		DbManagementConfig          *DatabaseManagementConfig                 `json:"dbManagementConfig"`
		InstanceDetails             []ExternalDatabaseInstance                `json:"instanceDetails"`
		FreeformTags                map[string]string                         `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{}         `json:"definedTags"`
		SystemTags                  map[string]map[string]interface{}         `json:"systemTags"`
		DbmgmtFeatureConfigs        []databasefeatureconfiguration            `json:"dbmgmtFeatureConfigs"`
		DatabaseVersion             *string                                   `json:"databaseVersion"`
		DatabasePlatformName        *string                                   `json:"databasePlatformName"`
		Id                          *string                                   `json:"id"`
		DisplayName                 *string                                   `json:"displayName"`
		CompartmentId               *string                                   `json:"compartmentId"`
		LifecycleState              ExternalDatabaseSummaryLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated                 *common.SDKTime                           `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DbUniqueName = model.DbUniqueName

	m.DatabaseType = model.DatabaseType

	m.DatabaseSubType = model.DatabaseSubType

	m.ExternalContainerDatabaseId = model.ExternalContainerDatabaseId

	m.ExternalDbHomeId = model.ExternalDbHomeId

	m.DbSystemInfo = model.DbSystemInfo

	m.DbManagementConfig = model.DbManagementConfig

	m.InstanceDetails = make([]ExternalDatabaseInstance, len(model.InstanceDetails))
	copy(m.InstanceDetails, model.InstanceDetails)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.DbmgmtFeatureConfigs = make([]DatabaseFeatureConfiguration, len(model.DbmgmtFeatureConfigs))
	for i, n := range model.DbmgmtFeatureConfigs {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DbmgmtFeatureConfigs[i] = nn.(DatabaseFeatureConfiguration)
		} else {
			m.DbmgmtFeatureConfigs[i] = nil
		}
	}
	m.DatabaseVersion = model.DatabaseVersion

	m.DatabasePlatformName = model.DatabasePlatformName

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	return
}

// ExternalDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type ExternalDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalDatabaseSummaryLifecycleStateEnum
const (
	ExternalDatabaseSummaryLifecycleStateCreating ExternalDatabaseSummaryLifecycleStateEnum = "CREATING"
	ExternalDatabaseSummaryLifecycleStateActive   ExternalDatabaseSummaryLifecycleStateEnum = "ACTIVE"
	ExternalDatabaseSummaryLifecycleStateInactive ExternalDatabaseSummaryLifecycleStateEnum = "INACTIVE"
	ExternalDatabaseSummaryLifecycleStateUpdating ExternalDatabaseSummaryLifecycleStateEnum = "UPDATING"
	ExternalDatabaseSummaryLifecycleStateDeleting ExternalDatabaseSummaryLifecycleStateEnum = "DELETING"
	ExternalDatabaseSummaryLifecycleStateDeleted  ExternalDatabaseSummaryLifecycleStateEnum = "DELETED"
	ExternalDatabaseSummaryLifecycleStateFailed   ExternalDatabaseSummaryLifecycleStateEnum = "FAILED"
)

var mappingExternalDatabaseSummaryLifecycleStateEnum = map[string]ExternalDatabaseSummaryLifecycleStateEnum{
	"CREATING": ExternalDatabaseSummaryLifecycleStateCreating,
	"ACTIVE":   ExternalDatabaseSummaryLifecycleStateActive,
	"INACTIVE": ExternalDatabaseSummaryLifecycleStateInactive,
	"UPDATING": ExternalDatabaseSummaryLifecycleStateUpdating,
	"DELETING": ExternalDatabaseSummaryLifecycleStateDeleting,
	"DELETED":  ExternalDatabaseSummaryLifecycleStateDeleted,
	"FAILED":   ExternalDatabaseSummaryLifecycleStateFailed,
}

var mappingExternalDatabaseSummaryLifecycleStateEnumLowerCase = map[string]ExternalDatabaseSummaryLifecycleStateEnum{
	"creating": ExternalDatabaseSummaryLifecycleStateCreating,
	"active":   ExternalDatabaseSummaryLifecycleStateActive,
	"inactive": ExternalDatabaseSummaryLifecycleStateInactive,
	"updating": ExternalDatabaseSummaryLifecycleStateUpdating,
	"deleting": ExternalDatabaseSummaryLifecycleStateDeleting,
	"deleted":  ExternalDatabaseSummaryLifecycleStateDeleted,
	"failed":   ExternalDatabaseSummaryLifecycleStateFailed,
}

// GetExternalDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for ExternalDatabaseSummaryLifecycleStateEnum
func GetExternalDatabaseSummaryLifecycleStateEnumValues() []ExternalDatabaseSummaryLifecycleStateEnum {
	values := make([]ExternalDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExternalDatabaseSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalDatabaseSummaryLifecycleStateEnum
func GetExternalDatabaseSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingExternalDatabaseSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseSummaryLifecycleStateEnum(val string) (ExternalDatabaseSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExternalDatabaseSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
