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

// OlvmIp Represents the IP configuration of a network interface.
type OlvmIp struct {

	// The text representation of the IP address.
	Address *string `mandatory:"false" json:"address"`

	// The address of the default gateway.
	Gateway *string `mandatory:"false" json:"gateway"`

	// The network mask.
	Netmask *string `mandatory:"false" json:"netmask"`

	// The version of the IP protocol.
	IpVersion OlvmIpIpVersionEnum `mandatory:"false" json:"ipVersion,omitempty"`
}

func (m OlvmIp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmIp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmIpIpVersionEnum(string(m.IpVersion)); !ok && m.IpVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpVersion: %s. Supported values are: %s.", m.IpVersion, strings.Join(GetOlvmIpIpVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmIpIpVersionEnum Enum with underlying type: string
type OlvmIpIpVersionEnum string

// Set of constants representing the allowable values for OlvmIpIpVersionEnum
const (
	OlvmIpIpVersionV4 OlvmIpIpVersionEnum = "V4"
	OlvmIpIpVersionV6 OlvmIpIpVersionEnum = "V6"
)

var mappingOlvmIpIpVersionEnum = map[string]OlvmIpIpVersionEnum{
	"V4": OlvmIpIpVersionV4,
	"V6": OlvmIpIpVersionV6,
}

var mappingOlvmIpIpVersionEnumLowerCase = map[string]OlvmIpIpVersionEnum{
	"v4": OlvmIpIpVersionV4,
	"v6": OlvmIpIpVersionV6,
}

// GetOlvmIpIpVersionEnumValues Enumerates the set of values for OlvmIpIpVersionEnum
func GetOlvmIpIpVersionEnumValues() []OlvmIpIpVersionEnum {
	values := make([]OlvmIpIpVersionEnum, 0)
	for _, v := range mappingOlvmIpIpVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmIpIpVersionEnumStringValues Enumerates the set of values in String for OlvmIpIpVersionEnum
func GetOlvmIpIpVersionEnumStringValues() []string {
	return []string{
		"V4",
		"V6",
	}
}

// GetMappingOlvmIpIpVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmIpIpVersionEnum(val string) (OlvmIpIpVersionEnum, bool) {
	enum, ok := mappingOlvmIpIpVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
