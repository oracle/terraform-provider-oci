// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// RepositoryRef Reference object with name and commit ID.
type RepositoryRef interface {

	// Unique reference name inside a repository.
	GetRefName() *string

	// Unique full reference name inside a repository.
	GetFullRefName() *string

	// The OCID of the repository containing the reference.
	GetRepositoryId() *string
}

type repositoryref struct {
	JsonData     []byte
	RefName      *string `mandatory:"true" json:"refName"`
	FullRefName  *string `mandatory:"true" json:"fullRefName"`
	RepositoryId *string `mandatory:"true" json:"repositoryId"`
	RefType      string  `json:"refType"`
}

// UnmarshalJSON unmarshals json
func (m *repositoryref) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrepositoryref repositoryref
	s := struct {
		Model Unmarshalerrepositoryref
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RefName = s.Model.RefName
	m.FullRefName = s.Model.FullRefName
	m.RepositoryId = s.Model.RepositoryId
	m.RefType = s.Model.RefType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *repositoryref) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RefType {
	case "BRANCH":
		mm := RepositoryBranch{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TAG":
		mm := RepositoryTag{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetRefName returns RefName
func (m repositoryref) GetRefName() *string {
	return m.RefName
}

//GetFullRefName returns FullRefName
func (m repositoryref) GetFullRefName() *string {
	return m.FullRefName
}

//GetRepositoryId returns RepositoryId
func (m repositoryref) GetRepositoryId() *string {
	return m.RepositoryId
}

func (m repositoryref) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m repositoryref) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RepositoryRefRefTypeEnum Enum with underlying type: string
type RepositoryRefRefTypeEnum string

// Set of constants representing the allowable values for RepositoryRefRefTypeEnum
const (
	RepositoryRefRefTypeBranch RepositoryRefRefTypeEnum = "BRANCH"
	RepositoryRefRefTypeTag    RepositoryRefRefTypeEnum = "TAG"
)

var mappingRepositoryRefRefTypeEnum = map[string]RepositoryRefRefTypeEnum{
	"BRANCH": RepositoryRefRefTypeBranch,
	"TAG":    RepositoryRefRefTypeTag,
}

// GetRepositoryRefRefTypeEnumValues Enumerates the set of values for RepositoryRefRefTypeEnum
func GetRepositoryRefRefTypeEnumValues() []RepositoryRefRefTypeEnum {
	values := make([]RepositoryRefRefTypeEnum, 0)
	for _, v := range mappingRepositoryRefRefTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryRefRefTypeEnumStringValues Enumerates the set of values in String for RepositoryRefRefTypeEnum
func GetRepositoryRefRefTypeEnumStringValues() []string {
	return []string{
		"BRANCH",
		"TAG",
	}
}

// GetMappingRepositoryRefRefTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryRefRefTypeEnum(val string) (RepositoryRefRefTypeEnum, bool) {
	mappingRepositoryRefRefTypeEnumIgnoreCase := make(map[string]RepositoryRefRefTypeEnum)
	for k, v := range mappingRepositoryRefRefTypeEnum {
		mappingRepositoryRefRefTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRepositoryRefRefTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
