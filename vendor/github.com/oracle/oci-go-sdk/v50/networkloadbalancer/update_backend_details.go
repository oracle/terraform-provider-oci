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

// UpdateBackendDetails The configuration details for updating a backend server.
type UpdateBackendDetails struct {

	// The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger
	// proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections
	// as a server weighted '1'.
	// For more information about load balancing policies, see
	// How Load Balancing Policies Work (https://docs.cloud.oracle.com/Content/Balance/Reference/lbpolicies.htm).
	// Example: `3`
	Weight *int `mandatory:"false" json:"weight"`

	// Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress
	// traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.
	// Example: `false`
	IsBackup *bool `mandatory:"false" json:"isBackup"`

	// Whether the network load balancer should drain this server. Servers marked "isDrain" receive no
	// incoming traffic.
	// Example: `false`
	IsDrain *bool `mandatory:"false" json:"isDrain"`

	// Whether the network load balancer should treat this server as offline. Offline servers receive no incoming
	// traffic.
	// Example: `false`
	IsOffline *bool `mandatory:"false" json:"isOffline"`
}

func (m UpdateBackendDetails) String() string {
	return common.PointerString(m)
}
