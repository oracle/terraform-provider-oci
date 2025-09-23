// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceAnalyticsInstance A ResourceAnalyticsInstance is an ADW housing analytics for all of a customers' OCI resources.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type ResourceAnalyticsInstance struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the ResourceAnalyticsInstance was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the ResourceAnalyticsInstance.
	LifecycleState ResourceAnalyticsInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A description of the ResourceAnalyticsInstance instance.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the created ADW instance.
	AdwId *string `mandatory:"false" json:"adwId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OAC enabled for the ResourceAnalyticsInstance.
	OacId *string `mandatory:"false" json:"oacId"`

	// The date and time the ResourceAnalyticsInstance was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the ResourceAnalyticsInstance in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ResourceAnalyticsInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceAnalyticsInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceAnalyticsInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourceAnalyticsInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceAnalyticsInstanceLifecycleStateEnum Enum with underlying type: string
type ResourceAnalyticsInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for ResourceAnalyticsInstanceLifecycleStateEnum
const (
	ResourceAnalyticsInstanceLifecycleStateCreating       ResourceAnalyticsInstanceLifecycleStateEnum = "CREATING"
	ResourceAnalyticsInstanceLifecycleStateUpdating       ResourceAnalyticsInstanceLifecycleStateEnum = "UPDATING"
	ResourceAnalyticsInstanceLifecycleStateActive         ResourceAnalyticsInstanceLifecycleStateEnum = "ACTIVE"
	ResourceAnalyticsInstanceLifecycleStateNeedsAttention ResourceAnalyticsInstanceLifecycleStateEnum = "NEEDS_ATTENTION"
	ResourceAnalyticsInstanceLifecycleStateDeleting       ResourceAnalyticsInstanceLifecycleStateEnum = "DELETING"
	ResourceAnalyticsInstanceLifecycleStateDeleted        ResourceAnalyticsInstanceLifecycleStateEnum = "DELETED"
	ResourceAnalyticsInstanceLifecycleStateFailed         ResourceAnalyticsInstanceLifecycleStateEnum = "FAILED"
)

var mappingResourceAnalyticsInstanceLifecycleStateEnum = map[string]ResourceAnalyticsInstanceLifecycleStateEnum{
	"CREATING":        ResourceAnalyticsInstanceLifecycleStateCreating,
	"UPDATING":        ResourceAnalyticsInstanceLifecycleStateUpdating,
	"ACTIVE":          ResourceAnalyticsInstanceLifecycleStateActive,
	"NEEDS_ATTENTION": ResourceAnalyticsInstanceLifecycleStateNeedsAttention,
	"DELETING":        ResourceAnalyticsInstanceLifecycleStateDeleting,
	"DELETED":         ResourceAnalyticsInstanceLifecycleStateDeleted,
	"FAILED":          ResourceAnalyticsInstanceLifecycleStateFailed,
}

var mappingResourceAnalyticsInstanceLifecycleStateEnumLowerCase = map[string]ResourceAnalyticsInstanceLifecycleStateEnum{
	"creating":        ResourceAnalyticsInstanceLifecycleStateCreating,
	"updating":        ResourceAnalyticsInstanceLifecycleStateUpdating,
	"active":          ResourceAnalyticsInstanceLifecycleStateActive,
	"needs_attention": ResourceAnalyticsInstanceLifecycleStateNeedsAttention,
	"deleting":        ResourceAnalyticsInstanceLifecycleStateDeleting,
	"deleted":         ResourceAnalyticsInstanceLifecycleStateDeleted,
	"failed":          ResourceAnalyticsInstanceLifecycleStateFailed,
}

// GetResourceAnalyticsInstanceLifecycleStateEnumValues Enumerates the set of values for ResourceAnalyticsInstanceLifecycleStateEnum
func GetResourceAnalyticsInstanceLifecycleStateEnumValues() []ResourceAnalyticsInstanceLifecycleStateEnum {
	values := make([]ResourceAnalyticsInstanceLifecycleStateEnum, 0)
	for _, v := range mappingResourceAnalyticsInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceAnalyticsInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for ResourceAnalyticsInstanceLifecycleStateEnum
func GetResourceAnalyticsInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingResourceAnalyticsInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceAnalyticsInstanceLifecycleStateEnum(val string) (ResourceAnalyticsInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingResourceAnalyticsInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
