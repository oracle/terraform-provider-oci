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

// BgpSessionInfo Information needed to establish a BGP Session on an interface.
type BgpSessionInfo struct {

	// This is the IPv4 Address used in the BGP peering session for the Oracle router. Example: 10.0.0.1/31
	OracleInterfaceIp *string `mandatory:"false" json:"oracleInterfaceIp"`

	// This is the IPv4 Address used in the BGP peering session for the non-Oracle router. Example: 10.0.0.2/31
	CustomerInterfaceIp *string `mandatory:"false" json:"customerInterfaceIp"`

	// This is the value of the Oracle Bgp ASN in asplain format, as a string. Example: 1587232876 (4 byte ASN) or 12345 (2 byte ASN)
	OracleBgpAsn *string `mandatory:"false" json:"oracleBgpAsn"`

	// This is the value of the remote Bgp ASN in asplain format, as a string. Example: 1587232876 (4 byte ASN) or 12345 (2 byte ASN)
	CustomerBgpAsn *string `mandatory:"false" json:"customerBgpAsn"`

	// the state of the BGP.
	BgpState BgpSessionInfoBgpStateEnum `mandatory:"false" json:"bgpState,omitempty"`
}

func (m BgpSessionInfo) String() string {
	return common.PointerString(m)
}

// BgpSessionInfoBgpStateEnum Enum with underlying type: string
type BgpSessionInfoBgpStateEnum string

// Set of constants representing the allowable values for BgpSessionInfoBgpStateEnum
const (
	BgpSessionInfoBgpStateUp   BgpSessionInfoBgpStateEnum = "UP"
	BgpSessionInfoBgpStateDown BgpSessionInfoBgpStateEnum = "DOWN"
)

var mappingBgpSessionInfoBgpState = map[string]BgpSessionInfoBgpStateEnum{
	"UP":   BgpSessionInfoBgpStateUp,
	"DOWN": BgpSessionInfoBgpStateDown,
}

// GetBgpSessionInfoBgpStateEnumValues Enumerates the set of values for BgpSessionInfoBgpStateEnum
func GetBgpSessionInfoBgpStateEnumValues() []BgpSessionInfoBgpStateEnum {
	values := make([]BgpSessionInfoBgpStateEnum, 0)
	for _, v := range mappingBgpSessionInfoBgpState {
		values = append(values, v)
	}
	return values
}
