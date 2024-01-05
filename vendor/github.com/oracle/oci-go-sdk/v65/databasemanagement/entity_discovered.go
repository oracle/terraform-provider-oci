// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EntityDiscovered The details of the base entity discovery.
type EntityDiscovered interface {

	// The name of the entity.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the entity discovered.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the agent used for monitoring.
	GetAgentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated connector.
	GetConnectorId() *string

	// The version of the entity.
	GetVersion() *string

	// The internal identifier of the entity.
	GetInternalId() *string

	// The status of the entity.
	GetStatus() *string

	// The status of the entity discovery.
	GetDiscoverStatus() EntityDiscoveredDiscoverStatusEnum

	// The error code of the discovery.
	GetDiscoverErrorCode() *string

	// The error message of the discovery.
	GetDiscoverErrorMsg() *string
}

type entitydiscovered struct {
	JsonData          []byte
	Id                *string                            `mandatory:"false" json:"id"`
	AgentId           *string                            `mandatory:"false" json:"agentId"`
	ConnectorId       *string                            `mandatory:"false" json:"connectorId"`
	Version           *string                            `mandatory:"false" json:"version"`
	InternalId        *string                            `mandatory:"false" json:"internalId"`
	Status            *string                            `mandatory:"false" json:"status"`
	DiscoverStatus    EntityDiscoveredDiscoverStatusEnum `mandatory:"false" json:"discoverStatus,omitempty"`
	DiscoverErrorCode *string                            `mandatory:"false" json:"discoverErrorCode"`
	DiscoverErrorMsg  *string                            `mandatory:"false" json:"discoverErrorMsg"`
	DisplayName       *string                            `mandatory:"true" json:"displayName"`
	EntityType        string                             `json:"entityType"`
}

// UnmarshalJSON unmarshals json
func (m *entitydiscovered) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerentitydiscovered entitydiscovered
	s := struct {
		Model Unmarshalerentitydiscovered
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Id = s.Model.Id
	m.AgentId = s.Model.AgentId
	m.ConnectorId = s.Model.ConnectorId
	m.Version = s.Model.Version
	m.InternalId = s.Model.InternalId
	m.Status = s.Model.Status
	m.DiscoverStatus = s.Model.DiscoverStatus
	m.DiscoverErrorCode = s.Model.DiscoverErrorCode
	m.DiscoverErrorMsg = s.Model.DiscoverErrorMsg
	m.EntityType = s.Model.EntityType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *entitydiscovered) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntityType {
	case "STORAGE_GRID_DISCOVER_SUMMARY":
		mm := ExternalStorageGridDiscoverySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INFRASTRUCTURE_DISCOVER":
		mm := ExternalExadataInfrastructureDiscovery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_SYSTEM_DISCOVER_SUMMARY":
		mm := ExternalDatabaseSystemDiscoverySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INFRASTRUCTURE_DISCOVER_SUMMARY":
		mm := ExternalExadataInfrastructureDiscoverySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STORAGE_SERVER_DISCOVER_SUMMARY":
		mm := ExternalStorageServerDiscoverySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for EntityDiscovered: %s.", m.EntityType)
		return *m, nil
	}
}

// GetId returns Id
func (m entitydiscovered) GetId() *string {
	return m.Id
}

// GetAgentId returns AgentId
func (m entitydiscovered) GetAgentId() *string {
	return m.AgentId
}

// GetConnectorId returns ConnectorId
func (m entitydiscovered) GetConnectorId() *string {
	return m.ConnectorId
}

// GetVersion returns Version
func (m entitydiscovered) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m entitydiscovered) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m entitydiscovered) GetStatus() *string {
	return m.Status
}

// GetDiscoverStatus returns DiscoverStatus
func (m entitydiscovered) GetDiscoverStatus() EntityDiscoveredDiscoverStatusEnum {
	return m.DiscoverStatus
}

// GetDiscoverErrorCode returns DiscoverErrorCode
func (m entitydiscovered) GetDiscoverErrorCode() *string {
	return m.DiscoverErrorCode
}

// GetDiscoverErrorMsg returns DiscoverErrorMsg
func (m entitydiscovered) GetDiscoverErrorMsg() *string {
	return m.DiscoverErrorMsg
}

// GetDisplayName returns DisplayName
func (m entitydiscovered) GetDisplayName() *string {
	return m.DisplayName
}

func (m entitydiscovered) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m entitydiscovered) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEntityDiscoveredDiscoverStatusEnum(string(m.DiscoverStatus)); !ok && m.DiscoverStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoverStatus: %s. Supported values are: %s.", m.DiscoverStatus, strings.Join(GetEntityDiscoveredDiscoverStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EntityDiscoveredDiscoverStatusEnum Enum with underlying type: string
type EntityDiscoveredDiscoverStatusEnum string

// Set of constants representing the allowable values for EntityDiscoveredDiscoverStatusEnum
const (
	EntityDiscoveredDiscoverStatusPrevDiscovered EntityDiscoveredDiscoverStatusEnum = "PREV_DISCOVERED"
	EntityDiscoveredDiscoverStatusNewDiscovered  EntityDiscoveredDiscoverStatusEnum = "NEW_DISCOVERED"
	EntityDiscoveredDiscoverStatusNotFound       EntityDiscoveredDiscoverStatusEnum = "NOT_FOUND"
	EntityDiscoveredDiscoverStatusDiscovering    EntityDiscoveredDiscoverStatusEnum = "DISCOVERING"
)

var mappingEntityDiscoveredDiscoverStatusEnum = map[string]EntityDiscoveredDiscoverStatusEnum{
	"PREV_DISCOVERED": EntityDiscoveredDiscoverStatusPrevDiscovered,
	"NEW_DISCOVERED":  EntityDiscoveredDiscoverStatusNewDiscovered,
	"NOT_FOUND":       EntityDiscoveredDiscoverStatusNotFound,
	"DISCOVERING":     EntityDiscoveredDiscoverStatusDiscovering,
}

var mappingEntityDiscoveredDiscoverStatusEnumLowerCase = map[string]EntityDiscoveredDiscoverStatusEnum{
	"prev_discovered": EntityDiscoveredDiscoverStatusPrevDiscovered,
	"new_discovered":  EntityDiscoveredDiscoverStatusNewDiscovered,
	"not_found":       EntityDiscoveredDiscoverStatusNotFound,
	"discovering":     EntityDiscoveredDiscoverStatusDiscovering,
}

// GetEntityDiscoveredDiscoverStatusEnumValues Enumerates the set of values for EntityDiscoveredDiscoverStatusEnum
func GetEntityDiscoveredDiscoverStatusEnumValues() []EntityDiscoveredDiscoverStatusEnum {
	values := make([]EntityDiscoveredDiscoverStatusEnum, 0)
	for _, v := range mappingEntityDiscoveredDiscoverStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityDiscoveredDiscoverStatusEnumStringValues Enumerates the set of values in String for EntityDiscoveredDiscoverStatusEnum
func GetEntityDiscoveredDiscoverStatusEnumStringValues() []string {
	return []string{
		"PREV_DISCOVERED",
		"NEW_DISCOVERED",
		"NOT_FOUND",
		"DISCOVERING",
	}
}

// GetMappingEntityDiscoveredDiscoverStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityDiscoveredDiscoverStatusEnum(val string) (EntityDiscoveredDiscoverStatusEnum, bool) {
	enum, ok := mappingEntityDiscoveredDiscoverStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EntityDiscoveredEntityTypeEnum Enum with underlying type: string
type EntityDiscoveredEntityTypeEnum string

// Set of constants representing the allowable values for EntityDiscoveredEntityTypeEnum
const (
	EntityDiscoveredEntityTypeStorageServerDiscoverSummary  EntityDiscoveredEntityTypeEnum = "STORAGE_SERVER_DISCOVER_SUMMARY"
	EntityDiscoveredEntityTypeStorageGridDiscoverSummary    EntityDiscoveredEntityTypeEnum = "STORAGE_GRID_DISCOVER_SUMMARY"
	EntityDiscoveredEntityTypeDatabaseSystemDiscoverSummary EntityDiscoveredEntityTypeEnum = "DATABASE_SYSTEM_DISCOVER_SUMMARY"
	EntityDiscoveredEntityTypeInfrastructureDiscoverSummary EntityDiscoveredEntityTypeEnum = "INFRASTRUCTURE_DISCOVER_SUMMARY"
	EntityDiscoveredEntityTypeInfrastructureDiscover        EntityDiscoveredEntityTypeEnum = "INFRASTRUCTURE_DISCOVER"
)

var mappingEntityDiscoveredEntityTypeEnum = map[string]EntityDiscoveredEntityTypeEnum{
	"STORAGE_SERVER_DISCOVER_SUMMARY":  EntityDiscoveredEntityTypeStorageServerDiscoverSummary,
	"STORAGE_GRID_DISCOVER_SUMMARY":    EntityDiscoveredEntityTypeStorageGridDiscoverSummary,
	"DATABASE_SYSTEM_DISCOVER_SUMMARY": EntityDiscoveredEntityTypeDatabaseSystemDiscoverSummary,
	"INFRASTRUCTURE_DISCOVER_SUMMARY":  EntityDiscoveredEntityTypeInfrastructureDiscoverSummary,
	"INFRASTRUCTURE_DISCOVER":          EntityDiscoveredEntityTypeInfrastructureDiscover,
}

var mappingEntityDiscoveredEntityTypeEnumLowerCase = map[string]EntityDiscoveredEntityTypeEnum{
	"storage_server_discover_summary":  EntityDiscoveredEntityTypeStorageServerDiscoverSummary,
	"storage_grid_discover_summary":    EntityDiscoveredEntityTypeStorageGridDiscoverSummary,
	"database_system_discover_summary": EntityDiscoveredEntityTypeDatabaseSystemDiscoverSummary,
	"infrastructure_discover_summary":  EntityDiscoveredEntityTypeInfrastructureDiscoverSummary,
	"infrastructure_discover":          EntityDiscoveredEntityTypeInfrastructureDiscover,
}

// GetEntityDiscoveredEntityTypeEnumValues Enumerates the set of values for EntityDiscoveredEntityTypeEnum
func GetEntityDiscoveredEntityTypeEnumValues() []EntityDiscoveredEntityTypeEnum {
	values := make([]EntityDiscoveredEntityTypeEnum, 0)
	for _, v := range mappingEntityDiscoveredEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityDiscoveredEntityTypeEnumStringValues Enumerates the set of values in String for EntityDiscoveredEntityTypeEnum
func GetEntityDiscoveredEntityTypeEnumStringValues() []string {
	return []string{
		"STORAGE_SERVER_DISCOVER_SUMMARY",
		"STORAGE_GRID_DISCOVER_SUMMARY",
		"DATABASE_SYSTEM_DISCOVER_SUMMARY",
		"INFRASTRUCTURE_DISCOVER_SUMMARY",
		"INFRASTRUCTURE_DISCOVER",
	}
}

// GetMappingEntityDiscoveredEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityDiscoveredEntityTypeEnum(val string) (EntityDiscoveredEntityTypeEnum, bool) {
	enum, ok := mappingEntityDiscoveredEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
