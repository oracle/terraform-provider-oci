// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AsmConnectionString The ASM instance connection string.
type AsmConnectionString struct {

	// The list of host names of the ASM instances.
	Hosts []string `mandatory:"true" json:"hosts"`

	// The port used to connect to the ASM instance.
	Port *int `mandatory:"true" json:"port"`

	// The service name of the ASM instance.
	Service *string `mandatory:"true" json:"service"`

	// The protocol used to connect to the ASM instance.
	Protocol AsmConnectionStringProtocolEnum `mandatory:"true" json:"protocol"`
}

func (m AsmConnectionString) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AsmConnectionString) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAsmConnectionStringProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetAsmConnectionStringProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AsmConnectionStringProtocolEnum Enum with underlying type: string
type AsmConnectionStringProtocolEnum string

// Set of constants representing the allowable values for AsmConnectionStringProtocolEnum
const (
	AsmConnectionStringProtocolTcp AsmConnectionStringProtocolEnum = "TCP"
)

var mappingAsmConnectionStringProtocolEnum = map[string]AsmConnectionStringProtocolEnum{
	"TCP": AsmConnectionStringProtocolTcp,
}

var mappingAsmConnectionStringProtocolEnumLowerCase = map[string]AsmConnectionStringProtocolEnum{
	"tcp": AsmConnectionStringProtocolTcp,
}

// GetAsmConnectionStringProtocolEnumValues Enumerates the set of values for AsmConnectionStringProtocolEnum
func GetAsmConnectionStringProtocolEnumValues() []AsmConnectionStringProtocolEnum {
	values := make([]AsmConnectionStringProtocolEnum, 0)
	for _, v := range mappingAsmConnectionStringProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetAsmConnectionStringProtocolEnumStringValues Enumerates the set of values in String for AsmConnectionStringProtocolEnum
func GetAsmConnectionStringProtocolEnumStringValues() []string {
	return []string{
		"TCP",
	}
}

// GetMappingAsmConnectionStringProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAsmConnectionStringProtocolEnum(val string) (AsmConnectionStringProtocolEnum, bool) {
	enum, ok := mappingAsmConnectionStringProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
