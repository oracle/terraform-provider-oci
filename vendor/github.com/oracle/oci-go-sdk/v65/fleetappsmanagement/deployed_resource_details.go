// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeployedResourceDetails The Filtered List Of Deployed Resources
type DeployedResourceDetails struct {

	// The mode of the resource. Example: "managed"
	Mode *string `mandatory:"true" json:"mode"`

	// The name of the resource
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The name of the Provider
	ResourceProvider *string `mandatory:"true" json:"resourceProvider"`

	// The provider resource type. Must be supported by the Oracle Cloud Infrastructure provider (https://registry.terraform.io/providers/oracle/oci/latest/docs).
	// Example: oci_core_instance
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Collection of InstanceSummary
	ResourceInstanceList []InstanceSummary `mandatory:"true" json:"resourceInstanceList"`

	// The drift status of the resource
	ResourceDriftStatus DeployedResourceDetailsResourceDriftStatusEnum `mandatory:"false" json:"resourceDriftStatus,omitempty"`

	// Key-value pair of the actual resource properties
	ActualProperties []KeyValueProperty `mandatory:"false" json:"actualProperties"`

	// Key-value pair of the expected resource properties
	ExpectedProperties []KeyValueProperty `mandatory:"false" json:"expectedProperties"`

	// A list of the modified properties
	ModifiedProperties []ModifiedProperty `mandatory:"false" json:"modifiedProperties"`

	// The time the drift has been checked. An RFC 3339 (https://tools.ietf.org/rfc/rfc3339) formatted datetime string
	TimeDriftChecked *common.SDKTime `mandatory:"false" json:"timeDriftChecked"`
}

func (m DeployedResourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployedResourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeployedResourceDetailsResourceDriftStatusEnum(string(m.ResourceDriftStatus)); !ok && m.ResourceDriftStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceDriftStatus: %s. Supported values are: %s.", m.ResourceDriftStatus, strings.Join(GetDeployedResourceDetailsResourceDriftStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeployedResourceDetailsResourceDriftStatusEnum Enum with underlying type: string
type DeployedResourceDetailsResourceDriftStatusEnum string

// Set of constants representing the allowable values for DeployedResourceDetailsResourceDriftStatusEnum
const (
	DeployedResourceDetailsResourceDriftStatusNotChecked DeployedResourceDetailsResourceDriftStatusEnum = "NOT_CHECKED"
	DeployedResourceDetailsResourceDriftStatusInSync     DeployedResourceDetailsResourceDriftStatusEnum = "IN_SYNC"
	DeployedResourceDetailsResourceDriftStatusModified   DeployedResourceDetailsResourceDriftStatusEnum = "MODIFIED"
	DeployedResourceDetailsResourceDriftStatusDeleted    DeployedResourceDetailsResourceDriftStatusEnum = "DELETED"
)

var mappingDeployedResourceDetailsResourceDriftStatusEnum = map[string]DeployedResourceDetailsResourceDriftStatusEnum{
	"NOT_CHECKED": DeployedResourceDetailsResourceDriftStatusNotChecked,
	"IN_SYNC":     DeployedResourceDetailsResourceDriftStatusInSync,
	"MODIFIED":    DeployedResourceDetailsResourceDriftStatusModified,
	"DELETED":     DeployedResourceDetailsResourceDriftStatusDeleted,
}

var mappingDeployedResourceDetailsResourceDriftStatusEnumLowerCase = map[string]DeployedResourceDetailsResourceDriftStatusEnum{
	"not_checked": DeployedResourceDetailsResourceDriftStatusNotChecked,
	"in_sync":     DeployedResourceDetailsResourceDriftStatusInSync,
	"modified":    DeployedResourceDetailsResourceDriftStatusModified,
	"deleted":     DeployedResourceDetailsResourceDriftStatusDeleted,
}

// GetDeployedResourceDetailsResourceDriftStatusEnumValues Enumerates the set of values for DeployedResourceDetailsResourceDriftStatusEnum
func GetDeployedResourceDetailsResourceDriftStatusEnumValues() []DeployedResourceDetailsResourceDriftStatusEnum {
	values := make([]DeployedResourceDetailsResourceDriftStatusEnum, 0)
	for _, v := range mappingDeployedResourceDetailsResourceDriftStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDeployedResourceDetailsResourceDriftStatusEnumStringValues Enumerates the set of values in String for DeployedResourceDetailsResourceDriftStatusEnum
func GetDeployedResourceDetailsResourceDriftStatusEnumStringValues() []string {
	return []string{
		"NOT_CHECKED",
		"IN_SYNC",
		"MODIFIED",
		"DELETED",
	}
}

// GetMappingDeployedResourceDetailsResourceDriftStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeployedResourceDetailsResourceDriftStatusEnum(val string) (DeployedResourceDetailsResourceDriftStatusEnum, bool) {
	enum, ok := mappingDeployedResourceDetailsResourceDriftStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
