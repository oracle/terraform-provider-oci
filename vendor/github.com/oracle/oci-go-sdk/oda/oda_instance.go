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

// OdaInstance Description of `OdaServiceInstance` object.
type OdaInstance struct {

	// Unique immutable identifier that was assigned when the instance was created.
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment that the instance belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Shape or size of the instance.
	ShapeName OdaInstanceShapeNameEnum `mandatory:"true" json:"shapeName"`

	// User-defined name for the Digital Assistant instance. Avoid entering confidential information.
	// You can change this value.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the Digital Assistant instance.
	Description *string `mandatory:"false" json:"description"`

	// URL for the Digital Assistant web application that's associated with the instance.
	WebAppUrl *string `mandatory:"false" json:"webAppUrl"`

	// URL for the connector's endpoint.
	ConnectorUrl *string `mandatory:"false" json:"connectorUrl"`

	// When the Digital Assistant instance was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// When the Digital Assistance instance was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Digital Assistant instance.
	LifecycleState OdaInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The current sub-state of the Digital Assistant instance.
	LifecycleSubState OdaInstanceLifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`

	// A message that describes the current state in more detail.
	// For example, actionable information about an instance that's in the `FAILED` state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OdaInstance) String() string {
	return common.PointerString(m)
}

// OdaInstanceShapeNameEnum Enum with underlying type: string
type OdaInstanceShapeNameEnum string

// Set of constants representing the allowable values for OdaInstanceShapeNameEnum
const (
	OdaInstanceShapeNameDevelopment OdaInstanceShapeNameEnum = "DEVELOPMENT"
	OdaInstanceShapeNameProduction  OdaInstanceShapeNameEnum = "PRODUCTION"
)

var mappingOdaInstanceShapeName = map[string]OdaInstanceShapeNameEnum{
	"DEVELOPMENT": OdaInstanceShapeNameDevelopment,
	"PRODUCTION":  OdaInstanceShapeNameProduction,
}

// GetOdaInstanceShapeNameEnumValues Enumerates the set of values for OdaInstanceShapeNameEnum
func GetOdaInstanceShapeNameEnumValues() []OdaInstanceShapeNameEnum {
	values := make([]OdaInstanceShapeNameEnum, 0)
	for _, v := range mappingOdaInstanceShapeName {
		values = append(values, v)
	}
	return values
}

// OdaInstanceLifecycleStateEnum Enum with underlying type: string
type OdaInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for OdaInstanceLifecycleStateEnum
const (
	OdaInstanceLifecycleStateCreating OdaInstanceLifecycleStateEnum = "CREATING"
	OdaInstanceLifecycleStateUpdating OdaInstanceLifecycleStateEnum = "UPDATING"
	OdaInstanceLifecycleStateActive   OdaInstanceLifecycleStateEnum = "ACTIVE"
	OdaInstanceLifecycleStateInactive OdaInstanceLifecycleStateEnum = "INACTIVE"
	OdaInstanceLifecycleStateDeleting OdaInstanceLifecycleStateEnum = "DELETING"
	OdaInstanceLifecycleStateDeleted  OdaInstanceLifecycleStateEnum = "DELETED"
	OdaInstanceLifecycleStateFailed   OdaInstanceLifecycleStateEnum = "FAILED"
)

var mappingOdaInstanceLifecycleState = map[string]OdaInstanceLifecycleStateEnum{
	"CREATING": OdaInstanceLifecycleStateCreating,
	"UPDATING": OdaInstanceLifecycleStateUpdating,
	"ACTIVE":   OdaInstanceLifecycleStateActive,
	"INACTIVE": OdaInstanceLifecycleStateInactive,
	"DELETING": OdaInstanceLifecycleStateDeleting,
	"DELETED":  OdaInstanceLifecycleStateDeleted,
	"FAILED":   OdaInstanceLifecycleStateFailed,
}

// GetOdaInstanceLifecycleStateEnumValues Enumerates the set of values for OdaInstanceLifecycleStateEnum
func GetOdaInstanceLifecycleStateEnumValues() []OdaInstanceLifecycleStateEnum {
	values := make([]OdaInstanceLifecycleStateEnum, 0)
	for _, v := range mappingOdaInstanceLifecycleState {
		values = append(values, v)
	}
	return values
}

// OdaInstanceLifecycleSubStateEnum Enum with underlying type: string
type OdaInstanceLifecycleSubStateEnum string

// Set of constants representing the allowable values for OdaInstanceLifecycleSubStateEnum
const (
	OdaInstanceLifecycleSubStateCreating            OdaInstanceLifecycleSubStateEnum = "CREATING"
	OdaInstanceLifecycleSubStateStarting            OdaInstanceLifecycleSubStateEnum = "STARTING"
	OdaInstanceLifecycleSubStateStopping            OdaInstanceLifecycleSubStateEnum = "STOPPING"
	OdaInstanceLifecycleSubStateChangingCompartment OdaInstanceLifecycleSubStateEnum = "CHANGING_COMPARTMENT"
	OdaInstanceLifecycleSubStateDeleting            OdaInstanceLifecycleSubStateEnum = "DELETING"
	OdaInstanceLifecycleSubStateDeletePending       OdaInstanceLifecycleSubStateEnum = "DELETE_PENDING"
	OdaInstanceLifecycleSubStateRecovering          OdaInstanceLifecycleSubStateEnum = "RECOVERING"
	OdaInstanceLifecycleSubStatePurging             OdaInstanceLifecycleSubStateEnum = "PURGING"
	OdaInstanceLifecycleSubStateQueued              OdaInstanceLifecycleSubStateEnum = "QUEUED"
)

var mappingOdaInstanceLifecycleSubState = map[string]OdaInstanceLifecycleSubStateEnum{
	"CREATING":             OdaInstanceLifecycleSubStateCreating,
	"STARTING":             OdaInstanceLifecycleSubStateStarting,
	"STOPPING":             OdaInstanceLifecycleSubStateStopping,
	"CHANGING_COMPARTMENT": OdaInstanceLifecycleSubStateChangingCompartment,
	"DELETING":             OdaInstanceLifecycleSubStateDeleting,
	"DELETE_PENDING":       OdaInstanceLifecycleSubStateDeletePending,
	"RECOVERING":           OdaInstanceLifecycleSubStateRecovering,
	"PURGING":              OdaInstanceLifecycleSubStatePurging,
	"QUEUED":               OdaInstanceLifecycleSubStateQueued,
}

// GetOdaInstanceLifecycleSubStateEnumValues Enumerates the set of values for OdaInstanceLifecycleSubStateEnum
func GetOdaInstanceLifecycleSubStateEnumValues() []OdaInstanceLifecycleSubStateEnum {
	values := make([]OdaInstanceLifecycleSubStateEnum, 0)
	for _, v := range mappingOdaInstanceLifecycleSubState {
		values = append(values, v)
	}
	return values
}
