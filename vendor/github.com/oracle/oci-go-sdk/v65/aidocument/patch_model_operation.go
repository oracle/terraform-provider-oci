// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchModelOperation The metadata which can be edited after model creation.
type PatchModelOperation struct {

	// The parameter of the resource to be changed.
	Path *string `mandatory:"false" json:"path"`

	// The value of the parameter to be updated.
	Value *string `mandatory:"false" json:"value"`

	// The value of the parameter to be updated.
	Operation PatchModelOperationOperationEnum `mandatory:"false" json:"operation,omitempty"`
}

func (m PatchModelOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchModelOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchModelOperationOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetPatchModelOperationOperationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchModelOperationOperationEnum Enum with underlying type: string
type PatchModelOperationOperationEnum string

// Set of constants representing the allowable values for PatchModelOperationOperationEnum
const (
	PatchModelOperationOperationDelete  PatchModelOperationOperationEnum = "DELETE"
	PatchModelOperationOperationAdd     PatchModelOperationOperationEnum = "ADD"
	PatchModelOperationOperationReplace PatchModelOperationOperationEnum = "REPLACE"
)

var mappingPatchModelOperationOperationEnum = map[string]PatchModelOperationOperationEnum{
	"DELETE":  PatchModelOperationOperationDelete,
	"ADD":     PatchModelOperationOperationAdd,
	"REPLACE": PatchModelOperationOperationReplace,
}

var mappingPatchModelOperationOperationEnumLowerCase = map[string]PatchModelOperationOperationEnum{
	"delete":  PatchModelOperationOperationDelete,
	"add":     PatchModelOperationOperationAdd,
	"replace": PatchModelOperationOperationReplace,
}

// GetPatchModelOperationOperationEnumValues Enumerates the set of values for PatchModelOperationOperationEnum
func GetPatchModelOperationOperationEnumValues() []PatchModelOperationOperationEnum {
	values := make([]PatchModelOperationOperationEnum, 0)
	for _, v := range mappingPatchModelOperationOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchModelOperationOperationEnumStringValues Enumerates the set of values in String for PatchModelOperationOperationEnum
func GetPatchModelOperationOperationEnumStringValues() []string {
	return []string{
		"DELETE",
		"ADD",
		"REPLACE",
	}
}

// GetMappingPatchModelOperationOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchModelOperationOperationEnum(val string) (PatchModelOperationOperationEnum, bool) {
	enum, ok := mappingPatchModelOperationOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
