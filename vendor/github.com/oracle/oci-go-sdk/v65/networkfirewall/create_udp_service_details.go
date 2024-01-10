// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateUdpServiceDetails Request for UDP Service used on the firewall policy rules.
type CreateUdpServiceDetails struct {

	// Name of the service
	Name *string `mandatory:"true" json:"name"`

	// List of port-ranges to be used.
	PortRanges []PortRange `mandatory:"true" json:"portRanges"`
}

// GetName returns Name
func (m CreateUdpServiceDetails) GetName() *string {
	return m.Name
}

func (m CreateUdpServiceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateUdpServiceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateUdpServiceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateUdpServiceDetails CreateUdpServiceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateUdpServiceDetails
	}{
		"UDP_SERVICE",
		(MarshalTypeCreateUdpServiceDetails)(m),
	}

	return json.Marshal(&s)
}
