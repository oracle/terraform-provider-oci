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

// AssociatedPatchSummary Summary of the Patches within a Patch Group.
type AssociatedPatchSummary struct {

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
	LifecycleState AssociatedPatchSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m AssociatedPatchSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedPatchSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAssociatedPatchSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssociatedPatchSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssociatedPatchSummaryLifecycleStateEnum Enum with underlying type: string
type AssociatedPatchSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AssociatedPatchSummaryLifecycleStateEnum
const (
	AssociatedPatchSummaryLifecycleStateActive   AssociatedPatchSummaryLifecycleStateEnum = "ACTIVE"
	AssociatedPatchSummaryLifecycleStateInactive AssociatedPatchSummaryLifecycleStateEnum = "INACTIVE"
	AssociatedPatchSummaryLifecycleStateDeleted  AssociatedPatchSummaryLifecycleStateEnum = "DELETED"
	AssociatedPatchSummaryLifecycleStateDeleting AssociatedPatchSummaryLifecycleStateEnum = "DELETING"
	AssociatedPatchSummaryLifecycleStateFailed   AssociatedPatchSummaryLifecycleStateEnum = "FAILED"
	AssociatedPatchSummaryLifecycleStateUpdating AssociatedPatchSummaryLifecycleStateEnum = "UPDATING"
)

var mappingAssociatedPatchSummaryLifecycleStateEnum = map[string]AssociatedPatchSummaryLifecycleStateEnum{
	"ACTIVE":   AssociatedPatchSummaryLifecycleStateActive,
	"INACTIVE": AssociatedPatchSummaryLifecycleStateInactive,
	"DELETED":  AssociatedPatchSummaryLifecycleStateDeleted,
	"DELETING": AssociatedPatchSummaryLifecycleStateDeleting,
	"FAILED":   AssociatedPatchSummaryLifecycleStateFailed,
	"UPDATING": AssociatedPatchSummaryLifecycleStateUpdating,
}

var mappingAssociatedPatchSummaryLifecycleStateEnumLowerCase = map[string]AssociatedPatchSummaryLifecycleStateEnum{
	"active":   AssociatedPatchSummaryLifecycleStateActive,
	"inactive": AssociatedPatchSummaryLifecycleStateInactive,
	"deleted":  AssociatedPatchSummaryLifecycleStateDeleted,
	"deleting": AssociatedPatchSummaryLifecycleStateDeleting,
	"failed":   AssociatedPatchSummaryLifecycleStateFailed,
	"updating": AssociatedPatchSummaryLifecycleStateUpdating,
}

// GetAssociatedPatchSummaryLifecycleStateEnumValues Enumerates the set of values for AssociatedPatchSummaryLifecycleStateEnum
func GetAssociatedPatchSummaryLifecycleStateEnumValues() []AssociatedPatchSummaryLifecycleStateEnum {
	values := make([]AssociatedPatchSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAssociatedPatchSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociatedPatchSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AssociatedPatchSummaryLifecycleStateEnum
func GetAssociatedPatchSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingAssociatedPatchSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociatedPatchSummaryLifecycleStateEnum(val string) (AssociatedPatchSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAssociatedPatchSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
