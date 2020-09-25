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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// SourceIpAddressCondition An access control rule condition that requires a match on the specified source IP address or address range.
// An access control rule condition that requires a match on the specified source IP address. The matching IP Address
// can be specified either as a single CIDR by value or referenced via a named CidrBlocks. If latter is used, then
// CidrBlocks must be created in the context of this Load Balancer.
type SourceIpAddressCondition struct {

	// An IPv4 or IPv6 address range that the source IP address of an incoming packet must match.
	// The service accepts only classless inter-domain routing (CIDR) format (x.x.x.x/y or x:x::x/y) strings.
	// Specify 0.0.0.0/0 or ::/0 to match all incoming traffic.
	// Besides IP ranges or IPs you can match against multiple CIDR blocks by creating a CidrBlocks resource with
	// a list of CIDR blocks and mentioning the name of CidrBlocks resource.
	// example: "192.168.0.0/16 or MySourceIPCidrBlocks"
	AttributeValue *string `mandatory:"true" json:"attributeValue"`
}

func (m SourceIpAddressCondition) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SourceIpAddressCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceIpAddressCondition SourceIpAddressCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypeSourceIpAddressCondition
	}{
		"SOURCE_IP_ADDRESS",
		(MarshalTypeSourceIpAddressCondition)(m),
	}

	return json.Marshal(&s)
}
