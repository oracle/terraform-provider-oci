// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchInstruction A single instruction to be included as part of Patch request content.
type PatchInstruction interface {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	GetSelection() *string
}

type patchinstruction struct {
	JsonData  []byte
	Selection *string `mandatory:"true" json:"selection"`
	Operation string  `json:"operation"`
}

// UnmarshalJSON unmarshals json
func (m *patchinstruction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpatchinstruction patchinstruction
	s := struct {
		Model Unmarshalerpatchinstruction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Selection = s.Model.Selection
	m.Operation = s.Model.Operation

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *patchinstruction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Operation {
	case "MERGE":
		mm := PatchMergeInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PatchInstruction: %s.", m.Operation)
		return *m, nil
	}
}

// GetSelection returns Selection
func (m patchinstruction) GetSelection() *string {
	return m.Selection
}

func (m patchinstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m patchinstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchInstructionOperationEnum Enum with underlying type: string
type PatchInstructionOperationEnum string

// Set of constants representing the allowable values for PatchInstructionOperationEnum
const (
	PatchInstructionOperationMerge PatchInstructionOperationEnum = "MERGE"
)

var mappingPatchInstructionOperationEnum = map[string]PatchInstructionOperationEnum{
	"MERGE": PatchInstructionOperationMerge,
}

var mappingPatchInstructionOperationEnumLowerCase = map[string]PatchInstructionOperationEnum{
	"merge": PatchInstructionOperationMerge,
}

// GetPatchInstructionOperationEnumValues Enumerates the set of values for PatchInstructionOperationEnum
func GetPatchInstructionOperationEnumValues() []PatchInstructionOperationEnum {
	values := make([]PatchInstructionOperationEnum, 0)
	for _, v := range mappingPatchInstructionOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchInstructionOperationEnumStringValues Enumerates the set of values in String for PatchInstructionOperationEnum
func GetPatchInstructionOperationEnumStringValues() []string {
	return []string{
		"MERGE",
	}
}

// GetMappingPatchInstructionOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchInstructionOperationEnum(val string) (PatchInstructionOperationEnum, bool) {
	enum, ok := mappingPatchInstructionOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
