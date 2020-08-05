// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Osn An Ordering Service Node details
type Osn struct {

	// OSN identifier
	OsnKey *string `mandatory:"true" json:"osnKey"`

	// Availability Domain of OSN
	Ad AvailabilityDomainAdsEnum `mandatory:"true" json:"ad"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"false" json:"ocpuAllocationParam"`

	// The current state of the OSN.
	LifecycleState OsnLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m Osn) String() string {
	return common.PointerString(m)
}

// OsnLifecycleStateEnum Enum with underlying type: string
type OsnLifecycleStateEnum string

// Set of constants representing the allowable values for OsnLifecycleStateEnum
const (
	OsnLifecycleStateActive   OsnLifecycleStateEnum = "ACTIVE"
	OsnLifecycleStateInactive OsnLifecycleStateEnum = "INACTIVE"
	OsnLifecycleStateFailed   OsnLifecycleStateEnum = "FAILED"
)

var mappingOsnLifecycleState = map[string]OsnLifecycleStateEnum{
	"ACTIVE":   OsnLifecycleStateActive,
	"INACTIVE": OsnLifecycleStateInactive,
	"FAILED":   OsnLifecycleStateFailed,
}

// GetOsnLifecycleStateEnumValues Enumerates the set of values for OsnLifecycleStateEnum
func GetOsnLifecycleStateEnumValues() []OsnLifecycleStateEnum {
	values := make([]OsnLifecycleStateEnum, 0)
	for _, v := range mappingOsnLifecycleState {
		values = append(values, v)
	}
	return values
}
