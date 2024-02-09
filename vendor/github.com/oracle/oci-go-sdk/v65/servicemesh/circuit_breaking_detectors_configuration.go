// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CircuitBreakingDetectorsConfiguration Configuration for detectors of circuit breaker.
type CircuitBreakingDetectorsConfiguration struct {

	// Count of errors with status code 5xx and locally originated errors.
	// In Split Mode, just errors with status code 5xx.
	// For TCP protocol it maps to TCP connection failures.
	TotalErrors *int `mandatory:"false" json:"totalErrors"`

	// Count of totalErrors related to gateway errors (502, 503 or 504 status code).
	// For TCP protocol it maps to TCP reset etc.
	GatewayErrors *int `mandatory:"false" json:"gatewayErrors"`

	// Taken into account only in Split Mode, number of locally originated errors.
	// FOR TCP it maps tcp reset, ICMP errors etc.
	LocalErrors *int `mandatory:"false" json:"localErrors"`
}

func (m CircuitBreakingDetectorsConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CircuitBreakingDetectorsConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
