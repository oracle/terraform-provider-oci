// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Capacity Capacity boundaries for the pool
type Capacity struct {

	// The maximum size the pool is allowed to increase to
	Max *int `mandatory:"true" json:"max"`

	// The minimum size the pool is allowed to decrease to
	Min *int `mandatory:"true" json:"min"`

	// The initial size of the pool
	Initial *int `mandatory:"true" json:"initial"`
}

func (m Capacity) String() string {
	return common.PointerString(m)
}
