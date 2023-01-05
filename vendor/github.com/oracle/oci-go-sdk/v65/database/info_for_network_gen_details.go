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

// InfoForNetworkGenDetails Parameters for generation of the client or backup network in a VM cluster network in an Exadata Cloud@Customer system.
type InfoForNetworkGenDetails struct {

	// The network type.
	NetworkType InfoForNetworkGenDetailsNetworkTypeEnum `mandatory:"true" json:"networkType"`

	// The network VLAN ID.
	VlanId *string `mandatory:"true" json:"vlanId"`

	// The cidr for the network.
	Cidr *string `mandatory:"true" json:"cidr"`

	// The network gateway.
	Gateway *string `mandatory:"true" json:"gateway"`

	// The network netmask.
	Netmask *string `mandatory:"true" json:"netmask"`

	// The network domain name.
	Domain *string `mandatory:"true" json:"domain"`

	// The network domain name.
	Prefix *string `mandatory:"true" json:"prefix"`
}

func (m InfoForNetworkGenDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InfoForNetworkGenDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInfoForNetworkGenDetailsNetworkTypeEnum(string(m.NetworkType)); !ok && m.NetworkType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkType: %s. Supported values are: %s.", m.NetworkType, strings.Join(GetInfoForNetworkGenDetailsNetworkTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InfoForNetworkGenDetailsNetworkTypeEnum Enum with underlying type: string
type InfoForNetworkGenDetailsNetworkTypeEnum string

// Set of constants representing the allowable values for InfoForNetworkGenDetailsNetworkTypeEnum
const (
	InfoForNetworkGenDetailsNetworkTypeClient InfoForNetworkGenDetailsNetworkTypeEnum = "CLIENT"
	InfoForNetworkGenDetailsNetworkTypeBackup InfoForNetworkGenDetailsNetworkTypeEnum = "BACKUP"
)

var mappingInfoForNetworkGenDetailsNetworkTypeEnum = map[string]InfoForNetworkGenDetailsNetworkTypeEnum{
	"CLIENT": InfoForNetworkGenDetailsNetworkTypeClient,
	"BACKUP": InfoForNetworkGenDetailsNetworkTypeBackup,
}

var mappingInfoForNetworkGenDetailsNetworkTypeEnumLowerCase = map[string]InfoForNetworkGenDetailsNetworkTypeEnum{
	"client": InfoForNetworkGenDetailsNetworkTypeClient,
	"backup": InfoForNetworkGenDetailsNetworkTypeBackup,
}

// GetInfoForNetworkGenDetailsNetworkTypeEnumValues Enumerates the set of values for InfoForNetworkGenDetailsNetworkTypeEnum
func GetInfoForNetworkGenDetailsNetworkTypeEnumValues() []InfoForNetworkGenDetailsNetworkTypeEnum {
	values := make([]InfoForNetworkGenDetailsNetworkTypeEnum, 0)
	for _, v := range mappingInfoForNetworkGenDetailsNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInfoForNetworkGenDetailsNetworkTypeEnumStringValues Enumerates the set of values in String for InfoForNetworkGenDetailsNetworkTypeEnum
func GetInfoForNetworkGenDetailsNetworkTypeEnumStringValues() []string {
	return []string{
		"CLIENT",
		"BACKUP",
	}
}

// GetMappingInfoForNetworkGenDetailsNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfoForNetworkGenDetailsNetworkTypeEnum(val string) (InfoForNetworkGenDetailsNetworkTypeEnum, bool) {
	enum, ok := mappingInfoForNetworkGenDetailsNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
