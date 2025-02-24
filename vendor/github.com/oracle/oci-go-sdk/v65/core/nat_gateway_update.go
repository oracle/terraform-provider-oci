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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NatGatewayUpdate Data plane information about a NAT gateway.  The `id` field contains the
// OCID for the NAT gateway.
// Note that a NAT gateway is only presented to the data plane through this
// protocol when it is active, meaning that its lifecycle state is
// `AVAILABLE` and blockTraffic is False. When its lifecycle state
// transitions from `AVAILABLE` to `TERMINATING`, or when blockTraffic
// transitions from False to True, this protocol signals the NAT gateway's
// deletion.
type NatGatewayUpdate struct {

	// Unique identifier for the primary resource affected by this update,
	// such as its OCID.
	Id *string `mandatory:"false" json:"id"`

	// True iff this update signals deletion of the identified resource.
	// If true, the type-specific fields of this object may be null.
	IsDelete *bool `mandatory:"false" json:"isDelete"`

	// The date and time that the API call was made that led to this Update.
	// The date and time format is defined by RFC3339.
	// Example: '2016-08-25T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// True iff this NAT gateway is currently configured to block all
	// traffic through it.
	BlockTraffic *bool `mandatory:"false" json:"blockTraffic"`

	// The OCID for the NAT gateway's compartment.
	CompartmentOcid *string `mandatory:"false" json:"compartmentOcid"`

	// The OCID for the NAT gateway's VCN.
	VcnOcid *string `mandatory:"false" json:"vcnOcid"`

	// The external-facing IPv4 address of this NAT gateway, in
	// dot-decimal notation, which source IPs are translated to (egress)
	// and destination IPs from (ingress) for traffic passing through the
	// gateway.
	NatIp *string `mandatory:"false" json:"natIp"`

	// The MPLS label which identifies this NAT gateway in encapsulated
	// traffic sent to either the NAT egress or ingress redirectors.
	// This label is scoped by the egress and ingress redirector IPs.
	Label *int `mandatory:"false" json:"label"`

	// The destination substrate IPv4 address, in dot-decimal notation,
	// used to encapsulate egress traffic routed to this NAT gateway
	// (e.g. from Caviums).
	EgressVip *string `mandatory:"false" json:"egressVip"`

	// The destination substrate IPv4 address, in dot-decimal notation,
	// used to encapsulate ingress traffic routed to this NAT gateway
	// (e.g. from the Internet ingress redirector).
	IngressVip *string `mandatory:"false" json:"ingressVip"`

	// How to encapsulate translated egress traffic sent out of this NAT
	// gateway to its substrate "next hop".
	NextHop *SubstrateRoute `mandatory:"false" json:"nextHop"`
}

// GetId returns Id
func (m NatGatewayUpdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m NatGatewayUpdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m NatGatewayUpdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m NatGatewayUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NatGatewayUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NatGatewayUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNatGatewayUpdate NatGatewayUpdate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNatGatewayUpdate
	}{
		"NatGatewayUpdate",
		(MarshalTypeNatGatewayUpdate)(m),
	}

	return json.Marshal(&s)
}
