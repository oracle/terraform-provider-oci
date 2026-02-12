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

// PatchesByConditionSummary Summary of the Patch based on condition.
type PatchesByConditionSummary struct {

	// The OCID of the resource.
	PatchId *string `mandatory:"true" json:"patchId"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the patch
	PatchName *string `mandatory:"true" json:"patchName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	PatchDescription *string `mandatory:"true" json:"patchDescription"`

	// Provide information on who defined the patch.
	// Example: For Custom Patches the value will be USER_DEFINED
	// For Oracle Defined Patches the value will be ORACLE_DEFINED
	Type PatchTypeEnum `mandatory:"true" json:"type"`

	PatchType *PatchTypeDetails `mandatory:"true" json:"patchType"`

	Product *PatchProductDetails `mandatory:"true" json:"product"`

	// Date when the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// Patch Severity.
	Severity PatchSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// The current state of the Patch.
	LifecycleState PatchesByConditionSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m PatchesByConditionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchesByConditionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchesByConditionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchesByConditionSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchesByConditionSummaryLifecycleStateEnum Enum with underlying type: string
type PatchesByConditionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for PatchesByConditionSummaryLifecycleStateEnum
const (
	PatchesByConditionSummaryLifecycleStateActive   PatchesByConditionSummaryLifecycleStateEnum = "ACTIVE"
	PatchesByConditionSummaryLifecycleStateInactive PatchesByConditionSummaryLifecycleStateEnum = "INACTIVE"
	PatchesByConditionSummaryLifecycleStateDeleted  PatchesByConditionSummaryLifecycleStateEnum = "DELETED"
	PatchesByConditionSummaryLifecycleStateDeleting PatchesByConditionSummaryLifecycleStateEnum = "DELETING"
	PatchesByConditionSummaryLifecycleStateFailed   PatchesByConditionSummaryLifecycleStateEnum = "FAILED"
	PatchesByConditionSummaryLifecycleStateUpdating PatchesByConditionSummaryLifecycleStateEnum = "UPDATING"
)

var mappingPatchesByConditionSummaryLifecycleStateEnum = map[string]PatchesByConditionSummaryLifecycleStateEnum{
	"ACTIVE":   PatchesByConditionSummaryLifecycleStateActive,
	"INACTIVE": PatchesByConditionSummaryLifecycleStateInactive,
	"DELETED":  PatchesByConditionSummaryLifecycleStateDeleted,
	"DELETING": PatchesByConditionSummaryLifecycleStateDeleting,
	"FAILED":   PatchesByConditionSummaryLifecycleStateFailed,
	"UPDATING": PatchesByConditionSummaryLifecycleStateUpdating,
}

var mappingPatchesByConditionSummaryLifecycleStateEnumLowerCase = map[string]PatchesByConditionSummaryLifecycleStateEnum{
	"active":   PatchesByConditionSummaryLifecycleStateActive,
	"inactive": PatchesByConditionSummaryLifecycleStateInactive,
	"deleted":  PatchesByConditionSummaryLifecycleStateDeleted,
	"deleting": PatchesByConditionSummaryLifecycleStateDeleting,
	"failed":   PatchesByConditionSummaryLifecycleStateFailed,
	"updating": PatchesByConditionSummaryLifecycleStateUpdating,
}

// GetPatchesByConditionSummaryLifecycleStateEnumValues Enumerates the set of values for PatchesByConditionSummaryLifecycleStateEnum
func GetPatchesByConditionSummaryLifecycleStateEnumValues() []PatchesByConditionSummaryLifecycleStateEnum {
	values := make([]PatchesByConditionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingPatchesByConditionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchesByConditionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for PatchesByConditionSummaryLifecycleStateEnum
func GetPatchesByConditionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingPatchesByConditionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchesByConditionSummaryLifecycleStateEnum(val string) (PatchesByConditionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingPatchesByConditionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
