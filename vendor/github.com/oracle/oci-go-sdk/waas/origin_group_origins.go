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

	// The reference string to the origin server.
	Origin *string `mandatory:"false" json:"origin"`

	// The weight of the origin used in load balancing. The higher the weight, the larger the proportion of client requests the server receives.
	Weight *int `mandatory:"false" json:"weight"`
}

func (m OriginGroupOrigins) String() string {
	return common.PointerString(m)
}
