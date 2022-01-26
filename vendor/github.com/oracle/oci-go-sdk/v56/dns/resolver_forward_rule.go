// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ResolverForwardRule The representation of ResolverForwardRule
type ResolverForwardRule struct {

	// A list of CIDR blocks. The query must come from a client within one of the blocks in order for the rule action
	// to apply.
	ClientAddressConditions []string `mandatory:"true" json:"clientAddressConditions"`

	// A list of domain names. The query must be covered by one of the domains in order for the rule action to apply.
	QnameCoverConditions []string `mandatory:"true" json:"qnameCoverConditions"`

	// IP addresses to which queries should be forwarded. Currently limited to a single address.
	DestinationAddresses []string `mandatory:"true" json:"destinationAddresses"`

	// Case-insensitive name of an endpoint, that is a sub-resource of the resolver, to use as the forwarding
	// interface. The endpoint must have isForwarding set to true.
	SourceEndpointName *string `mandatory:"false" json:"sourceEndpointName"`
}

//GetClientAddressConditions returns ClientAddressConditions
func (m ResolverForwardRule) GetClientAddressConditions() []string {
	return m.ClientAddressConditions
}

//GetQnameCoverConditions returns QnameCoverConditions
func (m ResolverForwardRule) GetQnameCoverConditions() []string {
	return m.QnameCoverConditions
}

func (m ResolverForwardRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ResolverForwardRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeResolverForwardRule ResolverForwardRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeResolverForwardRule
	}{
		"FORWARD",
		(MarshalTypeResolverForwardRule)(m),
	}

	return json.Marshal(&s)
}
