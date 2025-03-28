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

// CreateIpSecTunnelBgpSessionDetails The representation of CreateIpSecTunnelBgpSessionDetails
type CreateIpSecTunnelBgpSessionDetails struct {

	// The IP address for the Oracle end of the inside tunnel interface.
	// If the tunnel's `routing` attribute is set to `BGP`
	// (see IPSecConnectionTunnel), this IP address
	// is required and used for the tunnel's BGP session.
	// If `routing` is instead set to `STATIC`, this IP address is optional. You can set this IP
	// address to troubleshoot or monitor the tunnel.
	// The value must be a /30 or /31.
	// Example: `10.0.0.4/31`
	OracleInterfaceIp *string `mandatory:"false" json:"oracleInterfaceIp"`

	// The IP address for the CPE end of the inside tunnel interface.
	// If the tunnel's `routing` attribute is set to `BGP`
	// (see IPSecConnectionTunnel), this IP address
	// is required and used for the tunnel's BGP session.
	// If `routing` is instead set to `STATIC`, this IP address is optional. You can set this IP
	// address to troubleshoot or monitor the tunnel.
	// The value must be a /30 or /31.
	// Example: `10.0.0.5/31`
	CustomerInterfaceIp *string `mandatory:"false" json:"customerInterfaceIp"`

	// The IPv6 address for the Oracle end of the inside tunnel interface. This IP address is optional.
	// If the tunnel's `routing` attribute is set to `BGP`
	// (see IPSecConnectionTunnel), this IP address
	// is used for the tunnel's BGP session.
	// If `routing` is instead set to `STATIC`, you can set this IP
	// address to troubleshoot or monitor the tunnel.
	// Only subnet masks from /64 up to /127 are allowed.
	// Example: `2001:db8::1/64`
	OracleInterfaceIpv6 *string `mandatory:"false" json:"oracleInterfaceIpv6"`

	// The IPv6 address for the CPE end of the inside tunnel interface. This IP address is optional.
	// If the tunnel's `routing` attribute is set to `BGP`
	// (see IPSecConnectionTunnel), this IP address
	// is used for the tunnel's BGP session.
	// If `routing` is instead set to `STATIC`, you can set this IP
	// address to troubleshoot or monitor the tunnel.
	// Only subnet masks from /64 up to /127 are allowed.
	// Example: `2001:db8::1/64`
	CustomerInterfaceIpv6 *string `mandatory:"false" json:"customerInterfaceIpv6"`

	// If the tunnel's `routing` attribute is set to `BGP`
	// (see IPSecConnectionTunnel), this ASN
	// is required and used for the tunnel's BGP session. This is the ASN of the network on the
	// CPE end of the BGP session. Can be a 2-byte or 4-byte ASN. Uses "asplain" format.
	// If the tunnel's `routing` attribute is set to `STATIC`, the `customerBgpAsn` must be null.
	// Example: `12345` (2-byte) or `1587232876` (4-byte)
	CustomerBgpAsn *string `mandatory:"false" json:"customerBgpAsn"`
}

func (m CreateIpSecTunnelBgpSessionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateIpSecTunnelBgpSessionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
