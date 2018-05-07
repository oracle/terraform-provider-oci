// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
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
