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

// ComputeCapacityReservationSummary Summary information for an compute capacity reservation.
type ComputeCapacityReservationSummary struct {

	// The OCID of the instance reservation configuration.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the capacity reservation was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name for the capacity reservation.Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My Reservation`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the capacity reservation.
	LifecycleState ComputeCapacityReservationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The availability domain the capacity reservation is present in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The total number of instances that could reserved in this
	// compute capacity reservation. This number is the sum of reservedCount
	// of all the instance reservation config under this reservation.
	// This exists to calulate the percentage usage of the reservation.
	ReservedInstanceCount *int64 `mandatory:"false" json:"reservedInstanceCount"`

	// The total number of instances currently running using
	// this compute capacity reservation. This number is the sum of usedCount
	// of all the instance reservation config under this reservation.
	// This exists to calulate the percentage usage of the reservation.
	UsedInstanceCount *int64 `mandatory:"false" json:"usedInstanceCount"`

	// This shows whether this capacity reservation is default or not.
	IsDefaultReservation *bool `mandatory:"false" json:"isDefaultReservation"`
}

func (m ComputeCapacityReservationSummary) String() string {
	return common.PointerString(m)
}
