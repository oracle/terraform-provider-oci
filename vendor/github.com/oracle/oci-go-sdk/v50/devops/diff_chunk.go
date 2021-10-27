// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v50/common"
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

	// List of DiffSection.
	DiffSections []DiffSection `mandatory:"false" json:"diffSections"`
}

func (m DiffChunk) String() string {
	return common.PointerString(m)
}
