// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LoadBalancerPolicy A policy that determines how traffic is distributed among backend servers.
// For more information on load balancing policies, see
// [How Load Balancing Policies Work]({{DOC_SERVER_URL}}/Content/Balance/Reference/lbpolicies.htm).
type LoadBalancerPolicy struct {

	// The name of the load balancing policy.
	Name *string `mandatory:"true" json:"name"`
}

func (m LoadBalancerPolicy) String() string {
	return common.PointerString(m)
}
