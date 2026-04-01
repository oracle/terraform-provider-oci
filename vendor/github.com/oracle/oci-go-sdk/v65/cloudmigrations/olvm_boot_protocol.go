// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmBootProtocol Defines the options of the IP address assignment method to a NIC.
type OlvmBootProtocol struct {

	// IP address assignment methods to a NIC.
	Protocol OlvmBootProtocolProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m OlvmBootProtocol) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmBootProtocol) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmBootProtocolProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetOlvmBootProtocolProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmBootProtocolProtocolEnum Enum with underlying type: string
type OlvmBootProtocolProtocolEnum string

// Set of constants representing the allowable values for OlvmBootProtocolProtocolEnum
const (
	OlvmBootProtocolProtocolAutoconf         OlvmBootProtocolProtocolEnum = "AUTOCONF"
	OlvmBootProtocolProtocolDhcp             OlvmBootProtocolProtocolEnum = "DHCP"
	OlvmBootProtocolProtocolNone             OlvmBootProtocolProtocolEnum = "NONE"
	OlvmBootProtocolProtocolPolyDhcpAutoconf OlvmBootProtocolProtocolEnum = "POLY_DHCP_AUTOCONF"
	OlvmBootProtocolProtocolStatic           OlvmBootProtocolProtocolEnum = "STATIC"
)

var mappingOlvmBootProtocolProtocolEnum = map[string]OlvmBootProtocolProtocolEnum{
	"AUTOCONF":           OlvmBootProtocolProtocolAutoconf,
	"DHCP":               OlvmBootProtocolProtocolDhcp,
	"NONE":               OlvmBootProtocolProtocolNone,
	"POLY_DHCP_AUTOCONF": OlvmBootProtocolProtocolPolyDhcpAutoconf,
	"STATIC":             OlvmBootProtocolProtocolStatic,
}

var mappingOlvmBootProtocolProtocolEnumLowerCase = map[string]OlvmBootProtocolProtocolEnum{
	"autoconf":           OlvmBootProtocolProtocolAutoconf,
	"dhcp":               OlvmBootProtocolProtocolDhcp,
	"none":               OlvmBootProtocolProtocolNone,
	"poly_dhcp_autoconf": OlvmBootProtocolProtocolPolyDhcpAutoconf,
	"static":             OlvmBootProtocolProtocolStatic,
}

// GetOlvmBootProtocolProtocolEnumValues Enumerates the set of values for OlvmBootProtocolProtocolEnum
func GetOlvmBootProtocolProtocolEnumValues() []OlvmBootProtocolProtocolEnum {
	values := make([]OlvmBootProtocolProtocolEnum, 0)
	for _, v := range mappingOlvmBootProtocolProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmBootProtocolProtocolEnumStringValues Enumerates the set of values in String for OlvmBootProtocolProtocolEnum
func GetOlvmBootProtocolProtocolEnumStringValues() []string {
	return []string{
		"AUTOCONF",
		"DHCP",
		"NONE",
		"POLY_DHCP_AUTOCONF",
		"STATIC",
	}
}

// GetMappingOlvmBootProtocolProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmBootProtocolProtocolEnum(val string) (OlvmBootProtocolProtocolEnum, bool) {
	enum, ok := mappingOlvmBootProtocolProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
