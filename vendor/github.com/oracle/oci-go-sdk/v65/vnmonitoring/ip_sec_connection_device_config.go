// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpSecConnectionDeviceConfig Deprecated. For tunnel information, instead see:
//   - IPSecConnectionTunnel
//   - IPSecConnectionTunnelSharedSecret
type IpSecConnectionDeviceConfig struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the IPSec connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The IPSec connection's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// The date and time the IPSec connection was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Two TunnelConfig objects.
	Tunnels []TunnelConfig `mandatory:"false" json:"tunnels"`
}

func (m IpSecConnectionDeviceConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpSecConnectionDeviceConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
