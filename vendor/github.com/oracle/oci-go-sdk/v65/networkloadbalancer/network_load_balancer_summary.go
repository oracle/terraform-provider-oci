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

// NetworkLoadBalancerSummary Network load balancer object to be used for list operations.
type NetworkLoadBalancerSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network load balancer.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name, which does not have to be unique, and can be changed.
	// Example: `example_load_balancer`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the network load balancer.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the network load balancer was created, in the format defined by RFC3339.
	// Example: `2020-05-01T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// An array of IP addresses.
	IpAddresses []IpAddress `mandatory:"true" json:"ipAddresses"`

	// The subnet in which the network load balancer is spawned OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// IP version associated with the NLB.
	NlbIpVersion NlbIpVersionEnum `mandatory:"false" json:"nlbIpVersion,omitempty"`

	// The time the network load balancer was updated. An RFC3339 formatted date-time string.
	// Example: `2020-05-01T22:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Whether the network load balancer has a virtual cloud network-local (private) IP address.
	// If "true", then the service assigns a private IP address to the network load balancer.
	// If "false", then the service assigns a public IP address to the network load balancer.
	// A public network load balancer is accessible from the internet, depending the
	// security list rules (https://docs.cloud.oracle.com/Content/network/Concepts/securitylists.htm) for your virtual cloudn network. For more information about public and
	// private network load balancers,
	// see Network Load Balancer Types (https://docs.cloud.oracle.com/Content/NetworkLoadBalancer/introduction.htm#NetworkLoadBalancerTypes).
	// This value is true by default.
	// Example: `true`
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`

	// When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC.
	// Packets are sent to the backend set without any changes to the source and destination IP.
	IsPreserveSourceDestination *bool `mandatory:"false" json:"isPreserveSourceDestination"`

	// This can only be enabled when NLB is working in transparent mode with source destination header preservation enabled.
	// This removes the additional dependency from NLB backends(like Firewalls) to perform SNAT.
	IsSymmetricHashEnabled *bool `mandatory:"false" json:"isSymmetricHashEnabled"`

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
	Listeners map[string]Listener `mandatory:"false" json:"listeners"`

	// Backend sets associated with the network load balancer.
	BackendSets map[string]BackendSet `mandatory:"false" json:"backendSets"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// ZPR tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{ "oracle-zpr": { "td": { "value": "42", "mode": "audit" } } }`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Key-value pair representing system tags' keys and values scoped to a namespace.
	// Example: `{"bar-key": "value"}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m NetworkLoadBalancerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkLoadBalancerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingNlbIpVersionEnum(string(m.NlbIpVersion)); !ok && m.NlbIpVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NlbIpVersion: %s. Supported values are: %s.", m.NlbIpVersion, strings.Join(GetNlbIpVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
