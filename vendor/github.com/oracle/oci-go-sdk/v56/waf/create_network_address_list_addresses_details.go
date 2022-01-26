// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateNetworkAddressListAddressesDetails The information about new NetworkAddressListAddresses.
type CreateNetworkAddressListAddressesDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A list of IP address prefixes in CIDR notation.
	// To specify all addresses, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6.
	Addresses []string `mandatory:"true" json:"addresses"`

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
}

//GetDisplayName returns DisplayName
func (m CreateNetworkAddressListAddressesDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateNetworkAddressListAddressesDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m CreateNetworkAddressListAddressesDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateNetworkAddressListAddressesDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m CreateNetworkAddressListAddressesDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m CreateNetworkAddressListAddressesDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateNetworkAddressListAddressesDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateNetworkAddressListAddressesDetails CreateNetworkAddressListAddressesDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateNetworkAddressListAddressesDetails
	}{
		"ADDRESSES",
		(MarshalTypeCreateNetworkAddressListAddressesDetails)(m),
	}

	return json.Marshal(&s)
}
