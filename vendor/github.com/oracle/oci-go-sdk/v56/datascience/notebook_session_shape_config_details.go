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

// NotebookSessionShapeConfigDetails Details for the notebook session shape configuration.
type NotebookSessionShapeConfigDetails struct {

	// A notebook session instance of type VM.Standard.E3.Flex allows the ocpu count to be specified.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// A notebook session instance of type VM.Standard.E3.Flex allows memory to be specified. This specifies the size of the memory in GBs.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`
}

func (m NotebookSessionShapeConfigDetails) String() string {
	return common.PointerString(m)
}
