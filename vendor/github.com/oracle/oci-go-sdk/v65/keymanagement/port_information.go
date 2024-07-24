// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PortInformation Details of a single portInformation item include the PortNumber (an integer used as an identifier) and the PortType (this refers to either an enum value of Management Utility, Client Utility, or null)
type PortInformation struct {

	// The port number is a unique identifier which is typically used as the loadbalancer listener.
	PortNumber *int `mandatory:"true" json:"portNumber"`

	// Port type associated for the port number. The two port type enums are `CLIENTUTILITY` and `MANAGEMENTUTILITY`. The CLIENTUTILITY enum corresponds to a port which is used by the client daemon. The MANAGEMENTUTILITY enum corresponds to a port used by user management utility.
	PortType PortInformationPortTypeEnum `mandatory:"true" json:"portType"`
}

func (m PortInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PortInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPortInformationPortTypeEnum(string(m.PortType)); !ok && m.PortType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PortType: %s. Supported values are: %s.", m.PortType, strings.Join(GetPortInformationPortTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PortInformationPortTypeEnum Enum with underlying type: string
type PortInformationPortTypeEnum string

// Set of constants representing the allowable values for PortInformationPortTypeEnum
const (
	PortInformationPortTypeClientutility     PortInformationPortTypeEnum = "CLIENTUTILITY"
	PortInformationPortTypeManagementutility PortInformationPortTypeEnum = "MANAGEMENTUTILITY"
)

var mappingPortInformationPortTypeEnum = map[string]PortInformationPortTypeEnum{
	"CLIENTUTILITY":     PortInformationPortTypeClientutility,
	"MANAGEMENTUTILITY": PortInformationPortTypeManagementutility,
}

var mappingPortInformationPortTypeEnumLowerCase = map[string]PortInformationPortTypeEnum{
	"clientutility":     PortInformationPortTypeClientutility,
	"managementutility": PortInformationPortTypeManagementutility,
}

// GetPortInformationPortTypeEnumValues Enumerates the set of values for PortInformationPortTypeEnum
func GetPortInformationPortTypeEnumValues() []PortInformationPortTypeEnum {
	values := make([]PortInformationPortTypeEnum, 0)
	for _, v := range mappingPortInformationPortTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPortInformationPortTypeEnumStringValues Enumerates the set of values in String for PortInformationPortTypeEnum
func GetPortInformationPortTypeEnumStringValues() []string {
	return []string{
		"CLIENTUTILITY",
		"MANAGEMENTUTILITY",
	}
}

// GetMappingPortInformationPortTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPortInformationPortTypeEnum(val string) (PortInformationPortTypeEnum, bool) {
	enum, ok := mappingPortInformationPortTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
