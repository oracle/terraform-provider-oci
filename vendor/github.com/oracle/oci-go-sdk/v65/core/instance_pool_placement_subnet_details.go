// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstancePoolPlacementSubnetDetails Base details about the ipv6 subnet
type InstancePoolPlacementSubnetDetails struct {

	// The subnet OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Whether to allocate an IPv6 address at instance/Vnic creation from an IPv6 Enabled
	// Subnet. Default: False. When provided customer may optionally provide an IPv6 prefix
	// (ipv6SubnetCidr) of their choice to assign the IPv6 address from. If ipv6SubnetCidr
	// is not provided then IPv6 Prefix is chosen based on the logic described in "IPv6 Assignment
	// logic in IPv6BL logic" section.
	IsAssignIpv6Ip *bool `mandatory:"false" json:"isAssignIpv6Ip"`

	// A list of IPv6 CIDR ranges from which the Vnic should be assigned an IPv6 address.
	// The caller can provide only the CIDR ranges, in which case, OCI will select an available
	// address from the range.  The caller can optionally choose to leave the CIDR range empty
	// and instead provide the specific IPv6 address that should be used from within that range.
	Ipv6AddressIpv6SubnetCidrPairDetails []InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails `mandatory:"false" json:"ipv6AddressIpv6SubnetCidrPairDetails"`
}

func (m InstancePoolPlacementSubnetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstancePoolPlacementSubnetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
