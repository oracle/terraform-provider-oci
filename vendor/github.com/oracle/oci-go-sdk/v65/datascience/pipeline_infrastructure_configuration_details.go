// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PipelineInfrastructureConfigurationDetails The infrastructure configuration details of a pipeline or a step.
type PipelineInfrastructureConfigurationDetails struct {

	// The shape used to launch the instance for all step runs in the pipeline.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The size of the block storage volume to attach to the instance.
	BlockStorageSizeInGBs *int `mandatory:"true" json:"blockStorageSizeInGBs"`

	// The subnet to create a secondary vnic in to attach to the instance running the pipeline step.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	ShapeConfigDetails *PipelineShapeConfigDetails `mandatory:"false" json:"shapeConfigDetails"`
}

func (m PipelineInfrastructureConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineInfrastructureConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
