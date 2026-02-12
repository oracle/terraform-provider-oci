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

// PatchGroup Patch group.
type PatchGroup struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"true" json:"description"`

	// The current state of the Patch Group.
	LifecycleState PatchGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// create as tenancy level patch group.
	IsTenancyLevel *bool `mandatory:"false" json:"isTenancyLevel"`

	// Start time when the patch group is applicable for the patchGroup locked fleets.
	TimeApplicableFrom *common.SDKTime `mandatory:"false" json:"timeApplicableFrom"`

	// End time when the patch group is not applicable to any patchGroup locked fleets.
	TimeApplicableTo *common.SDKTime `mandatory:"false" json:"timeApplicableTo"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PatchGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchGroupLifecycleStateEnum Enum with underlying type: string
type PatchGroupLifecycleStateEnum string

// Set of constants representing the allowable values for PatchGroupLifecycleStateEnum
const (
	PatchGroupLifecycleStateActive   PatchGroupLifecycleStateEnum = "ACTIVE"
	PatchGroupLifecycleStateInactive PatchGroupLifecycleStateEnum = "INACTIVE"
	PatchGroupLifecycleStateDeleted  PatchGroupLifecycleStateEnum = "DELETED"
	PatchGroupLifecycleStateDeleting PatchGroupLifecycleStateEnum = "DELETING"
	PatchGroupLifecycleStateFailed   PatchGroupLifecycleStateEnum = "FAILED"
	PatchGroupLifecycleStateUpdating PatchGroupLifecycleStateEnum = "UPDATING"
)

var mappingPatchGroupLifecycleStateEnum = map[string]PatchGroupLifecycleStateEnum{
	"ACTIVE":   PatchGroupLifecycleStateActive,
	"INACTIVE": PatchGroupLifecycleStateInactive,
	"DELETED":  PatchGroupLifecycleStateDeleted,
	"DELETING": PatchGroupLifecycleStateDeleting,
	"FAILED":   PatchGroupLifecycleStateFailed,
	"UPDATING": PatchGroupLifecycleStateUpdating,
}

var mappingPatchGroupLifecycleStateEnumLowerCase = map[string]PatchGroupLifecycleStateEnum{
	"active":   PatchGroupLifecycleStateActive,
	"inactive": PatchGroupLifecycleStateInactive,
	"deleted":  PatchGroupLifecycleStateDeleted,
	"deleting": PatchGroupLifecycleStateDeleting,
	"failed":   PatchGroupLifecycleStateFailed,
	"updating": PatchGroupLifecycleStateUpdating,
}

// GetPatchGroupLifecycleStateEnumValues Enumerates the set of values for PatchGroupLifecycleStateEnum
func GetPatchGroupLifecycleStateEnumValues() []PatchGroupLifecycleStateEnum {
	values := make([]PatchGroupLifecycleStateEnum, 0)
	for _, v := range mappingPatchGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchGroupLifecycleStateEnumStringValues Enumerates the set of values in String for PatchGroupLifecycleStateEnum
func GetPatchGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingPatchGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchGroupLifecycleStateEnum(val string) (PatchGroupLifecycleStateEnum, bool) {
	enum, ok := mappingPatchGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
