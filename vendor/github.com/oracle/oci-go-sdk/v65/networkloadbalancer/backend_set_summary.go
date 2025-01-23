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

// BackendSetSummary The configuration of a network load balancer backend set.
// For more information about backend set configuration, see
// Backend Sets for Network Load Balancers (https://docs.cloud.oracle.com/Content/NetworkLoadBalancer/BackendSets/backend-set-management.htm).
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type BackendSetSummary struct {

	// A user-friendly name for the backend set that must be unique and cannot be changed.
	// Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot
	// contain spaces. Avoid entering confidential information.
	// Example: `example_backend_set`
	Name *string `mandatory:"true" json:"name"`

	// The network load balancer policy for the backend set.
	// Example: `FIVE_TUPLE`
	Policy NetworkLoadBalancingPolicyEnum `mandatory:"true" json:"policy"`

	// An array of backends.
	Backends []Backend `mandatory:"true" json:"backends"`

	HealthChecker *HealthChecker `mandatory:"true" json:"healthChecker"`

	// If this parameter is enabled, the network load balancer preserves the source IP of the packet forwarded to the backend servers.
	// Backend servers see the original source IP. If the `isPreserveSourceDestination` parameter is enabled for the network load balancer resource, this parameter cannot be disabled.
	// The value is true by default.
	IsPreserveSource *bool `mandatory:"false" json:"isPreserveSource"`

	// If enabled, the network load balancer will continue to distribute traffic in the configured distribution in the event all backends are unhealthy.
	// The value is false by default.
	IsFailOpen *bool `mandatory:"false" json:"isFailOpen"`

	// If enabled existing connections will be forwarded to an alternative healthy backend as soon as current backend becomes unhealthy.
	IsInstantFailoverEnabled *bool `mandatory:"false" json:"isInstantFailoverEnabled"`

	// If enabled along with instant failover, the network load balancer will send TCP RST to the clients for the existing connections instead of failing over to a healthy backend. This only applies when using the instant failover. By default, TCP RST is enabled.
	IsInstantFailoverTcpResetEnabled *bool `mandatory:"false" json:"isInstantFailoverTcpResetEnabled"`

	// If enabled, NLB supports active-standby backends. The standby backend takes over the traffic when the active node fails, and continues to serve the traffic even when the old active node is back healthy.
	AreOperationallyActiveBackendsPreferred *bool `mandatory:"false" json:"areOperationallyActiveBackendsPreferred"`

	// IP version associated with the backend set.
	IpVersion IpVersionEnum `mandatory:"false" json:"ipVersion,omitempty"`
}

func (m BackendSetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackendSetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkLoadBalancingPolicyEnum(string(m.Policy)); !ok && m.Policy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Policy: %s. Supported values are: %s.", m.Policy, strings.Join(GetNetworkLoadBalancingPolicyEnumStringValues(), ",")))
	}

	if _, ok := GetMappingIpVersionEnum(string(m.IpVersion)); !ok && m.IpVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpVersion: %s. Supported values are: %s.", m.IpVersion, strings.Join(GetIpVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
