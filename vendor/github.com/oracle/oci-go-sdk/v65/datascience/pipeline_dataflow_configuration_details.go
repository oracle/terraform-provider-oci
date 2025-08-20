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

// PipelineDataflowConfigurationDetails The configuration details of a Dataflow step.
type PipelineDataflowConfigurationDetails struct {

	// The Spark configuration passed to the running process.
	Configuration *interface{} `mandatory:"false" json:"configuration"`

	// The VM shape for the driver.
	DriverShape *string `mandatory:"false" json:"driverShape"`

	DriverShapeConfigDetails *PipelineShapeConfigDetails `mandatory:"false" json:"driverShapeConfigDetails"`

	// The VM shape for the executors.
	ExecutorShape *string `mandatory:"false" json:"executorShape"`

	ExecutorShapeConfigDetails *PipelineShapeConfigDetails `mandatory:"false" json:"executorShapeConfigDetails"`

	// The number of executor VMs requested.
	NumExecutors *int `mandatory:"false" json:"numExecutors"`

	// An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs.
	WarehouseBucketUri *string `mandatory:"false" json:"warehouseBucketUri"`

	// An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded.
	LogsBucketUri *string `mandatory:"false" json:"logsBucketUri"`
}

func (m PipelineDataflowConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineDataflowConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
