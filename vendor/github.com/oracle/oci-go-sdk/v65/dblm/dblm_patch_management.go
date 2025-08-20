// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DblmPatchManagement Description of PatchManagement.
type DblmPatchManagement struct {

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// resources objects
	Resources []ResourceInfo `mandatory:"false" json:"resources"`

	// A message describing the status of the feature's state
	Message *string `mandatory:"false" json:"message"`

	// The time the Vulnerability was enabled. An RFC3339 formatted datetime string.
	TimeEnabled *common.SDKTime `mandatory:"false" json:"timeEnabled"`

	// The current state of the feature.
	LifecycleState DblmPatchManagementLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Summary of patch operations.
	PatchOperationsSummary *interface{} `mandatory:"false" json:"patchOperationsSummary"`

	// Summary of image patches recommended to install.
	ImagesPatchRecommendationSummary *interface{} `mandatory:"false" json:"imagesPatchRecommendationSummary"`

	// Summary of image patches to be compliant to install.
	ResourcesPatchComplianceSummary *interface{} `mandatory:"false" json:"resourcesPatchComplianceSummary"`

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

func (m DblmPatchManagement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DblmPatchManagement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDblmPatchManagementLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDblmPatchManagementLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DblmPatchManagementLifecycleStateEnum Enum with underlying type: string
type DblmPatchManagementLifecycleStateEnum string

// Set of constants representing the allowable values for DblmPatchManagementLifecycleStateEnum
const (
	DblmPatchManagementLifecycleStateCreating       DblmPatchManagementLifecycleStateEnum = "CREATING"
	DblmPatchManagementLifecycleStateActive         DblmPatchManagementLifecycleStateEnum = "ACTIVE"
	DblmPatchManagementLifecycleStateFailed         DblmPatchManagementLifecycleStateEnum = "FAILED"
	DblmPatchManagementLifecycleStateNeedsAttention DblmPatchManagementLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDblmPatchManagementLifecycleStateEnum = map[string]DblmPatchManagementLifecycleStateEnum{
	"CREATING":        DblmPatchManagementLifecycleStateCreating,
	"ACTIVE":          DblmPatchManagementLifecycleStateActive,
	"FAILED":          DblmPatchManagementLifecycleStateFailed,
	"NEEDS_ATTENTION": DblmPatchManagementLifecycleStateNeedsAttention,
}

var mappingDblmPatchManagementLifecycleStateEnumLowerCase = map[string]DblmPatchManagementLifecycleStateEnum{
	"creating":        DblmPatchManagementLifecycleStateCreating,
	"active":          DblmPatchManagementLifecycleStateActive,
	"failed":          DblmPatchManagementLifecycleStateFailed,
	"needs_attention": DblmPatchManagementLifecycleStateNeedsAttention,
}

// GetDblmPatchManagementLifecycleStateEnumValues Enumerates the set of values for DblmPatchManagementLifecycleStateEnum
func GetDblmPatchManagementLifecycleStateEnumValues() []DblmPatchManagementLifecycleStateEnum {
	values := make([]DblmPatchManagementLifecycleStateEnum, 0)
	for _, v := range mappingDblmPatchManagementLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDblmPatchManagementLifecycleStateEnumStringValues Enumerates the set of values in String for DblmPatchManagementLifecycleStateEnum
func GetDblmPatchManagementLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDblmPatchManagementLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDblmPatchManagementLifecycleStateEnum(val string) (DblmPatchManagementLifecycleStateEnum, bool) {
	enum, ok := mappingDblmPatchManagementLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
