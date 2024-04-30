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

// ExternalDbSystemDiscoveryConnector The connector details used to connect to the external DB system component.
type ExternalDbSystemDiscoveryConnector interface {

	// The user-friendly name for the external connector. The name does not have to be unique.
	GetDisplayName() *string

	// The status of connectivity to the external DB system component.
	GetConnectionStatus() *string

	// The error message indicating the reason for connection failure or `null` if
	// the connection was successful.
	GetConnectionFailureMessage() *string

	// The date and time the connectionStatus of the external DB system connector was last updated.
	GetTimeConnectionStatusLastUpdated() *common.SDKTime
}

type externaldbsystemdiscoveryconnector struct {
	JsonData                        []byte
	ConnectionStatus                *string         `mandatory:"false" json:"connectionStatus"`
	ConnectionFailureMessage        *string         `mandatory:"false" json:"connectionFailureMessage"`
	TimeConnectionStatusLastUpdated *common.SDKTime `mandatory:"false" json:"timeConnectionStatusLastUpdated"`
	DisplayName                     *string         `mandatory:"true" json:"displayName"`
	ConnectorType                   string          `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *externaldbsystemdiscoveryconnector) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexternaldbsystemdiscoveryconnector externaldbsystemdiscoveryconnector
	s := struct {
		Model Unmarshalerexternaldbsystemdiscoveryconnector
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.ConnectionStatus = s.Model.ConnectionStatus
	m.ConnectionFailureMessage = s.Model.ConnectionFailureMessage
	m.TimeConnectionStatusLastUpdated = s.Model.TimeConnectionStatusLastUpdated
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *externaldbsystemdiscoveryconnector) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := ExternalDbSystemDiscoveryMacsConnector{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ExternalDbSystemDiscoveryConnector: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetConnectionStatus returns ConnectionStatus
func (m externaldbsystemdiscoveryconnector) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

// GetConnectionFailureMessage returns ConnectionFailureMessage
func (m externaldbsystemdiscoveryconnector) GetConnectionFailureMessage() *string {
	return m.ConnectionFailureMessage
}

// GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m externaldbsystemdiscoveryconnector) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

// GetDisplayName returns DisplayName
func (m externaldbsystemdiscoveryconnector) GetDisplayName() *string {
	return m.DisplayName
}

func (m externaldbsystemdiscoveryconnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m externaldbsystemdiscoveryconnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDbSystemDiscoveryConnectorConnectorTypeEnum Enum with underlying type: string
type ExternalDbSystemDiscoveryConnectorConnectorTypeEnum string

// Set of constants representing the allowable values for ExternalDbSystemDiscoveryConnectorConnectorTypeEnum
const (
	ExternalDbSystemDiscoveryConnectorConnectorTypeMacs ExternalDbSystemDiscoveryConnectorConnectorTypeEnum = "MACS"
)

var mappingExternalDbSystemDiscoveryConnectorConnectorTypeEnum = map[string]ExternalDbSystemDiscoveryConnectorConnectorTypeEnum{
	"MACS": ExternalDbSystemDiscoveryConnectorConnectorTypeMacs,
}

var mappingExternalDbSystemDiscoveryConnectorConnectorTypeEnumLowerCase = map[string]ExternalDbSystemDiscoveryConnectorConnectorTypeEnum{
	"macs": ExternalDbSystemDiscoveryConnectorConnectorTypeMacs,
}

// GetExternalDbSystemDiscoveryConnectorConnectorTypeEnumValues Enumerates the set of values for ExternalDbSystemDiscoveryConnectorConnectorTypeEnum
func GetExternalDbSystemDiscoveryConnectorConnectorTypeEnumValues() []ExternalDbSystemDiscoveryConnectorConnectorTypeEnum {
	values := make([]ExternalDbSystemDiscoveryConnectorConnectorTypeEnum, 0)
	for _, v := range mappingExternalDbSystemDiscoveryConnectorConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemDiscoveryConnectorConnectorTypeEnumStringValues Enumerates the set of values in String for ExternalDbSystemDiscoveryConnectorConnectorTypeEnum
func GetExternalDbSystemDiscoveryConnectorConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingExternalDbSystemDiscoveryConnectorConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemDiscoveryConnectorConnectorTypeEnum(val string) (ExternalDbSystemDiscoveryConnectorConnectorTypeEnum, bool) {
	enum, ok := mappingExternalDbSystemDiscoveryConnectorConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
