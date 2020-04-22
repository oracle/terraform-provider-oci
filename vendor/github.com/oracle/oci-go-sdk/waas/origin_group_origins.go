// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
