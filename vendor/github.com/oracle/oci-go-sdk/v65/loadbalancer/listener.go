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

// Listener The listener's configuration.
// For more information on backend set configuration, see
// Managing Load Balancer Listeners (https://docs.oracle.com/iaas/Content/Balance/Tasks/managinglisteners.htm).
type Listener struct {

	// A friendly name for the listener. It must be unique and it cannot be changed.
	// Example: `example_listener`
	Name *string `mandatory:"true" json:"name"`

	// The name of the associated backend set.
	// Example: `example_backend_set`
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName"`

	// The communication port for the listener.
	// Example: `80`
	Port *int `mandatory:"true" json:"port"`

	// The protocol on which the listener accepts connection requests. The supported protocols are HTTP, HTTP2, TCP, and GRPC.
	// You can also use the ListProtocols operation to get a list of valid protocols.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol"`

	// An array of hostname resource names.
	HostnameNames []string `mandatory:"false" json:"hostnameNames"`

	// Deprecated. Please use `routingPolicies` instead.
	// The name of the set of path-based routing rules, PathRouteSet,
	// applied to this listener's traffic.
	// Example: `example_path_route_set`
	PathRouteSetName *string `mandatory:"false" json:"pathRouteSetName"`

	SslConfiguration *SslConfiguration `mandatory:"false" json:"sslConfiguration"`

	ConnectionConfiguration *ConnectionConfiguration `mandatory:"false" json:"connectionConfiguration"`

	// The names of the RuleSet to apply to the listener.
	// Example: ["example_rule_set"]
	RuleSetNames []string `mandatory:"false" json:"ruleSetNames"`

	// The name of the routing policy applied to this listener's traffic.
	// Example: `example_routing_policy_name`
	RoutingPolicyName *string `mandatory:"false" json:"routingPolicyName"`
}

func (m Listener) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Listener) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
