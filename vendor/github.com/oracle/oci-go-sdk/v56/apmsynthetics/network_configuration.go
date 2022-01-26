// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// NetworkConfiguration Details of the network configuration.
type NetworkConfiguration struct {

	// Number of hops.
	NumberOfHops *int `mandatory:"false" json:"numberOfHops"`

	// Number of probes per hop.
	ProbePerHop *int `mandatory:"false" json:"probePerHop"`

	// Number of probe packets sent out simultaneously.
	TransmissionRate *int `mandatory:"false" json:"transmissionRate"`

	// Type of protocol.
	Protocol ProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	// Type of probe mode when TCP protocol is selected.
	ProbeMode ProbeModeEnum `mandatory:"false" json:"probeMode,omitempty"`
}

func (m NetworkConfiguration) String() string {
	return common.PointerString(m)
}
