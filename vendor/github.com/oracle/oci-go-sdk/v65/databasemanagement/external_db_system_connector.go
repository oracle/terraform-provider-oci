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

// ExternalDbSystemConnector The details of an external DB system connector.
type ExternalDbSystemConnector interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system connector.
	GetId() *string

	// The user-friendly name for the external connector. The name does not have to be unique.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the connector is a part of.
	GetExternalDbSystemId() *string

	// The current lifecycle state of the external DB system connector.
	GetLifecycleState() ExternalDbSystemConnectorLifecycleStateEnum

	// The date and time the external DB system connector was created.
	GetTimeCreated() *common.SDKTime

	// The date and time the external DB system connector was last updated.
	GetTimeUpdated() *common.SDKTime

	// The status of connectivity to the external DB system component.
	GetConnectionStatus() *string

	// The error message indicating the reason for connection failure or `null` if
	// the connection was successful.
	GetConnectionFailureMessage() *string

	// Additional information about the current lifecycle state.
	GetLifecycleDetails() *string

	// The date and time the connectionStatus of the external DB system connector was last updated.
	GetTimeConnectionStatusLastUpdated() *common.SDKTime
}

type externaldbsystemconnector struct {
	JsonData                        []byte
	ConnectionStatus                *string                                     `mandatory:"false" json:"connectionStatus"`
	ConnectionFailureMessage        *string                                     `mandatory:"false" json:"connectionFailureMessage"`
	LifecycleDetails                *string                                     `mandatory:"false" json:"lifecycleDetails"`
	TimeConnectionStatusLastUpdated *common.SDKTime                             `mandatory:"false" json:"timeConnectionStatusLastUpdated"`
	Id                              *string                                     `mandatory:"true" json:"id"`
	DisplayName                     *string                                     `mandatory:"true" json:"displayName"`
	CompartmentId                   *string                                     `mandatory:"true" json:"compartmentId"`
	ExternalDbSystemId              *string                                     `mandatory:"true" json:"externalDbSystemId"`
	LifecycleState                  ExternalDbSystemConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	TimeCreated                     *common.SDKTime                             `mandatory:"true" json:"timeCreated"`
	TimeUpdated                     *common.SDKTime                             `mandatory:"true" json:"timeUpdated"`
	ConnectorType                   string                                      `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *externaldbsystemconnector) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexternaldbsystemconnector externaldbsystemconnector
	s := struct {
		Model Unmarshalerexternaldbsystemconnector
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.ExternalDbSystemId = s.Model.ExternalDbSystemId
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.ConnectionStatus = s.Model.ConnectionStatus
	m.ConnectionFailureMessage = s.Model.ConnectionFailureMessage
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.TimeConnectionStatusLastUpdated = s.Model.TimeConnectionStatusLastUpdated
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *externaldbsystemconnector) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := ExternalDbSystemMacsConnector{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ExternalDbSystemConnector: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetConnectionStatus returns ConnectionStatus
func (m externaldbsystemconnector) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

// GetConnectionFailureMessage returns ConnectionFailureMessage
func (m externaldbsystemconnector) GetConnectionFailureMessage() *string {
	return m.ConnectionFailureMessage
}

// GetLifecycleDetails returns LifecycleDetails
func (m externaldbsystemconnector) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m externaldbsystemconnector) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

// GetId returns Id
func (m externaldbsystemconnector) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m externaldbsystemconnector) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m externaldbsystemconnector) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetExternalDbSystemId returns ExternalDbSystemId
func (m externaldbsystemconnector) GetExternalDbSystemId() *string {
	return m.ExternalDbSystemId
}

// GetLifecycleState returns LifecycleState
func (m externaldbsystemconnector) GetLifecycleState() ExternalDbSystemConnectorLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m externaldbsystemconnector) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m externaldbsystemconnector) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m externaldbsystemconnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m externaldbsystemconnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDbSystemConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDbSystemConnectorLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDbSystemConnectorLifecycleStateEnum Enum with underlying type: string
type ExternalDbSystemConnectorLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalDbSystemConnectorLifecycleStateEnum
const (
	ExternalDbSystemConnectorLifecycleStateCreating     ExternalDbSystemConnectorLifecycleStateEnum = "CREATING"
	ExternalDbSystemConnectorLifecycleStateNotConnected ExternalDbSystemConnectorLifecycleStateEnum = "NOT_CONNECTED"
	ExternalDbSystemConnectorLifecycleStateActive       ExternalDbSystemConnectorLifecycleStateEnum = "ACTIVE"
	ExternalDbSystemConnectorLifecycleStateInactive     ExternalDbSystemConnectorLifecycleStateEnum = "INACTIVE"
	ExternalDbSystemConnectorLifecycleStateUpdating     ExternalDbSystemConnectorLifecycleStateEnum = "UPDATING"
	ExternalDbSystemConnectorLifecycleStateDeleting     ExternalDbSystemConnectorLifecycleStateEnum = "DELETING"
	ExternalDbSystemConnectorLifecycleStateDeleted      ExternalDbSystemConnectorLifecycleStateEnum = "DELETED"
	ExternalDbSystemConnectorLifecycleStateFailed       ExternalDbSystemConnectorLifecycleStateEnum = "FAILED"
)

var mappingExternalDbSystemConnectorLifecycleStateEnum = map[string]ExternalDbSystemConnectorLifecycleStateEnum{
	"CREATING":      ExternalDbSystemConnectorLifecycleStateCreating,
	"NOT_CONNECTED": ExternalDbSystemConnectorLifecycleStateNotConnected,
	"ACTIVE":        ExternalDbSystemConnectorLifecycleStateActive,
	"INACTIVE":      ExternalDbSystemConnectorLifecycleStateInactive,
	"UPDATING":      ExternalDbSystemConnectorLifecycleStateUpdating,
	"DELETING":      ExternalDbSystemConnectorLifecycleStateDeleting,
	"DELETED":       ExternalDbSystemConnectorLifecycleStateDeleted,
	"FAILED":        ExternalDbSystemConnectorLifecycleStateFailed,
}

var mappingExternalDbSystemConnectorLifecycleStateEnumLowerCase = map[string]ExternalDbSystemConnectorLifecycleStateEnum{
	"creating":      ExternalDbSystemConnectorLifecycleStateCreating,
	"not_connected": ExternalDbSystemConnectorLifecycleStateNotConnected,
	"active":        ExternalDbSystemConnectorLifecycleStateActive,
	"inactive":      ExternalDbSystemConnectorLifecycleStateInactive,
	"updating":      ExternalDbSystemConnectorLifecycleStateUpdating,
	"deleting":      ExternalDbSystemConnectorLifecycleStateDeleting,
	"deleted":       ExternalDbSystemConnectorLifecycleStateDeleted,
	"failed":        ExternalDbSystemConnectorLifecycleStateFailed,
}

// GetExternalDbSystemConnectorLifecycleStateEnumValues Enumerates the set of values for ExternalDbSystemConnectorLifecycleStateEnum
func GetExternalDbSystemConnectorLifecycleStateEnumValues() []ExternalDbSystemConnectorLifecycleStateEnum {
	values := make([]ExternalDbSystemConnectorLifecycleStateEnum, 0)
	for _, v := range mappingExternalDbSystemConnectorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemConnectorLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalDbSystemConnectorLifecycleStateEnum
func GetExternalDbSystemConnectorLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NOT_CONNECTED",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingExternalDbSystemConnectorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemConnectorLifecycleStateEnum(val string) (ExternalDbSystemConnectorLifecycleStateEnum, bool) {
	enum, ok := mappingExternalDbSystemConnectorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalDbSystemConnectorConnectorTypeEnum Enum with underlying type: string
type ExternalDbSystemConnectorConnectorTypeEnum string

// Set of constants representing the allowable values for ExternalDbSystemConnectorConnectorTypeEnum
const (
	ExternalDbSystemConnectorConnectorTypeMacs ExternalDbSystemConnectorConnectorTypeEnum = "MACS"
)

var mappingExternalDbSystemConnectorConnectorTypeEnum = map[string]ExternalDbSystemConnectorConnectorTypeEnum{
	"MACS": ExternalDbSystemConnectorConnectorTypeMacs,
}

var mappingExternalDbSystemConnectorConnectorTypeEnumLowerCase = map[string]ExternalDbSystemConnectorConnectorTypeEnum{
	"macs": ExternalDbSystemConnectorConnectorTypeMacs,
}

// GetExternalDbSystemConnectorConnectorTypeEnumValues Enumerates the set of values for ExternalDbSystemConnectorConnectorTypeEnum
func GetExternalDbSystemConnectorConnectorTypeEnumValues() []ExternalDbSystemConnectorConnectorTypeEnum {
	values := make([]ExternalDbSystemConnectorConnectorTypeEnum, 0)
	for _, v := range mappingExternalDbSystemConnectorConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemConnectorConnectorTypeEnumStringValues Enumerates the set of values in String for ExternalDbSystemConnectorConnectorTypeEnum
func GetExternalDbSystemConnectorConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingExternalDbSystemConnectorConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemConnectorConnectorTypeEnum(val string) (ExternalDbSystemConnectorConnectorTypeEnum, bool) {
	enum, ok := mappingExternalDbSystemConnectorConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
