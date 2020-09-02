// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LoadBalancer The properties that define a load balancer. For more information, see
// Managing a Load Balancer (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingloadbalancer.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// About the API (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// SDKS and Other Tools (https://docs.cloud.oracle.com/Content/API/Concepts/sdks.htm).
type LoadBalancer struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the load balancer.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `example_load_balancer`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the load balancer.
	LifecycleState LoadBalancerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the load balancer was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A template that determines the total pre-provisioned bandwidth (ingress plus egress).
	// To get a list of available shapes, use the ListShapes
	// operation.
	// Example: `100Mbps`
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// An array of IP addresses.
	IpAddresses []IpAddress `mandatory:"false" json:"ipAddresses"`

	// Whether the load balancer has a VCN-local (private) IP address.
	// If "true", the service assigns a private IP address to the load balancer.
	// If "false", the service assigns a public IP address to the load balancer.
	// A public load balancer is accessible from the internet, depending on your VCN's
	// security list rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securitylists.htm). For more information about public and
	// private load balancers, see How Load Balancing Works (https://docs.cloud.oracle.com/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).
	// Example: `true`
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`

	// An array of subnet OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SubnetIds []string `mandatory:"false" json:"subnetIds"`

	// An array of NSG OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) associated with the load
	// balancer.
	// During the load balancer's creation, the service adds the new load balancer to the specified NSGs.
	// The benefits of associating the load balancer with NSGs include:
	// *  NSGs define network security rules to govern ingress and egress traffic for the load balancer.
	// *  The network security rules of other resources can reference the NSGs associated with the load balancer
	//    to ensure access.
	// Example: ["ocid1.nsg.oc1.phx.unique_ID"]
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`

	Listeners map[string]Listener `mandatory:"false" json:"listeners"`

	Hostnames map[string]Hostname `mandatory:"false" json:"hostnames"`

	SslCipherSuites map[string]SslCipherSuite `mandatory:"false" json:"sslCipherSuites"`

	Certificates map[string]Certificate `mandatory:"false" json:"certificates"`

	BackendSets map[string]BackendSet `mandatory:"false" json:"backendSets"`

	PathRouteSets map[string]PathRouteSet `mandatory:"false" json:"pathRouteSets"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	RuleSets map[string]RuleSet `mandatory:"false" json:"ruleSets"`
}

func (m LoadBalancer) String() string {
	return common.PointerString(m)
}

// LoadBalancerLifecycleStateEnum Enum with underlying type: string
type LoadBalancerLifecycleStateEnum string

// Set of constants representing the allowable values for LoadBalancerLifecycleStateEnum
const (
	LoadBalancerLifecycleStateCreating LoadBalancerLifecycleStateEnum = "CREATING"
	LoadBalancerLifecycleStateFailed   LoadBalancerLifecycleStateEnum = "FAILED"
	LoadBalancerLifecycleStateActive   LoadBalancerLifecycleStateEnum = "ACTIVE"
	LoadBalancerLifecycleStateDeleting LoadBalancerLifecycleStateEnum = "DELETING"
	LoadBalancerLifecycleStateDeleted  LoadBalancerLifecycleStateEnum = "DELETED"
)

var mappingLoadBalancerLifecycleState = map[string]LoadBalancerLifecycleStateEnum{
	"CREATING": LoadBalancerLifecycleStateCreating,
	"FAILED":   LoadBalancerLifecycleStateFailed,
	"ACTIVE":   LoadBalancerLifecycleStateActive,
	"DELETING": LoadBalancerLifecycleStateDeleting,
	"DELETED":  LoadBalancerLifecycleStateDeleted,
}

// GetLoadBalancerLifecycleStateEnumValues Enumerates the set of values for LoadBalancerLifecycleStateEnum
func GetLoadBalancerLifecycleStateEnumValues() []LoadBalancerLifecycleStateEnum {
	values := make([]LoadBalancerLifecycleStateEnum, 0)
	for _, v := range mappingLoadBalancerLifecycleState {
		values = append(values, v)
	}
	return values
}
