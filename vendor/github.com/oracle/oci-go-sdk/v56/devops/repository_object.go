// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// RepositoryObjectTypeEnum Enum with underlying type: string
type RepositoryObjectTypeEnum string

// Set of constants representing the allowable values for RepositoryObjectTypeEnum
const (
	RepositoryObjectTypeBlob   RepositoryObjectTypeEnum = "BLOB"
	RepositoryObjectTypeTree   RepositoryObjectTypeEnum = "TREE"
	RepositoryObjectTypeCommit RepositoryObjectTypeEnum = "COMMIT"
)

var mappingRepositoryObjectType = map[string]RepositoryObjectTypeEnum{
	"BLOB":   RepositoryObjectTypeBlob,
	"TREE":   RepositoryObjectTypeTree,
	"COMMIT": RepositoryObjectTypeCommit,
}

// GetRepositoryObjectTypeEnumValues Enumerates the set of values for RepositoryObjectTypeEnum
func GetRepositoryObjectTypeEnumValues() []RepositoryObjectTypeEnum {
	values := make([]RepositoryObjectTypeEnum, 0)
	for _, v := range mappingRepositoryObjectType {
		values = append(values, v)
	}
	return values
}
