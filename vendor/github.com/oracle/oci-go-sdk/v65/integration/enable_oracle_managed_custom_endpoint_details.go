// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnableOracleManagedCustomEndpointDetails Details for enabling Oracle Managed custom endpoint
type EnableOracleManagedCustomEndpointDetails struct {

	// Oracle managed custom hostname
	Hostname *string `mandatory:"true" json:"hostname"`

	// Type of DNS.
	DnsType EnableOracleManagedCustomEndpointDetailsDnsTypeEnum `mandatory:"false" json:"dnsType,omitempty"`

	// DNS Zone name
	DnsZoneName *string `mandatory:"false" json:"dnsZoneName"`
}

func (m EnableOracleManagedCustomEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableOracleManagedCustomEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEnableOracleManagedCustomEndpointDetailsDnsTypeEnum(string(m.DnsType)); !ok && m.DnsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsType: %s. Supported values are: %s.", m.DnsType, strings.Join(GetEnableOracleManagedCustomEndpointDetailsDnsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnableOracleManagedCustomEndpointDetailsDnsTypeEnum Enum with underlying type: string
type EnableOracleManagedCustomEndpointDetailsDnsTypeEnum string

// Set of constants representing the allowable values for EnableOracleManagedCustomEndpointDetailsDnsTypeEnum
const (
	EnableOracleManagedCustomEndpointDetailsDnsTypeOci EnableOracleManagedCustomEndpointDetailsDnsTypeEnum = "OCI"
)

var mappingEnableOracleManagedCustomEndpointDetailsDnsTypeEnum = map[string]EnableOracleManagedCustomEndpointDetailsDnsTypeEnum{
	"OCI": EnableOracleManagedCustomEndpointDetailsDnsTypeOci,
}

var mappingEnableOracleManagedCustomEndpointDetailsDnsTypeEnumLowerCase = map[string]EnableOracleManagedCustomEndpointDetailsDnsTypeEnum{
	"oci": EnableOracleManagedCustomEndpointDetailsDnsTypeOci,
}

// GetEnableOracleManagedCustomEndpointDetailsDnsTypeEnumValues Enumerates the set of values for EnableOracleManagedCustomEndpointDetailsDnsTypeEnum
func GetEnableOracleManagedCustomEndpointDetailsDnsTypeEnumValues() []EnableOracleManagedCustomEndpointDetailsDnsTypeEnum {
	values := make([]EnableOracleManagedCustomEndpointDetailsDnsTypeEnum, 0)
	for _, v := range mappingEnableOracleManagedCustomEndpointDetailsDnsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableOracleManagedCustomEndpointDetailsDnsTypeEnumStringValues Enumerates the set of values in String for EnableOracleManagedCustomEndpointDetailsDnsTypeEnum
func GetEnableOracleManagedCustomEndpointDetailsDnsTypeEnumStringValues() []string {
	return []string{
		"OCI",
	}
}

// GetMappingEnableOracleManagedCustomEndpointDetailsDnsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableOracleManagedCustomEndpointDetailsDnsTypeEnum(val string) (EnableOracleManagedCustomEndpointDetailsDnsTypeEnum, bool) {
	enum, ok := mappingEnableOracleManagedCustomEndpointDetailsDnsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
