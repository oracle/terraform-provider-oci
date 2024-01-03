// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateNetworkAddressListVcnAddressesDetails The information to be updated for NetworkAddressListVcnAddresses.
type UpdateNetworkAddressListVcnAddressesDetails struct {

	// NetworkAddressList display name, can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A list of private address prefixes, each associated with a particular VCN.
	// To specify all addresses in a VCN, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6.
	VcnAddresses []PrivateAddresses `mandatory:"false" json:"vcnAddresses"`
}

// GetDisplayName returns DisplayName
func (m UpdateNetworkAddressListVcnAddressesDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m UpdateNetworkAddressListVcnAddressesDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateNetworkAddressListVcnAddressesDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m UpdateNetworkAddressListVcnAddressesDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m UpdateNetworkAddressListVcnAddressesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateNetworkAddressListVcnAddressesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateNetworkAddressListVcnAddressesDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateNetworkAddressListVcnAddressesDetails UpdateNetworkAddressListVcnAddressesDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateNetworkAddressListVcnAddressesDetails
	}{
		"VCN_ADDRESSES",
		(MarshalTypeUpdateNetworkAddressListVcnAddressesDetails)(m),
	}

	return json.Marshal(&s)
}
