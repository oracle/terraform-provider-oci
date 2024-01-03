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

// CreateWebAppAccelerationLoadBalancerDetails The information about new WebAppAccelerationLoadBalancer.
type CreateWebAppAccelerationLoadBalancerDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of WebAppAccelerationPolicy, which is attached to the resource.
	WebAppAccelerationPolicyId *string `mandatory:"true" json:"webAppAccelerationPolicyId"`

	// LoadBalancer OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) to which the WebAppAccelerationPolicy is attached to.
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

	// WebAppAcceleration display name, can be renamed.
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

// GetDisplayName returns DisplayName
func (m CreateWebAppAccelerationLoadBalancerDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateWebAppAccelerationLoadBalancerDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetWebAppAccelerationPolicyId returns WebAppAccelerationPolicyId
func (m CreateWebAppAccelerationLoadBalancerDetails) GetWebAppAccelerationPolicyId() *string {
	return m.WebAppAccelerationPolicyId
}

// GetFreeformTags returns FreeformTags
func (m CreateWebAppAccelerationLoadBalancerDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateWebAppAccelerationLoadBalancerDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m CreateWebAppAccelerationLoadBalancerDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m CreateWebAppAccelerationLoadBalancerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateWebAppAccelerationLoadBalancerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateWebAppAccelerationLoadBalancerDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateWebAppAccelerationLoadBalancerDetails CreateWebAppAccelerationLoadBalancerDetails
	s := struct {
		DiscriminatorParam string `json:"backendType"`
		MarshalTypeCreateWebAppAccelerationLoadBalancerDetails
	}{
		"LOAD_BALANCER",
		(MarshalTypeCreateWebAppAccelerationLoadBalancerDetails)(m),
	}

	return json.Marshal(&s)
}
