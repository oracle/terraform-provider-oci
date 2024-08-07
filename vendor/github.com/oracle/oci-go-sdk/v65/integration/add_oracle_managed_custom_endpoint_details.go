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

// AddOracleManagedCustomEndpointDetails Details for enabling Oracle Managed custom endpoint
type AddOracleManagedCustomEndpointDetails struct {

	// Oracle managed custom hostname
	Hostname *string `mandatory:"true" json:"hostname"`

	// Type of DNS.
	DnsType AddOracleManagedCustomEndpointDetailsDnsTypeEnum `mandatory:"false" json:"dnsType,omitempty"`

	// DNS Zone name
	DnsZoneName *string `mandatory:"false" json:"dnsZoneName"`
}

func (m AddOracleManagedCustomEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddOracleManagedCustomEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAddOracleManagedCustomEndpointDetailsDnsTypeEnum(string(m.DnsType)); !ok && m.DnsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsType: %s. Supported values are: %s.", m.DnsType, strings.Join(GetAddOracleManagedCustomEndpointDetailsDnsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddOracleManagedCustomEndpointDetailsDnsTypeEnum Enum with underlying type: string
type AddOracleManagedCustomEndpointDetailsDnsTypeEnum string

// Set of constants representing the allowable values for AddOracleManagedCustomEndpointDetailsDnsTypeEnum
const (
	AddOracleManagedCustomEndpointDetailsDnsTypeOci AddOracleManagedCustomEndpointDetailsDnsTypeEnum = "OCI"
)

var mappingAddOracleManagedCustomEndpointDetailsDnsTypeEnum = map[string]AddOracleManagedCustomEndpointDetailsDnsTypeEnum{
	"OCI": AddOracleManagedCustomEndpointDetailsDnsTypeOci,
}

var mappingAddOracleManagedCustomEndpointDetailsDnsTypeEnumLowerCase = map[string]AddOracleManagedCustomEndpointDetailsDnsTypeEnum{
	"oci": AddOracleManagedCustomEndpointDetailsDnsTypeOci,
}

// GetAddOracleManagedCustomEndpointDetailsDnsTypeEnumValues Enumerates the set of values for AddOracleManagedCustomEndpointDetailsDnsTypeEnum
func GetAddOracleManagedCustomEndpointDetailsDnsTypeEnumValues() []AddOracleManagedCustomEndpointDetailsDnsTypeEnum {
	values := make([]AddOracleManagedCustomEndpointDetailsDnsTypeEnum, 0)
	for _, v := range mappingAddOracleManagedCustomEndpointDetailsDnsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddOracleManagedCustomEndpointDetailsDnsTypeEnumStringValues Enumerates the set of values in String for AddOracleManagedCustomEndpointDetailsDnsTypeEnum
func GetAddOracleManagedCustomEndpointDetailsDnsTypeEnumStringValues() []string {
	return []string{
		"OCI",
	}
}

// GetMappingAddOracleManagedCustomEndpointDetailsDnsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddOracleManagedCustomEndpointDetailsDnsTypeEnum(val string) (AddOracleManagedCustomEndpointDetailsDnsTypeEnum, bool) {
	enum, ok := mappingAddOracleManagedCustomEndpointDetailsDnsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
