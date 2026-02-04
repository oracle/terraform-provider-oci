// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NatConfigurationRequest Request to configure Network Address Translation (NAT) on a firewall.
// To perform NAT on traffic passing the private NAT IPs to the firewall, the attached network firewall policy must also have NAT rules and NAT configuration must be enabled. If NAT configuration is enabled and the attached firewall policy does not contain NAT rule then NAT IPs will get allocated but NAT will not be performed on any traffic.
type NatConfigurationRequest struct {

	// The value of this field must be set to true if the network firewall policy being applied contains NAT rules.
	// The value of this field can be set to false if the network firewall policy being applied or the currently attached firewall policy doesn't contain NAT rules.
	MustEnablePrivateNat *bool `mandatory:"true" json:"mustEnablePrivateNat"`
}

func (m NatConfigurationRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NatConfigurationRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
