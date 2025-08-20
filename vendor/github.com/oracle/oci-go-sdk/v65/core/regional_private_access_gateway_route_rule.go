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

// RegionalPrivateAccessGatewayRouteRule Route Rule for regional Private Access Gateway (PAGW) Route Table
type RegionalPrivateAccessGatewayRouteRule struct {

	// Represents the range of IP addresses to match against when routing traffic.
	// Potential values:
	//    * An IP address range (IPv4 or IPv6) in CIDR notation. For example: `192.168.1.0/24`
	//    or `2001:0db8:0123:45::/56`.
	Destination *string `mandatory:"true" json:"destination"`

	// Next hop IP (IPv4 or IPv6) for a given destination
	NextHopIpAddress *string `mandatory:"true" json:"nextHopIpAddress"`
}

func (m RegionalPrivateAccessGatewayRouteRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegionalPrivateAccessGatewayRouteRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
