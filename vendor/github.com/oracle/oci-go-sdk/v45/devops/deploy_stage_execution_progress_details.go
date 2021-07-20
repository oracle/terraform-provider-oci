// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// DeployStageExecutionProgressDetails Details about stage execution for each target environment.
type DeployStageExecutionProgressDetails struct {

	// The function ID, instance ID or the cluster ID. For Wait stage it will be the stage ID.
	TargetId *string `mandatory:"false" json:"targetId"`

	// Group for the target environment for example, the batch number for an Instance Group deployment.
	TargetGroup *string `mandatory:"false" json:"targetGroup"`

	// Details about all the steps for one target environment.
	Steps []DeployStageExecutionStep `mandatory:"false" json:"steps"`

	// Details about all the rollback steps for one target environment.
	RollbackSteps []DeployStageExecutionStep `mandatory:"false" json:"rollbackSteps"`
}

func (m DeployStageExecutionProgressDetails) String() string {
	return common.PointerString(m)
}
