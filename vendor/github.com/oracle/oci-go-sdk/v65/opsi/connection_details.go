// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectionDetails Connection details to connect to the database. HostName, protocol, and port should be specified.
type ConnectionDetails struct {

	// Name of the listener host that will be used to create the connect string to the database.
	HostName *string `mandatory:"true" json:"hostName"`

	// Protocol used for connection requests.
	Protocol ConnectionDetailsProtocolEnum `mandatory:"true" json:"protocol"`

	// Listener port number used for connection requests.
	Port *int `mandatory:"true" json:"port"`

	// Database service name used for connection requests.
	ServiceName *string `mandatory:"true" json:"serviceName"`
}

func (m ConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionDetailsProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetConnectionDetailsProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionDetailsProtocolEnum Enum with underlying type: string
type ConnectionDetailsProtocolEnum string

// Set of constants representing the allowable values for ConnectionDetailsProtocolEnum
const (
	ConnectionDetailsProtocolTcp  ConnectionDetailsProtocolEnum = "TCP"
	ConnectionDetailsProtocolTcps ConnectionDetailsProtocolEnum = "TCPS"
)

var mappingConnectionDetailsProtocolEnum = map[string]ConnectionDetailsProtocolEnum{
	"TCP":  ConnectionDetailsProtocolTcp,
	"TCPS": ConnectionDetailsProtocolTcps,
}

var mappingConnectionDetailsProtocolEnumLowerCase = map[string]ConnectionDetailsProtocolEnum{
	"tcp":  ConnectionDetailsProtocolTcp,
	"tcps": ConnectionDetailsProtocolTcps,
}

// GetConnectionDetailsProtocolEnumValues Enumerates the set of values for ConnectionDetailsProtocolEnum
func GetConnectionDetailsProtocolEnumValues() []ConnectionDetailsProtocolEnum {
	values := make([]ConnectionDetailsProtocolEnum, 0)
	for _, v := range mappingConnectionDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionDetailsProtocolEnumStringValues Enumerates the set of values in String for ConnectionDetailsProtocolEnum
func GetConnectionDetailsProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingConnectionDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionDetailsProtocolEnum(val string) (ConnectionDetailsProtocolEnum, bool) {
	enum, ok := mappingConnectionDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
