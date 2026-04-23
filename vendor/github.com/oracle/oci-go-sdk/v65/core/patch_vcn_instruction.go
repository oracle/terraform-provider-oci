// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// PatchVcnInstruction A single instruction to be included as part of PatchVcn request content.
type PatchVcnInstruction interface {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation
	// against the VCN resource representation.
	// The PatchVcn operation restricts supported selections (see PatchVcn documentation).
	// Example: "ipv6PrivateCidrBlocks"
	GetSelection() *string
}

type patchvcninstruction struct {
	JsonData  []byte
	Selection *string `mandatory:"true" json:"selection"`
	Operation string  `json:"operation"`
}

// UnmarshalJSON unmarshals json
func (m *patchvcninstruction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpatchvcninstruction patchvcninstruction
	s := struct {
		Model Unmarshalerpatchvcninstruction
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
func (m *patchvcninstruction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Operation {
	case "REPLACE":
		mm := PatchVcnReplaceInstruction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PatchVcnInstruction: %s.", m.Operation)
		return *m, nil
	}
}

// GetSelection returns Selection
func (m patchvcninstruction) GetSelection() *string {
	return m.Selection
}

func (m patchvcninstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m patchvcninstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchVcnInstructionOperationEnum Enum with underlying type: string
type PatchVcnInstructionOperationEnum string

// Set of constants representing the allowable values for PatchVcnInstructionOperationEnum
const (
	PatchVcnInstructionOperationReplace PatchVcnInstructionOperationEnum = "REPLACE"
)

var mappingPatchVcnInstructionOperationEnum = map[string]PatchVcnInstructionOperationEnum{
	"REPLACE": PatchVcnInstructionOperationReplace,
}

var mappingPatchVcnInstructionOperationEnumLowerCase = map[string]PatchVcnInstructionOperationEnum{
	"replace": PatchVcnInstructionOperationReplace,
}

// GetPatchVcnInstructionOperationEnumValues Enumerates the set of values for PatchVcnInstructionOperationEnum
func GetPatchVcnInstructionOperationEnumValues() []PatchVcnInstructionOperationEnum {
	values := make([]PatchVcnInstructionOperationEnum, 0)
	for _, v := range mappingPatchVcnInstructionOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchVcnInstructionOperationEnumStringValues Enumerates the set of values in String for PatchVcnInstructionOperationEnum
func GetPatchVcnInstructionOperationEnumStringValues() []string {
	return []string{
		"REPLACE",
	}
}

// GetMappingPatchVcnInstructionOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchVcnInstructionOperationEnum(val string) (PatchVcnInstructionOperationEnum, bool) {
	enum, ok := mappingPatchVcnInstructionOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
