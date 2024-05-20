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

// PatchInsertMultipleInstruction An operation that inserts multiple consecutive values into an array, shifting array items as necessary and handling NOT_FOUND exceptions by creating the implied containing structure.
type PatchInsertMultipleInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	Selection *string `mandatory:"true" json:"selection"`

	// A list of consecutive values to be inserted into the target.
	Values []interface{} `mandatory:"true" json:"values"`

	// A selection to be evaluated against the array for identifying a particular reference item within it, with the same format and semantics as `selection`.
	SelectedItem *string `mandatory:"false" json:"selectedItem"`

	// Where to insert the values, relative to the first item matched by `selectedItem`.
	// If `selectedItem` is unspecified, then "BEFORE" specifies insertion at the first position in an array and "AFTER" specifies insertion at the last position.
	// If `selectedItem` is specified but results in an empty selection, then both values specify insertion at the last position.
	Position PatchInsertMultipleInstructionPositionEnum `mandatory:"false" json:"position,omitempty"`
}

// GetSelection returns Selection
func (m PatchInsertMultipleInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchInsertMultipleInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchInsertMultipleInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchInsertMultipleInstructionPositionEnum(string(m.Position)); !ok && m.Position != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Position: %s. Supported values are: %s.", m.Position, strings.Join(GetPatchInsertMultipleInstructionPositionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchInsertMultipleInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchInsertMultipleInstruction PatchInsertMultipleInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchInsertMultipleInstruction
	}{
		"INSERT_MULTIPLE",
		(MarshalTypePatchInsertMultipleInstruction)(m),
	}

	return json.Marshal(&s)
}

// PatchInsertMultipleInstructionPositionEnum Enum with underlying type: string
type PatchInsertMultipleInstructionPositionEnum string

// Set of constants representing the allowable values for PatchInsertMultipleInstructionPositionEnum
const (
	PatchInsertMultipleInstructionPositionBefore PatchInsertMultipleInstructionPositionEnum = "BEFORE"
	PatchInsertMultipleInstructionPositionAfter  PatchInsertMultipleInstructionPositionEnum = "AFTER"
)

var mappingPatchInsertMultipleInstructionPositionEnum = map[string]PatchInsertMultipleInstructionPositionEnum{
	"BEFORE": PatchInsertMultipleInstructionPositionBefore,
	"AFTER":  PatchInsertMultipleInstructionPositionAfter,
}

var mappingPatchInsertMultipleInstructionPositionEnumLowerCase = map[string]PatchInsertMultipleInstructionPositionEnum{
	"before": PatchInsertMultipleInstructionPositionBefore,
	"after":  PatchInsertMultipleInstructionPositionAfter,
}

// GetPatchInsertMultipleInstructionPositionEnumValues Enumerates the set of values for PatchInsertMultipleInstructionPositionEnum
func GetPatchInsertMultipleInstructionPositionEnumValues() []PatchInsertMultipleInstructionPositionEnum {
	values := make([]PatchInsertMultipleInstructionPositionEnum, 0)
	for _, v := range mappingPatchInsertMultipleInstructionPositionEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchInsertMultipleInstructionPositionEnumStringValues Enumerates the set of values in String for PatchInsertMultipleInstructionPositionEnum
func GetPatchInsertMultipleInstructionPositionEnumStringValues() []string {
	return []string{
		"BEFORE",
		"AFTER",
	}
}

// GetMappingPatchInsertMultipleInstructionPositionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchInsertMultipleInstructionPositionEnum(val string) (PatchInsertMultipleInstructionPositionEnum, bool) {
	enum, ok := mappingPatchInsertMultipleInstructionPositionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
