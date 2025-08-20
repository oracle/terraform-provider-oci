// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpAnycastPatchInsertInstruction An operation that inserts a value into an array, shifting array items as necessary and handling NOT_FOUND exceptions by creating the implied containing structure.
type IpAnycastPatchInsertInstruction struct {

	// A value to be inserted into the target.
	Value *interface{} `mandatory:"true" json:"value"`

	// The set of value to be passed for each operation
	Selection *string `mandatory:"false" json:"selection"`

	// Check if the value of byoipId matches with the byoipRanges present
	SelectedItem *string `mandatory:"false" json:"selectedItem"`

	// Where to insert the value, relative to the first item matched by byoipRangeId. If byoipRangeId is unspecified, then "BEFORE" specifies insertion at the first position in an array and "AFTER" specifies insertion at the last position. If byoipRangeId is specified but results in an empty selection, then both values specify insertion at the last position.
	Position IpAnycastPatchInsertInstructionPositionEnum `mandatory:"false" json:"position,omitempty"`
}

// GetSelection returns Selection
func (m IpAnycastPatchInsertInstruction) GetSelection() *string {
	return m.Selection
}

func (m IpAnycastPatchInsertInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpAnycastPatchInsertInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIpAnycastPatchInsertInstructionPositionEnum(string(m.Position)); !ok && m.Position != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Position: %s. Supported values are: %s.", m.Position, strings.Join(GetIpAnycastPatchInsertInstructionPositionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IpAnycastPatchInsertInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIpAnycastPatchInsertInstruction IpAnycastPatchInsertInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeIpAnycastPatchInsertInstruction
	}{
		"INSERT",
		(MarshalTypeIpAnycastPatchInsertInstruction)(m),
	}

	return json.Marshal(&s)
}

// IpAnycastPatchInsertInstructionPositionEnum Enum with underlying type: string
type IpAnycastPatchInsertInstructionPositionEnum string

// Set of constants representing the allowable values for IpAnycastPatchInsertInstructionPositionEnum
const (
	IpAnycastPatchInsertInstructionPositionBefore IpAnycastPatchInsertInstructionPositionEnum = "BEFORE"
	IpAnycastPatchInsertInstructionPositionAfter  IpAnycastPatchInsertInstructionPositionEnum = "AFTER"
)

var mappingIpAnycastPatchInsertInstructionPositionEnum = map[string]IpAnycastPatchInsertInstructionPositionEnum{
	"BEFORE": IpAnycastPatchInsertInstructionPositionBefore,
	"AFTER":  IpAnycastPatchInsertInstructionPositionAfter,
}

var mappingIpAnycastPatchInsertInstructionPositionEnumLowerCase = map[string]IpAnycastPatchInsertInstructionPositionEnum{
	"before": IpAnycastPatchInsertInstructionPositionBefore,
	"after":  IpAnycastPatchInsertInstructionPositionAfter,
}

// GetIpAnycastPatchInsertInstructionPositionEnumValues Enumerates the set of values for IpAnycastPatchInsertInstructionPositionEnum
func GetIpAnycastPatchInsertInstructionPositionEnumValues() []IpAnycastPatchInsertInstructionPositionEnum {
	values := make([]IpAnycastPatchInsertInstructionPositionEnum, 0)
	for _, v := range mappingIpAnycastPatchInsertInstructionPositionEnum {
		values = append(values, v)
	}
	return values
}

// GetIpAnycastPatchInsertInstructionPositionEnumStringValues Enumerates the set of values in String for IpAnycastPatchInsertInstructionPositionEnum
func GetIpAnycastPatchInsertInstructionPositionEnumStringValues() []string {
	return []string{
		"BEFORE",
		"AFTER",
	}
}

// GetMappingIpAnycastPatchInsertInstructionPositionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIpAnycastPatchInsertInstructionPositionEnum(val string) (IpAnycastPatchInsertInstructionPositionEnum, bool) {
	enum, ok := mappingIpAnycastPatchInsertInstructionPositionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
