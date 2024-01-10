// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RepositoryObject Object containing information about files and directories in a repository.
type RepositoryObject struct {

	// The type of git object.
	Type RepositoryObjectTypeEnum `mandatory:"true" json:"type"`

	// Size in bytes.
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`

	// SHA-1 hash of git object.
	Sha *string `mandatory:"true" json:"sha"`

	// Flag to determine if the object contains binary file content or not.
	IsBinary *bool `mandatory:"false" json:"isBinary"`
}

func (m RepositoryObject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryObject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryObjectTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRepositoryObjectTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RepositoryObjectTypeEnum Enum with underlying type: string
type RepositoryObjectTypeEnum string

// Set of constants representing the allowable values for RepositoryObjectTypeEnum
const (
	RepositoryObjectTypeBlob   RepositoryObjectTypeEnum = "BLOB"
	RepositoryObjectTypeTree   RepositoryObjectTypeEnum = "TREE"
	RepositoryObjectTypeCommit RepositoryObjectTypeEnum = "COMMIT"
)

var mappingRepositoryObjectTypeEnum = map[string]RepositoryObjectTypeEnum{
	"BLOB":   RepositoryObjectTypeBlob,
	"TREE":   RepositoryObjectTypeTree,
	"COMMIT": RepositoryObjectTypeCommit,
}

var mappingRepositoryObjectTypeEnumLowerCase = map[string]RepositoryObjectTypeEnum{
	"blob":   RepositoryObjectTypeBlob,
	"tree":   RepositoryObjectTypeTree,
	"commit": RepositoryObjectTypeCommit,
}

// GetRepositoryObjectTypeEnumValues Enumerates the set of values for RepositoryObjectTypeEnum
func GetRepositoryObjectTypeEnumValues() []RepositoryObjectTypeEnum {
	values := make([]RepositoryObjectTypeEnum, 0)
	for _, v := range mappingRepositoryObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryObjectTypeEnumStringValues Enumerates the set of values in String for RepositoryObjectTypeEnum
func GetRepositoryObjectTypeEnumStringValues() []string {
	return []string{
		"BLOB",
		"TREE",
		"COMMIT",
	}
}

// GetMappingRepositoryObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryObjectTypeEnum(val string) (RepositoryObjectTypeEnum, bool) {
	enum, ok := mappingRepositoryObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
