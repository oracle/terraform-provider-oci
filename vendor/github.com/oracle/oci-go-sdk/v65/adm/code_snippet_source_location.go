// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CodeSnippetSourceLocation Source location of a code snippet.
type CodeSnippetSourceLocation interface {
}

type codesnippetsourcelocation struct {
	JsonData []byte
	Origin   string `json:"origin"`
}

// UnmarshalJSON unmarshals json
func (m *codesnippetsourcelocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercodesnippetsourcelocation codesnippetsourcelocation
	s := struct {
		Model Unmarshalercodesnippetsourcelocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Origin = s.Model.Origin

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *codesnippetsourcelocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Origin {
	case "GITHUB":
		mm := GithubCodeSnippetSourceLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CodeSnippetSourceLocation: %s.", m.Origin)
		return *m, nil
	}
}

func (m codesnippetsourcelocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m codesnippetsourcelocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CodeSnippetSourceLocationOriginEnum Enum with underlying type: string
type CodeSnippetSourceLocationOriginEnum string

// Set of constants representing the allowable values for CodeSnippetSourceLocationOriginEnum
const (
	CodeSnippetSourceLocationOriginGithub CodeSnippetSourceLocationOriginEnum = "GITHUB"
)

var mappingCodeSnippetSourceLocationOriginEnum = map[string]CodeSnippetSourceLocationOriginEnum{
	"GITHUB": CodeSnippetSourceLocationOriginGithub,
}

var mappingCodeSnippetSourceLocationOriginEnumLowerCase = map[string]CodeSnippetSourceLocationOriginEnum{
	"github": CodeSnippetSourceLocationOriginGithub,
}

// GetCodeSnippetSourceLocationOriginEnumValues Enumerates the set of values for CodeSnippetSourceLocationOriginEnum
func GetCodeSnippetSourceLocationOriginEnumValues() []CodeSnippetSourceLocationOriginEnum {
	values := make([]CodeSnippetSourceLocationOriginEnum, 0)
	for _, v := range mappingCodeSnippetSourceLocationOriginEnum {
		values = append(values, v)
	}
	return values
}

// GetCodeSnippetSourceLocationOriginEnumStringValues Enumerates the set of values in String for CodeSnippetSourceLocationOriginEnum
func GetCodeSnippetSourceLocationOriginEnumStringValues() []string {
	return []string{
		"GITHUB",
	}
}

// GetMappingCodeSnippetSourceLocationOriginEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCodeSnippetSourceLocationOriginEnum(val string) (CodeSnippetSourceLocationOriginEnum, bool) {
	enum, ok := mappingCodeSnippetSourceLocationOriginEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
