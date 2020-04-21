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

// UpdateNetworkSecurityGroupsDetails An object representing an updated list of network security groups (NSGs) that overwrites the existing list of NSGs.
// *  If the load balancer has no NSGs configured, it uses the NSGs in this list.
// *  If the load balancer has a list of NSGs configured, this list replaces the existing list.
// *  If the load balancer has a list of NSGs configured and this list is empty, the operation removes all of the load balancer's NSG associations.
type UpdateNetworkSecurityGroupsDetails struct {

	// An array of NSG OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) associated with the load
	// balancer.
	// During the load balancer's creation, the service adds the new load balancer to the specified NSGs.
	// The benefits of associating the load balancer with NSGs include:
	// *  NSGs define network security rules to govern ingress and egress traffic for the load balancer.
	// *  The network security rules of other resources can reference the NSGs associated with the load balancer
	//    to ensure access.
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`
}

func (m UpdateNetworkSecurityGroupsDetails) String() string {
	return common.PointerString(m)
}
