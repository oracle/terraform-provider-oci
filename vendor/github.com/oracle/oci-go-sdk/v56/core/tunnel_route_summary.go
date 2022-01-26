// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TunnelRouteSummary The routes advertised to the Customer and the routes received from the Customer
type TunnelRouteSummary struct {

	// BGP Network Layer Reachability Information
	Prefix *string `mandatory:"false" json:"prefix"`

	// The age of the route
	Age *int64 `mandatory:"false" json:"age"`

	// Is this the best route
	IsBestPath *bool `mandatory:"false" json:"isBestPath"`

	// List of ASNs in AS Path
	AsPath []int `mandatory:"false" json:"asPath"`

	// Route advertiser
	Advertiser TunnelRouteSummaryAdvertiserEnum `mandatory:"false" json:"advertiser,omitempty"`
}

func (m TunnelRouteSummary) String() string {
	return common.PointerString(m)
}

// TunnelRouteSummaryAdvertiserEnum Enum with underlying type: string
type TunnelRouteSummaryAdvertiserEnum string

// Set of constants representing the allowable values for TunnelRouteSummaryAdvertiserEnum
const (
	TunnelRouteSummaryAdvertiserCustomer TunnelRouteSummaryAdvertiserEnum = "CUSTOMER"
	TunnelRouteSummaryAdvertiserOracle   TunnelRouteSummaryAdvertiserEnum = "ORACLE"
)

var mappingTunnelRouteSummaryAdvertiser = map[string]TunnelRouteSummaryAdvertiserEnum{
	"CUSTOMER": TunnelRouteSummaryAdvertiserCustomer,
	"ORACLE":   TunnelRouteSummaryAdvertiserOracle,
}

// GetTunnelRouteSummaryAdvertiserEnumValues Enumerates the set of values for TunnelRouteSummaryAdvertiserEnum
func GetTunnelRouteSummaryAdvertiserEnumValues() []TunnelRouteSummaryAdvertiserEnum {
	values := make([]TunnelRouteSummaryAdvertiserEnum, 0)
	for _, v := range mappingTunnelRouteSummaryAdvertiser {
		values = append(values, v)
	}
	return values
}
