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

// UpdateNetworkLoadBalancerDetails Configuration details to update a network load balancer.
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateNetworkLoadBalancerDetails struct {

	// The user-friendly display name for the network load balancer, which does not have to be unique and can be changed.
	// Avoid entering confidential information.
	// Example: `example_network_load_balancer`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically
	// enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact.
	IsPreserveSourceDestination *bool `mandatory:"false" json:"isPreserveSourceDestination"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateNetworkLoadBalancerDetails) String() string {
	return common.PointerString(m)
}
