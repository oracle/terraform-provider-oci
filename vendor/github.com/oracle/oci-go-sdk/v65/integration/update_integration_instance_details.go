// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateIntegrationInstanceDetails The information to be updated.
// Some properties may not be applicable to specific integration types,
// see Differences in Instance Management (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/application-integration&id=INTOO-GUID-931B5E33-4FE6-4997-93E5-8748516F46AA__GUID-176E43D5-4116-4828-8120-B929DF2A6B5E)
// for details.
type UpdateIntegrationInstanceDetails struct {

	// Integration Instance Identifier.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Standard or Enterprise type,
	// Oracle Integration Generation 2 uses ENTERPRISE and STANDARD,
	// Oracle Integration 3 uses ENTERPRISEX and STANDARDX
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateIntegrationInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum(string(m.IntegrationInstanceType)); !ok && m.IntegrationInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IntegrationInstanceType: %s. Supported values are: %s.", m.IntegrationInstanceType, strings.Join(GetUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum Enum with underlying type: string
type UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum string

// Set of constants representing the allowable values for UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
const (
	UpdateIntegrationInstanceDetailsIntegrationInstanceTypeStandard    UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "STANDARD"
	UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise  UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "ENTERPRISE"
	UpdateIntegrationInstanceDetailsIntegrationInstanceTypeStandardx   UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "STANDARDX"
	UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprisex UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "ENTERPRISEX"
)

var mappingUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = map[string]UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum{
	"STANDARD":    UpdateIntegrationInstanceDetailsIntegrationInstanceTypeStandard,
	"ENTERPRISE":  UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise,
	"STANDARDX":   UpdateIntegrationInstanceDetailsIntegrationInstanceTypeStandardx,
	"ENTERPRISEX": UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprisex,
}

var mappingUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumLowerCase = map[string]UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum{
	"standard":    UpdateIntegrationInstanceDetailsIntegrationInstanceTypeStandard,
	"enterprise":  UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise,
	"standardx":   UpdateIntegrationInstanceDetailsIntegrationInstanceTypeStandardx,
	"enterprisex": UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprisex,
}

// GetUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues Enumerates the set of values for UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
func GetUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues() []UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum {
	values := make([]UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumStringValues Enumerates the set of values in String for UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
func GetUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"ENTERPRISE",
		"STANDARDX",
		"ENTERPRISEX",
	}
}

// GetMappingUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum(val string) (UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum, bool) {
	enum, ok := mappingUpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
