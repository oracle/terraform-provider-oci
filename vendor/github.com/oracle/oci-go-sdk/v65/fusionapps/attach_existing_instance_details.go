// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AttachExistingInstanceDetails Information about the service attachment.
type AttachExistingInstanceDetails struct {

	// The service instance OCID of the instance being attached
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Type of the ServiceInstance being attached.
	ServiceInstanceType AttachExistingInstanceDetailsServiceInstanceTypeEnum `mandatory:"true" json:"serviceInstanceType"`
}

func (m AttachExistingInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttachExistingInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttachExistingInstanceDetailsServiceInstanceTypeEnum(string(m.ServiceInstanceType)); !ok && m.ServiceInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceInstanceType: %s. Supported values are: %s.", m.ServiceInstanceType, strings.Join(GetAttachExistingInstanceDetailsServiceInstanceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AttachExistingInstanceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAttachExistingInstanceDetails AttachExistingInstanceDetails
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeAttachExistingInstanceDetails
	}{
		"ATTACH_EXISTING_INSTANCE",
		(MarshalTypeAttachExistingInstanceDetails)(m),
	}

	return json.Marshal(&s)
}

// AttachExistingInstanceDetailsServiceInstanceTypeEnum Enum with underlying type: string
type AttachExistingInstanceDetailsServiceInstanceTypeEnum string

// Set of constants representing the allowable values for AttachExistingInstanceDetailsServiceInstanceTypeEnum
const (
	AttachExistingInstanceDetailsServiceInstanceTypeIntegrationCloud   AttachExistingInstanceDetailsServiceInstanceTypeEnum = "INTEGRATION_CLOUD"
	AttachExistingInstanceDetailsServiceInstanceTypeAnalyticsWarehouse AttachExistingInstanceDetailsServiceInstanceTypeEnum = "ANALYTICS_WAREHOUSE"
)

var mappingAttachExistingInstanceDetailsServiceInstanceTypeEnum = map[string]AttachExistingInstanceDetailsServiceInstanceTypeEnum{
	"INTEGRATION_CLOUD":   AttachExistingInstanceDetailsServiceInstanceTypeIntegrationCloud,
	"ANALYTICS_WAREHOUSE": AttachExistingInstanceDetailsServiceInstanceTypeAnalyticsWarehouse,
}

var mappingAttachExistingInstanceDetailsServiceInstanceTypeEnumLowerCase = map[string]AttachExistingInstanceDetailsServiceInstanceTypeEnum{
	"integration_cloud":   AttachExistingInstanceDetailsServiceInstanceTypeIntegrationCloud,
	"analytics_warehouse": AttachExistingInstanceDetailsServiceInstanceTypeAnalyticsWarehouse,
}

// GetAttachExistingInstanceDetailsServiceInstanceTypeEnumValues Enumerates the set of values for AttachExistingInstanceDetailsServiceInstanceTypeEnum
func GetAttachExistingInstanceDetailsServiceInstanceTypeEnumValues() []AttachExistingInstanceDetailsServiceInstanceTypeEnum {
	values := make([]AttachExistingInstanceDetailsServiceInstanceTypeEnum, 0)
	for _, v := range mappingAttachExistingInstanceDetailsServiceInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttachExistingInstanceDetailsServiceInstanceTypeEnumStringValues Enumerates the set of values in String for AttachExistingInstanceDetailsServiceInstanceTypeEnum
func GetAttachExistingInstanceDetailsServiceInstanceTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_CLOUD",
		"ANALYTICS_WAREHOUSE",
	}
}

// GetMappingAttachExistingInstanceDetailsServiceInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttachExistingInstanceDetailsServiceInstanceTypeEnum(val string) (AttachExistingInstanceDetailsServiceInstanceTypeEnum, bool) {
	enum, ok := mappingAttachExistingInstanceDetailsServiceInstanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
