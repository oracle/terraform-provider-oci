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

// PeerRole Peer role
type PeerRole struct {

	// Peer role names
	Role PeerRoleRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m PeerRole) String() string {
	return common.PointerString(m)
}

// PeerRoleRoleEnum Enum with underlying type: string
type PeerRoleRoleEnum string

// Set of constants representing the allowable values for PeerRoleRoleEnum
const (
	PeerRoleRoleMember PeerRoleRoleEnum = "MEMBER"
	PeerRoleRoleAdmin  PeerRoleRoleEnum = "ADMIN"
)

var mappingPeerRoleRole = map[string]PeerRoleRoleEnum{
	"MEMBER": PeerRoleRoleMember,
	"ADMIN":  PeerRoleRoleAdmin,
}

// GetPeerRoleRoleEnumValues Enumerates the set of values for PeerRoleRoleEnum
func GetPeerRoleRoleEnumValues() []PeerRoleRoleEnum {
	values := make([]PeerRoleRoleEnum, 0)
	for _, v := range mappingPeerRoleRole {
		values = append(values, v)
	}
	return values
}
