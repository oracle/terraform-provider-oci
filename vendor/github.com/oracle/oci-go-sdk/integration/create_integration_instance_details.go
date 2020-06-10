// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateIntegrationInstanceDetails The information about new IntegrationInstance.
type CreateIntegrationInstanceDetails struct {

	// Integration Instance Identifier.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Standard or Enterprise type
	IntegrationInstanceType CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum `mandatory:"true" json:"integrationInstanceType"`

	// Bring your own license.
	IsByol *bool `mandatory:"true" json:"isByol"`

	// The number of configured message packs
	MessagePacks *int `mandatory:"true" json:"messagePacks"`

	// Simple key-value pair that is applied without any predefined name,
	// type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to
	// namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// IDCS Authentication token. This is is required for pre-UCPIS cloud accounts, but not UCPIS, hence not a required parameter
	IdcsAt *string `mandatory:"false" json:"idcsAt"`

	// The file server is enabled or not.
	IsFileServerEnabled *bool `mandatory:"false" json:"isFileServerEnabled"`
}

func (m CreateIntegrationInstanceDetails) String() string {
	return common.PointerString(m)
}

// CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum Enum with underlying type: string
type CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum string

// Set of constants representing the allowable values for CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
const (
	CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandard   CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "STANDARD"
	CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "ENTERPRISE"
)

var mappingCreateIntegrationInstanceDetailsIntegrationInstanceType = map[string]CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum{
	"STANDARD":   CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandard,
	"ENTERPRISE": CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise,
}

// GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues Enumerates the set of values for CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
func GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues() []CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum {
	values := make([]CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingCreateIntegrationInstanceDetailsIntegrationInstanceType {
		values = append(values, v)
	}
	return values
}
