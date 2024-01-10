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

// PeerRole Peer role
type PeerRole struct {

	// Peer role names
	Role PeerRoleRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m PeerRole) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeerRole) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPeerRoleRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetPeerRoleRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PeerRoleRoleEnum Enum with underlying type: string
type PeerRoleRoleEnum string

// Set of constants representing the allowable values for PeerRoleRoleEnum
const (
	PeerRoleRoleMember PeerRoleRoleEnum = "MEMBER"
	PeerRoleRoleAdmin  PeerRoleRoleEnum = "ADMIN"
)

var mappingPeerRoleRoleEnum = map[string]PeerRoleRoleEnum{
	"MEMBER": PeerRoleRoleMember,
	"ADMIN":  PeerRoleRoleAdmin,
}

var mappingPeerRoleRoleEnumLowerCase = map[string]PeerRoleRoleEnum{
	"member": PeerRoleRoleMember,
	"admin":  PeerRoleRoleAdmin,
}

// GetPeerRoleRoleEnumValues Enumerates the set of values for PeerRoleRoleEnum
func GetPeerRoleRoleEnumValues() []PeerRoleRoleEnum {
	values := make([]PeerRoleRoleEnum, 0)
	for _, v := range mappingPeerRoleRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetPeerRoleRoleEnumStringValues Enumerates the set of values in String for PeerRoleRoleEnum
func GetPeerRoleRoleEnumStringValues() []string {
	return []string{
		"MEMBER",
		"ADMIN",
	}
}

// GetMappingPeerRoleRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPeerRoleRoleEnum(val string) (PeerRoleRoleEnum, bool) {
	enum, ok := mappingPeerRoleRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
