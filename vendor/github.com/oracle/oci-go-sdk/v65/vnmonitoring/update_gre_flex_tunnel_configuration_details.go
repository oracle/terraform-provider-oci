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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateGreFlexTunnelConfigurationDetails The gre tunnel configuration update details.
type UpdateGreFlexTunnelConfigurationDetails struct {

	// IP address of your end of the tunnel.
	CustomerTunnelIp *string `mandatory:"false" json:"customerTunnelIp"`

	// IP address of the oracle end of the tunnel.
	OracleTunnelIp *string `mandatory:"false" json:"oracleTunnelIp"`

	BgpSession *UpdateFlexTunnelBgpSessionDetails `mandatory:"false" json:"bgpSession"`
}

// GetCustomerTunnelIp returns CustomerTunnelIp
func (m UpdateGreFlexTunnelConfigurationDetails) GetCustomerTunnelIp() *string {
	return m.CustomerTunnelIp
}

// GetOracleTunnelIp returns OracleTunnelIp
func (m UpdateGreFlexTunnelConfigurationDetails) GetOracleTunnelIp() *string {
	return m.OracleTunnelIp
}

// GetBgpSession returns BgpSession
func (m UpdateGreFlexTunnelConfigurationDetails) GetBgpSession() *UpdateFlexTunnelBgpSessionDetails {
	return m.BgpSession
}

func (m UpdateGreFlexTunnelConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateGreFlexTunnelConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateGreFlexTunnelConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateGreFlexTunnelConfigurationDetails UpdateGreFlexTunnelConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"tunnelType"`
		MarshalTypeUpdateGreFlexTunnelConfigurationDetails
	}{
		"GRE",
		(MarshalTypeUpdateGreFlexTunnelConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
