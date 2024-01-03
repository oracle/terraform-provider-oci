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

// MergeCheck Merge Check summary
type MergeCheck interface {
}

type mergecheck struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *mergecheck) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermergecheck mergecheck
	s := struct {
		Model Unmarshalermergecheck
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *mergecheck) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "BUILD":
		mm := BuildMergeCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONFLICT":
		mm := ConflictMergeCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPROVAL_RULE":
		mm := ApprovalRuleMergeCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MergeCheck: %s.", m.Type)
		return *m, nil
	}
}

func (m mergecheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m mergecheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MergeCheckTypeEnum Enum with underlying type: string
type MergeCheckTypeEnum string

// Set of constants representing the allowable values for MergeCheckTypeEnum
const (
	MergeCheckTypeConflict     MergeCheckTypeEnum = "CONFLICT"
	MergeCheckTypeApprovalRule MergeCheckTypeEnum = "APPROVAL_RULE"
	MergeCheckTypeBuild        MergeCheckTypeEnum = "BUILD"
)

var mappingMergeCheckTypeEnum = map[string]MergeCheckTypeEnum{
	"CONFLICT":      MergeCheckTypeConflict,
	"APPROVAL_RULE": MergeCheckTypeApprovalRule,
	"BUILD":         MergeCheckTypeBuild,
}

var mappingMergeCheckTypeEnumLowerCase = map[string]MergeCheckTypeEnum{
	"conflict":      MergeCheckTypeConflict,
	"approval_rule": MergeCheckTypeApprovalRule,
	"build":         MergeCheckTypeBuild,
}

// GetMergeCheckTypeEnumValues Enumerates the set of values for MergeCheckTypeEnum
func GetMergeCheckTypeEnumValues() []MergeCheckTypeEnum {
	values := make([]MergeCheckTypeEnum, 0)
	for _, v := range mappingMergeCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMergeCheckTypeEnumStringValues Enumerates the set of values in String for MergeCheckTypeEnum
func GetMergeCheckTypeEnumStringValues() []string {
	return []string{
		"CONFLICT",
		"APPROVAL_RULE",
		"BUILD",
	}
}

// GetMappingMergeCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMergeCheckTypeEnum(val string) (MergeCheckTypeEnum, bool) {
	enum, ok := mappingMergeCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
