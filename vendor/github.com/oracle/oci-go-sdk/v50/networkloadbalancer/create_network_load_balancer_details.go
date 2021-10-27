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

// CreateNetworkLoadBalancerDetails The properties that define a network load balancer. For more information, see
// Managing a network load balancer (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingloadbalancer.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, then
// contact an administrator. If you are an administrator who writes policies to give users access, then see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// About the API (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// SDKS and Other Tools (https://docs.cloud.oracle.com/Content/API/Concepts/sdks.htm).
type CreateNetworkLoadBalancerDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Network load balancer identifier, which can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The subnet in which the network load balancer is spawned OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically
	// enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact.
	IsPreserveSourceDestination *bool `mandatory:"false" json:"isPreserveSourceDestination"`

	// An array of reserved Ips.
	ReservedIps []ReservedIp `mandatory:"false" json:"reservedIps"`

	// Whether the network load balancer has a virtual cloud network-local (private) IP address.
	// If "true", then the service assigns a private IP address to the network load balancer.
	// If "false", then the service assigns a public IP address to the network load balancer.
	// A public network load balancer is accessible from the internet, depending on the
	// security list rules (https://docs.cloud.oracle.com/Content/network/Concepts/securitylists.htm) for your virtual cloud network. For more information about public and
	// private network load balancers,
	// see How Network Load Balancing Works (https://docs.cloud.oracle.com/Content/Balance/Concepts/balanceoverview.htm#how-network-load-balancing-works).
	// This value is true by default.
	// Example: `true`
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`

	// An array of network security groups OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) associated with the network load
	// balancer.
	// During the creation of the network load balancer, the service adds the new load balancer to the specified network security groups.
	// The benefits of associating the network load balancer with network security groups include:
	// *  Network security groups define network security rules to govern ingress and egress traffic for the network load balancer.
	// *  The network security rules of other resources can reference the network security groups associated with the network load balancer
	//    to ensure access.
	// Example: ["ocid1.nsg.oc1.phx.unique_ID"]
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`

	// Listeners associated with the network load balancer.
	Listeners map[string]ListenerDetails `mandatory:"false" json:"listeners"`

	// Backend sets associated with the network load balancer.
	BackendSets map[string]BackendSetDetails `mandatory:"false" json:"backendSets"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateNetworkLoadBalancerDetails) String() string {
	return common.PointerString(m)
}
