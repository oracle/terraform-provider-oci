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

// UpdateIcmpApplicationDetails Request for updating ICMP Application used on the firewall policy rules.
type UpdateIcmpApplicationDetails struct {

	// The value of the ICMP message Type field as defined by RFC 792 (https://www.rfc-editor.org/rfc/rfc792.html).
	IcmpType *int `mandatory:"true" json:"icmpType"`

	// The value of the ICMP message Code (subtype) field as defined by RFC 792 (https://www.rfc-editor.org/rfc/rfc792.html).
	IcmpCode *int `mandatory:"false" json:"icmpCode"`
}

func (m UpdateIcmpApplicationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateIcmpApplicationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateIcmpApplicationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateIcmpApplicationDetails UpdateIcmpApplicationDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateIcmpApplicationDetails
	}{
		"ICMP",
		(MarshalTypeUpdateIcmpApplicationDetails)(m),
	}

	return json.Marshal(&s)
}
