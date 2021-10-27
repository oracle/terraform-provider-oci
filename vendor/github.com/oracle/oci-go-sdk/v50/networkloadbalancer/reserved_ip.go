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

// ReservedIp An object representing a reserved IP address to be attached or that is already attached to a network load balancer.
type ReservedIp struct {

	// OCID of the reserved public IP address created with the virtual cloud network.
	// Reserved public IP addresses are IP addresses that are registered using the virtual cloud network API.
	// Create a reserved public IP address. When you create the network load balancer, enter the OCID of the reserved public IP address in the
	// reservedIp field to attach the IP address to the network load balancer. This task configures the network load balancer to listen to traffic on this IP address.
	// Reserved public IP addresses are not deleted when the network load balancer is deleted. The IP addresses become unattached from the network load balancer.
	// Example: "ocid1.publicip.oc1.phx.unique_ID"
	Id *string `mandatory:"false" json:"id"`
}

func (m ReservedIp) String() string {
	return common.PointerString(m)
}
