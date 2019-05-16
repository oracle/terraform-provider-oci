// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateIpSecTunnelBgpSessionDetails Details to modify an IPSec Tunnel's BGP session paramaters.
type UpdateIpSecTunnelBgpSessionDetails struct {

	// The IPv4 Address used in the BGP peering session for the Oracle router. Example: 10.0.0.1/31.
	OracleInterfaceIp *string `mandatory:"false" json:"oracleInterfaceIp"`

	// The IPv4 Address used in the BGP peering session for the non-Oracle router. Example: 10.0.0.2/31.
	CustomerInterfaceIp *string `mandatory:"false" json:"customerInterfaceIp"`

	// The value of the remote BGP ASN in asplain format, as a string. Example: 1587232876 (4 byte ASN) or 12345 (2 byte ASN).
	CustomerBgpAsn *string `mandatory:"false" json:"customerBgpAsn"`
}

func (m UpdateIpSecTunnelBgpSessionDetails) String() string {
	return common.PointerString(m)
}
