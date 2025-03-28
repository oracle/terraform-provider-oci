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

// Backend The configuration of a backend server that is a member of a load balancer backend set.
// For more information, see Managing Backend Servers (https://docs.oracle.com/iaas/Content/Balance/Tasks/managingbackendservers.htm).
type Backend struct {

	// A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.
	// Example: `10.0.0.3:8080`
	Name *string `mandatory:"true" json:"name"`

	// The IP address of the backend server.
	// Example: `10.0.0.3`
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The communication port for the backend server.
	// Example: `8080`
	Port *int `mandatory:"true" json:"port"`

	// The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger
	// proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections
	// as a server weighted '1'.
	// For more information on load balancing policies, see
	// How Load Balancing Policies Work (https://docs.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).
	// Example: `3`
	Weight *int `mandatory:"true" json:"weight"`

	// Whether the load balancer should drain this server. Servers marked "drain" receive no new
	// incoming traffic.
	// Example: `false`
	Drain *bool `mandatory:"true" json:"drain"`

	// Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress
	// traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.
	// **Note:** You cannot add a backend server marked as `backup` to a backend set that uses the IP Hash policy.
	// Example: `false`
	Backup *bool `mandatory:"true" json:"backup"`

	// Whether the load balancer should treat this server as offline. Offline servers receive no incoming
	// traffic.
	// Example: `false`
	Offline *bool `mandatory:"true" json:"offline"`

	// The maximum number of simultaneous connections the load balancer can make to the backend.
	// If this is not set then the maximum number of simultaneous connections the load balancer
	// can make to the backend is unlimited.
	// Example: `300`
	MaxConnections *int `mandatory:"false" json:"maxConnections"`
}

func (m Backend) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Backend) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
