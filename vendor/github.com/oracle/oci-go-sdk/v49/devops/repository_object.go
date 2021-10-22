// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v49/common"
)

// RepositoryObject Object containing information about files and directories in a repository
type RepositoryObject struct {

	// The type of git object.
	Type RepositoryObjectTypeEnum `mandatory:"true" json:"type"`

	// Size in Bytes
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`

	// SHA-1 hash of git object
	Sha *string `mandatory:"true" json:"sha"`

	// flag to determine is the object contains binary file content or not.
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
