// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Demand Signal API
//
// Use the OCI Control Center Demand Signal API to manage Demand Signals.
//

package demandsignal

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchInsertInstruction An operation that inserts a value into an array, shifting array items as necessary and handling NOT_FOUND exceptions by creating the implied containing structure.
type PatchInsertInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	Selection *string `mandatory:"true" json:"selection"`

	// A value to be inserted into the target.
	Value *interface{} `mandatory:"true" json:"value"`

	// A selection to be evaluated against the array for identifying a particular reference item within it, with the same format and semantics as `selection`.
	SelectedItem *string `mandatory:"false" json:"selectedItem"`

	// Where to insert the value, relative to the first item matched by `selectedItem`.
	// If `selectedItem` is unspecified, then "BEFORE" specifies insertion at the first position in an array and "AFTER" specifies insertion at the last position.
	// If `selectedItem` is specified but results in an empty selection, then both values specify insertion at the last position.
	Position PatchInsertInstructionPositionEnum `mandatory:"false" json:"position,omitempty"`
}

// GetSelection returns Selection
func (m PatchInsertInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchInsertInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchInsertInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchInsertInstructionPositionEnum(string(m.Position)); !ok && m.Position != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Position: %s. Supported values are: %s.", m.Position, strings.Join(GetPatchInsertInstructionPositionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchInsertInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchInsertInstruction PatchInsertInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchInsertInstruction
	}{
		"INSERT",
		(MarshalTypePatchInsertInstruction)(m),
	}

	return json.Marshal(&s)
}

// PatchInsertInstructionPositionEnum Enum with underlying type: string
type PatchInsertInstructionPositionEnum string

// Set of constants representing the allowable values for PatchInsertInstructionPositionEnum
const (
	PatchInsertInstructionPositionBefore PatchInsertInstructionPositionEnum = "BEFORE"
	PatchInsertInstructionPositionAfter  PatchInsertInstructionPositionEnum = "AFTER"
)

var mappingPatchInsertInstructionPositionEnum = map[string]PatchInsertInstructionPositionEnum{
	"BEFORE": PatchInsertInstructionPositionBefore,
	"AFTER":  PatchInsertInstructionPositionAfter,
}

var mappingPatchInsertInstructionPositionEnumLowerCase = map[string]PatchInsertInstructionPositionEnum{
	"before": PatchInsertInstructionPositionBefore,
	"after":  PatchInsertInstructionPositionAfter,
}

// GetPatchInsertInstructionPositionEnumValues Enumerates the set of values for PatchInsertInstructionPositionEnum
func GetPatchInsertInstructionPositionEnumValues() []PatchInsertInstructionPositionEnum {
	values := make([]PatchInsertInstructionPositionEnum, 0)
	for _, v := range mappingPatchInsertInstructionPositionEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchInsertInstructionPositionEnumStringValues Enumerates the set of values in String for PatchInsertInstructionPositionEnum
func GetPatchInsertInstructionPositionEnumStringValues() []string {
	return []string{
		"BEFORE",
		"AFTER",
	}
}

// GetMappingPatchInsertInstructionPositionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchInsertInstructionPositionEnum(val string) (PatchInsertInstructionPositionEnum, bool) {
	enum, ok := mappingPatchInsertInstructionPositionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
