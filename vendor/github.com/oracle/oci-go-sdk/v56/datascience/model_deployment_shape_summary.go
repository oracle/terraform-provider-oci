// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ModelDeploymentShapeSummary The compute shape used to launch a model deployment compute instance.
type ModelDeploymentShapeSummary struct {

	// The name of the model deployment shape.
	Name *string `mandatory:"true" json:"name"`

	// The number of cores associated with this model deployment shape.
	CoreCount *int `mandatory:"true" json:"coreCount"`

	// The amount of memory in GBs associated with this model deployment shape.
	MemoryInGBs *int `mandatory:"true" json:"memoryInGBs"`
}

func (m ModelDeploymentShapeSummary) String() string {
	return common.PointerString(m)
}
