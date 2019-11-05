// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Capacity Service instance capacity metadata (e.g.: OLPU count, number of users, ...etc...).
type Capacity struct {

	// The capacity model to use.
	CapacityType CapacityTypeEnum `mandatory:"true" json:"capacityType"`

	// The capacity value selected (OLPU count, number of users, ...etc...). This parameter affects the
	// number of CPUs, amount of memory or other resources allocated to the instance.
	CapacityValue *int `mandatory:"true" json:"capacityValue"`
}

func (m Capacity) String() string {
	return common.PointerString(m)
}
