// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchReference Details of the Patch to be added/removed to/from a Patch Group.
type PatchReference struct {

	// The OCID of the resource.
	PatchId *string `mandatory:"true" json:"patchId"`

	// Type of operation to be done against the given patchId.
	// ADD - Add patch.
	// REMOVE - Remove patch.
	Operation PatchReferenceOperationEnum `mandatory:"true" json:"operation"`
}

func (m PatchReference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchReference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchReferenceOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetPatchReferenceOperationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchReferenceOperationEnum Enum with underlying type: string
type PatchReferenceOperationEnum string

// Set of constants representing the allowable values for PatchReferenceOperationEnum
const (
	PatchReferenceOperationAdd    PatchReferenceOperationEnum = "ADD"
	PatchReferenceOperationRemove PatchReferenceOperationEnum = "REMOVE"
)

var mappingPatchReferenceOperationEnum = map[string]PatchReferenceOperationEnum{
	"ADD":    PatchReferenceOperationAdd,
	"REMOVE": PatchReferenceOperationRemove,
}

var mappingPatchReferenceOperationEnumLowerCase = map[string]PatchReferenceOperationEnum{
	"add":    PatchReferenceOperationAdd,
	"remove": PatchReferenceOperationRemove,
}

// GetPatchReferenceOperationEnumValues Enumerates the set of values for PatchReferenceOperationEnum
func GetPatchReferenceOperationEnumValues() []PatchReferenceOperationEnum {
	values := make([]PatchReferenceOperationEnum, 0)
	for _, v := range mappingPatchReferenceOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchReferenceOperationEnumStringValues Enumerates the set of values in String for PatchReferenceOperationEnum
func GetPatchReferenceOperationEnumStringValues() []string {
	return []string{
		"ADD",
		"REMOVE",
	}
}

// GetMappingPatchReferenceOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchReferenceOperationEnum(val string) (PatchReferenceOperationEnum, bool) {
	enum, ok := mappingPatchReferenceOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
