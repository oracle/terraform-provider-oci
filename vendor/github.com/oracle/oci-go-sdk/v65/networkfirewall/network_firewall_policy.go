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

// NetworkFirewallPolicy Description of NetworkFirewall Policy.
type NetworkFirewallPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource - Network Firewall Policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the NetworkFirewall Policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly optional name for the firewall policy. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time instant at which the Network Firewall Policy was created in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time instant at which the Network Firewall Policy was updated in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Network Firewall Policy.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// To determine if any Network Firewall is associated with this Network Firewall Policy.
	IsFirewallAttached *bool `mandatory:"true" json:"isFirewallAttached"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m NetworkFirewallPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkFirewallPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *NetworkFirewallPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LifecycleDetails   *string                           `json:"lifecycleDetails"`
		MappedSecrets      map[string]mappedsecret           `json:"mappedSecrets"`
		ApplicationLists   map[string][]Application          `json:"applicationLists"`
		UrlLists           map[string][]UrlPattern           `json:"urlLists"`
		IpAddressLists     map[string][]string               `json:"ipAddressLists"`
		SecurityRules      []SecurityRule                    `json:"securityRules"`
		DecryptionRules    []DecryptionRule                  `json:"decryptionRules"`
		DecryptionProfiles map[string]decryptionprofile      `json:"decryptionProfiles"`
		SystemTags         map[string]map[string]interface{} `json:"systemTags"`
		Id                 *string                           `json:"id"`
		CompartmentId      *string                           `json:"compartmentId"`
		DisplayName        *string                           `json:"displayName"`
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState     LifecycleStateEnum                `json:"lifecycleState"`
		IsFirewallAttached *bool                             `json:"isFirewallAttached"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LifecycleDetails = model.LifecycleDetails

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

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.IsFirewallAttached = model.IsFirewallAttached

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
