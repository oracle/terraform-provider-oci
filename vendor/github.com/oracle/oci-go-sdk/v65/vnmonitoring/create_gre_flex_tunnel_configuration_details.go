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

// CreateGreFlexTunnelConfigurationDetails The gre tunnel configuration create details.
type CreateGreFlexTunnelConfigurationDetails struct {

	// IP address of your end of the tunnel.
	CustomerTunnelIp *string `mandatory:"true" json:"customerTunnelIp"`

	// IP address of the oracle end of the tunnel.
	OracleTunnelIp *string `mandatory:"true" json:"oracleTunnelIp"`

	BgpSession *CreateFlexTunnelBgpSessionDetails `mandatory:"true" json:"bgpSession"`
}

// GetCustomerTunnelIp returns CustomerTunnelIp
func (m CreateGreFlexTunnelConfigurationDetails) GetCustomerTunnelIp() *string {
	return m.CustomerTunnelIp
}

// GetOracleTunnelIp returns OracleTunnelIp
func (m CreateGreFlexTunnelConfigurationDetails) GetOracleTunnelIp() *string {
	return m.OracleTunnelIp
}

// GetBgpSession returns BgpSession
func (m CreateGreFlexTunnelConfigurationDetails) GetBgpSession() *CreateFlexTunnelBgpSessionDetails {
	return m.BgpSession
}

func (m CreateGreFlexTunnelConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGreFlexTunnelConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateGreFlexTunnelConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateGreFlexTunnelConfigurationDetails CreateGreFlexTunnelConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"tunnelType"`
		MarshalTypeCreateGreFlexTunnelConfigurationDetails
	}{
		"GRE",
		(MarshalTypeCreateGreFlexTunnelConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
