// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateWebAppAccelerationDetails The information about new WebAppAcceleration.
type CreateWebAppAccelerationDetails interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of WebAppAccelerationPolicy, which is attached to the resource.
	GetWebAppAccelerationPolicyId() *string

	// WebAppAcceleration display name, can be renamed.
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

type createwebappaccelerationdetails struct {
	JsonData                   []byte
	DisplayName                *string                           `mandatory:"false" json:"displayName"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags                 map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	CompartmentId              *string                           `mandatory:"true" json:"compartmentId"`
	WebAppAccelerationPolicyId *string                           `mandatory:"true" json:"webAppAccelerationPolicyId"`
	BackendType                string                            `json:"backendType"`
}

// UnmarshalJSON unmarshals json
func (m *createwebappaccelerationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatewebappaccelerationdetails createwebappaccelerationdetails
	s := struct {
		Model Unmarshalercreatewebappaccelerationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.WebAppAccelerationPolicyId = s.Model.WebAppAccelerationPolicyId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.BackendType = s.Model.BackendType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createwebappaccelerationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BackendType {
	case "LOAD_BALANCER":
		mm := CreateWebAppAccelerationLoadBalancerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateWebAppAccelerationDetails: %s.", m.BackendType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createwebappaccelerationdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m createwebappaccelerationdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createwebappaccelerationdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m createwebappaccelerationdetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetCompartmentId returns CompartmentId
func (m createwebappaccelerationdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetWebAppAccelerationPolicyId returns WebAppAccelerationPolicyId
func (m createwebappaccelerationdetails) GetWebAppAccelerationPolicyId() *string {
	return m.WebAppAccelerationPolicyId
}

func (m createwebappaccelerationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createwebappaccelerationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
