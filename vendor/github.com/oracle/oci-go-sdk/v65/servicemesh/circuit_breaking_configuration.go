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

// CircuitBreakingConfiguration This configuration determines when to open and close the circuits to hosts.
type CircuitBreakingConfiguration struct {

	// Time interval between ejection sweep analysis.
	EjectionIntervalInSeconds *int `mandatory:"false" json:"ejectionIntervalInSeconds"`

	// A host will remain ejected for a period equal to the product of minimum ejection duration and the
	// number of times the host has been ejected.
	// This technique allows the system to automatically increase the ejection period for unhealthy upstream servers.
	MinEjectionTimeInSeconds *int `mandatory:"false" json:"minEjectionTimeInSeconds"`

	// The maximum time that a host is ejected for.
	// If not specified, the default value 300s or minEjectionTime value is applied, whatever is larger.
	MaxEjectionTimeInSeconds *int `mandatory:"false" json:"maxEjectionTimeInSeconds"`

	// Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.
	MaxEjectionPercent *int `mandatory:"false" json:"maxEjectionPercent"`

	// Determines whether to distinguish local origin failures from external errors.
	// If set to true consecutive_local_origin_failure is taken into account for outlier detection calculations.
	IsSplitEnabled *bool `mandatory:"false" json:"isSplitEnabled"`

	Detectors *CircuitBreakingDetectorsConfiguration `mandatory:"false" json:"detectors"`

	Threshold CircuitBreakingThresholdConfiguration `mandatory:"false" json:"threshold"`
}

func (m CircuitBreakingConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CircuitBreakingConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CircuitBreakingConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		EjectionIntervalInSeconds *int                                   `json:"ejectionIntervalInSeconds"`
		MinEjectionTimeInSeconds  *int                                   `json:"minEjectionTimeInSeconds"`
		MaxEjectionTimeInSeconds  *int                                   `json:"maxEjectionTimeInSeconds"`
		MaxEjectionPercent        *int                                   `json:"maxEjectionPercent"`
		IsSplitEnabled            *bool                                  `json:"isSplitEnabled"`
		Detectors                 *CircuitBreakingDetectorsConfiguration `json:"detectors"`
		Threshold                 circuitbreakingthresholdconfiguration  `json:"threshold"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.EjectionIntervalInSeconds = model.EjectionIntervalInSeconds

	m.MinEjectionTimeInSeconds = model.MinEjectionTimeInSeconds

	m.MaxEjectionTimeInSeconds = model.MaxEjectionTimeInSeconds

	m.MaxEjectionPercent = model.MaxEjectionPercent

	m.IsSplitEnabled = model.IsSplitEnabled

	m.Detectors = model.Detectors

	nn, e = model.Threshold.UnmarshalPolymorphicJSON(model.Threshold.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Threshold = nn.(CircuitBreakingThresholdConfiguration)
	} else {
		m.Threshold = nil
	}

	return
}
