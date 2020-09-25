// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreateCidrBlocksDetails The configuration details of the CidrBlocks.
// CidrBlocks contains a name and list of CIDR block. Each of the CidrBlocks should have unique name
// within the load balancer. CidrBlocks resource name can be used in rule conditions.
// Example:
//  "name" : `ClientRealIpCidrBlocks`
//  "items" : `["129.213.176.0/24","150.136.187.0/24", "2002::1234:abcd:ffff:c0a8:101/64"]`
// **Warning:** No confidential information should be passed in this API.
type CreateCidrBlocksDetails struct {

	// A friendly name for the CidrBlocks.
	// Example: `SourceIpCidrBlocks`
	Name *string `mandatory:"true" json:"name"`

	// Each element in the list should be valid IPv4 or IPv6 CIDR Block address.
	// Example: '["129.213.176.0/24", "150.136.187.0/24", "2002::1234:abcd:ffff:c0a8:101/64"]'
	Items []string `mandatory:"true" json:"items"`
}

func (m CreateCidrBlocksDetails) String() string {
	return common.PointerString(m)
}
