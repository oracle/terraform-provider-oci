// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmNetworkDetails Details of the client or backup networks in an Exadata VM cluster network. Applies to Exadata Cloud@Customer instances only.
type VmNetworkDetails struct {

	// The network type.
	NetworkType VmNetworkDetailsNetworkTypeEnum `mandatory:"true" json:"networkType"`

	// The list of node details.
	Nodes []NodeDetails `mandatory:"true" json:"nodes"`

	// The network VLAN ID.
	VlanId *string `mandatory:"false" json:"vlanId"`

	// The network netmask.
	Netmask *string `mandatory:"false" json:"netmask"`

	// The network gateway.
	Gateway *string `mandatory:"false" json:"gateway"`

	// The network domain name.
	DomainName *string `mandatory:"false" json:"domainName"`
}

func (m VmNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmNetworkDetailsNetworkTypeEnum(string(m.NetworkType)); !ok && m.NetworkType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkType: %s. Supported values are: %s.", m.NetworkType, strings.Join(GetVmNetworkDetailsNetworkTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmNetworkDetailsNetworkTypeEnum Enum with underlying type: string
type VmNetworkDetailsNetworkTypeEnum string

// Set of constants representing the allowable values for VmNetworkDetailsNetworkTypeEnum
const (
	VmNetworkDetailsNetworkTypeClient VmNetworkDetailsNetworkTypeEnum = "CLIENT"
	VmNetworkDetailsNetworkTypeBackup VmNetworkDetailsNetworkTypeEnum = "BACKUP"
)

var mappingVmNetworkDetailsNetworkTypeEnum = map[string]VmNetworkDetailsNetworkTypeEnum{
	"CLIENT": VmNetworkDetailsNetworkTypeClient,
	"BACKUP": VmNetworkDetailsNetworkTypeBackup,
}

var mappingVmNetworkDetailsNetworkTypeEnumLowerCase = map[string]VmNetworkDetailsNetworkTypeEnum{
	"client": VmNetworkDetailsNetworkTypeClient,
	"backup": VmNetworkDetailsNetworkTypeBackup,
}

// GetVmNetworkDetailsNetworkTypeEnumValues Enumerates the set of values for VmNetworkDetailsNetworkTypeEnum
func GetVmNetworkDetailsNetworkTypeEnumValues() []VmNetworkDetailsNetworkTypeEnum {
	values := make([]VmNetworkDetailsNetworkTypeEnum, 0)
	for _, v := range mappingVmNetworkDetailsNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmNetworkDetailsNetworkTypeEnumStringValues Enumerates the set of values in String for VmNetworkDetailsNetworkTypeEnum
func GetVmNetworkDetailsNetworkTypeEnumStringValues() []string {
	return []string{
		"CLIENT",
		"BACKUP",
	}
}

// GetMappingVmNetworkDetailsNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmNetworkDetailsNetworkTypeEnum(val string) (VmNetworkDetailsNetworkTypeEnum, bool) {
	enum, ok := mappingVmNetworkDetailsNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
