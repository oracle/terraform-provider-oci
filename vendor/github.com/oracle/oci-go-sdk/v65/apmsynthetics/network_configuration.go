// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkConfiguration Details of the network configuration. For NETWORK monitor type, NetworkConfiguration is mandatory.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProbeModeEnum(string(m.ProbeMode)); !ok && m.ProbeMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProbeMode: %s. Supported values are: %s.", m.ProbeMode, strings.Join(GetProbeModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
