// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// NatConfigurationResponse Nat Configuration response.
type NatConfigurationResponse struct {

	// To allocate private NAT IPs to the firewall. The attached network firewall policy must also have NAT rules to enable NAT on any traffic passing through the firewall.
	MustEnablePrivateNat *bool `mandatory:"false" json:"mustEnablePrivateNat"`

	// An array of NAT IP addresses that are associated with the Network Firewall. These IPs are reserved for NAT and shouldn't be used for any other purpose in the subnet.
	NatIpAddressList []string `mandatory:"false" json:"natIpAddressList"`
}

func (m NatConfigurationResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NatConfigurationResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
