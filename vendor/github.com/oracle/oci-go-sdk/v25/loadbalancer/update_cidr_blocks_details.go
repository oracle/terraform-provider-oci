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

// UpdateCidrBlocksDetails The configuration details for updating CidrBlocks. If an empty array is sent the request will be disallowed.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateCidrBlocksDetails struct {

	// Each element in the list should be valid IPv4 or IPv6 CIDR Block address.
	// Example: '["129.213.176.0/24", "150.136.187.0/24", "2002::1234:abcd:ffff:c0a8:101/64"]'
	Items []string `mandatory:"true" json:"items"`
}

func (m UpdateCidrBlocksDetails) String() string {
	return common.PointerString(m)
}
