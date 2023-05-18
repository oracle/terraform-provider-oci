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

// LearnInternalPrivateIpDetails Details to attach a private IP to a VNIC
type LearnInternalPrivateIpDetails struct {

	// Unique identifier of a VNIC to map the private IP to
	VnicId *string `mandatory:"true" json:"vnicId"`

	// Unique identifier of the Subnet in which the private IP Address exists
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The private IP address that needs to be mapped to the VNIC. This IP address will become one of the secondary
	// private IPs on the VNIC.
	// Example: `129.146.2.1`
	IpAddress *string `mandatory:"true" json:"ipAddress"`
}

func (m LearnInternalPrivateIpDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LearnInternalPrivateIpDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
