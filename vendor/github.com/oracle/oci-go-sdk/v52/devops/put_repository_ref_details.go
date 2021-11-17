// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// PutRepositoryRefDetails The information needed to create a ref. If the ref already exists, it can be used to update it.
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
		return *m, nil
	}
}

func (m putrepositoryrefdetails) String() string {
	return common.PointerString(m)
}

// PutRepositoryRefDetailsRefTypeEnum Enum with underlying type: string
type PutRepositoryRefDetailsRefTypeEnum string

// Set of constants representing the allowable values for PutRepositoryRefDetailsRefTypeEnum
const (
	PutRepositoryRefDetailsRefTypeBranch PutRepositoryRefDetailsRefTypeEnum = "BRANCH"
	PutRepositoryRefDetailsRefTypeTag    PutRepositoryRefDetailsRefTypeEnum = "TAG"
)

var mappingPutRepositoryRefDetailsRefType = map[string]PutRepositoryRefDetailsRefTypeEnum{
	"BRANCH": PutRepositoryRefDetailsRefTypeBranch,
	"TAG":    PutRepositoryRefDetailsRefTypeTag,
}

// GetPutRepositoryRefDetailsRefTypeEnumValues Enumerates the set of values for PutRepositoryRefDetailsRefTypeEnum
func GetPutRepositoryRefDetailsRefTypeEnumValues() []PutRepositoryRefDetailsRefTypeEnum {
	values := make([]PutRepositoryRefDetailsRefTypeEnum, 0)
	for _, v := range mappingPutRepositoryRefDetailsRefType {
		values = append(values, v)
	}
	return values
}
