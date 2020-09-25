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

// InstanceReservationConfig The data to define instance reservation config.
type InstanceReservationConfig struct {

	// The shape which the customer wants to reserve. The shape determines the number of CPUs, amount of memory,
	// and other resources allocated to the instance.
	// You can enumerate all available shapes by calling ListComputeCapacityReservationInstanceShapes.
	InstanceShape *string `mandatory:"true" json:"instanceShape"`

	// The number of instances the customer wants to reserve under this reservation config.
	ReservedCount *int64 `mandatory:"true" json:"reservedCount"`

	// The number of instances the customer is using in this reservation out of the reservedCount.
	UsedCount *int64 `mandatory:"true" json:"usedCount"`

	// The fault domain the reservation config is intented for. If not supplied this config is applicable to all fault domains in the specified AD.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	InstanceShapeConfig *InstanceReservationShapeConfigDetails `mandatory:"false" json:"instanceShapeConfig"`
}

func (m InstanceReservationConfig) String() string {
	return common.PointerString(m)
}
