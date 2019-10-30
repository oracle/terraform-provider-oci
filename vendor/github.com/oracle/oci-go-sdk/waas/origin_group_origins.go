// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OriginGroupOrigins The representation of OriginGroupOrigins
type OriginGroupOrigins struct {

	// The IP address or CIDR notation of the origin server.
	Origin *string `mandatory:"false" json:"origin"`

	// The weight of the origin used in load balancing. Origins with higher weights will receive larger proportions of client requests.
	Weight *int `mandatory:"false" json:"weight"`
}

func (m OriginGroupOrigins) String() string {
	return common.PointerString(m)
}
