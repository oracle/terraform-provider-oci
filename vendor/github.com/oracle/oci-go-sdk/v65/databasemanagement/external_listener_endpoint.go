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

// ExternalListenerEndpoint The protocol address that an external listener is configured to listen on.
type ExternalListenerEndpoint interface {

	// The list of services registered with the listener.
	GetServices() []string
}

type externallistenerendpoint struct {
	JsonData []byte
	Services []string `mandatory:"false" json:"services"`
	Protocol string   `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *externallistenerendpoint) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexternallistenerendpoint externallistenerendpoint
	s := struct {
		Model Unmarshalerexternallistenerendpoint
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Services = s.Model.Services
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *externallistenerendpoint) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "TCP":
		mm := ExternalListenerTcpEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TCPS":
		mm := ExternalListenerTcpsEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IPC":
		mm := ExternalListenerIpcEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ExternalListenerEndpoint: %s.", m.Protocol)
		return *m, nil
	}
}

// GetServices returns Services
func (m externallistenerendpoint) GetServices() []string {
	return m.Services
}

func (m externallistenerendpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m externallistenerendpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalListenerEndpointProtocolEnum Enum with underlying type: string
type ExternalListenerEndpointProtocolEnum string

// Set of constants representing the allowable values for ExternalListenerEndpointProtocolEnum
const (
	ExternalListenerEndpointProtocolIpc  ExternalListenerEndpointProtocolEnum = "IPC"
	ExternalListenerEndpointProtocolTcp  ExternalListenerEndpointProtocolEnum = "TCP"
	ExternalListenerEndpointProtocolTcps ExternalListenerEndpointProtocolEnum = "TCPS"
)

var mappingExternalListenerEndpointProtocolEnum = map[string]ExternalListenerEndpointProtocolEnum{
	"IPC":  ExternalListenerEndpointProtocolIpc,
	"TCP":  ExternalListenerEndpointProtocolTcp,
	"TCPS": ExternalListenerEndpointProtocolTcps,
}

var mappingExternalListenerEndpointProtocolEnumLowerCase = map[string]ExternalListenerEndpointProtocolEnum{
	"ipc":  ExternalListenerEndpointProtocolIpc,
	"tcp":  ExternalListenerEndpointProtocolTcp,
	"tcps": ExternalListenerEndpointProtocolTcps,
}

// GetExternalListenerEndpointProtocolEnumValues Enumerates the set of values for ExternalListenerEndpointProtocolEnum
func GetExternalListenerEndpointProtocolEnumValues() []ExternalListenerEndpointProtocolEnum {
	values := make([]ExternalListenerEndpointProtocolEnum, 0)
	for _, v := range mappingExternalListenerEndpointProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalListenerEndpointProtocolEnumStringValues Enumerates the set of values in String for ExternalListenerEndpointProtocolEnum
func GetExternalListenerEndpointProtocolEnumStringValues() []string {
	return []string{
		"IPC",
		"TCP",
		"TCPS",
	}
}

// GetMappingExternalListenerEndpointProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalListenerEndpointProtocolEnum(val string) (ExternalListenerEndpointProtocolEnum, bool) {
	enum, ok := mappingExternalListenerEndpointProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
