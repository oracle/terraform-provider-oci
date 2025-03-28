// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackendDetails The network network load balancing configuration details of a backend server.
type BackendDetails struct {

	// The communication port for the backend server.
	// Example: `8080`
	Port *int `mandatory:"true" json:"port"`

	// A read-only field showing the IP address/OCID and port that uniquely identify this backend server in the backend set.
	// Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0`
	Name *string `mandatory:"false" json:"name"`

	// The IP address of the backend server.
	// Example: `10.0.0.3`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The IP OCID/Instance OCID associated with the backend server.
	// Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>`
	TargetId *string `mandatory:"false" json:"targetId"`

	// The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger
	// proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections
	// as a server weighted '1'.
	// For more information about load balancing policies, see
	// Network Load Balancer Policies (https://docs.oracle.com/iaas/Content/NetworkLoadBalancer/introduction.htm#Policies).
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

func (m BackendDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackendDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
