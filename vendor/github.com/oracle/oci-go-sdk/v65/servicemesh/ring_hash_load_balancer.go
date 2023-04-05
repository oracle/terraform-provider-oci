// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RingHashLoadBalancer The ring/modulo hash load balancer implements consistent hashing to upstream hosts. MinimumRingSize must not
// be greater than maximumRingSize.
type RingHashLoadBalancer struct {

	// Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided
	// host) the better the request distribution.
	MinimumRingSize *int64 `mandatory:"false" json:"minimumRingSize"`

	// Maximum hash ring size. Lower this number to constrain resource usage if applicable.
	MaximumRingSize *int64 `mandatory:"false" json:"maximumRingSize"`

	HashKey *HashKey `mandatory:"false" json:"hashKey"`
}

func (m RingHashLoadBalancer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RingHashLoadBalancer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RingHashLoadBalancer) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRingHashLoadBalancer RingHashLoadBalancer
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRingHashLoadBalancer
	}{
		"RING_HASH",
		(MarshalTypeRingHashLoadBalancer)(m),
	}

	return json.Marshal(&s)
}
