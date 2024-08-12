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

	// If enabled, the network load balancer will continue to distribute traffic in the configured distribution in the event all backends are unhealthy.
	// The value is false by default.
	IsFailOpen *bool `mandatory:"false" json:"isFailOpen"`

	// If enabled existing connections will be forwarded to an alternative healthy backend as soon as current backend becomes unhealthy.
	IsInstantFailoverEnabled *bool `mandatory:"false" json:"isInstantFailoverEnabled"`

	// The IP version associated with the backend set.
	IpVersion IpVersionEnum `mandatory:"false" json:"ipVersion,omitempty"`

	// An array of backends associated with the backend set.
	Backends []BackendDetails `mandatory:"false" json:"backends"`

	HealthChecker *HealthCheckerDetails `mandatory:"false" json:"healthChecker"`
}

func (m UpdateBackendSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateBackendSetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIpVersionEnum(string(m.IpVersion)); !ok && m.IpVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpVersion: %s. Supported values are: %s.", m.IpVersion, strings.Join(GetIpVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
