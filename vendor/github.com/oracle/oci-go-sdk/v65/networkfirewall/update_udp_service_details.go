// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateUdpServiceDetails Request for updating UDP Service used on the firewall policy rules.
type UpdateUdpServiceDetails struct {

	// List of port-ranges to be used.
	PortRanges []PortRange `mandatory:"true" json:"portRanges"`
}

func (m UpdateUdpServiceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateUdpServiceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateUdpServiceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateUdpServiceDetails UpdateUdpServiceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateUdpServiceDetails
	}{
		"UDP_SERVICE",
		(MarshalTypeUpdateUdpServiceDetails)(m),
	}

	return json.Marshal(&s)
}
