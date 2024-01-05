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

// WebAppFirewall A resource connecting a WebAppFirewallPolicy to a backend of particular type, applying that policy's coverage to the backend.
type WebAppFirewall interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WebAppFirewall.
	GetId() *string

	// WebAppFirewall display name, can be renamed.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of WebAppFirewallPolicy, which is attached to the resource.
	GetWebAppFirewallPolicyId() *string

	// The time the WebAppFirewall was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the WebAppFirewall.
	GetLifecycleState() WebAppFirewallLifecycleStateEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the WebAppFirewall was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in FAILED state.
	GetLifecycleDetails() *string
}

type webappfirewall struct {
	JsonData               []byte
	TimeUpdated            *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails       *string                           `mandatory:"false" json:"lifecycleDetails"`
	Id                     *string                           `mandatory:"true" json:"id"`
	DisplayName            *string                           `mandatory:"true" json:"displayName"`
	CompartmentId          *string                           `mandatory:"true" json:"compartmentId"`
	WebAppFirewallPolicyId *string                           `mandatory:"true" json:"webAppFirewallPolicyId"`
	TimeCreated            *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState         WebAppFirewallLifecycleStateEnum  `mandatory:"true" json:"lifecycleState"`
	FreeformTags           map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags            map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	SystemTags             map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`
	BackendType            string                            `json:"backendType"`
}

// UnmarshalJSON unmarshals json
func (m *webappfirewall) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerwebappfirewall webappfirewall
	s := struct {
		Model Unmarshalerwebappfirewall
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.WebAppFirewallPolicyId = s.Model.WebAppFirewallPolicyId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.BackendType = s.Model.BackendType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *webappfirewall) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BackendType {
	case "LOAD_BALANCER":
		mm := WebAppFirewallLoadBalancer{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for WebAppFirewall: %s.", m.BackendType)
		return *m, nil
	}
}

// GetTimeUpdated returns TimeUpdated
func (m webappfirewall) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m webappfirewall) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m webappfirewall) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m webappfirewall) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m webappfirewall) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetWebAppFirewallPolicyId returns WebAppFirewallPolicyId
func (m webappfirewall) GetWebAppFirewallPolicyId() *string {
	return m.WebAppFirewallPolicyId
}

// GetTimeCreated returns TimeCreated
func (m webappfirewall) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m webappfirewall) GetLifecycleState() WebAppFirewallLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m webappfirewall) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m webappfirewall) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m webappfirewall) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m webappfirewall) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m webappfirewall) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWebAppFirewallLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWebAppFirewallLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WebAppFirewallLifecycleStateEnum Enum with underlying type: string
type WebAppFirewallLifecycleStateEnum string

// Set of constants representing the allowable values for WebAppFirewallLifecycleStateEnum
const (
	WebAppFirewallLifecycleStateCreating WebAppFirewallLifecycleStateEnum = "CREATING"
	WebAppFirewallLifecycleStateUpdating WebAppFirewallLifecycleStateEnum = "UPDATING"
	WebAppFirewallLifecycleStateActive   WebAppFirewallLifecycleStateEnum = "ACTIVE"
	WebAppFirewallLifecycleStateDeleting WebAppFirewallLifecycleStateEnum = "DELETING"
	WebAppFirewallLifecycleStateDeleted  WebAppFirewallLifecycleStateEnum = "DELETED"
	WebAppFirewallLifecycleStateFailed   WebAppFirewallLifecycleStateEnum = "FAILED"
)

var mappingWebAppFirewallLifecycleStateEnum = map[string]WebAppFirewallLifecycleStateEnum{
	"CREATING": WebAppFirewallLifecycleStateCreating,
	"UPDATING": WebAppFirewallLifecycleStateUpdating,
	"ACTIVE":   WebAppFirewallLifecycleStateActive,
	"DELETING": WebAppFirewallLifecycleStateDeleting,
	"DELETED":  WebAppFirewallLifecycleStateDeleted,
	"FAILED":   WebAppFirewallLifecycleStateFailed,
}

var mappingWebAppFirewallLifecycleStateEnumLowerCase = map[string]WebAppFirewallLifecycleStateEnum{
	"creating": WebAppFirewallLifecycleStateCreating,
	"updating": WebAppFirewallLifecycleStateUpdating,
	"active":   WebAppFirewallLifecycleStateActive,
	"deleting": WebAppFirewallLifecycleStateDeleting,
	"deleted":  WebAppFirewallLifecycleStateDeleted,
	"failed":   WebAppFirewallLifecycleStateFailed,
}

// GetWebAppFirewallLifecycleStateEnumValues Enumerates the set of values for WebAppFirewallLifecycleStateEnum
func GetWebAppFirewallLifecycleStateEnumValues() []WebAppFirewallLifecycleStateEnum {
	values := make([]WebAppFirewallLifecycleStateEnum, 0)
	for _, v := range mappingWebAppFirewallLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetWebAppFirewallLifecycleStateEnumStringValues Enumerates the set of values in String for WebAppFirewallLifecycleStateEnum
func GetWebAppFirewallLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingWebAppFirewallLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWebAppFirewallLifecycleStateEnum(val string) (WebAppFirewallLifecycleStateEnum, bool) {
	enum, ok := mappingWebAppFirewallLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
