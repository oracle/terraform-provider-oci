// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Sdk Information about the SDK.
type Sdk struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of API resource
	ApiId *string `mandatory:"true" json:"apiId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The string representing the target programming language for generating the SDK.
	TargetLanguage *string `mandatory:"true" json:"targetLanguage"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// File location for generated SDK.
	ArtifactUrl *string `mandatory:"false" json:"artifactUrl"`

	// Expiry of artifact url.
	TimeArtifactUrlExpiresAt *common.SDKTime `mandatory:"false" json:"timeArtifactUrlExpiresAt"`

	// The current state of the SDK.
	// - The SDK will be in CREATING state if the SDK creation is in progress.
	// - The SDK will be in ACTIVE state if create is successful.
	// - The SDK will be in FAILED state if the create, or delete fails.
	// - The SDK will be in DELETING state if the deletion in in progress.
	// - The SDK will be in DELETED state if the delete is successful.
	LifecycleState SdkLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a
	// resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Additional optional configurations passed.
	// The applicable config keys are listed under "parameters" when "/sdkLanguageTypes" is called.
	// Example: `{"configName": "configValue"}`
	Parameters map[string]string `mandatory:"false" json:"parameters"`
}

func (m Sdk) String() string {
	return common.PointerString(m)
}

// SdkLifecycleStateEnum Enum with underlying type: string
type SdkLifecycleStateEnum string

// Set of constants representing the allowable values for SdkLifecycleStateEnum
const (
	SdkLifecycleStateCreating SdkLifecycleStateEnum = "CREATING"
	SdkLifecycleStateActive   SdkLifecycleStateEnum = "ACTIVE"
	SdkLifecycleStateFailed   SdkLifecycleStateEnum = "FAILED"
	SdkLifecycleStateDeleting SdkLifecycleStateEnum = "DELETING"
	SdkLifecycleStateDeleted  SdkLifecycleStateEnum = "DELETED"
)

var mappingSdkLifecycleState = map[string]SdkLifecycleStateEnum{
	"CREATING": SdkLifecycleStateCreating,
	"ACTIVE":   SdkLifecycleStateActive,
	"FAILED":   SdkLifecycleStateFailed,
	"DELETING": SdkLifecycleStateDeleting,
	"DELETED":  SdkLifecycleStateDeleted,
}

// GetSdkLifecycleStateEnumValues Enumerates the set of values for SdkLifecycleStateEnum
func GetSdkLifecycleStateEnumValues() []SdkLifecycleStateEnum {
	values := make([]SdkLifecycleStateEnum, 0)
	for _, v := range mappingSdkLifecycleState {
		values = append(values, v)
	}
	return values
}
