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

// Device Device Path corresponding to the block devices attached to instances having a name and isAvailable flag.
type Device struct {

	// The device name.
	Name *string `mandatory:"true" json:"name"`

	// The flag denoting whether device is available.
	IsAvailable *bool `mandatory:"true" json:"isAvailable"`
}

func (m Device) String() string {
	return common.PointerString(m)
}
