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

// DbmResource The base Exadata resource.
type DbmResource interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	GetId() *string

	// The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	GetDisplayName() *string

	// The version of the Exadata resource.
	GetVersion() *string

	// The internal ID of the Exadata resource.
	GetInternalId() *string

	// The status of the Exadata resource.
	GetStatus() *string

	// The current lifecycle state of the database resource.
	GetLifecycleState() DbmResourceLifecycleStateEnum

	// The timestamp of the creation of the Exadata resource.
	GetTimeCreated() *common.SDKTime

	// The timestamp of the last update of the Exadata resource.
	GetTimeUpdated() *common.SDKTime

	// The details of the lifecycle state of the Exadata resource.
	GetLifecycleDetails() *string

	// The additional details of the resource defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	GetAdditionalDetails() map[string]string
}

type dbmresource struct {
	JsonData          []byte
	Version           *string                       `mandatory:"false" json:"version"`
	InternalId        *string                       `mandatory:"false" json:"internalId"`
	Status            *string                       `mandatory:"false" json:"status"`
	LifecycleState    DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	TimeCreated       *common.SDKTime               `mandatory:"false" json:"timeCreated"`
	TimeUpdated       *common.SDKTime               `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails  *string                       `mandatory:"false" json:"lifecycleDetails"`
	AdditionalDetails map[string]string             `mandatory:"false" json:"additionalDetails"`
	Id                *string                       `mandatory:"true" json:"id"`
	DisplayName       *string                       `mandatory:"true" json:"displayName"`
	ResourceType      string                        `json:"resourceType"`
}

// UnmarshalJSON unmarshals json
func (m *dbmresource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdbmresource dbmresource
	s := struct {
		Model Unmarshalerdbmresource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.Version = s.Model.Version
	m.InternalId = s.Model.InternalId
	m.Status = s.Model.Status
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.AdditionalDetails = s.Model.AdditionalDetails
	m.ResourceType = s.Model.ResourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dbmresource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ResourceType {
	case "STORAGE_CONNECTOR":
		mm := ExternalExadataStorageConnector{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STORAGE_GRID_SUMMARY":
		mm := ExternalExadataStorageGridSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STORAGE_SERVER":
		mm := ExternalExadataStorageServer{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INFRASTRUCTURE":
		mm := ExternalExadataInfrastructure{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STORAGE_GRID":
		mm := ExternalExadataStorageGrid{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INFRASTRUCTURE_SUMMARY":
		mm := ExternalExadataInfrastructureSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_SYSTEM_SUMMARY":
		mm := ExternalExadataDatabaseSystemSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STORAGE_CONNECTOR_SUMMARY":
		mm := ExternalExadataStorageConnectorSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STORAGE_SERVER_SUMMARY":
		mm := ExternalExadataStorageServerSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DbmResource: %s.", m.ResourceType)
		return *m, nil
	}
}

// GetVersion returns Version
func (m dbmresource) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m dbmresource) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m dbmresource) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m dbmresource) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m dbmresource) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m dbmresource) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m dbmresource) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m dbmresource) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

// GetId returns Id
func (m dbmresource) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m dbmresource) GetDisplayName() *string {
	return m.DisplayName
}

func (m dbmresource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dbmresource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbmResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbmResourceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbmResourceLifecycleStateEnum Enum with underlying type: string
type DbmResourceLifecycleStateEnum string

// Set of constants representing the allowable values for DbmResourceLifecycleStateEnum
const (
	DbmResourceLifecycleStateCreating DbmResourceLifecycleStateEnum = "CREATING"
	DbmResourceLifecycleStateActive   DbmResourceLifecycleStateEnum = "ACTIVE"
	DbmResourceLifecycleStateInactive DbmResourceLifecycleStateEnum = "INACTIVE"
	DbmResourceLifecycleStateUpdating DbmResourceLifecycleStateEnum = "UPDATING"
	DbmResourceLifecycleStateDeleting DbmResourceLifecycleStateEnum = "DELETING"
	DbmResourceLifecycleStateDeleted  DbmResourceLifecycleStateEnum = "DELETED"
	DbmResourceLifecycleStateFailed   DbmResourceLifecycleStateEnum = "FAILED"
)

var mappingDbmResourceLifecycleStateEnum = map[string]DbmResourceLifecycleStateEnum{
	"CREATING": DbmResourceLifecycleStateCreating,
	"ACTIVE":   DbmResourceLifecycleStateActive,
	"INACTIVE": DbmResourceLifecycleStateInactive,
	"UPDATING": DbmResourceLifecycleStateUpdating,
	"DELETING": DbmResourceLifecycleStateDeleting,
	"DELETED":  DbmResourceLifecycleStateDeleted,
	"FAILED":   DbmResourceLifecycleStateFailed,
}

var mappingDbmResourceLifecycleStateEnumLowerCase = map[string]DbmResourceLifecycleStateEnum{
	"creating": DbmResourceLifecycleStateCreating,
	"active":   DbmResourceLifecycleStateActive,
	"inactive": DbmResourceLifecycleStateInactive,
	"updating": DbmResourceLifecycleStateUpdating,
	"deleting": DbmResourceLifecycleStateDeleting,
	"deleted":  DbmResourceLifecycleStateDeleted,
	"failed":   DbmResourceLifecycleStateFailed,
}

// GetDbmResourceLifecycleStateEnumValues Enumerates the set of values for DbmResourceLifecycleStateEnum
func GetDbmResourceLifecycleStateEnumValues() []DbmResourceLifecycleStateEnum {
	values := make([]DbmResourceLifecycleStateEnum, 0)
	for _, v := range mappingDbmResourceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbmResourceLifecycleStateEnumStringValues Enumerates the set of values in String for DbmResourceLifecycleStateEnum
func GetDbmResourceLifecycleStateEnumStringValues() []string {
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

// GetMappingDbmResourceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbmResourceLifecycleStateEnum(val string) (DbmResourceLifecycleStateEnum, bool) {
	enum, ok := mappingDbmResourceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbmResourceResourceTypeEnum Enum with underlying type: string
type DbmResourceResourceTypeEnum string

// Set of constants representing the allowable values for DbmResourceResourceTypeEnum
const (
	DbmResourceResourceTypeInfrastructureSummary   DbmResourceResourceTypeEnum = "INFRASTRUCTURE_SUMMARY"
	DbmResourceResourceTypeInfrastructure          DbmResourceResourceTypeEnum = "INFRASTRUCTURE"
	DbmResourceResourceTypeStorageServerSummary    DbmResourceResourceTypeEnum = "STORAGE_SERVER_SUMMARY"
	DbmResourceResourceTypeStorageServer           DbmResourceResourceTypeEnum = "STORAGE_SERVER"
	DbmResourceResourceTypeStorageGridSummary      DbmResourceResourceTypeEnum = "STORAGE_GRID_SUMMARY"
	DbmResourceResourceTypeStorageGrid             DbmResourceResourceTypeEnum = "STORAGE_GRID"
	DbmResourceResourceTypeStorageConnectorSummary DbmResourceResourceTypeEnum = "STORAGE_CONNECTOR_SUMMARY"
	DbmResourceResourceTypeStorageConnector        DbmResourceResourceTypeEnum = "STORAGE_CONNECTOR"
	DbmResourceResourceTypeDatabaseSystemSummary   DbmResourceResourceTypeEnum = "DATABASE_SYSTEM_SUMMARY"
	DbmResourceResourceTypeDatabaseSummary         DbmResourceResourceTypeEnum = "DATABASE_SUMMARY"
)

var mappingDbmResourceResourceTypeEnum = map[string]DbmResourceResourceTypeEnum{
	"INFRASTRUCTURE_SUMMARY":    DbmResourceResourceTypeInfrastructureSummary,
	"INFRASTRUCTURE":            DbmResourceResourceTypeInfrastructure,
	"STORAGE_SERVER_SUMMARY":    DbmResourceResourceTypeStorageServerSummary,
	"STORAGE_SERVER":            DbmResourceResourceTypeStorageServer,
	"STORAGE_GRID_SUMMARY":      DbmResourceResourceTypeStorageGridSummary,
	"STORAGE_GRID":              DbmResourceResourceTypeStorageGrid,
	"STORAGE_CONNECTOR_SUMMARY": DbmResourceResourceTypeStorageConnectorSummary,
	"STORAGE_CONNECTOR":         DbmResourceResourceTypeStorageConnector,
	"DATABASE_SYSTEM_SUMMARY":   DbmResourceResourceTypeDatabaseSystemSummary,
	"DATABASE_SUMMARY":          DbmResourceResourceTypeDatabaseSummary,
}

var mappingDbmResourceResourceTypeEnumLowerCase = map[string]DbmResourceResourceTypeEnum{
	"infrastructure_summary":    DbmResourceResourceTypeInfrastructureSummary,
	"infrastructure":            DbmResourceResourceTypeInfrastructure,
	"storage_server_summary":    DbmResourceResourceTypeStorageServerSummary,
	"storage_server":            DbmResourceResourceTypeStorageServer,
	"storage_grid_summary":      DbmResourceResourceTypeStorageGridSummary,
	"storage_grid":              DbmResourceResourceTypeStorageGrid,
	"storage_connector_summary": DbmResourceResourceTypeStorageConnectorSummary,
	"storage_connector":         DbmResourceResourceTypeStorageConnector,
	"database_system_summary":   DbmResourceResourceTypeDatabaseSystemSummary,
	"database_summary":          DbmResourceResourceTypeDatabaseSummary,
}

// GetDbmResourceResourceTypeEnumValues Enumerates the set of values for DbmResourceResourceTypeEnum
func GetDbmResourceResourceTypeEnumValues() []DbmResourceResourceTypeEnum {
	values := make([]DbmResourceResourceTypeEnum, 0)
	for _, v := range mappingDbmResourceResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbmResourceResourceTypeEnumStringValues Enumerates the set of values in String for DbmResourceResourceTypeEnum
func GetDbmResourceResourceTypeEnumStringValues() []string {
	return []string{
		"INFRASTRUCTURE_SUMMARY",
		"INFRASTRUCTURE",
		"STORAGE_SERVER_SUMMARY",
		"STORAGE_SERVER",
		"STORAGE_GRID_SUMMARY",
		"STORAGE_GRID",
		"STORAGE_CONNECTOR_SUMMARY",
		"STORAGE_CONNECTOR",
		"DATABASE_SYSTEM_SUMMARY",
		"DATABASE_SUMMARY",
	}
}

// GetMappingDbmResourceResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbmResourceResourceTypeEnum(val string) (DbmResourceResourceTypeEnum, bool) {
	enum, ok := mappingDbmResourceResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
