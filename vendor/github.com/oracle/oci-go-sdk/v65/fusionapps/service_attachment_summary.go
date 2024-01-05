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

// ServiceAttachmentSummary Summary of the ServiceInstance.
type ServiceAttachmentSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// ServiceInstance Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the service.
	ServiceInstanceType ServiceAttachmentServiceInstanceTypeEnum `mandatory:"true" json:"serviceInstanceType"`

	// The current state of the ServiceInstance.
	LifecycleState ServiceAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Whether this service is provisioned due to the customer being subscribed to a specific SKU
	IsSkuBased *bool `mandatory:"true" json:"isSkuBased"`

	// The ID of the service instance created that can be used to identify this on the service control plane
	ServiceInstanceId *string `mandatory:"false" json:"serviceInstanceId"`

	// Service URL of the instance
	ServiceUrl *string `mandatory:"false" json:"serviceUrl"`

	// The time the service instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the serivce instance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ServiceAttachmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceAttachmentSummary) ValidateEnumValue() (bool, error) {
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
