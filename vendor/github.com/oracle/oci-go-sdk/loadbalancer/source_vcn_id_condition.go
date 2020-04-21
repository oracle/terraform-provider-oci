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
	"github.com/oracle/oci-go-sdk/common"
)

// SourceVcnIdCondition An access control rule condition that requires a match on the specified source VCN OCID.
type SourceVcnIdCondition struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the originating VCN that an incoming packet
	// must match.
	// You can use this condition in conjunction with `SourceVcnIpAddressCondition`.
	// **NOTE:** If you define this condition for a rule without a `SourceVcnIpAddressCondition`, this condition
	// matches all incoming traffic in the specified VCN.
	AttributeValue *string `mandatory:"true" json:"attributeValue"`
}

func (m SourceVcnIdCondition) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SourceVcnIdCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceVcnIdCondition SourceVcnIdCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypeSourceVcnIdCondition
	}{
		"SOURCE_VCN_ID",
		(MarshalTypeSourceVcnIdCondition)(m),
	}

	return json.Marshal(&s)
}
