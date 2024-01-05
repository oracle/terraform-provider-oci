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

// CreateWebAppFirewallDetails The information about new Web App Firewall.
type CreateWebAppFirewallDetails interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of WebAppFirewallPolicy, which is attached to the resource.
	GetWebAppFirewallPolicyId() *string

	// WebAppFirewall display name, can be renamed.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type createwebappfirewalldetails struct {
	JsonData               []byte
	DisplayName            *string                           `mandatory:"false" json:"displayName"`
	FreeformTags           map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags            map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags             map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	CompartmentId          *string                           `mandatory:"true" json:"compartmentId"`
	WebAppFirewallPolicyId *string                           `mandatory:"true" json:"webAppFirewallPolicyId"`
	BackendType            string                            `json:"backendType"`
}

// UnmarshalJSON unmarshals json
func (m *createwebappfirewalldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatewebappfirewalldetails createwebappfirewalldetails
	s := struct {
		Model Unmarshalercreatewebappfirewalldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.WebAppFirewallPolicyId = s.Model.WebAppFirewallPolicyId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.BackendType = s.Model.BackendType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createwebappfirewalldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BackendType {
	case "LOAD_BALANCER":
		mm := CreateWebAppFirewallLoadBalancerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateWebAppFirewallDetails: %s.", m.BackendType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createwebappfirewalldetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m createwebappfirewalldetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createwebappfirewalldetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m createwebappfirewalldetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetCompartmentId returns CompartmentId
func (m createwebappfirewalldetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetWebAppFirewallPolicyId returns WebAppFirewallPolicyId
func (m createwebappfirewalldetails) GetWebAppFirewallPolicyId() *string {
	return m.WebAppFirewallPolicyId
}

func (m createwebappfirewalldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createwebappfirewalldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
