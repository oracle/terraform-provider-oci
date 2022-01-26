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

// DeployPipelineEnvironment Environment used in the pipeline.
type DeployPipelineEnvironment struct {

	// The OCID of an Environment
	DeployEnvironmentId *string `mandatory:"true" json:"deployEnvironmentId"`

	DeployPipelineStages *DeployPipelineStageCollection `mandatory:"true" json:"deployPipelineStages"`

	// Display name of the environment. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m DeployPipelineEnvironment) String() string {
	return common.PointerString(m)
}
