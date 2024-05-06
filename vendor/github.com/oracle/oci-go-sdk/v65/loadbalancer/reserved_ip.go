// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReservedIp The representation of ReservedIp
type ReservedIp struct {

	// Ocid of the Reserved IP/Public Ip created with VCN.
	// Reserved IPs are IPs which already registered using VCN API.
	// Create a reserved Public IP and then while creating the load balancer pass the ocid of the reserved IP in this
	// field reservedIp to attach the Ip to Load balancer. Load balancer will be configured to listen to traffic on this IP.
	// Reserved IPs will not be deleted when the Load balancer is deleted. They will be unattached from the Load balancer.
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
