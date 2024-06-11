// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MdsMySqlDatabaseInsight Database insight resource.
type MdsMySqlDatabaseInsight struct {

	// Database insight identifier
	Id *string `mandatory:"true" json:"id"`

	// Compartment identifier of the database
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the the database insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// Name of database
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// OCI database resource type
	DatabaseResourceType *string `mandatory:"true" json:"databaseResourceType"`

	// Ops Insights internal representation of the database type.
	DatabaseType *string `mandatory:"false" json:"databaseType"`

	// The version of the database.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time the database insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A message describing the status of the database connection of this resource. For example, it can be used to provide actionable information about the permission and content validity of the database connection.
	DatabaseConnectionStatusDetails *string `mandatory:"false" json:"databaseConnectionStatusDetails"`

	// Display name of database
	DatabaseDisplayName *string `mandatory:"false" json:"databaseDisplayName"`

	// Specifies if MYSQL DB System is highly available.
	IsHighlyAvailable *bool `mandatory:"false" json:"isHighlyAvailable"`

	// Specifies if MYSQL DB System has heatwave cluster attached.
	IsHeatWaveClusterAttached *bool `mandatory:"false" json:"isHeatWaveClusterAttached"`

	// Additional details of a db system in JSON format.
	// For MySQL DB System, this is the DbSystem object serialized as a JSON string as defined in https://docs.oracle.com/en-us/iaas/api/#/en/mysql/20190415/DbSystem/.
	DbAdditionalDetails *interface{} `mandatory:"false" json:"dbAdditionalDetails"`

	// Indicates the status of a database insight in Operations Insights
	Status ResourceStatusEnum `mandatory:"true" json:"status"`

	// The current state of the database.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m MdsMySqlDatabaseInsight) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m MdsMySqlDatabaseInsight) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetStatus returns Status
func (m MdsMySqlDatabaseInsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetDatabaseType returns DatabaseType
func (m MdsMySqlDatabaseInsight) GetDatabaseType() *string {
	return m.DatabaseType
}

// GetDatabaseVersion returns DatabaseVersion
func (m MdsMySqlDatabaseInsight) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

// GetProcessorCount returns ProcessorCount
func (m MdsMySqlDatabaseInsight) GetProcessorCount() *int {
	return m.ProcessorCount
}

// GetFreeformTags returns FreeformTags
func (m MdsMySqlDatabaseInsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MdsMySqlDatabaseInsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MdsMySqlDatabaseInsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeCreated returns TimeCreated
func (m MdsMySqlDatabaseInsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MdsMySqlDatabaseInsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m MdsMySqlDatabaseInsight) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MdsMySqlDatabaseInsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDatabaseConnectionStatusDetails returns DatabaseConnectionStatusDetails
func (m MdsMySqlDatabaseInsight) GetDatabaseConnectionStatusDetails() *string {
	return m.DatabaseConnectionStatusDetails
}

func (m MdsMySqlDatabaseInsight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MdsMySqlDatabaseInsight) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetResourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MdsMySqlDatabaseInsight) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMdsMySqlDatabaseInsight MdsMySqlDatabaseInsight
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeMdsMySqlDatabaseInsight
	}{
		"MDS_MYSQL_DATABASE_SYSTEM",
		(MarshalTypeMdsMySqlDatabaseInsight)(m),
	}

	return json.Marshal(&s)
}
