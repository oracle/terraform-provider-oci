// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiffChunk Details about a group of changes.
type DiffChunk struct {

	// Line number in base version where changes begin.
	BaseLine *int `mandatory:"false" json:"baseLine"`

	// Number of lines chunk spans in base version.
	BaseSpan *int `mandatory:"false" json:"baseSpan"`

	// Line number in target version where changes begin.
	TargetLine *int `mandatory:"false" json:"targetLine"`

	// Number of lines chunk spans in target version.
	TargetSpan *int `mandatory:"false" json:"targetSpan"`

	// List of difference section.
	DiffSections []DiffSection `mandatory:"false" json:"diffSections"`
}

func (m DiffChunk) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiffChunk) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
