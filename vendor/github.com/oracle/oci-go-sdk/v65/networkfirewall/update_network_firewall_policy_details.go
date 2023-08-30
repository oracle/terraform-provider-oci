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

// UpdateNetworkFirewallPolicyDetails The request details to be updated in the firewall policy.
type UpdateNetworkFirewallPolicyDetails struct {

	// A user-friendly name for the firewall. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Map defining secrets of the policy.
	// The value of an entry is a "mapped secret" consisting of a purpose and source.
	// The associated key is the identifier by which the mapped secret is referenced.
	MappedSecrets map[string]MappedSecret `mandatory:"false" json:"mappedSecrets"`

	// Map defining application lists of the policy.
	// The value of an entry is a list of "applications", each consisting of a protocol identifier (such as TCP, UDP, or ICMP) and protocol-specific parameters (such as a port range).
	// The associated key is the identifier by which the application list is referenced.
	ApplicationLists map[string][]Application `mandatory:"false" json:"applicationLists"`

	// Map defining URL pattern lists of the policy.
	// The value of an entry is a list of URL patterns.
	// The associated key is the identifier by which the URL pattern list is referenced.
	UrlLists map[string][]UrlPattern `mandatory:"false" json:"urlLists"`

	// Map defining IP address lists of the policy.
	// The value of an entry is a list of IP addresses or prefixes in CIDR notation.
	// The associated key is the identifier by which the IP address list is referenced.
	IpAddressLists map[string][]string `mandatory:"false" json:"ipAddressLists"`

	// List of Security Rules defining the behavior of the policy.
	// The first rule with a matching condition determines the action taken upon network traffic.
	SecurityRules []SecurityRule `mandatory:"false" json:"securityRules"`

	// List of Decryption Rules defining the behavior of the policy.
	// The first rule with a matching condition determines the action taken upon network traffic.
	DecryptionRules []DecryptionRule `mandatory:"false" json:"decryptionRules"`

	// Map defining decryption profiles of the policy.
	// The value of an entry is a decryption profile.
	// The associated key is the identifier by which the decryption profile is referenced.
	DecryptionProfiles map[string]DecryptionProfile `mandatory:"false" json:"decryptionProfiles"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateNetworkFirewallPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateNetworkFirewallPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateNetworkFirewallPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                           `json:"displayName"`
		MappedSecrets      map[string]mappedsecret           `json:"mappedSecrets"`
		ApplicationLists   map[string][]Application          `json:"applicationLists"`
		UrlLists           map[string][]UrlPattern           `json:"urlLists"`
		IpAddressLists     map[string][]string               `json:"ipAddressLists"`
		SecurityRules      []SecurityRule                    `json:"securityRules"`
		DecryptionRules    []DecryptionRule                  `json:"decryptionRules"`
		DecryptionProfiles map[string]decryptionprofile      `json:"decryptionProfiles"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.MappedSecrets = make(map[string]MappedSecret)
	for k, v := range model.MappedSecrets {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.MappedSecrets[k] = nn.(MappedSecret)
		} else {
			m.MappedSecrets[k] = nil
		}
	}

	m.ApplicationLists = model.ApplicationLists

	m.UrlLists = model.UrlLists

	m.IpAddressLists = model.IpAddressLists

	m.SecurityRules = make([]SecurityRule, len(model.SecurityRules))
	copy(m.SecurityRules, model.SecurityRules)
	m.DecryptionRules = make([]DecryptionRule, len(model.DecryptionRules))
	copy(m.DecryptionRules, model.DecryptionRules)
	m.DecryptionProfiles = make(map[string]DecryptionProfile)
	for k, v := range model.DecryptionProfiles {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DecryptionProfiles[k] = nn.(DecryptionProfile)
		} else {
			m.DecryptionProfiles[k] = nil
		}
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
