// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PeComanagedDatabaseConnectionDetails Connection details of the private endpoints.
type PeComanagedDatabaseConnectionDetails struct {

	// List of hosts and port for private endpoint accessed database resource.
	Hosts []PeComanagedDatabaseHostDetails `mandatory:"true" json:"hosts"`

	// Protocol used for connection requests for private endpoint accssed database resource.
	Protocol PeComanagedDatabaseConnectionDetailsProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	// Database service name used for connection requests.
	ServiceName *string `mandatory:"false" json:"serviceName"`
}

func (m PeComanagedDatabaseConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeComanagedDatabaseConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPeComanagedDatabaseConnectionDetailsProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetPeComanagedDatabaseConnectionDetailsProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PeComanagedDatabaseConnectionDetailsProtocolEnum Enum with underlying type: string
type PeComanagedDatabaseConnectionDetailsProtocolEnum string

// Set of constants representing the allowable values for PeComanagedDatabaseConnectionDetailsProtocolEnum
const (
	PeComanagedDatabaseConnectionDetailsProtocolTcp  PeComanagedDatabaseConnectionDetailsProtocolEnum = "TCP"
	PeComanagedDatabaseConnectionDetailsProtocolTcps PeComanagedDatabaseConnectionDetailsProtocolEnum = "TCPS"
)

var mappingPeComanagedDatabaseConnectionDetailsProtocolEnum = map[string]PeComanagedDatabaseConnectionDetailsProtocolEnum{
	"TCP":  PeComanagedDatabaseConnectionDetailsProtocolTcp,
	"TCPS": PeComanagedDatabaseConnectionDetailsProtocolTcps,
}

var mappingPeComanagedDatabaseConnectionDetailsProtocolEnumLowerCase = map[string]PeComanagedDatabaseConnectionDetailsProtocolEnum{
	"tcp":  PeComanagedDatabaseConnectionDetailsProtocolTcp,
	"tcps": PeComanagedDatabaseConnectionDetailsProtocolTcps,
}

// GetPeComanagedDatabaseConnectionDetailsProtocolEnumValues Enumerates the set of values for PeComanagedDatabaseConnectionDetailsProtocolEnum
func GetPeComanagedDatabaseConnectionDetailsProtocolEnumValues() []PeComanagedDatabaseConnectionDetailsProtocolEnum {
	values := make([]PeComanagedDatabaseConnectionDetailsProtocolEnum, 0)
	for _, v := range mappingPeComanagedDatabaseConnectionDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetPeComanagedDatabaseConnectionDetailsProtocolEnumStringValues Enumerates the set of values in String for PeComanagedDatabaseConnectionDetailsProtocolEnum
func GetPeComanagedDatabaseConnectionDetailsProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingPeComanagedDatabaseConnectionDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPeComanagedDatabaseConnectionDetailsProtocolEnum(val string) (PeComanagedDatabaseConnectionDetailsProtocolEnum, bool) {
	enum, ok := mappingPeComanagedDatabaseConnectionDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
