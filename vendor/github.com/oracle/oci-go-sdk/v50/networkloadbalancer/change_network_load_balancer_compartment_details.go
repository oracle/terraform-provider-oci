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

// ChangeNetworkLoadBalancerCompartmentDetails The configuration details for moving a network load balancer to a different compartment.
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ChangeNetworkLoadBalancerCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which to move the network load balancer.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeNetworkLoadBalancerCompartmentDetails) String() string {
	return common.PointerString(m)
}
