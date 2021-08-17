// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// DeployStagePredecessor Metadata for defining a stage's predecessor.
type DeployStagePredecessor struct {

	// The OCID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's OCID.
	Id *string `mandatory:"true" json:"id"`
}

func (m DeployStagePredecessor) String() string {
	return common.PointerString(m)
}
