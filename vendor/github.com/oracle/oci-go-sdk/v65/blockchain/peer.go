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

// Peer A Peer details
type Peer struct {

	// peer identifier
	PeerKey *string `mandatory:"true" json:"peerKey"`

	// Peer role
	Role PeerRoleRoleEnum `mandatory:"true" json:"role"`

	// Host on which the Peer exists
	Host *string `mandatory:"true" json:"host"`

	// Availability Domain of peer
	Ad AvailabilityDomainAdsEnum `mandatory:"true" json:"ad"`

	// peer alias
	Alias *string `mandatory:"false" json:"alias"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"false" json:"ocpuAllocationParam"`

	// The current state of the peer.
	LifecycleState PeerLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m Peer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Peer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPeerRoleRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetPeerRoleRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAvailabilityDomainAdsEnum(string(m.Ad)); !ok && m.Ad != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Ad: %s. Supported values are: %s.", m.Ad, strings.Join(GetAvailabilityDomainAdsEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPeerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPeerLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PeerLifecycleStateEnum Enum with underlying type: string
type PeerLifecycleStateEnum string

// Set of constants representing the allowable values for PeerLifecycleStateEnum
const (
	PeerLifecycleStateActive   PeerLifecycleStateEnum = "ACTIVE"
	PeerLifecycleStateInactive PeerLifecycleStateEnum = "INACTIVE"
	PeerLifecycleStateFailed   PeerLifecycleStateEnum = "FAILED"
)

var mappingPeerLifecycleStateEnum = map[string]PeerLifecycleStateEnum{
	"ACTIVE":   PeerLifecycleStateActive,
	"INACTIVE": PeerLifecycleStateInactive,
	"FAILED":   PeerLifecycleStateFailed,
}

var mappingPeerLifecycleStateEnumLowerCase = map[string]PeerLifecycleStateEnum{
	"active":   PeerLifecycleStateActive,
	"inactive": PeerLifecycleStateInactive,
	"failed":   PeerLifecycleStateFailed,
}

// GetPeerLifecycleStateEnumValues Enumerates the set of values for PeerLifecycleStateEnum
func GetPeerLifecycleStateEnumValues() []PeerLifecycleStateEnum {
	values := make([]PeerLifecycleStateEnum, 0)
	for _, v := range mappingPeerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPeerLifecycleStateEnumStringValues Enumerates the set of values in String for PeerLifecycleStateEnum
func GetPeerLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"FAILED",
	}
}

// GetMappingPeerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPeerLifecycleStateEnum(val string) (PeerLifecycleStateEnum, bool) {
	enum, ok := mappingPeerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
