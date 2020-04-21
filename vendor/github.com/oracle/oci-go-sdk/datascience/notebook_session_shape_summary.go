// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
)

// NotebookSessionShapeSummary The compute shape used to launch a notebook session compute instance.
type NotebookSessionShapeSummary struct {

	// The name of the notebook session shape.
	Name *string `mandatory:"true" json:"name"`

	// The number of cores associated with this notebook session shape.
	CoreCount *int `mandatory:"true" json:"coreCount"`

	// The amount of memory in GBs associated with this notebook session shape.
	MemoryInGBs *int `mandatory:"true" json:"memoryInGBs"`
}

func (m NotebookSessionShapeSummary) String() string {
	return common.PointerString(m)
}
