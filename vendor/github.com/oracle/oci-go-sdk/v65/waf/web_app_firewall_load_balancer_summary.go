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

// WebAppFirewallLoadBalancerSummary Summary of the WebAppFirewallLoadBalancer.
type WebAppFirewallLoadBalancerSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WebAppFirewall.
	Id *string `mandatory:"true" json:"id"`

	// WebAppFirewall display name, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of WebAppFirewallPolicy, which is attached to the resource.
	WebAppFirewallPolicyId *string `mandatory:"true" json:"webAppFirewallPolicyId"`

	// The time the WebAppFirewall was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// LoadBalancer OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) to which the WebAppFirewallPolicy is attached to.
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

	// The time the WebAppFirewall was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in FAILED state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the WebAppFirewall.
	LifecycleState WebAppFirewallLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m WebAppFirewallLoadBalancerSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m WebAppFirewallLoadBalancerSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m WebAppFirewallLoadBalancerSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetWebAppFirewallPolicyId returns WebAppFirewallPolicyId
func (m WebAppFirewallLoadBalancerSummary) GetWebAppFirewallPolicyId() *string {
	return m.WebAppFirewallPolicyId
}

// GetTimeCreated returns TimeCreated
func (m WebAppFirewallLoadBalancerSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m WebAppFirewallLoadBalancerSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m WebAppFirewallLoadBalancerSummary) GetLifecycleState() WebAppFirewallLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m WebAppFirewallLoadBalancerSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m WebAppFirewallLoadBalancerSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m WebAppFirewallLoadBalancerSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m WebAppFirewallLoadBalancerSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m WebAppFirewallLoadBalancerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WebAppFirewallLoadBalancerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWebAppFirewallLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWebAppFirewallLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m WebAppFirewallLoadBalancerSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeWebAppFirewallLoadBalancerSummary WebAppFirewallLoadBalancerSummary
	s := struct {
		DiscriminatorParam string `json:"backendType"`
		MarshalTypeWebAppFirewallLoadBalancerSummary
	}{
		"LOAD_BALANCER",
		(MarshalTypeWebAppFirewallLoadBalancerSummary)(m),
	}

	return json.Marshal(&s)
}
