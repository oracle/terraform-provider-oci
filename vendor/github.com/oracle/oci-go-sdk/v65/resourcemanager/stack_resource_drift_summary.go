// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StackResourceDriftSummary Drift status details for the indicated resource and stack. Includes actual and expected (defined) properties.
type StackResourceDriftSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack.
	StackId *string `mandatory:"false" json:"stackId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the stack is located.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the resource as defined in the stack.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource provisioned by Terraform.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The provider resource type.
	// Must be supported by the Oracle Cloud Infrastructure provider (https://www.terraform.io/docs/providers/oci/index.html).
	// Example: `oci_core_instance`
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// The drift status of the resource.
	// A drift status value indicates whether or not the actual state of the resource differs from the expected (defined) state for that resource.
	ResourceDriftStatus StackResourceDriftSummaryResourceDriftStatusEnum `mandatory:"false" json:"resourceDriftStatus,omitempty"`

	// Actual values of properties that the stack defines for the indicated resource.
	// Each property and value is provided as a key-value pair.
	// The following example shows actual values for the resource's display name and server type:
	// `{"display_name": "tf-default-dhcp-options-new", "options.0.server_type": "VcnLocalPlusInternet"}`
	ActualProperties map[string]string `mandatory:"false" json:"actualProperties"`

	// Expected values of properties that the stack defines for the indicated resource.
	// Each property and value is provided as a key-value pair.
	// The following example shows expected (defined) values for the resource's display name and server type:
	// `{"display_name": "tf-default-dhcp-options", "options.0.server_type": "VcnLocalPlusInternet"}`
	ExpectedProperties map[string]string `mandatory:"false" json:"expectedProperties"`

	// The date and time when the drift detection was executed.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeDriftChecked *common.SDKTime `mandatory:"false" json:"timeDriftChecked"`
}

func (m StackResourceDriftSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StackResourceDriftSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStackResourceDriftSummaryResourceDriftStatusEnum(string(m.ResourceDriftStatus)); !ok && m.ResourceDriftStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceDriftStatus: %s. Supported values are: %s.", m.ResourceDriftStatus, strings.Join(GetStackResourceDriftSummaryResourceDriftStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StackResourceDriftSummaryResourceDriftStatusEnum Enum with underlying type: string
type StackResourceDriftSummaryResourceDriftStatusEnum string

// Set of constants representing the allowable values for StackResourceDriftSummaryResourceDriftStatusEnum
const (
	StackResourceDriftSummaryResourceDriftStatusNotChecked StackResourceDriftSummaryResourceDriftStatusEnum = "NOT_CHECKED"
	StackResourceDriftSummaryResourceDriftStatusInSync     StackResourceDriftSummaryResourceDriftStatusEnum = "IN_SYNC"
	StackResourceDriftSummaryResourceDriftStatusModified   StackResourceDriftSummaryResourceDriftStatusEnum = "MODIFIED"
	StackResourceDriftSummaryResourceDriftStatusDeleted    StackResourceDriftSummaryResourceDriftStatusEnum = "DELETED"
)

var mappingStackResourceDriftSummaryResourceDriftStatusEnum = map[string]StackResourceDriftSummaryResourceDriftStatusEnum{
	"NOT_CHECKED": StackResourceDriftSummaryResourceDriftStatusNotChecked,
	"IN_SYNC":     StackResourceDriftSummaryResourceDriftStatusInSync,
	"MODIFIED":    StackResourceDriftSummaryResourceDriftStatusModified,
	"DELETED":     StackResourceDriftSummaryResourceDriftStatusDeleted,
}

var mappingStackResourceDriftSummaryResourceDriftStatusEnumLowerCase = map[string]StackResourceDriftSummaryResourceDriftStatusEnum{
	"not_checked": StackResourceDriftSummaryResourceDriftStatusNotChecked,
	"in_sync":     StackResourceDriftSummaryResourceDriftStatusInSync,
	"modified":    StackResourceDriftSummaryResourceDriftStatusModified,
	"deleted":     StackResourceDriftSummaryResourceDriftStatusDeleted,
}

// GetStackResourceDriftSummaryResourceDriftStatusEnumValues Enumerates the set of values for StackResourceDriftSummaryResourceDriftStatusEnum
func GetStackResourceDriftSummaryResourceDriftStatusEnumValues() []StackResourceDriftSummaryResourceDriftStatusEnum {
	values := make([]StackResourceDriftSummaryResourceDriftStatusEnum, 0)
	for _, v := range mappingStackResourceDriftSummaryResourceDriftStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetStackResourceDriftSummaryResourceDriftStatusEnumStringValues Enumerates the set of values in String for StackResourceDriftSummaryResourceDriftStatusEnum
func GetStackResourceDriftSummaryResourceDriftStatusEnumStringValues() []string {
	return []string{
		"NOT_CHECKED",
		"IN_SYNC",
		"MODIFIED",
		"DELETED",
	}
}

// GetMappingStackResourceDriftSummaryResourceDriftStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStackResourceDriftSummaryResourceDriftStatusEnum(val string) (StackResourceDriftSummaryResourceDriftStatusEnum, bool) {
	enum, ok := mappingStackResourceDriftSummaryResourceDriftStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
