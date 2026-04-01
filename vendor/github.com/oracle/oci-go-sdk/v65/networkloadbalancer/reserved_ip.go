// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

	// Ocid of the Reserved IP (Public IP, Private IP or IPv6) created with VCN.
	// Reserved IPs are IPs which are already registered using VCN API.
	// For public Network load balancers, customer can create a reserved Public IP and/or reserved private IP and/or reserved IPv6 and pass the OCID's in the
	// reservedIps array field to attach the IP addresses to the network load balancer during create
	// For private Network load balancers, customer can create a reserved Private IP and/or reserved IPv6 and pass the OCID's in the
	// reservedIps array field to attach the IP addresses to the network load balancer during create
	// Reserved IPs will not be deleted when the Network Load balancer is deleted. They will be detached from the Network Load balancer.
	// Public IP Example: "ocid1.publicip.oc1.phx.unique_ID"
	// Private IP Example: "ocid1.privateip.oc1.phx.unique_ID"
	// IPV6 example: "ocid1.ipv6.oc1.phx.unique_ID"
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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
