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

// PutRepositoryRefDetails The information needed to create a reference. If the reference already exists, then it can be used to update the reference.
type PutRepositoryRefDetails interface {
}

type putrepositoryrefdetails struct {
	JsonData []byte
	RefType  string `json:"refType"`
}

// UnmarshalJSON unmarshals json
func (m *putrepositoryrefdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerputrepositoryrefdetails putrepositoryrefdetails
	s := struct {
		Model Unmarshalerputrepositoryrefdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RefType = s.Model.RefType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *putrepositoryrefdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RefType {
	case "TAG":
		mm := PutRepositoryTagDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BRANCH":
		mm := PutRepositoryBranchDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PutRepositoryRefDetails: %s.", m.RefType)
		return *m, nil
	}
}

func (m putrepositoryrefdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m putrepositoryrefdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PutRepositoryRefDetailsRefTypeEnum Enum with underlying type: string
type PutRepositoryRefDetailsRefTypeEnum string

// Set of constants representing the allowable values for PutRepositoryRefDetailsRefTypeEnum
const (
	PutRepositoryRefDetailsRefTypeBranch PutRepositoryRefDetailsRefTypeEnum = "BRANCH"
	PutRepositoryRefDetailsRefTypeTag    PutRepositoryRefDetailsRefTypeEnum = "TAG"
)

var mappingPutRepositoryRefDetailsRefTypeEnum = map[string]PutRepositoryRefDetailsRefTypeEnum{
	"BRANCH": PutRepositoryRefDetailsRefTypeBranch,
	"TAG":    PutRepositoryRefDetailsRefTypeTag,
}

var mappingPutRepositoryRefDetailsRefTypeEnumLowerCase = map[string]PutRepositoryRefDetailsRefTypeEnum{
	"branch": PutRepositoryRefDetailsRefTypeBranch,
	"tag":    PutRepositoryRefDetailsRefTypeTag,
}

// GetPutRepositoryRefDetailsRefTypeEnumValues Enumerates the set of values for PutRepositoryRefDetailsRefTypeEnum
func GetPutRepositoryRefDetailsRefTypeEnumValues() []PutRepositoryRefDetailsRefTypeEnum {
	values := make([]PutRepositoryRefDetailsRefTypeEnum, 0)
	for _, v := range mappingPutRepositoryRefDetailsRefTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPutRepositoryRefDetailsRefTypeEnumStringValues Enumerates the set of values in String for PutRepositoryRefDetailsRefTypeEnum
func GetPutRepositoryRefDetailsRefTypeEnumStringValues() []string {
	return []string{
		"BRANCH",
		"TAG",
	}
}

// GetMappingPutRepositoryRefDetailsRefTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPutRepositoryRefDetailsRefTypeEnum(val string) (PutRepositoryRefDetailsRefTypeEnum, bool) {
	enum, ok := mappingPutRepositoryRefDetailsRefTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
