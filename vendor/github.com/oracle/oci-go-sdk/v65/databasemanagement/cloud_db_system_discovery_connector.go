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

// CloudDbSystemDiscoveryConnector The connector details used to connect to the cloud DB system component.
type CloudDbSystemDiscoveryConnector interface {

	// The user-friendly name for the cloud connector. The name does not have to be unique.
	GetDisplayName() *string

	// The status of connectivity to the cloud DB system component.
	GetConnectionStatus() *string

	// The error message indicating the reason for connection failure or `null` if
	// the connection was successful.
	GetConnectionFailureMessage() *string

	// The date and time the connectionStatus of the cloud DB system connector was last updated.
	GetTimeConnectionStatusLastUpdated() *common.SDKTime
}

type clouddbsystemdiscoveryconnector struct {
	JsonData                        []byte
	ConnectionStatus                *string         `mandatory:"false" json:"connectionStatus"`
	ConnectionFailureMessage        *string         `mandatory:"false" json:"connectionFailureMessage"`
	TimeConnectionStatusLastUpdated *common.SDKTime `mandatory:"false" json:"timeConnectionStatusLastUpdated"`
	DisplayName                     *string         `mandatory:"true" json:"displayName"`
	ConnectorType                   string          `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *clouddbsystemdiscoveryconnector) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerclouddbsystemdiscoveryconnector clouddbsystemdiscoveryconnector
	s := struct {
		Model Unmarshalerclouddbsystemdiscoveryconnector
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
func (m *clouddbsystemdiscoveryconnector) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := CloudDbSystemDiscoveryMacsConnector{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CloudDbSystemDiscoveryConnector: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetConnectionStatus returns ConnectionStatus
func (m clouddbsystemdiscoveryconnector) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

// GetConnectionFailureMessage returns ConnectionFailureMessage
func (m clouddbsystemdiscoveryconnector) GetConnectionFailureMessage() *string {
	return m.ConnectionFailureMessage
}

// GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m clouddbsystemdiscoveryconnector) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

// GetDisplayName returns DisplayName
func (m clouddbsystemdiscoveryconnector) GetDisplayName() *string {
	return m.DisplayName
}

func (m clouddbsystemdiscoveryconnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m clouddbsystemdiscoveryconnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudDbSystemDiscoveryConnectorConnectorTypeEnum Enum with underlying type: string
type CloudDbSystemDiscoveryConnectorConnectorTypeEnum string

// Set of constants representing the allowable values for CloudDbSystemDiscoveryConnectorConnectorTypeEnum
const (
	CloudDbSystemDiscoveryConnectorConnectorTypeMacs CloudDbSystemDiscoveryConnectorConnectorTypeEnum = "MACS"
)

var mappingCloudDbSystemDiscoveryConnectorConnectorTypeEnum = map[string]CloudDbSystemDiscoveryConnectorConnectorTypeEnum{
	"MACS": CloudDbSystemDiscoveryConnectorConnectorTypeMacs,
}

var mappingCloudDbSystemDiscoveryConnectorConnectorTypeEnumLowerCase = map[string]CloudDbSystemDiscoveryConnectorConnectorTypeEnum{
	"macs": CloudDbSystemDiscoveryConnectorConnectorTypeMacs,
}

// GetCloudDbSystemDiscoveryConnectorConnectorTypeEnumValues Enumerates the set of values for CloudDbSystemDiscoveryConnectorConnectorTypeEnum
func GetCloudDbSystemDiscoveryConnectorConnectorTypeEnumValues() []CloudDbSystemDiscoveryConnectorConnectorTypeEnum {
	values := make([]CloudDbSystemDiscoveryConnectorConnectorTypeEnum, 0)
	for _, v := range mappingCloudDbSystemDiscoveryConnectorConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbSystemDiscoveryConnectorConnectorTypeEnumStringValues Enumerates the set of values in String for CloudDbSystemDiscoveryConnectorConnectorTypeEnum
func GetCloudDbSystemDiscoveryConnectorConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingCloudDbSystemDiscoveryConnectorConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbSystemDiscoveryConnectorConnectorTypeEnum(val string) (CloudDbSystemDiscoveryConnectorConnectorTypeEnum, bool) {
	enum, ok := mappingCloudDbSystemDiscoveryConnectorConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
