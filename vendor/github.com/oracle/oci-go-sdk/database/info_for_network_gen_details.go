// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// InfoForNetworkGenDetails Parameters for generation of the client or backup network in a VM cluster network.
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

// InfoForNetworkGenDetailsNetworkTypeEnum Enum with underlying type: string
type InfoForNetworkGenDetailsNetworkTypeEnum string

// Set of constants representing the allowable values for InfoForNetworkGenDetailsNetworkTypeEnum
const (
	InfoForNetworkGenDetailsNetworkTypeClient InfoForNetworkGenDetailsNetworkTypeEnum = "CLIENT"
	InfoForNetworkGenDetailsNetworkTypeBackup InfoForNetworkGenDetailsNetworkTypeEnum = "BACKUP"
)

var mappingInfoForNetworkGenDetailsNetworkType = map[string]InfoForNetworkGenDetailsNetworkTypeEnum{
	"CLIENT": InfoForNetworkGenDetailsNetworkTypeClient,
	"BACKUP": InfoForNetworkGenDetailsNetworkTypeBackup,
}

// GetInfoForNetworkGenDetailsNetworkTypeEnumValues Enumerates the set of values for InfoForNetworkGenDetailsNetworkTypeEnum
func GetInfoForNetworkGenDetailsNetworkTypeEnumValues() []InfoForNetworkGenDetailsNetworkTypeEnum {
	values := make([]InfoForNetworkGenDetailsNetworkTypeEnum, 0)
	for _, v := range mappingInfoForNetworkGenDetailsNetworkType {
		values = append(values, v)
	}
	return values
}
