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
	case "MOVE":
		mm := PatchMoveInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOVE":
		mm := PatchRemoveInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REPLACE":
		mm := PatchReplaceInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INSERT":
		mm := PatchInsertInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REQUIRE":
		mm := PatchRequireInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MERGE":
		mm := PatchMergeInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PROHIBIT":
		mm := PatchProhibitInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PatchInstruction: %s.", m.Operation)
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
	PatchInstructionOperationRequire  PatchInstructionOperationEnum = "REQUIRE"
	PatchInstructionOperationProhibit PatchInstructionOperationEnum = "PROHIBIT"
	PatchInstructionOperationReplace  PatchInstructionOperationEnum = "REPLACE"
	PatchInstructionOperationInsert   PatchInstructionOperationEnum = "INSERT"
	PatchInstructionOperationRemove   PatchInstructionOperationEnum = "REMOVE"
	PatchInstructionOperationMove     PatchInstructionOperationEnum = "MOVE"
	PatchInstructionOperationMerge    PatchInstructionOperationEnum = "MERGE"
)

var mappingPatchInstructionOperationEnum = map[string]PatchInstructionOperationEnum{
	"REQUIRE":  PatchInstructionOperationRequire,
	"PROHIBIT": PatchInstructionOperationProhibit,
	"REPLACE":  PatchInstructionOperationReplace,
	"INSERT":   PatchInstructionOperationInsert,
	"REMOVE":   PatchInstructionOperationRemove,
	"MOVE":     PatchInstructionOperationMove,
	"MERGE":    PatchInstructionOperationMerge,
}

var mappingPatchInstructionOperationEnumLowerCase = map[string]PatchInstructionOperationEnum{
	"require":  PatchInstructionOperationRequire,
	"prohibit": PatchInstructionOperationProhibit,
	"replace":  PatchInstructionOperationReplace,
	"insert":   PatchInstructionOperationInsert,
	"remove":   PatchInstructionOperationRemove,
	"move":     PatchInstructionOperationMove,
	"merge":    PatchInstructionOperationMerge,
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
		"REQUIRE",
		"PROHIBIT",
		"REPLACE",
		"INSERT",
		"REMOVE",
		"MOVE",
		"MERGE",
	}
}

// GetMappingPatchInstructionOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchInstructionOperationEnum(val string) (PatchInstructionOperationEnum, bool) {
	enum, ok := mappingPatchInstructionOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
