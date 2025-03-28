// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceVcnIpAddressCondition An access control rule condition that requires a match on the specified source VCN and IP address range.
// This condition must be used only in conjunction with `SourceVcnIdCondition`.
type SourceVcnIpAddressCondition struct {

	// An IPv4 address range that the original client IP address (in the context of the specified VCN) of an
	// incoming packet must match.
	// The service accepts only classless inter-domain routing (CIDR) format (x.x.x.x/y) strings.
	// Specify 0.0.0.0/0 to match all incoming traffic in the customer VCN.
	// example: "10.10.1.0/24"
	AttributeValue *string `mandatory:"true" json:"attributeValue"`
}

func (m SourceVcnIpAddressCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SourceVcnIpAddressCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SourceVcnIpAddressCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceVcnIpAddressCondition SourceVcnIpAddressCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypeSourceVcnIpAddressCondition
	}{
		"SOURCE_VCN_IP_ADDRESS",
		(MarshalTypeSourceVcnIpAddressCondition)(m),
	}

	return json.Marshal(&s)
}
