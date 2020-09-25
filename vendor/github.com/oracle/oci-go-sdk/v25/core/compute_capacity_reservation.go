// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// ComputeCapacityReservation A capacity reservation is a template that defines the settings to use when creating capacity reservations.
type ComputeCapacityReservation struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment
	// containing the compute capacity reservation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity reservation.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the compute capacity reservation.
	LifecycleState ComputeCapacityReservationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the compute capacity reservation was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The availability domain where this compute capacity reservation lives.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the compute capacity reservation.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// This shows whether this compute capacity reservation is default or not.
	IsDefaultReservation *bool `mandatory:"false" json:"isDefaultReservation"`

	// The reservation configurations for the capacity reservation.
	// To use the reservation for the desired shape, specify the shape and count and
	// optional fault domain where you want this configuration.
	InstanceReservationConfigs []InstanceReservationConfig `mandatory:"false" json:"instanceReservationConfigs"`

	// The total number of instances that could reserved in this
	// compute capacity reservation. This number is the sum of reservedCount
	// of all the instance reservation config under this reservation.
	// This exists to calulate the percentage usage of the reservation.
	ReservedInstanceCount *int64 `mandatory:"false" json:"reservedInstanceCount"`

	// The date and time the compute capacity reservation was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The total number of instances currently running using
	// this compute capacity reservation. This number is the sum of usedCount
	// of all the instance reservation config under this reservation.
	// This exists to calulate the percentage usage of the reservation.
	UsedInstanceCount *int64 `mandatory:"false" json:"usedInstanceCount"`
}

func (m ComputeCapacityReservation) String() string {
	return common.PointerString(m)
}

// ComputeCapacityReservationLifecycleStateEnum Enum with underlying type: string
type ComputeCapacityReservationLifecycleStateEnum string

// Set of constants representing the allowable values for ComputeCapacityReservationLifecycleStateEnum
const (
	ComputeCapacityReservationLifecycleStateActive   ComputeCapacityReservationLifecycleStateEnum = "ACTIVE"
	ComputeCapacityReservationLifecycleStateCreating ComputeCapacityReservationLifecycleStateEnum = "CREATING"
	ComputeCapacityReservationLifecycleStateUpdating ComputeCapacityReservationLifecycleStateEnum = "UPDATING"
	ComputeCapacityReservationLifecycleStateMoving   ComputeCapacityReservationLifecycleStateEnum = "MOVING"
	ComputeCapacityReservationLifecycleStateDeleted  ComputeCapacityReservationLifecycleStateEnum = "DELETED"
	ComputeCapacityReservationLifecycleStateDeleting ComputeCapacityReservationLifecycleStateEnum = "DELETING"
)

var mappingComputeCapacityReservationLifecycleState = map[string]ComputeCapacityReservationLifecycleStateEnum{
	"ACTIVE":   ComputeCapacityReservationLifecycleStateActive,
	"CREATING": ComputeCapacityReservationLifecycleStateCreating,
	"UPDATING": ComputeCapacityReservationLifecycleStateUpdating,
	"MOVING":   ComputeCapacityReservationLifecycleStateMoving,
	"DELETED":  ComputeCapacityReservationLifecycleStateDeleted,
	"DELETING": ComputeCapacityReservationLifecycleStateDeleting,
}

// GetComputeCapacityReservationLifecycleStateEnumValues Enumerates the set of values for ComputeCapacityReservationLifecycleStateEnum
func GetComputeCapacityReservationLifecycleStateEnumValues() []ComputeCapacityReservationLifecycleStateEnum {
	values := make([]ComputeCapacityReservationLifecycleStateEnum, 0)
	for _, v := range mappingComputeCapacityReservationLifecycleState {
		values = append(values, v)
	}
	return values
}
