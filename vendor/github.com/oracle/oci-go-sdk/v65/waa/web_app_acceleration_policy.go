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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WebAppAccelerationPolicy The details of WebAppAccelerationPolicy. A policy is comprised of rules, which allows enablement of Caching
// and Compression of HTTP response.
// Caching can be enabled for a particular path
// Compression is enabled at global level
type WebAppAccelerationPolicy struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebAppAccelerationPolicy.
	Id *string `mandatory:"true" json:"id"`

	// WebAppAccelerationPolicy display name, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the WebAppAccelerationPolicy was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the WebAppAccelerationPolicy.
	LifecycleState WebAppAccelerationPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The time the WebAppAccelerationPolicy was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in FAILED state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	ResponseCachingPolicy *ResponseCachingPolicy `mandatory:"false" json:"responseCachingPolicy"`

	ResponseCompressionPolicy *ResponseCompressionPolicy `mandatory:"false" json:"responseCompressionPolicy"`
}

func (m WebAppAccelerationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WebAppAccelerationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWebAppAccelerationPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWebAppAccelerationPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WebAppAccelerationPolicyLifecycleStateEnum Enum with underlying type: string
type WebAppAccelerationPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for WebAppAccelerationPolicyLifecycleStateEnum
const (
	WebAppAccelerationPolicyLifecycleStateCreating WebAppAccelerationPolicyLifecycleStateEnum = "CREATING"
	WebAppAccelerationPolicyLifecycleStateUpdating WebAppAccelerationPolicyLifecycleStateEnum = "UPDATING"
	WebAppAccelerationPolicyLifecycleStateActive   WebAppAccelerationPolicyLifecycleStateEnum = "ACTIVE"
	WebAppAccelerationPolicyLifecycleStateDeleting WebAppAccelerationPolicyLifecycleStateEnum = "DELETING"
	WebAppAccelerationPolicyLifecycleStateDeleted  WebAppAccelerationPolicyLifecycleStateEnum = "DELETED"
	WebAppAccelerationPolicyLifecycleStateFailed   WebAppAccelerationPolicyLifecycleStateEnum = "FAILED"
)

var mappingWebAppAccelerationPolicyLifecycleStateEnum = map[string]WebAppAccelerationPolicyLifecycleStateEnum{
	"CREATING": WebAppAccelerationPolicyLifecycleStateCreating,
	"UPDATING": WebAppAccelerationPolicyLifecycleStateUpdating,
	"ACTIVE":   WebAppAccelerationPolicyLifecycleStateActive,
	"DELETING": WebAppAccelerationPolicyLifecycleStateDeleting,
	"DELETED":  WebAppAccelerationPolicyLifecycleStateDeleted,
	"FAILED":   WebAppAccelerationPolicyLifecycleStateFailed,
}

var mappingWebAppAccelerationPolicyLifecycleStateEnumLowerCase = map[string]WebAppAccelerationPolicyLifecycleStateEnum{
	"creating": WebAppAccelerationPolicyLifecycleStateCreating,
	"updating": WebAppAccelerationPolicyLifecycleStateUpdating,
	"active":   WebAppAccelerationPolicyLifecycleStateActive,
	"deleting": WebAppAccelerationPolicyLifecycleStateDeleting,
	"deleted":  WebAppAccelerationPolicyLifecycleStateDeleted,
	"failed":   WebAppAccelerationPolicyLifecycleStateFailed,
}

// GetWebAppAccelerationPolicyLifecycleStateEnumValues Enumerates the set of values for WebAppAccelerationPolicyLifecycleStateEnum
func GetWebAppAccelerationPolicyLifecycleStateEnumValues() []WebAppAccelerationPolicyLifecycleStateEnum {
	values := make([]WebAppAccelerationPolicyLifecycleStateEnum, 0)
	for _, v := range mappingWebAppAccelerationPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetWebAppAccelerationPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for WebAppAccelerationPolicyLifecycleStateEnum
func GetWebAppAccelerationPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingWebAppAccelerationPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWebAppAccelerationPolicyLifecycleStateEnum(val string) (WebAppAccelerationPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingWebAppAccelerationPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
