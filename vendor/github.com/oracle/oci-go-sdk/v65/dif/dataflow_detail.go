// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataflowDetail Details to create an OCI Dataflow resources.
type DataflowDetail struct {

	// Id for dataflow instance
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The Spark version utilized to run the application.
	SparkVersion *string `mandatory:"true" json:"sparkVersion"`

	// The VM shape for the driver. Sets the driver cores and memory.
	DriverShape *string `mandatory:"true" json:"driverShape"`

	// The VM shape for the executors. Sets the executor cores and memory.
	ExecutorShape *string `mandatory:"true" json:"executorShape"`

	// The number of executor VMs requested.
	NumExecutors *int `mandatory:"true" json:"numExecutors"`

	// InstanceId of log bucket created as part of objectstorage service in stack. Used for storing application run logs.
	LogBucketInstanceId *string `mandatory:"true" json:"logBucketInstanceId"`

	DriverShapeConfig *ShapeConfig `mandatory:"false" json:"driverShapeConfig"`

	ExecutorShapeConfig *ShapeConfig `mandatory:"false" json:"executorShapeConfig"`

	// OCID of the already provisioned dataflow private endpoint.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	Connections *DataflowConnections `mandatory:"false" json:"connections"`

	// InstanceId of warehouse bucket created as part of objectstorage service in stack. Mandatory for SQL applications.
	WarehouseBucketInstanceId *string `mandatory:"false" json:"warehouseBucketInstanceId"`
}

func (m DataflowDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataflowDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
