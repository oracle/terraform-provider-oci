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

// CreateListenerDetails The configuration of the listener.
// For more information about backend set configuration, see
// Managing Load Balancer Listeners (https://docs.cloud.oracle.com/Content/Balance/Tasks/managinglisteners.htm).
type CreateListenerDetails struct {

	// A friendly name for the listener. It must be unique and it cannot be changed.
	// Example: `example_listener`
	Name *string `mandatory:"true" json:"name"`

	// The name of the associated backend set.
	// Example: `example_backend_set`
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName"`

	// The communication port for the listener.
	// Example: `80`
	Port *int `mandatory:"true" json:"port"`

	// The protocol on which the listener accepts connection requests.
	// For public network load balancers, ANY protocol refers to TCP/UDP.
	// For private network load balancers, ANY protocol refers to TCP/UDP/ICMP (note that ICMP requires isPreserveSourceDestination to be set to true).
	// To get a list of valid protocols, use the ListNetworkLoadBalancersProtocols
	// operation.
	// Example: `TCP`
	Protocol ListenerProtocolsEnum `mandatory:"true" json:"protocol"`
}

func (m CreateListenerDetails) String() string {
	return common.PointerString(m)
}
