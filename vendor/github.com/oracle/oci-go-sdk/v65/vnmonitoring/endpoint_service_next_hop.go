// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EndpointServiceNextHop Information of a particular service's next hop
type EndpointServiceNextHop struct {

	// An IP address that handles requests to the endpoint service.
	ServiceIp *string `mandatory:"true" json:"serviceIp"`

	// An Internal IP address that handles requests to the substrate anycast of endpoint service.
	NextHopIp *string `mandatory:"true" json:"nextHopIp"`

	// MPLS label that identifies the substrate endpoint service
	NextHopSlotId *int `mandatory:"true" json:"nextHopSlotId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint service.
	EndpointServiceId *string `mandatory:"false" json:"endpointServiceId"`
}

func (m EndpointServiceNextHop) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EndpointServiceNextHop) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
