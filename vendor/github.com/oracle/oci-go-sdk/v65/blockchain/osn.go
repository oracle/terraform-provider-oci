// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Osn) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAvailabilityDomainAdsEnum(string(m.Ad)); !ok && m.Ad != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Ad: %s. Supported values are: %s.", m.Ad, strings.Join(GetAvailabilityDomainAdsEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOsnLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOsnLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OsnLifecycleStateEnum Enum with underlying type: string
type OsnLifecycleStateEnum string

// Set of constants representing the allowable values for OsnLifecycleStateEnum
const (
	OsnLifecycleStateActive   OsnLifecycleStateEnum = "ACTIVE"
	OsnLifecycleStateInactive OsnLifecycleStateEnum = "INACTIVE"
	OsnLifecycleStateFailed   OsnLifecycleStateEnum = "FAILED"
)

var mappingOsnLifecycleStateEnum = map[string]OsnLifecycleStateEnum{
	"ACTIVE":   OsnLifecycleStateActive,
	"INACTIVE": OsnLifecycleStateInactive,
	"FAILED":   OsnLifecycleStateFailed,
}

var mappingOsnLifecycleStateEnumLowerCase = map[string]OsnLifecycleStateEnum{
	"active":   OsnLifecycleStateActive,
	"inactive": OsnLifecycleStateInactive,
	"failed":   OsnLifecycleStateFailed,
}

// GetOsnLifecycleStateEnumValues Enumerates the set of values for OsnLifecycleStateEnum
func GetOsnLifecycleStateEnumValues() []OsnLifecycleStateEnum {
	values := make([]OsnLifecycleStateEnum, 0)
	for _, v := range mappingOsnLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOsnLifecycleStateEnumStringValues Enumerates the set of values in String for OsnLifecycleStateEnum
func GetOsnLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"FAILED",
	}
}

// GetMappingOsnLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsnLifecycleStateEnum(val string) (OsnLifecycleStateEnum, bool) {
	enum, ok := mappingOsnLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
