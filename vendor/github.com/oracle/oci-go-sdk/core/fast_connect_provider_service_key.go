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

// FastConnectProviderServiceKey A provider service key and its details. A provider service key is an identifier for a provider's
// virtual circuit.
type FastConnectProviderServiceKey struct {

	// The service key that the provider gives you when you set up a virtual circuit connection
	// from the provider to Oracle Cloud Infrastructure. Use this value as the `providerServiceKeyName`
	// query parameter for
	// GetFastConnectProviderServiceKey.
	Name *string `mandatory:"false" json:"name"`

	// The provisioned data rate of the connection.  To get a list of the
	// available bandwidth levels (that is, shapes), see
	// ListFastConnectProviderVirtualCircuitBandwidthShapes.
	// Example: `10 Gbps`
	BandwidthShapeName *string `mandatory:"false" json:"bandwidthShapeName"`

	// The provider's peering location.
	PeeringLocation *string `mandatory:"false" json:"peeringLocation"`
}

func (m FastConnectProviderServiceKey) String() string {
	return common.PointerString(m)
}
