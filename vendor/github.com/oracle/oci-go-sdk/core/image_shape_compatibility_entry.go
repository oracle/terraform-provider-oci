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

// ImageShapeCompatibilityEntry An image and shape that are compatible.
type ImageShapeCompatibilityEntry struct {

	// The image OCID.
	ImageId *string `mandatory:"true" json:"imageId"`

	// The shape name.
	Shape *string `mandatory:"true" json:"shape"`
}

func (m ImageShapeCompatibilityEntry) String() string {
	return common.PointerString(m)
}
