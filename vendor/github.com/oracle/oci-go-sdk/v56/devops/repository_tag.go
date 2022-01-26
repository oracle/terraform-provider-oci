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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RepositoryTag The information needed to create a lightweight tag.
type RepositoryTag struct {

	// Unique reference name inside a repository.
	RefName *string `mandatory:"true" json:"refName"`

	// Unique full reference name inside a repository.
	FullRefName *string `mandatory:"true" json:"fullRefName"`

	// The OCID of the repository containing the reference.
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// SHA-1 hash value of the object pointed to by the tag.
	ObjectId *string `mandatory:"true" json:"objectId"`
}

//GetRefName returns RefName
func (m RepositoryTag) GetRefName() *string {
	return m.RefName
}

//GetFullRefName returns FullRefName
func (m RepositoryTag) GetFullRefName() *string {
	return m.FullRefName
}

//GetRepositoryId returns RepositoryId
func (m RepositoryTag) GetRepositoryId() *string {
	return m.RepositoryId
}

func (m RepositoryTag) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RepositoryTag) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRepositoryTag RepositoryTag
	s := struct {
		DiscriminatorParam string `json:"refType"`
		MarshalTypeRepositoryTag
	}{
		"TAG",
		(MarshalTypeRepositoryTag)(m),
	}

	return json.Marshal(&s)
}
