// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ConnectionDetails Connection details to connect to the database. HostName, protocol, and port should be specified.
type ConnectionDetails struct {

	// Name of the listener host that will be used to create the connect string to the database.
	HostName *string `mandatory:"true" json:"hostName"`

	// Protocol used for connection requests.
	Protocol ConnectionDetailsProtocolEnum `mandatory:"true" json:"protocol"`

	// Listener port number used for connection requests.
	Port *int `mandatory:"true" json:"port"`

	// Service name used for connection requests.
	ServiceName *string `mandatory:"true" json:"serviceName"`
}

func (m ConnectionDetails) String() string {
	return common.PointerString(m)
}

// ConnectionDetailsProtocolEnum Enum with underlying type: string
type ConnectionDetailsProtocolEnum string

// Set of constants representing the allowable values for ConnectionDetailsProtocolEnum
const (
	ConnectionDetailsProtocolTcp  ConnectionDetailsProtocolEnum = "TCP"
	ConnectionDetailsProtocolTcps ConnectionDetailsProtocolEnum = "TCPS"
)

var mappingConnectionDetailsProtocol = map[string]ConnectionDetailsProtocolEnum{
	"TCP":  ConnectionDetailsProtocolTcp,
	"TCPS": ConnectionDetailsProtocolTcps,
}

// GetConnectionDetailsProtocolEnumValues Enumerates the set of values for ConnectionDetailsProtocolEnum
func GetConnectionDetailsProtocolEnumValues() []ConnectionDetailsProtocolEnum {
	values := make([]ConnectionDetailsProtocolEnum, 0)
	for _, v := range mappingConnectionDetailsProtocol {
		values = append(values, v)
	}
	return values
}
