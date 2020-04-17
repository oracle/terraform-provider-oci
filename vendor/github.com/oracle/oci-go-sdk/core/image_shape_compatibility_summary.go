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

// ImageShapeCompatibilitySummary Summary information for an image shape compatibility entry.
type ImageShapeCompatibilitySummary struct {

	// The image OCID.
	ImageId *string `mandatory:"true" json:"imageId"`

	// The shape name.
	Shape *string `mandatory:"true" json:"shape"`

	OcpuConstraints *ImageOcpuConstraints `mandatory:"false" json:"ocpuConstraints"`
}

func (m ImageShapeCompatibilitySummary) String() string {
	return common.PointerString(m)
}
