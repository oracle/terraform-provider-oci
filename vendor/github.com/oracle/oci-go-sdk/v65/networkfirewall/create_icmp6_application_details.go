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

// CreateIcmp6ApplicationDetails Request for ICMP6 Application used on the firewall policy rules.
type CreateIcmp6ApplicationDetails struct {

	// Name of the application
	Name *string `mandatory:"true" json:"name"`

	// The value of the ICMP6 message Type field as defined by RFC 4443 (https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
	IcmpType *int `mandatory:"true" json:"icmpType"`

	// The value of the ICMP6 message Code (subtype) field as defined by RFC 4443 (https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
	IcmpCode *int `mandatory:"false" json:"icmpCode"`
}

// GetName returns Name
func (m CreateIcmp6ApplicationDetails) GetName() *string {
	return m.Name
}

func (m CreateIcmp6ApplicationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateIcmp6ApplicationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateIcmp6ApplicationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateIcmp6ApplicationDetails CreateIcmp6ApplicationDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateIcmp6ApplicationDetails
	}{
		"ICMP_V6",
		(MarshalTypeCreateIcmp6ApplicationDetails)(m),
	}

	return json.Marshal(&s)
}
