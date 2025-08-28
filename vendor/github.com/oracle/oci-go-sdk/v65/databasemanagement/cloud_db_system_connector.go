// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudDbSystemConnector The details of a cloud DB system connector.
type CloudDbSystemConnector interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system connector.
	GetId() *string

	// The user-friendly name for the cloud connector. The name does not have to be unique.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the connector is a part of.
	GetCloudDbSystemId() *string

	// The current lifecycle state of the cloud DB system connector.
	GetLifecycleState() CloudDbSystemConnectorLifecycleStateEnum

	// The date and time the cloud DB system connector was created.
	GetTimeCreated() *common.SDKTime

	// The date and time the cloud DB system connector was last updated.
	GetTimeUpdated() *common.SDKTime

	// The status of connectivity to the cloud DB system component.
	GetConnectionStatus() *string

	// The error message indicating the reason for connection failure or `null` if
	// the connection was successful.
	GetConnectionFailureMessage() *string

	// Additional information about the current lifecycle state.
	GetLifecycleDetails() *string

	// The date and time the connectionStatus of the cloud DB system connector was last updated.
	GetTimeConnectionStatusLastUpdated() *common.SDKTime
}

type clouddbsystemconnector struct {
	JsonData                        []byte
	ConnectionStatus                *string                                  `mandatory:"false" json:"connectionStatus"`
	ConnectionFailureMessage        *string                                  `mandatory:"false" json:"connectionFailureMessage"`
	LifecycleDetails                *string                                  `mandatory:"false" json:"lifecycleDetails"`
	TimeConnectionStatusLastUpdated *common.SDKTime                          `mandatory:"false" json:"timeConnectionStatusLastUpdated"`
	Id                              *string                                  `mandatory:"true" json:"id"`
	DisplayName                     *string                                  `mandatory:"true" json:"displayName"`
	CompartmentId                   *string                                  `mandatory:"true" json:"compartmentId"`
	CloudDbSystemId                 *string                                  `mandatory:"true" json:"cloudDbSystemId"`
	LifecycleState                  CloudDbSystemConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	TimeCreated                     *common.SDKTime                          `mandatory:"true" json:"timeCreated"`
	TimeUpdated                     *common.SDKTime                          `mandatory:"true" json:"timeUpdated"`
	ConnectorType                   string                                   `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *clouddbsystemconnector) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerclouddbsystemconnector clouddbsystemconnector
	s := struct {
		Model Unmarshalerclouddbsystemconnector
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.CloudDbSystemId = s.Model.CloudDbSystemId
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
func (m *clouddbsystemconnector) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := CloudDbSystemMacsConnector{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CloudDbSystemConnector: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetConnectionStatus returns ConnectionStatus
func (m clouddbsystemconnector) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

// GetConnectionFailureMessage returns ConnectionFailureMessage
func (m clouddbsystemconnector) GetConnectionFailureMessage() *string {
	return m.ConnectionFailureMessage
}

// GetLifecycleDetails returns LifecycleDetails
func (m clouddbsystemconnector) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m clouddbsystemconnector) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

// GetId returns Id
func (m clouddbsystemconnector) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m clouddbsystemconnector) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m clouddbsystemconnector) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetCloudDbSystemId returns CloudDbSystemId
func (m clouddbsystemconnector) GetCloudDbSystemId() *string {
	return m.CloudDbSystemId
}

// GetLifecycleState returns LifecycleState
func (m clouddbsystemconnector) GetLifecycleState() CloudDbSystemConnectorLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m clouddbsystemconnector) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m clouddbsystemconnector) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m clouddbsystemconnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m clouddbsystemconnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudDbSystemConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudDbSystemConnectorLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudDbSystemConnectorLifecycleStateEnum Enum with underlying type: string
type CloudDbSystemConnectorLifecycleStateEnum string

// Set of constants representing the allowable values for CloudDbSystemConnectorLifecycleStateEnum
const (
	CloudDbSystemConnectorLifecycleStateCreating     CloudDbSystemConnectorLifecycleStateEnum = "CREATING"
	CloudDbSystemConnectorLifecycleStateNotConnected CloudDbSystemConnectorLifecycleStateEnum = "NOT_CONNECTED"
	CloudDbSystemConnectorLifecycleStateActive       CloudDbSystemConnectorLifecycleStateEnum = "ACTIVE"
	CloudDbSystemConnectorLifecycleStateInactive     CloudDbSystemConnectorLifecycleStateEnum = "INACTIVE"
	CloudDbSystemConnectorLifecycleStateUpdating     CloudDbSystemConnectorLifecycleStateEnum = "UPDATING"
	CloudDbSystemConnectorLifecycleStateDeleting     CloudDbSystemConnectorLifecycleStateEnum = "DELETING"
	CloudDbSystemConnectorLifecycleStateDeleted      CloudDbSystemConnectorLifecycleStateEnum = "DELETED"
	CloudDbSystemConnectorLifecycleStateFailed       CloudDbSystemConnectorLifecycleStateEnum = "FAILED"
)

var mappingCloudDbSystemConnectorLifecycleStateEnum = map[string]CloudDbSystemConnectorLifecycleStateEnum{
	"CREATING":      CloudDbSystemConnectorLifecycleStateCreating,
	"NOT_CONNECTED": CloudDbSystemConnectorLifecycleStateNotConnected,
	"ACTIVE":        CloudDbSystemConnectorLifecycleStateActive,
	"INACTIVE":      CloudDbSystemConnectorLifecycleStateInactive,
	"UPDATING":      CloudDbSystemConnectorLifecycleStateUpdating,
	"DELETING":      CloudDbSystemConnectorLifecycleStateDeleting,
	"DELETED":       CloudDbSystemConnectorLifecycleStateDeleted,
	"FAILED":        CloudDbSystemConnectorLifecycleStateFailed,
}

var mappingCloudDbSystemConnectorLifecycleStateEnumLowerCase = map[string]CloudDbSystemConnectorLifecycleStateEnum{
	"creating":      CloudDbSystemConnectorLifecycleStateCreating,
	"not_connected": CloudDbSystemConnectorLifecycleStateNotConnected,
	"active":        CloudDbSystemConnectorLifecycleStateActive,
	"inactive":      CloudDbSystemConnectorLifecycleStateInactive,
	"updating":      CloudDbSystemConnectorLifecycleStateUpdating,
	"deleting":      CloudDbSystemConnectorLifecycleStateDeleting,
	"deleted":       CloudDbSystemConnectorLifecycleStateDeleted,
	"failed":        CloudDbSystemConnectorLifecycleStateFailed,
}

// GetCloudDbSystemConnectorLifecycleStateEnumValues Enumerates the set of values for CloudDbSystemConnectorLifecycleStateEnum
func GetCloudDbSystemConnectorLifecycleStateEnumValues() []CloudDbSystemConnectorLifecycleStateEnum {
	values := make([]CloudDbSystemConnectorLifecycleStateEnum, 0)
	for _, v := range mappingCloudDbSystemConnectorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbSystemConnectorLifecycleStateEnumStringValues Enumerates the set of values in String for CloudDbSystemConnectorLifecycleStateEnum
func GetCloudDbSystemConnectorLifecycleStateEnumStringValues() []string {
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

// GetMappingCloudDbSystemConnectorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbSystemConnectorLifecycleStateEnum(val string) (CloudDbSystemConnectorLifecycleStateEnum, bool) {
	enum, ok := mappingCloudDbSystemConnectorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudDbSystemConnectorConnectorTypeEnum Enum with underlying type: string
type CloudDbSystemConnectorConnectorTypeEnum string

// Set of constants representing the allowable values for CloudDbSystemConnectorConnectorTypeEnum
const (
	CloudDbSystemConnectorConnectorTypeMacs CloudDbSystemConnectorConnectorTypeEnum = "MACS"
)

var mappingCloudDbSystemConnectorConnectorTypeEnum = map[string]CloudDbSystemConnectorConnectorTypeEnum{
	"MACS": CloudDbSystemConnectorConnectorTypeMacs,
}

var mappingCloudDbSystemConnectorConnectorTypeEnumLowerCase = map[string]CloudDbSystemConnectorConnectorTypeEnum{
	"macs": CloudDbSystemConnectorConnectorTypeMacs,
}

// GetCloudDbSystemConnectorConnectorTypeEnumValues Enumerates the set of values for CloudDbSystemConnectorConnectorTypeEnum
func GetCloudDbSystemConnectorConnectorTypeEnumValues() []CloudDbSystemConnectorConnectorTypeEnum {
	values := make([]CloudDbSystemConnectorConnectorTypeEnum, 0)
	for _, v := range mappingCloudDbSystemConnectorConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbSystemConnectorConnectorTypeEnumStringValues Enumerates the set of values in String for CloudDbSystemConnectorConnectorTypeEnum
func GetCloudDbSystemConnectorConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingCloudDbSystemConnectorConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbSystemConnectorConnectorTypeEnum(val string) (CloudDbSystemConnectorConnectorTypeEnum, bool) {
	enum, ok := mappingCloudDbSystemConnectorConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
