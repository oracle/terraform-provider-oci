// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// UpdateSubnetDetails The representation of UpdateSubnetDetails
type UpdateSubnetDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID of the set of DHCP options the subnet will use.
	DhcpOptionsId *string `mandatory:"false" json:"dhcpOptionsId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the route table the subnet will use.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The OCIDs of the security list or lists the subnet will use. This
	// replaces the entire current set of security lists. Remember that
	// security lists are associated *with the subnet*, but the rules are
	// applied to the individual VNICs in the subnet.
	SecurityListIds []string `mandatory:"false" json:"securityListIds"`

	// The CIDR IP address block of the Subnet. The CIDR must maintain the following rules -
	// a. The CIDR block is valid and correctly formatted.
	// b. The new range is within one of the parent VCN ranges.
	// c. The old and new CIDR ranges both use the same base address. Example: 10.0.0.0/25 and 10.0.0.0/24.
	// d. The new CIDR range contains all previously allocated private IP addresses in the old CIDR range.
	// e. No previously allocated IP address overlaps the broadcast address (the last IP of a subnet CIDR range) of the new CIDR range.
	// Example: `172.16.0.0/16`
	CidrBlock *string `mandatory:"false" json:"cidrBlock"`
}

func (m UpdateSubnetDetails) String() string {
	return common.PointerString(m)
}
