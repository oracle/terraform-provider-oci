// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WlsDomain WLSDomain is representative of a WebLogic Domain running on one or more managed instances.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type WlsDomain struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name that does not have to be unique and is changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the WebLogic service domain.
	LifecycleState WlsDomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The middleware type on the administration server of the WebLogic domain.
	MiddlewareType *string `mandatory:"false" json:"middlewareType"`

	// The version of the WebLogic domain.
	WeblogicVersion *string `mandatory:"false" json:"weblogicVersion"`

	// The patch readiness status of the WebLogic domain.
	PatchReadinessStatus PatchReadinessStatusEnum `mandatory:"false" json:"patchReadinessStatus,omitempty"`

	// Whether or not the terms of use agreement has been accepted for the WebLogic domain.
	IsAcceptedTermsAndConditions *bool `mandatory:"false" json:"isAcceptedTermsAndConditions"`

	// A message that describes the current state of the WebLogic domain in more detail. For example,
	// it can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	Configuration *WlsDomainConfiguration `mandatory:"false" json:"configuration"`

	// The date and time the WebLogic domain was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the WebLogic domain was updated (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m WlsDomain) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WlsDomain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWlsDomainLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWlsDomainLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchReadinessStatusEnum(string(m.PatchReadinessStatus)); !ok && m.PatchReadinessStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchReadinessStatus: %s. Supported values are: %s.", m.PatchReadinessStatus, strings.Join(GetPatchReadinessStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WlsDomainLifecycleStateEnum Enum with underlying type: string
type WlsDomainLifecycleStateEnum string

// Set of constants representing the allowable values for WlsDomainLifecycleStateEnum
const (
	WlsDomainLifecycleStateActive         WlsDomainLifecycleStateEnum = "ACTIVE"
	WlsDomainLifecycleStateCreating       WlsDomainLifecycleStateEnum = "CREATING"
	WlsDomainLifecycleStateDeleted        WlsDomainLifecycleStateEnum = "DELETED"
	WlsDomainLifecycleStateDeleting       WlsDomainLifecycleStateEnum = "DELETING"
	WlsDomainLifecycleStateFailed         WlsDomainLifecycleStateEnum = "FAILED"
	WlsDomainLifecycleStateNeedsAttention WlsDomainLifecycleStateEnum = "NEEDS_ATTENTION"
	WlsDomainLifecycleStateUpdating       WlsDomainLifecycleStateEnum = "UPDATING"
)

var mappingWlsDomainLifecycleStateEnum = map[string]WlsDomainLifecycleStateEnum{
	"ACTIVE":          WlsDomainLifecycleStateActive,
	"CREATING":        WlsDomainLifecycleStateCreating,
	"DELETED":         WlsDomainLifecycleStateDeleted,
	"DELETING":        WlsDomainLifecycleStateDeleting,
	"FAILED":          WlsDomainLifecycleStateFailed,
	"NEEDS_ATTENTION": WlsDomainLifecycleStateNeedsAttention,
	"UPDATING":        WlsDomainLifecycleStateUpdating,
}

var mappingWlsDomainLifecycleStateEnumLowerCase = map[string]WlsDomainLifecycleStateEnum{
	"active":          WlsDomainLifecycleStateActive,
	"creating":        WlsDomainLifecycleStateCreating,
	"deleted":         WlsDomainLifecycleStateDeleted,
	"deleting":        WlsDomainLifecycleStateDeleting,
	"failed":          WlsDomainLifecycleStateFailed,
	"needs_attention": WlsDomainLifecycleStateNeedsAttention,
	"updating":        WlsDomainLifecycleStateUpdating,
}

// GetWlsDomainLifecycleStateEnumValues Enumerates the set of values for WlsDomainLifecycleStateEnum
func GetWlsDomainLifecycleStateEnumValues() []WlsDomainLifecycleStateEnum {
	values := make([]WlsDomainLifecycleStateEnum, 0)
	for _, v := range mappingWlsDomainLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetWlsDomainLifecycleStateEnumStringValues Enumerates the set of values in String for WlsDomainLifecycleStateEnum
func GetWlsDomainLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"NEEDS_ATTENTION",
		"UPDATING",
	}
}

// GetMappingWlsDomainLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWlsDomainLifecycleStateEnum(val string) (WlsDomainLifecycleStateEnum, bool) {
	enum, ok := mappingWlsDomainLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
