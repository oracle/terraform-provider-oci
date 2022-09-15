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

// InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails Optional. Used to specify from which subnet CIDRs an IPv6 address should be allocated, or to assign valid available IPv6 addresses.
type InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails struct {

	// Optional. Used to disambiguate which Subnet CIDR should be used to create an IPv6 allocation.
	Ipv6SubnetCidr *string `mandatory:"false" json:"ipv6SubnetCidr"`

	// Optional. An available IPv6 address of your subnet from a valid IPv6 CIDR on the subnet (otherwise the IP address is automatically assigned).
	Ipv6Address *string `mandatory:"false" json:"ipv6Address"`
}

func (m InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
