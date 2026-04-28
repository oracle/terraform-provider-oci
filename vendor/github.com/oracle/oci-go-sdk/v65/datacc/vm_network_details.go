// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmNetworkDetails Details of the client or backup networks in an VM cluster.
type VmNetworkDetails struct {

	// The network type.
	NetworkType VmClusterNetworkTypeEnum `mandatory:"true" json:"networkType"`

	// The network netmask.
	Netmask *string `mandatory:"true" json:"netmask"`

	// The network gateway.
	Gateway *string `mandatory:"true" json:"gateway"`

	// The network domain name.
	DomainName *string `mandatory:"true" json:"domainName"`

	// The list of node details.
	Nodes []NodeDetails `mandatory:"true" json:"nodes"`

	// The network VLAN ID.
	VlanId *string `mandatory:"false" json:"vlanId"`

	// The network domain name prefix.
	Prefix *string `mandatory:"false" json:"prefix"`
}

func (m VmNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmClusterNetworkTypeEnum(string(m.NetworkType)); !ok && m.NetworkType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkType: %s. Supported values are: %s.", m.NetworkType, strings.Join(GetVmClusterNetworkTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
