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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReservedIp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
