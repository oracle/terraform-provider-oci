// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DelayFaultConfiguration Delay requests before forwarding, emulating various failures such as network issues, overloaded upstream service, etc.
type DelayFaultConfiguration struct {

	// The duration during which the requests will be delayed before forwarding.
	DelayInMs *int64 `mandatory:"true" json:"delayInMs"`

	// Percentage of requests on which the delay will be injected.
	// Percentage values range from [0-100].
	// If left unspecified, all requests will be delayed.
	PercentageOfRequests *int `mandatory:"false" json:"percentageOfRequests"`
}

func (m DelayFaultConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DelayFaultConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
