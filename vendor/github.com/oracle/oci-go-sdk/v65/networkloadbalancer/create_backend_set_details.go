// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateBackendSetDetails The configuration details for creating a backend set in a network load balancer.
// For more information about backend set configuration, see
// Managing Backend Sets (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingbackendsets.htm).
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateBackendSetDetails struct {

	// A user-friendly name for the backend set that must be unique and cannot be changed.
	// Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot
	// contain spaces. Avoid entering confidential information.
	// Example: `example_backend_set`
	Name *string `mandatory:"true" json:"name"`

	// The network load balancer policy for the backend set.
	// Example: `FIVE_TUPLE``
	Policy NetworkLoadBalancingPolicyEnum `mandatory:"true" json:"policy"`

	HealthChecker *HealthCheckerDetails `mandatory:"true" json:"healthChecker"`

	// If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends.
	// Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled.
	// The value is true by default.
	IsPreserveSource *bool `mandatory:"false" json:"isPreserveSource"`

	// If enabled, the network load balancer will continue to distribute traffic in the configured distribution in the event all backends are unhealthy.
	// The value is false by default.
	IsFailOpen *bool `mandatory:"false" json:"isFailOpen"`

	// If enabled existing connections will be forwarded to an alternative healthy backend as soon as current backend becomes unhealthy.
	IsInstantFailoverEnabled *bool `mandatory:"false" json:"isInstantFailoverEnabled"`

	// IP version associated with the backend set.
	IpVersion IpVersionEnum `mandatory:"false" json:"ipVersion,omitempty"`

	// An array of backends to be associated with the backend set.
	Backends []BackendDetails `mandatory:"false" json:"backends"`
}

func (m CreateBackendSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBackendSetDetails) ValidateEnumValue() (bool, error) {
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
