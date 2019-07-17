// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

// ChangeLoadBalancerCompartmentDetails The configuration details for moving a load balancer to a different compartment.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ChangeLoadBalancerCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to move the load balancer to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeLoadBalancerCompartmentDetails) String() string {
	return common.PointerString(m)
}
