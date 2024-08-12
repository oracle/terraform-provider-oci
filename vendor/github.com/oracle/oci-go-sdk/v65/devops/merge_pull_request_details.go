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

// MergePullRequestDetails determines if this is a merge or a validation.
type MergePullRequestDetails interface {
}

type mergepullrequestdetails struct {
	JsonData   []byte
	ActionType string `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *mergepullrequestdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermergepullrequestdetails mergepullrequestdetails
	s := struct {
		Model Unmarshalermergepullrequestdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *mergepullrequestdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "EXECUTE":
		mm := ExecuteMergePullRequestDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VALIDATE":
		mm := ValidateMergePullRequestDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MergePullRequestDetails: %s.", m.ActionType)
		return *m, nil
	}
}

func (m mergepullrequestdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m mergepullrequestdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MergePullRequestDetailsActionTypeEnum Enum with underlying type: string
type MergePullRequestDetailsActionTypeEnum string

// Set of constants representing the allowable values for MergePullRequestDetailsActionTypeEnum
const (
	MergePullRequestDetailsActionTypeExecute  MergePullRequestDetailsActionTypeEnum = "EXECUTE"
	MergePullRequestDetailsActionTypeValidate MergePullRequestDetailsActionTypeEnum = "VALIDATE"
)

var mappingMergePullRequestDetailsActionTypeEnum = map[string]MergePullRequestDetailsActionTypeEnum{
	"EXECUTE":  MergePullRequestDetailsActionTypeExecute,
	"VALIDATE": MergePullRequestDetailsActionTypeValidate,
}

var mappingMergePullRequestDetailsActionTypeEnumLowerCase = map[string]MergePullRequestDetailsActionTypeEnum{
	"execute":  MergePullRequestDetailsActionTypeExecute,
	"validate": MergePullRequestDetailsActionTypeValidate,
}

// GetMergePullRequestDetailsActionTypeEnumValues Enumerates the set of values for MergePullRequestDetailsActionTypeEnum
func GetMergePullRequestDetailsActionTypeEnumValues() []MergePullRequestDetailsActionTypeEnum {
	values := make([]MergePullRequestDetailsActionTypeEnum, 0)
	for _, v := range mappingMergePullRequestDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMergePullRequestDetailsActionTypeEnumStringValues Enumerates the set of values in String for MergePullRequestDetailsActionTypeEnum
func GetMergePullRequestDetailsActionTypeEnumStringValues() []string {
	return []string{
		"EXECUTE",
		"VALIDATE",
	}
}

// GetMappingMergePullRequestDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMergePullRequestDetailsActionTypeEnum(val string) (MergePullRequestDetailsActionTypeEnum, bool) {
	enum, ok := mappingMergePullRequestDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
