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

// FlexTunnelBgpSession Boarder Gateway Protocol (BGP) fields.
type FlexTunnelBgpSession struct {

	// The BGP ASN of the network on your end of the BGP session.
	CustomerBgpAsn *string `mandatory:"true" json:"customerBgpAsn"`

	// The Oracle BPG ASN number.
	OracleBgpAsn *string `mandatory:"true" json:"oracleBgpAsn"`

	// This IPv4 CIDR block is for your end of the inside tunnel interface.
	CustomerBgpIp *string `mandatory:"true" json:"customerBgpIp"`

	// The IPv4 CIDR block for the Oracle end of the inside tunnel interface.
	OracleBgpIp *string `mandatory:"true" json:"oracleBgpIp"`

	// This IPv6 prefix is for your end of the inside tunnel interface.
	CustomerBgpIpv6 *string `mandatory:"false" json:"customerBgpIpv6"`

	// The IPv6 prefix block for the Oracle end of the inside tunnel interface.
	OracleBgpIpv6 *string `mandatory:"false" json:"oracleBgpIpv6"`
}

func (m FlexTunnelBgpSession) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlexTunnelBgpSession) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
