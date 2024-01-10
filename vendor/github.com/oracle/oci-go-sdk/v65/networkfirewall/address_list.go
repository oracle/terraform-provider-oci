// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AddressList List of addresses with a reference name.
// The value of an entry is a list of IP addresses or prefixes in CIDR notation or FQDNs.
// The associated key is the identifier by which the IP address list is referenced.
type AddressList struct {

	// Unique name to identify the group of addresses to be used in the policy rules.
	Name *string `mandatory:"true" json:"name"`

	// Type of address List. The accepted values are - * FQDN * IP
	Type AddressListTypeEnum `mandatory:"true" json:"type"`

	// List of addresses.
	Addresses []string `mandatory:"true" json:"addresses"`

	// Count of total Addresses in the AddressList
	TotalAddresses *int `mandatory:"true" json:"totalAddresses"`

	// OCID of the Network Firewall Policy this Address List belongs to.
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`
}

func (m AddressList) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddressList) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddressListTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAddressListTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
