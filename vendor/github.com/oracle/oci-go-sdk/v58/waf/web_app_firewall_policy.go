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

// WebAppFirewallPolicy The details of WebAppFirewallPolicy. A policy is comprised of rules, which allows executing inspections of
// incoming/outgoing HTTP message parameters and execution of actions, based on results of rules execution.
// In policy, rules are grouped into modules by their functionality. Modules can be further divided by the type
// of HTTP messages they handle:
//   Modules that inspect incoming HTTP request. These modules are executed in the order they are enumerated here:
//     * requestAccessControl
//     * requestRateLimiting
//     * requestProtection
//  Modules that inspect outgoing HTTP responses. These modules are executed in the order they are enumerated here:
//    * responseAccessControl
//    * responseProtection
type WebAppFirewallPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WebAppFirewallPolicy.
	Id *string `mandatory:"true" json:"id"`

	// WebAppFirewallPolicy display name, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the WebAppFirewallPolicy was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the WebAppFirewallPolicy.
	LifecycleState WebAppFirewallPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The time the WebAppFirewallPolicy was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in FAILED state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Predefined actions for use in multiple different rules. Not all actions are supported in every module.
	// Some actions terminate further execution of modules and rules in a module and some do not.
	// Actions names must be unique within this array.
	Actions []Action `mandatory:"false" json:"actions"`

	RequestAccessControl *RequestAccessControl `mandatory:"false" json:"requestAccessControl"`

	RequestRateLimiting *RequestRateLimiting `mandatory:"false" json:"requestRateLimiting"`

	RequestProtection *RequestProtection `mandatory:"false" json:"requestProtection"`

	ResponseAccessControl *ResponseAccessControl `mandatory:"false" json:"responseAccessControl"`

	ResponseProtection *ResponseProtection `mandatory:"false" json:"responseProtection"`
}

func (m WebAppFirewallPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WebAppFirewallPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWebAppFirewallPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWebAppFirewallPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *WebAppFirewallPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeUpdated           *common.SDKTime                        `json:"timeUpdated"`
		LifecycleDetails      *string                                `json:"lifecycleDetails"`
		Actions               []action                               `json:"actions"`
		RequestAccessControl  *RequestAccessControl                  `json:"requestAccessControl"`
		RequestRateLimiting   *RequestRateLimiting                   `json:"requestRateLimiting"`
		RequestProtection     *RequestProtection                     `json:"requestProtection"`
		ResponseAccessControl *ResponseAccessControl                 `json:"responseAccessControl"`
		ResponseProtection    *ResponseProtection                    `json:"responseProtection"`
		Id                    *string                                `json:"id"`
		DisplayName           *string                                `json:"displayName"`
		CompartmentId         *string                                `json:"compartmentId"`
		TimeCreated           *common.SDKTime                        `json:"timeCreated"`
		LifecycleState        WebAppFirewallPolicyLifecycleStateEnum `json:"lifecycleState"`
		FreeformTags          map[string]string                      `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}      `json:"definedTags"`
		SystemTags            map[string]map[string]interface{}      `json:"systemTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

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

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	return
}

// WebAppFirewallPolicyLifecycleStateEnum Enum with underlying type: string
type WebAppFirewallPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for WebAppFirewallPolicyLifecycleStateEnum
const (
	WebAppFirewallPolicyLifecycleStateCreating WebAppFirewallPolicyLifecycleStateEnum = "CREATING"
	WebAppFirewallPolicyLifecycleStateUpdating WebAppFirewallPolicyLifecycleStateEnum = "UPDATING"
	WebAppFirewallPolicyLifecycleStateActive   WebAppFirewallPolicyLifecycleStateEnum = "ACTIVE"
	WebAppFirewallPolicyLifecycleStateDeleting WebAppFirewallPolicyLifecycleStateEnum = "DELETING"
	WebAppFirewallPolicyLifecycleStateDeleted  WebAppFirewallPolicyLifecycleStateEnum = "DELETED"
	WebAppFirewallPolicyLifecycleStateFailed   WebAppFirewallPolicyLifecycleStateEnum = "FAILED"
)

var mappingWebAppFirewallPolicyLifecycleStateEnum = map[string]WebAppFirewallPolicyLifecycleStateEnum{
	"CREATING": WebAppFirewallPolicyLifecycleStateCreating,
	"UPDATING": WebAppFirewallPolicyLifecycleStateUpdating,
	"ACTIVE":   WebAppFirewallPolicyLifecycleStateActive,
	"DELETING": WebAppFirewallPolicyLifecycleStateDeleting,
	"DELETED":  WebAppFirewallPolicyLifecycleStateDeleted,
	"FAILED":   WebAppFirewallPolicyLifecycleStateFailed,
}

// GetWebAppFirewallPolicyLifecycleStateEnumValues Enumerates the set of values for WebAppFirewallPolicyLifecycleStateEnum
func GetWebAppFirewallPolicyLifecycleStateEnumValues() []WebAppFirewallPolicyLifecycleStateEnum {
	values := make([]WebAppFirewallPolicyLifecycleStateEnum, 0)
	for _, v := range mappingWebAppFirewallPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetWebAppFirewallPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for WebAppFirewallPolicyLifecycleStateEnum
func GetWebAppFirewallPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingWebAppFirewallPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWebAppFirewallPolicyLifecycleStateEnum(val string) (WebAppFirewallPolicyLifecycleStateEnum, bool) {
	mappingWebAppFirewallPolicyLifecycleStateEnumIgnoreCase := make(map[string]WebAppFirewallPolicyLifecycleStateEnum)
	for k, v := range mappingWebAppFirewallPolicyLifecycleStateEnum {
		mappingWebAppFirewallPolicyLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWebAppFirewallPolicyLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
