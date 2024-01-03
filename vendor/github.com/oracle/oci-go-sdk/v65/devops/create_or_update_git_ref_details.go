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

// CreateOrUpdateGitRefDetails The information needed to create a reference. If the reference already exists, then it can be used to update the reference.
type CreateOrUpdateGitRefDetails interface {

	// The name of the reference to create or update.
	GetRefName() *string
}

type createorupdategitrefdetails struct {
	JsonData []byte
	RefName  *string `mandatory:"true" json:"refName"`
	RefType  string  `json:"refType"`
}

// UnmarshalJSON unmarshals json
func (m *createorupdategitrefdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateorupdategitrefdetails createorupdategitrefdetails
	s := struct {
		Model Unmarshalercreateorupdategitrefdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RefName = s.Model.RefName
	m.RefType = s.Model.RefType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createorupdategitrefdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RefType {
	case "BRANCH":
		mm := CreateOrUpdateGitBranchDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TAG":
		mm := CreateOrUpdateGitTagDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateOrUpdateGitRefDetails: %s.", m.RefType)
		return *m, nil
	}
}

// GetRefName returns RefName
func (m createorupdategitrefdetails) GetRefName() *string {
	return m.RefName
}

func (m createorupdategitrefdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createorupdategitrefdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOrUpdateGitRefDetailsRefTypeEnum Enum with underlying type: string
type CreateOrUpdateGitRefDetailsRefTypeEnum string

// Set of constants representing the allowable values for CreateOrUpdateGitRefDetailsRefTypeEnum
const (
	CreateOrUpdateGitRefDetailsRefTypeBranch CreateOrUpdateGitRefDetailsRefTypeEnum = "BRANCH"
	CreateOrUpdateGitRefDetailsRefTypeTag    CreateOrUpdateGitRefDetailsRefTypeEnum = "TAG"
)

var mappingCreateOrUpdateGitRefDetailsRefTypeEnum = map[string]CreateOrUpdateGitRefDetailsRefTypeEnum{
	"BRANCH": CreateOrUpdateGitRefDetailsRefTypeBranch,
	"TAG":    CreateOrUpdateGitRefDetailsRefTypeTag,
}

var mappingCreateOrUpdateGitRefDetailsRefTypeEnumLowerCase = map[string]CreateOrUpdateGitRefDetailsRefTypeEnum{
	"branch": CreateOrUpdateGitRefDetailsRefTypeBranch,
	"tag":    CreateOrUpdateGitRefDetailsRefTypeTag,
}

// GetCreateOrUpdateGitRefDetailsRefTypeEnumValues Enumerates the set of values for CreateOrUpdateGitRefDetailsRefTypeEnum
func GetCreateOrUpdateGitRefDetailsRefTypeEnumValues() []CreateOrUpdateGitRefDetailsRefTypeEnum {
	values := make([]CreateOrUpdateGitRefDetailsRefTypeEnum, 0)
	for _, v := range mappingCreateOrUpdateGitRefDetailsRefTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOrUpdateGitRefDetailsRefTypeEnumStringValues Enumerates the set of values in String for CreateOrUpdateGitRefDetailsRefTypeEnum
func GetCreateOrUpdateGitRefDetailsRefTypeEnumStringValues() []string {
	return []string{
		"BRANCH",
		"TAG",
	}
}

// GetMappingCreateOrUpdateGitRefDetailsRefTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOrUpdateGitRefDetailsRefTypeEnum(val string) (CreateOrUpdateGitRefDetailsRefTypeEnum, bool) {
	enum, ok := mappingCreateOrUpdateGitRefDetailsRefTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
