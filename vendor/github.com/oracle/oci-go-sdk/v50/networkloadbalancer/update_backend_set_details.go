// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// UpdateBackendSetDetails The configuration details for updating a load balancer backend set.
// For more information about backend set configuration, see
// Managing Backend Sets (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingbackendsets.htm).
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateBackendSetDetails struct {

	// The network load balancer policy for the backend set. To get a list of available policies, use the
	// ListNetworkLoadBalancersPolicies operation.
	// Example: `FIVE_TUPLE`
	Policy *string `mandatory:"false" json:"policy"`

	// If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends.
	// Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled.
	// The value is true by default.
	IsPreserveSource *bool `mandatory:"false" json:"isPreserveSource"`

	// An array of backends associated with the backend set.
	Backends []BackendDetails `mandatory:"false" json:"backends"`

	HealthChecker *HealthCheckerDetails `mandatory:"false" json:"healthChecker"`
}

func (m UpdateBackendSetDetails) String() string {
	return common.PointerString(m)
}
