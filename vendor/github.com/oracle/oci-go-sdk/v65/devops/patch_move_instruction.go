// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchMoveInstruction An operation that "puts" values from elsewhere in the target, functionally equivalent to a single add and then a remove.
// The first item of the selection is replaced, or created if the selection is empty.
// NOT_FOUND exceptions in the selection are handled by creating the implied containing structure.
// This operation fails if the `from` selection yields any exceptions, or if an item is moved to any of its descendants.
type PatchMoveInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	Selection *string `mandatory:"true" json:"selection"`

	// The selection that is to be moved, with the same format and semantics as `selection`.
	From *string `mandatory:"true" json:"from"`

	// Where to insert the value in an array, relative to the first item in the selection.
	// If there is no such item, then "BEFORE" specifies insertion at the first position in an array and "AFTER" specifies insertion at the last position.
	// If the first item in the selection is not the child of an array, then this field has no effect.
	Position PatchMoveInstructionPositionEnum `mandatory:"false" json:"position,omitempty"`
}

// GetSelection returns Selection
func (m PatchMoveInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchMoveInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchMoveInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchMoveInstructionPositionEnum(string(m.Position)); !ok && m.Position != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Position: %s. Supported values are: %s.", m.Position, strings.Join(GetPatchMoveInstructionPositionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchMoveInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchMoveInstruction PatchMoveInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchMoveInstruction
	}{
		"MOVE",
		(MarshalTypePatchMoveInstruction)(m),
	}

	return json.Marshal(&s)
}

// PatchMoveInstructionPositionEnum Enum with underlying type: string
type PatchMoveInstructionPositionEnum string

// Set of constants representing the allowable values for PatchMoveInstructionPositionEnum
const (
	PatchMoveInstructionPositionAt     PatchMoveInstructionPositionEnum = "AT"
	PatchMoveInstructionPositionBefore PatchMoveInstructionPositionEnum = "BEFORE"
	PatchMoveInstructionPositionAfter  PatchMoveInstructionPositionEnum = "AFTER"
)

var mappingPatchMoveInstructionPositionEnum = map[string]PatchMoveInstructionPositionEnum{
	"AT":     PatchMoveInstructionPositionAt,
	"BEFORE": PatchMoveInstructionPositionBefore,
	"AFTER":  PatchMoveInstructionPositionAfter,
}

var mappingPatchMoveInstructionPositionEnumLowerCase = map[string]PatchMoveInstructionPositionEnum{
	"at":     PatchMoveInstructionPositionAt,
	"before": PatchMoveInstructionPositionBefore,
	"after":  PatchMoveInstructionPositionAfter,
}

// GetPatchMoveInstructionPositionEnumValues Enumerates the set of values for PatchMoveInstructionPositionEnum
func GetPatchMoveInstructionPositionEnumValues() []PatchMoveInstructionPositionEnum {
	values := make([]PatchMoveInstructionPositionEnum, 0)
	for _, v := range mappingPatchMoveInstructionPositionEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchMoveInstructionPositionEnumStringValues Enumerates the set of values in String for PatchMoveInstructionPositionEnum
func GetPatchMoveInstructionPositionEnumStringValues() []string {
	return []string{
		"AT",
		"BEFORE",
		"AFTER",
	}
}

// GetMappingPatchMoveInstructionPositionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchMoveInstructionPositionEnum(val string) (PatchMoveInstructionPositionEnum, bool) {
	enum, ok := mappingPatchMoveInstructionPositionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
