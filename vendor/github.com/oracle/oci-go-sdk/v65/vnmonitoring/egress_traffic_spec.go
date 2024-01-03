// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EgressTrafficSpec Defines the traffic configuration that leaves the traffic node.
type EgressTrafficSpec struct {

	// The IP protocol to use for the traffic path analysis.
	Protocol *int `mandatory:"true" json:"protocol"`

	// The IPv4 address of the source node.
	SourceAddress *string `mandatory:"true" json:"sourceAddress"`

	// The IPv4 address of the destination node.
	DestinationAddress *string `mandatory:"true" json:"destinationAddress"`

	TrafficProtocolParameters TrafficProtocolParameters `mandatory:"false" json:"trafficProtocolParameters"`
}

func (m EgressTrafficSpec) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EgressTrafficSpec) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *EgressTrafficSpec) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TrafficProtocolParameters trafficprotocolparameters `json:"trafficProtocolParameters"`
		Protocol                  *int                      `json:"protocol"`
		SourceAddress             *string                   `json:"sourceAddress"`
		DestinationAddress        *string                   `json:"destinationAddress"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.TrafficProtocolParameters.UnmarshalPolymorphicJSON(model.TrafficProtocolParameters.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrafficProtocolParameters = nn.(TrafficProtocolParameters)
	} else {
		m.TrafficProtocolParameters = nil
	}

	m.Protocol = model.Protocol

	m.SourceAddress = model.SourceAddress

	m.DestinationAddress = model.DestinationAddress

	return
}
