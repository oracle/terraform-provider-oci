// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FileDiffResponse Response object for showing differences for a file between two commits.
type FileDiffResponse struct {

	// List of changed section in the file.
	Changes []DiffChunk `mandatory:"true" json:"changes"`

	// The path on the base version to the changed object.
	OldPath *string `mandatory:"false" json:"oldPath"`

	// The path on the target version to the changed object.
	NewPath *string `mandatory:"false" json:"newPath"`

	// The ID of the changed object on the base version.
	OldId *string `mandatory:"false" json:"oldId"`

	// The ID of the changed object on the target version.
	NewId *string `mandatory:"false" json:"newId"`

	// Indicates whether the changed file contains conflicts.
	AreConflictsInFile *bool `mandatory:"false" json:"areConflictsInFile"`

	// Indicates whether the file is large.
	IsLarge *bool `mandatory:"false" json:"isLarge"`

	// Indicates whether the file is binary.
	IsBinary *bool `mandatory:"false" json:"isBinary"`
}

func (m FileDiffResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileDiffResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
