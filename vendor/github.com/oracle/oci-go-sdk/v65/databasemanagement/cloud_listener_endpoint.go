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

// CloudListenerEndpoint The protocol address that a cloud listener is configured to listen on.
type CloudListenerEndpoint interface {

	// The list of services registered with the listener.
	GetServices() []string
}

type cloudlistenerendpoint struct {
	JsonData []byte
	Services []string `mandatory:"false" json:"services"`
	Protocol string   `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *cloudlistenerendpoint) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercloudlistenerendpoint cloudlistenerendpoint
	s := struct {
		Model Unmarshalercloudlistenerendpoint
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
func (m *cloudlistenerendpoint) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "TCPS":
		mm := CloudListenerTcpsEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IPC":
		mm := CloudListenerIpcEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TCP":
		mm := CloudListenerTcpEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CloudListenerEndpoint: %s.", m.Protocol)
		return *m, nil
	}
}

// GetServices returns Services
func (m cloudlistenerendpoint) GetServices() []string {
	return m.Services
}

func (m cloudlistenerendpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m cloudlistenerendpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudListenerEndpointProtocolEnum Enum with underlying type: string
type CloudListenerEndpointProtocolEnum string

// Set of constants representing the allowable values for CloudListenerEndpointProtocolEnum
const (
	CloudListenerEndpointProtocolIpc  CloudListenerEndpointProtocolEnum = "IPC"
	CloudListenerEndpointProtocolTcp  CloudListenerEndpointProtocolEnum = "TCP"
	CloudListenerEndpointProtocolTcps CloudListenerEndpointProtocolEnum = "TCPS"
)

var mappingCloudListenerEndpointProtocolEnum = map[string]CloudListenerEndpointProtocolEnum{
	"IPC":  CloudListenerEndpointProtocolIpc,
	"TCP":  CloudListenerEndpointProtocolTcp,
	"TCPS": CloudListenerEndpointProtocolTcps,
}

var mappingCloudListenerEndpointProtocolEnumLowerCase = map[string]CloudListenerEndpointProtocolEnum{
	"ipc":  CloudListenerEndpointProtocolIpc,
	"tcp":  CloudListenerEndpointProtocolTcp,
	"tcps": CloudListenerEndpointProtocolTcps,
}

// GetCloudListenerEndpointProtocolEnumValues Enumerates the set of values for CloudListenerEndpointProtocolEnum
func GetCloudListenerEndpointProtocolEnumValues() []CloudListenerEndpointProtocolEnum {
	values := make([]CloudListenerEndpointProtocolEnum, 0)
	for _, v := range mappingCloudListenerEndpointProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudListenerEndpointProtocolEnumStringValues Enumerates the set of values in String for CloudListenerEndpointProtocolEnum
func GetCloudListenerEndpointProtocolEnumStringValues() []string {
	return []string{
		"IPC",
		"TCP",
		"TCPS",
	}
}

// GetMappingCloudListenerEndpointProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudListenerEndpointProtocolEnum(val string) (CloudListenerEndpointProtocolEnum, bool) {
	enum, ok := mappingCloudListenerEndpointProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
