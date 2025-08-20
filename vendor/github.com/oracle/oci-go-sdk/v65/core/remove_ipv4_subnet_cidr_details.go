// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveIpv4SubnetCidrDetails Details object for removing an IPv4 prefix from a subnet.
type RemoveIpv4SubnetCidrDetails struct {

	// This field should only be specified when removing an IPv4 prefix from a subnet's IPv4 address space.
	// The CIDR must maintain the following rules -
	// a. The CIDR block is valid and correctly formatted.
	// b. The CIDR range is within one of the parent VCN ranges.
	// c. The CIDR range to be removed is already present in the list of ipv4CidrBlocks
	// Example: `10.0.1.0/24`
	Ipv4CidrBlock *string `mandatory:"true" json:"ipv4CidrBlock"`
}

func (m RemoveIpv4SubnetCidrDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveIpv4SubnetCidrDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
