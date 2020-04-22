// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// VmNetworkDetails Details of the client or backup networks in a VM cluster network.
type VmNetworkDetails struct {

	// The network VLAN ID.
	VlanId *string `mandatory:"true" json:"vlanId"`

	// The network type.
	NetworkType VmNetworkDetailsNetworkTypeEnum `mandatory:"true" json:"networkType"`

	// The network netmask.
	Netmask *string `mandatory:"true" json:"netmask"`

	// The network gateway.
	Gateway *string `mandatory:"true" json:"gateway"`

	// The network domain name.
	DomainName *string `mandatory:"true" json:"domainName"`

	// The list of node details.
	Nodes []NodeDetails `mandatory:"true" json:"nodes"`
}

func (m VmNetworkDetails) String() string {
	return common.PointerString(m)
}

// VmNetworkDetailsNetworkTypeEnum Enum with underlying type: string
type VmNetworkDetailsNetworkTypeEnum string

// Set of constants representing the allowable values for VmNetworkDetailsNetworkTypeEnum
const (
	VmNetworkDetailsNetworkTypeClient VmNetworkDetailsNetworkTypeEnum = "CLIENT"
	VmNetworkDetailsNetworkTypeBackup VmNetworkDetailsNetworkTypeEnum = "BACKUP"
)

var mappingVmNetworkDetailsNetworkType = map[string]VmNetworkDetailsNetworkTypeEnum{
	"CLIENT": VmNetworkDetailsNetworkTypeClient,
	"BACKUP": VmNetworkDetailsNetworkTypeBackup,
}

// GetVmNetworkDetailsNetworkTypeEnumValues Enumerates the set of values for VmNetworkDetailsNetworkTypeEnum
func GetVmNetworkDetailsNetworkTypeEnumValues() []VmNetworkDetailsNetworkTypeEnum {
	values := make([]VmNetworkDetailsNetworkTypeEnum, 0)
	for _, v := range mappingVmNetworkDetailsNetworkType {
		values = append(values, v)
	}
	return values
}
