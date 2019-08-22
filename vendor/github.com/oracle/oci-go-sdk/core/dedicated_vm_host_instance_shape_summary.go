// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
)

// DedicatedVmHostInstanceShapeSummary The shape used to launch instances associated with the dedicated VM host.
type DedicatedVmHostInstanceShapeSummary struct {

	// The name of the virtual machine instance shapes that can be launched on a dedicated VM host.
	InstanceShapeName *string `mandatory:"true" json:"instanceShapeName"`

	// The shape's availability domain.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`
}

func (m DedicatedVmHostInstanceShapeSummary) String() string {
	return common.PointerString(m)
}
