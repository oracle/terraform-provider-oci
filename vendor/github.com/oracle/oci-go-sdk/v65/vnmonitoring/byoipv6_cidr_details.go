// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Byoipv6CidrDetails The list of one or more BYOIPv6 CIDR blocks for the VCN that meets the following criteria:
// - The CIDR must from a BYOIPv6 range.
// - The IPv6 CIDR blocks must be valid.
// - Multiple CIDR blocks must not overlap each other or the on-premises network CIDR block.
// - The number of CIDR blocks must not exceed the limit of IPv6 CIDR blocks allowed to a VCN.
type Byoipv6CidrDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `ByoipRange` resource to which the CIDR block belongs.
	Byoipv6RangeId *string `mandatory:"true" json:"byoipv6RangeId"`

	// An IPv6 CIDR block required to create a VCN with a BYOIP prefix. It could be the whole CIDR block identified in `byoipv6RangeId`, or a subrange.
	// Example: `2001:0db8:0123::/48`
	Ipv6CidrBlock *string `mandatory:"true" json:"ipv6CidrBlock"`
}

func (m Byoipv6CidrDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Byoipv6CidrDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
