// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSubnetDetails The representation of UpdateSubnetDetails
type UpdateSubnetDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the set of DHCP options the subnet will use.
	DhcpOptionsId *string `mandatory:"false" json:"dhcpOptionsId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the subnet will use.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The OCIDs of the security list or lists the subnet will use. This
	// replaces the entire current set of security lists. Remember that
	// security lists are associated *with the subnet*, but the rules are
	// applied to the individual VNICs in the subnet.
	SecurityListIds []string `mandatory:"false" json:"securityListIds"`

	// The CIDR block of the subnet. The new CIDR block must meet the following criteria:
	// - Must be valid.
	// - The CIDR block's IP range must be completely within one of the VCN's CIDR block ranges.
	// - The old and new CIDR block ranges must use the same network address. Example: `10.0.0.0/25` and `10.0.0.0/24`.
	// - Must contain all IP addresses in use in the old CIDR range.
	// - The new CIDR range's broadcast address (last IP address of CIDR range) must not be an IP address in use in the old CIDR range.
	// **Note:** If you are changing the CIDR block, you cannot create VNICs or private IPs for this resource while the update is in progress.
	// Example: `172.16.0.0/16`
	CidrBlock *string `mandatory:"false" json:"cidrBlock"`

	// This is the IPv6 CIDR block for the subnet's IP address space.
	// The subnet size is always /64.
	// See IPv6 Addresses (https://docs.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	// The provided CIDR must maintain the following rules -
	// a. The IPv6 CIDR block is valid and correctly formatted.
	// b. The IPv6 CIDR is within the parent VCN IPv6 range.
	// Example: `2001:0db8:0123:1111::/64`
	Ipv6CidrBlock *string `mandatory:"false" json:"ipv6CidrBlock"`

	// The list of all IPv6 CIDR blocks (Oracle allocated IPv6 GUA, ULA or private IPv6 CIDR blocks, BYOIPv6 CIDR blocks) for the subnet that meets the following criteria:
	// - The CIDR blocks must be valid.
	// - Multiple CIDR blocks must not overlap each other or the on-premises network CIDR block.
	// - The number of CIDR blocks must not exceed the limit of IPv6 CIDR blocks allowed to a subnet.
	Ipv6CidrBlocks []string `mandatory:"false" json:"ipv6CidrBlocks"`
}

func (m UpdateSubnetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSubnetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
