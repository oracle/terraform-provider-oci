// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration (WAA) API
//
// API for the Web Application Acceleration service.
// Use this API to manage regional Web App Acceleration policies such as Caching and Compression
// for accelerating HTTP services.
//

package waa

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WebAppAccelerationSummary Summary of the WebAppAcceleration.
type WebAppAccelerationSummary interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebAppAcceleration.
	GetId() *string

	// WebAppAcceleration display name, can be renamed.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of WebAppAccelerationPolicy, which is attached to the resource.
	GetWebAppAccelerationPolicyId() *string

	// The time the WebAppAcceleration was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the WebAppAcceleration.
	GetLifecycleState() WebAppAccelerationLifecycleStateEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the WebAppAcceleration was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in FAILED state.
	GetLifecycleDetails() *string
}

type webappaccelerationsummary struct {
	JsonData                   []byte
	TimeUpdated                *common.SDKTime                      `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails           *string                              `mandatory:"false" json:"lifecycleDetails"`
	Id                         *string                              `mandatory:"true" json:"id"`
	DisplayName                *string                              `mandatory:"true" json:"displayName"`
	CompartmentId              *string                              `mandatory:"true" json:"compartmentId"`
	WebAppAccelerationPolicyId *string                              `mandatory:"true" json:"webAppAccelerationPolicyId"`
	TimeCreated                *common.SDKTime                      `mandatory:"true" json:"timeCreated"`
	LifecycleState             WebAppAccelerationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	FreeformTags               map[string]string                    `mandatory:"true" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{}    `mandatory:"true" json:"definedTags"`
	SystemTags                 map[string]map[string]interface{}    `mandatory:"true" json:"systemTags"`
	BackendType                string                               `json:"backendType"`
}

// UnmarshalJSON unmarshals json
func (m *webappaccelerationsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerwebappaccelerationsummary webappaccelerationsummary
	s := struct {
		Model Unmarshalerwebappaccelerationsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.WebAppAccelerationPolicyId = s.Model.WebAppAccelerationPolicyId
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
func (m *webappaccelerationsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BackendType {
	case "LOAD_BALANCER":
		mm := WebAppAccelerationLoadBalancerSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for WebAppAccelerationSummary: %s.", m.BackendType)
		return *m, nil
	}
}

// GetTimeUpdated returns TimeUpdated
func (m webappaccelerationsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m webappaccelerationsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m webappaccelerationsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m webappaccelerationsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m webappaccelerationsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetWebAppAccelerationPolicyId returns WebAppAccelerationPolicyId
func (m webappaccelerationsummary) GetWebAppAccelerationPolicyId() *string {
	return m.WebAppAccelerationPolicyId
}

// GetTimeCreated returns TimeCreated
func (m webappaccelerationsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m webappaccelerationsummary) GetLifecycleState() WebAppAccelerationLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m webappaccelerationsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m webappaccelerationsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m webappaccelerationsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m webappaccelerationsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m webappaccelerationsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWebAppAccelerationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWebAppAccelerationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
