// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OdaInstanceSummary Summary of the Digital Assistant instance.
type OdaInstanceSummary struct {

	// Unique identifier of the Digital Assistant instance.
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment that the instance belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the instance.
	LifecycleState OdaInstanceSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// User-defined name for the Digital Assistant instance. You can change this value.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the Digital Assistant instance.
	Description *string `mandatory:"false" json:"description"`

	// Shape or size of the instance.
	ShapeName OdaInstanceSummaryShapeNameEnum `mandatory:"false" json:"shapeName,omitempty"`

	// When the Digital Assistant instance was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// When the Digital Assistant instance was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current sub-state of the Digital Assistant instance.
	LifecycleSubState OdaInstanceSummaryLifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`

	// A message describing the current state in more detail. For example, actionable
	// information about an instance that's in the `FAILED` state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OdaInstanceSummary) String() string {
	return common.PointerString(m)
}

// OdaInstanceSummaryShapeNameEnum Enum with underlying type: string
type OdaInstanceSummaryShapeNameEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryShapeNameEnum
const (
	OdaInstanceSummaryShapeNameDevelopment OdaInstanceSummaryShapeNameEnum = "DEVELOPMENT"
	OdaInstanceSummaryShapeNameProduction  OdaInstanceSummaryShapeNameEnum = "PRODUCTION"
)

var mappingOdaInstanceSummaryShapeName = map[string]OdaInstanceSummaryShapeNameEnum{
	"DEVELOPMENT": OdaInstanceSummaryShapeNameDevelopment,
	"PRODUCTION":  OdaInstanceSummaryShapeNameProduction,
}

// GetOdaInstanceSummaryShapeNameEnumValues Enumerates the set of values for OdaInstanceSummaryShapeNameEnum
func GetOdaInstanceSummaryShapeNameEnumValues() []OdaInstanceSummaryShapeNameEnum {
	values := make([]OdaInstanceSummaryShapeNameEnum, 0)
	for _, v := range mappingOdaInstanceSummaryShapeName {
		values = append(values, v)
	}
	return values
}

// OdaInstanceSummaryLifecycleStateEnum Enum with underlying type: string
type OdaInstanceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryLifecycleStateEnum
const (
	OdaInstanceSummaryLifecycleStateCreating OdaInstanceSummaryLifecycleStateEnum = "CREATING"
	OdaInstanceSummaryLifecycleStateUpdating OdaInstanceSummaryLifecycleStateEnum = "UPDATING"
	OdaInstanceSummaryLifecycleStateActive   OdaInstanceSummaryLifecycleStateEnum = "ACTIVE"
	OdaInstanceSummaryLifecycleStateInactive OdaInstanceSummaryLifecycleStateEnum = "INACTIVE"
	OdaInstanceSummaryLifecycleStateDeleting OdaInstanceSummaryLifecycleStateEnum = "DELETING"
	OdaInstanceSummaryLifecycleStateDeleted  OdaInstanceSummaryLifecycleStateEnum = "DELETED"
	OdaInstanceSummaryLifecycleStateFailed   OdaInstanceSummaryLifecycleStateEnum = "FAILED"
)

var mappingOdaInstanceSummaryLifecycleState = map[string]OdaInstanceSummaryLifecycleStateEnum{
	"CREATING": OdaInstanceSummaryLifecycleStateCreating,
	"UPDATING": OdaInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   OdaInstanceSummaryLifecycleStateActive,
	"INACTIVE": OdaInstanceSummaryLifecycleStateInactive,
	"DELETING": OdaInstanceSummaryLifecycleStateDeleting,
	"DELETED":  OdaInstanceSummaryLifecycleStateDeleted,
	"FAILED":   OdaInstanceSummaryLifecycleStateFailed,
}

// GetOdaInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for OdaInstanceSummaryLifecycleStateEnum
func GetOdaInstanceSummaryLifecycleStateEnumValues() []OdaInstanceSummaryLifecycleStateEnum {
	values := make([]OdaInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingOdaInstanceSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// OdaInstanceSummaryLifecycleSubStateEnum Enum with underlying type: string
type OdaInstanceSummaryLifecycleSubStateEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryLifecycleSubStateEnum
const (
	OdaInstanceSummaryLifecycleSubStateCreating            OdaInstanceSummaryLifecycleSubStateEnum = "CREATING"
	OdaInstanceSummaryLifecycleSubStateStarting            OdaInstanceSummaryLifecycleSubStateEnum = "STARTING"
	OdaInstanceSummaryLifecycleSubStateStopping            OdaInstanceSummaryLifecycleSubStateEnum = "STOPPING"
	OdaInstanceSummaryLifecycleSubStateChangingCompartment OdaInstanceSummaryLifecycleSubStateEnum = "CHANGING_COMPARTMENT"
	OdaInstanceSummaryLifecycleSubStateDeleting            OdaInstanceSummaryLifecycleSubStateEnum = "DELETING"
	OdaInstanceSummaryLifecycleSubStateDeletePending       OdaInstanceSummaryLifecycleSubStateEnum = "DELETE_PENDING"
	OdaInstanceSummaryLifecycleSubStateRecovering          OdaInstanceSummaryLifecycleSubStateEnum = "RECOVERING"
	OdaInstanceSummaryLifecycleSubStatePurging             OdaInstanceSummaryLifecycleSubStateEnum = "PURGING"
	OdaInstanceSummaryLifecycleSubStateQueued              OdaInstanceSummaryLifecycleSubStateEnum = "QUEUED"
)

var mappingOdaInstanceSummaryLifecycleSubState = map[string]OdaInstanceSummaryLifecycleSubStateEnum{
	"CREATING":             OdaInstanceSummaryLifecycleSubStateCreating,
	"STARTING":             OdaInstanceSummaryLifecycleSubStateStarting,
	"STOPPING":             OdaInstanceSummaryLifecycleSubStateStopping,
	"CHANGING_COMPARTMENT": OdaInstanceSummaryLifecycleSubStateChangingCompartment,
	"DELETING":             OdaInstanceSummaryLifecycleSubStateDeleting,
	"DELETE_PENDING":       OdaInstanceSummaryLifecycleSubStateDeletePending,
	"RECOVERING":           OdaInstanceSummaryLifecycleSubStateRecovering,
	"PURGING":              OdaInstanceSummaryLifecycleSubStatePurging,
	"QUEUED":               OdaInstanceSummaryLifecycleSubStateQueued,
}

// GetOdaInstanceSummaryLifecycleSubStateEnumValues Enumerates the set of values for OdaInstanceSummaryLifecycleSubStateEnum
func GetOdaInstanceSummaryLifecycleSubStateEnumValues() []OdaInstanceSummaryLifecycleSubStateEnum {
	values := make([]OdaInstanceSummaryLifecycleSubStateEnum, 0)
	for _, v := range mappingOdaInstanceSummaryLifecycleSubState {
		values = append(values, v)
	}
	return values
}
