// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateIntegrationInstanceDetails The information to be updated.
type UpdateIntegrationInstanceDetails struct {

	// Integration Instance Identifier.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Standard or Enterprise type
	IntegrationInstanceType UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum `mandatory:"false" json:"integrationInstanceType,omitempty"`

	// Simple key-value pair that is applied without any predefined name,
	// type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to
	// namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Bring your own license.
	IsByol *bool `mandatory:"false" json:"isByol"`

	// The number of configured message packs
	MessagePacks *int `mandatory:"false" json:"messagePacks"`

	// The file server is enabled or not.
	IsFileServerEnabled *bool `mandatory:"false" json:"isFileServerEnabled"`

	// Visual Builder is enabled or not.
	IsVisualBuilderEnabled *bool `mandatory:"false" json:"isVisualBuilderEnabled"`

	CustomEndpoint *UpdateCustomEndpointDetails `mandatory:"false" json:"customEndpoint"`

	// A list of alternate custom endpoints to be used for the integration instance URL
	// (contact Oracle for alternateCustomEndpoints availability for a specific instance).
	AlternateCustomEndpoints []UpdateCustomEndpointDetails `mandatory:"false" json:"alternateCustomEndpoints"`
}

func (m UpdateIntegrationInstanceDetails) String() string {
	return common.PointerString(m)
}

// UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum Enum with underlying type: string
type UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum string

// Set of constants representing the allowable values for UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
const (
	UpdateIntegrationInstanceDetailsIntegrationInstanceTypeStandard   UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "STANDARD"
	UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "ENTERPRISE"
)

var mappingUpdateIntegrationInstanceDetailsIntegrationInstanceType = map[string]UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum{
	"STANDARD":   UpdateIntegrationInstanceDetailsIntegrationInstanceTypeStandard,
	"ENTERPRISE": UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise,
}

// GetUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues Enumerates the set of values for UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
func GetUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues() []UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum {
	values := make([]UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingUpdateIntegrationInstanceDetailsIntegrationInstanceType {
		values = append(values, v)
	}
	return values
}
