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

// ConflictMergeCheck The status of the merge conflict.
type ConflictMergeCheck struct {

	// The status of the conflict.
	Status ConflictMergeCheckStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m ConflictMergeCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConflictMergeCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConflictMergeCheckStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetConflictMergeCheckStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ConflictMergeCheck) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConflictMergeCheck ConflictMergeCheck
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeConflictMergeCheck
	}{
		"CONFLICT",
		(MarshalTypeConflictMergeCheck)(m),
	}

	return json.Marshal(&s)
}

// ConflictMergeCheckStatusEnum Enum with underlying type: string
type ConflictMergeCheckStatusEnum string

// Set of constants representing the allowable values for ConflictMergeCheckStatusEnum
const (
	ConflictMergeCheckStatusSucceeded ConflictMergeCheckStatusEnum = "SUCCEEDED"
	ConflictMergeCheckStatusFailed    ConflictMergeCheckStatusEnum = "FAILED"
)

var mappingConflictMergeCheckStatusEnum = map[string]ConflictMergeCheckStatusEnum{
	"SUCCEEDED": ConflictMergeCheckStatusSucceeded,
	"FAILED":    ConflictMergeCheckStatusFailed,
}

var mappingConflictMergeCheckStatusEnumLowerCase = map[string]ConflictMergeCheckStatusEnum{
	"succeeded": ConflictMergeCheckStatusSucceeded,
	"failed":    ConflictMergeCheckStatusFailed,
}

// GetConflictMergeCheckStatusEnumValues Enumerates the set of values for ConflictMergeCheckStatusEnum
func GetConflictMergeCheckStatusEnumValues() []ConflictMergeCheckStatusEnum {
	values := make([]ConflictMergeCheckStatusEnum, 0)
	for _, v := range mappingConflictMergeCheckStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetConflictMergeCheckStatusEnumStringValues Enumerates the set of values in String for ConflictMergeCheckStatusEnum
func GetConflictMergeCheckStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingConflictMergeCheckStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConflictMergeCheckStatusEnum(val string) (ConflictMergeCheckStatusEnum, bool) {
	enum, ok := mappingConflictMergeCheckStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
