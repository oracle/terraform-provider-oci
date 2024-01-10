// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ServiceAttachment Description of ServiceAttachment.
type ServiceAttachment struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Service Attachment Display name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the serviceInstance.
	ServiceInstanceType ServiceAttachmentServiceInstanceTypeEnum `mandatory:"true" json:"serviceInstanceType"`

	// The current state of the ServiceInstance.
	LifecycleState ServiceAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Whether this service is provisioned due to the customer being subscribed to a specific SKU
	IsSkuBased *bool `mandatory:"true" json:"isSkuBased"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The ID of the service instance created that can be used to identify this on the service control plane
	ServiceInstanceId *string `mandatory:"false" json:"serviceInstanceId"`

	// Public URL
	ServiceUrl *string `mandatory:"false" json:"serviceUrl"`

	// The time the the ServiceInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the ServiceInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ServiceAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceAttachmentServiceInstanceTypeEnum(string(m.ServiceInstanceType)); !ok && m.ServiceInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceInstanceType: %s. Supported values are: %s.", m.ServiceInstanceType, strings.Join(GetServiceAttachmentServiceInstanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingServiceAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetServiceAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ServiceAttachmentServiceInstanceTypeEnum Enum with underlying type: string
type ServiceAttachmentServiceInstanceTypeEnum string

// Set of constants representing the allowable values for ServiceAttachmentServiceInstanceTypeEnum
const (
	ServiceAttachmentServiceInstanceTypeDigitalAssistant    ServiceAttachmentServiceInstanceTypeEnum = "DIGITAL_ASSISTANT"
	ServiceAttachmentServiceInstanceTypeIntegrationCloud    ServiceAttachmentServiceInstanceTypeEnum = "INTEGRATION_CLOUD"
	ServiceAttachmentServiceInstanceTypeAnalyticsWarehouse  ServiceAttachmentServiceInstanceTypeEnum = "ANALYTICS_WAREHOUSE"
	ServiceAttachmentServiceInstanceTypeVbcs                ServiceAttachmentServiceInstanceTypeEnum = "VBCS"
	ServiceAttachmentServiceInstanceTypeVisualBuilderStudio ServiceAttachmentServiceInstanceTypeEnum = "VISUAL_BUILDER_STUDIO"
)

var mappingServiceAttachmentServiceInstanceTypeEnum = map[string]ServiceAttachmentServiceInstanceTypeEnum{
	"DIGITAL_ASSISTANT":     ServiceAttachmentServiceInstanceTypeDigitalAssistant,
	"INTEGRATION_CLOUD":     ServiceAttachmentServiceInstanceTypeIntegrationCloud,
	"ANALYTICS_WAREHOUSE":   ServiceAttachmentServiceInstanceTypeAnalyticsWarehouse,
	"VBCS":                  ServiceAttachmentServiceInstanceTypeVbcs,
	"VISUAL_BUILDER_STUDIO": ServiceAttachmentServiceInstanceTypeVisualBuilderStudio,
}

var mappingServiceAttachmentServiceInstanceTypeEnumLowerCase = map[string]ServiceAttachmentServiceInstanceTypeEnum{
	"digital_assistant":     ServiceAttachmentServiceInstanceTypeDigitalAssistant,
	"integration_cloud":     ServiceAttachmentServiceInstanceTypeIntegrationCloud,
	"analytics_warehouse":   ServiceAttachmentServiceInstanceTypeAnalyticsWarehouse,
	"vbcs":                  ServiceAttachmentServiceInstanceTypeVbcs,
	"visual_builder_studio": ServiceAttachmentServiceInstanceTypeVisualBuilderStudio,
}

// GetServiceAttachmentServiceInstanceTypeEnumValues Enumerates the set of values for ServiceAttachmentServiceInstanceTypeEnum
func GetServiceAttachmentServiceInstanceTypeEnumValues() []ServiceAttachmentServiceInstanceTypeEnum {
	values := make([]ServiceAttachmentServiceInstanceTypeEnum, 0)
	for _, v := range mappingServiceAttachmentServiceInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceAttachmentServiceInstanceTypeEnumStringValues Enumerates the set of values in String for ServiceAttachmentServiceInstanceTypeEnum
func GetServiceAttachmentServiceInstanceTypeEnumStringValues() []string {
	return []string{
		"DIGITAL_ASSISTANT",
		"INTEGRATION_CLOUD",
		"ANALYTICS_WAREHOUSE",
		"VBCS",
		"VISUAL_BUILDER_STUDIO",
	}
}

// GetMappingServiceAttachmentServiceInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceAttachmentServiceInstanceTypeEnum(val string) (ServiceAttachmentServiceInstanceTypeEnum, bool) {
	enum, ok := mappingServiceAttachmentServiceInstanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ServiceAttachmentLifecycleStateEnum Enum with underlying type: string
type ServiceAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for ServiceAttachmentLifecycleStateEnum
const (
	ServiceAttachmentLifecycleStateCreating ServiceAttachmentLifecycleStateEnum = "CREATING"
	ServiceAttachmentLifecycleStateUpdating ServiceAttachmentLifecycleStateEnum = "UPDATING"
	ServiceAttachmentLifecycleStateActive   ServiceAttachmentLifecycleStateEnum = "ACTIVE"
	ServiceAttachmentLifecycleStateDeleting ServiceAttachmentLifecycleStateEnum = "DELETING"
	ServiceAttachmentLifecycleStateDeleted  ServiceAttachmentLifecycleStateEnum = "DELETED"
	ServiceAttachmentLifecycleStateFailed   ServiceAttachmentLifecycleStateEnum = "FAILED"
)

var mappingServiceAttachmentLifecycleStateEnum = map[string]ServiceAttachmentLifecycleStateEnum{
	"CREATING": ServiceAttachmentLifecycleStateCreating,
	"UPDATING": ServiceAttachmentLifecycleStateUpdating,
	"ACTIVE":   ServiceAttachmentLifecycleStateActive,
	"DELETING": ServiceAttachmentLifecycleStateDeleting,
	"DELETED":  ServiceAttachmentLifecycleStateDeleted,
	"FAILED":   ServiceAttachmentLifecycleStateFailed,
}

var mappingServiceAttachmentLifecycleStateEnumLowerCase = map[string]ServiceAttachmentLifecycleStateEnum{
	"creating": ServiceAttachmentLifecycleStateCreating,
	"updating": ServiceAttachmentLifecycleStateUpdating,
	"active":   ServiceAttachmentLifecycleStateActive,
	"deleting": ServiceAttachmentLifecycleStateDeleting,
	"deleted":  ServiceAttachmentLifecycleStateDeleted,
	"failed":   ServiceAttachmentLifecycleStateFailed,
}

// GetServiceAttachmentLifecycleStateEnumValues Enumerates the set of values for ServiceAttachmentLifecycleStateEnum
func GetServiceAttachmentLifecycleStateEnumValues() []ServiceAttachmentLifecycleStateEnum {
	values := make([]ServiceAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingServiceAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for ServiceAttachmentLifecycleStateEnum
func GetServiceAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingServiceAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceAttachmentLifecycleStateEnum(val string) (ServiceAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingServiceAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
