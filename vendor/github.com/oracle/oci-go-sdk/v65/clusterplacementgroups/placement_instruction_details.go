// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cluster Placement Groups API
//
// API for managing cluster placement groups.
//

package clusterplacementgroups

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PlacementInstructionDetails Details that inform cluster placement group provisioning.
type PlacementInstructionDetails struct {

	// The type of placement instruction.
	Type PlacementInstructionDetailsTypeEnum `mandatory:"true" json:"type"`

	// The value of the token designated for placement of the cluster placement group upon creation.
	Value *string `mandatory:"true" json:"value"`
}

func (m PlacementInstructionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PlacementInstructionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPlacementInstructionDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPlacementInstructionDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PlacementInstructionDetailsTypeEnum Enum with underlying type: string
type PlacementInstructionDetailsTypeEnum string

// Set of constants representing the allowable values for PlacementInstructionDetailsTypeEnum
const (
	PlacementInstructionDetailsTypeToken PlacementInstructionDetailsTypeEnum = "TOKEN"
)

var mappingPlacementInstructionDetailsTypeEnum = map[string]PlacementInstructionDetailsTypeEnum{
	"TOKEN": PlacementInstructionDetailsTypeToken,
}

var mappingPlacementInstructionDetailsTypeEnumLowerCase = map[string]PlacementInstructionDetailsTypeEnum{
	"token": PlacementInstructionDetailsTypeToken,
}

// GetPlacementInstructionDetailsTypeEnumValues Enumerates the set of values for PlacementInstructionDetailsTypeEnum
func GetPlacementInstructionDetailsTypeEnumValues() []PlacementInstructionDetailsTypeEnum {
	values := make([]PlacementInstructionDetailsTypeEnum, 0)
	for _, v := range mappingPlacementInstructionDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPlacementInstructionDetailsTypeEnumStringValues Enumerates the set of values in String for PlacementInstructionDetailsTypeEnum
func GetPlacementInstructionDetailsTypeEnumStringValues() []string {
	return []string{
		"TOKEN",
	}
}

// GetMappingPlacementInstructionDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPlacementInstructionDetailsTypeEnum(val string) (PlacementInstructionDetailsTypeEnum, bool) {
	enum, ok := mappingPlacementInstructionDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
