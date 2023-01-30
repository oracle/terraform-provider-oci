// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// TcpApplication TCP Application used on the firewall policy rules.
type TcpApplication struct {

	// The minimum port in the range (inclusive), or the sole port of a single-port range.
	MinimumPort *int `mandatory:"true" json:"minimumPort"`

	// The maximum port in the range (inclusive), which may be absent for a single-port range.
	MaximumPort *int `mandatory:"false" json:"maximumPort"`
}

func (m TcpApplication) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TcpApplication) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TcpApplication) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTcpApplication TcpApplication
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTcpApplication
	}{
		"TCP",
		(MarshalTypeTcpApplication)(m),
	}

	return json.Marshal(&s)
}
