// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceNetworkInterface Describes a network interface.
type InstanceNetworkInterface struct {
	Association *InstanceNetworkInterfaceAssociation `mandatory:"false" json:"association"`

	Attachment *InstanceNetworkInterfaceAttachment `mandatory:"false" json:"attachment"`

	// The description.
	Description *string `mandatory:"false" json:"description"`

	// The security groups.
	SecurityGroups []GroupIdentifier `mandatory:"false" json:"securityGroups"`

	// The type of network interface.
	InterfaceType *string `mandatory:"false" json:"interfaceType"`

	// The IPv4 delegated prefixes that are assigned to the network interface.
	Ipv4Prefixes []string `mandatory:"false" json:"ipv4Prefixes"`

	// The IPv6 addresses associated with the network interface.
	Ipv6Addresses []string `mandatory:"false" json:"ipv6Addresses"`

	// The IPv6 delegated prefixes that are assigned to the network interface.
	Ipv6Prefixes []string `mandatory:"false" json:"ipv6Prefixes"`

	// The MAC address.
	MacAddress *string `mandatory:"false" json:"macAddress"`

	// The ID of the network interface.
	NetworkInterfaceKey *string `mandatory:"false" json:"networkInterfaceKey"`

	// The ID of the AWS account that created the network interface.
	OwnerKey *string `mandatory:"false" json:"ownerKey"`

	// The private IPv4 addresses associated with the network interface.
	PrivateIpAddresses []InstancePrivateIpAddress `mandatory:"false" json:"privateIpAddresses"`

	// Indicates whether source/destination checking is enabled.
	IsSourceDestCheck *bool `mandatory:"false" json:"isSourceDestCheck"`

	// The status of the network interface.
	Status *string `mandatory:"false" json:"status"`

	// The ID of the subnet.
	SubnetKey *string `mandatory:"false" json:"subnetKey"`
}

func (m InstanceNetworkInterface) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceNetworkInterface) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
