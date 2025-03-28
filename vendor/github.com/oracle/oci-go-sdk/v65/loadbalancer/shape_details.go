// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShapeDetails The configuration details to update load balancer to a different shape.
type ShapeDetails struct {

	// Bandwidth in Mbps that determines the total pre-provisioned bandwidth (ingress plus egress).
	// The values must be between 10 and the maximumBandwidthInMbps.
	// Example: `150`
	MinimumBandwidthInMbps *int `mandatory:"true" json:"minimumBandwidthInMbps"`

	// Bandwidth in Mbps that determines the maximum bandwidth (ingress plus egress) that the load balancer can
	// achieve. This bandwidth cannot be always guaranteed. For a guaranteed bandwidth use the minimumBandwidthInMbps
	// parameter.
	// The values must be between minimumBandwidthInMbps and 8000 (8Gbps).
	// Example: `1500`
	MaximumBandwidthInMbps *int `mandatory:"true" json:"maximumBandwidthInMbps"`
}

func (m ShapeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShapeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
