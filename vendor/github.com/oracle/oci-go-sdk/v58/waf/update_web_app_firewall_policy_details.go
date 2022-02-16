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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateWebAppFirewallPolicyDetails The information to be updated.
// When updating WebAppFirewallPolicy, shallow merge is used for all top-level fields,
// meaning that top-level fields with defined values are completely overwritten and
// top-level fields without defined values are unchanged.
type UpdateWebAppFirewallPolicyDetails struct {

	// WebAppFirewallPolicy display name, can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Predefined actions for use in multiple different rules. Not all actions are supported in every module.
	// Some actions terminate further execution of modules and rules in a module and some do not.
	// Actions names must be unique within this array.
	Actions []Action `mandatory:"false" json:"actions"`

	RequestAccessControl *RequestAccessControl `mandatory:"false" json:"requestAccessControl"`

	RequestRateLimiting *RequestRateLimiting `mandatory:"false" json:"requestRateLimiting"`

	RequestProtection *RequestProtection `mandatory:"false" json:"requestProtection"`

	ResponseAccessControl *ResponseAccessControl `mandatory:"false" json:"responseAccessControl"`

	ResponseProtection *ResponseProtection `mandatory:"false" json:"responseProtection"`

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

func (m UpdateWebAppFirewallPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateWebAppFirewallPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateWebAppFirewallPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                           `json:"displayName"`
		Actions               []action                          `json:"actions"`
		RequestAccessControl  *RequestAccessControl             `json:"requestAccessControl"`
		RequestRateLimiting   *RequestRateLimiting              `json:"requestRateLimiting"`
		RequestProtection     *RequestProtection                `json:"requestProtection"`
		ResponseAccessControl *ResponseAccessControl            `json:"responseAccessControl"`
		ResponseProtection    *ResponseProtection               `json:"responseProtection"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		SystemTags            map[string]map[string]interface{} `json:"systemTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Actions = make([]Action, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(Action)
		} else {
			m.Actions[i] = nil
		}
	}

	m.RequestAccessControl = model.RequestAccessControl

	m.RequestRateLimiting = model.RequestRateLimiting

	m.RequestProtection = model.RequestProtection

	m.ResponseAccessControl = model.ResponseAccessControl

	m.ResponseProtection = model.ResponseProtection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	return
}
