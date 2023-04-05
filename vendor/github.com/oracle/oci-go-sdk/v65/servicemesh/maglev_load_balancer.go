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

// MaglevLoadBalancer The Maglev load balancer implements consistent hashing to upstream hosts.
type MaglevLoadBalancer struct {

	// The table size for Maglev hashing. The table size must be prime number. Increasing the table size
	// reduces the amount of disruption, and a connection will likely be sent to the same upstream host.
	TableSize *int64 `mandatory:"false" json:"tableSize"`

	HashKey *HashKey `mandatory:"false" json:"hashKey"`
}

func (m MaglevLoadBalancer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaglevLoadBalancer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MaglevLoadBalancer) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMaglevLoadBalancer MaglevLoadBalancer
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeMaglevLoadBalancer
	}{
		"MAGLEV",
		(MarshalTypeMaglevLoadBalancer)(m),
	}

	return json.Marshal(&s)
}
