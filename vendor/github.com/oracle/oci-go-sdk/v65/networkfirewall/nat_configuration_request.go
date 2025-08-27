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

// NatConfigurationRequest Nat Configuration request to use Nat feature on firewall.
type NatConfigurationRequest struct {

	// To allocate private NAT IPs to the firewall. The attached network firewall policy must also have NAT rules to enable NAT on any traffic passing through the firewall. The value of this field can not be false to release the NAT IPs given that the attached network firewall policy does not contains any NAT rules. The value of this field should be set to true if the network firewall policy being applied contains NAT rules.
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
