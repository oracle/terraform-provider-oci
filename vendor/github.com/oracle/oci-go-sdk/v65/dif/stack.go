// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Stack A Stack is a collection of a templates and services.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type Stack struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Stack.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of templates to be onboarded for the stack.
	StackTemplates []StackTemplateEnum `mandatory:"true" json:"stackTemplates"`

	// List of services to be onboarded for the stack.
	Services []ServiceEnum `mandatory:"true" json:"services"`

	// The date and time the Stack was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Stack.
	LifecycleState StackLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// email id to which the stack notifications would be sent.
	NotificationEmail *string `mandatory:"false" json:"notificationEmail"`

	// ADB details if adb is included in the services.
	Adb []AdbDetail `mandatory:"false" json:"adb"`

	// GGCS details if ggcs is included in the services.
	Ggcs []GgcsDetail `mandatory:"false" json:"ggcs"`

	// DATAFLOW details if dataflow is included in the services.
	Dataflow []DataflowDetail `mandatory:"false" json:"dataflow"`

	// Object Storage Details if object storage is included in services.
	Objectstorage []ObjectStorageDetail `mandatory:"false" json:"objectstorage"`

	// GenAI Details if genai is included in services.
	Genai []GenAiDetail `mandatory:"false" json:"genai"`

	// Details of the service onboarded for the data intelligence stack.
	ServiceDetails []ServiceDetailResponse `mandatory:"false" json:"serviceDetails"`

	// The date and time the Stack was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the Stack in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Stack) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Stack) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.StackTemplates {
		if _, ok := GetMappingStackTemplateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StackTemplates: %s. Supported values are: %s.", val, strings.Join(GetStackTemplateEnumStringValues(), ",")))
		}
	}

	for _, val := range m.Services {
		if _, ok := GetMappingServiceEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Services: %s. Supported values are: %s.", val, strings.Join(GetServiceEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingStackLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStackLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StackLifecycleStateEnum Enum with underlying type: string
type StackLifecycleStateEnum string

// Set of constants representing the allowable values for StackLifecycleStateEnum
const (
	StackLifecycleStateCreating       StackLifecycleStateEnum = "CREATING"
	StackLifecycleStateUpdating       StackLifecycleStateEnum = "UPDATING"
	StackLifecycleStateActive         StackLifecycleStateEnum = "ACTIVE"
	StackLifecycleStateDeleting       StackLifecycleStateEnum = "DELETING"
	StackLifecycleStateDeleted        StackLifecycleStateEnum = "DELETED"
	StackLifecycleStateFailed         StackLifecycleStateEnum = "FAILED"
	StackLifecycleStateNeedsAttention StackLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingStackLifecycleStateEnum = map[string]StackLifecycleStateEnum{
	"CREATING":        StackLifecycleStateCreating,
	"UPDATING":        StackLifecycleStateUpdating,
	"ACTIVE":          StackLifecycleStateActive,
	"DELETING":        StackLifecycleStateDeleting,
	"DELETED":         StackLifecycleStateDeleted,
	"FAILED":          StackLifecycleStateFailed,
	"NEEDS_ATTENTION": StackLifecycleStateNeedsAttention,
}

var mappingStackLifecycleStateEnumLowerCase = map[string]StackLifecycleStateEnum{
	"creating":        StackLifecycleStateCreating,
	"updating":        StackLifecycleStateUpdating,
	"active":          StackLifecycleStateActive,
	"deleting":        StackLifecycleStateDeleting,
	"deleted":         StackLifecycleStateDeleted,
	"failed":          StackLifecycleStateFailed,
	"needs_attention": StackLifecycleStateNeedsAttention,
}

// GetStackLifecycleStateEnumValues Enumerates the set of values for StackLifecycleStateEnum
func GetStackLifecycleStateEnumValues() []StackLifecycleStateEnum {
	values := make([]StackLifecycleStateEnum, 0)
	for _, v := range mappingStackLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStackLifecycleStateEnumStringValues Enumerates the set of values in String for StackLifecycleStateEnum
func GetStackLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingStackLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStackLifecycleStateEnum(val string) (StackLifecycleStateEnum, bool) {
	enum, ok := mappingStackLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
