// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// ResourceConfiguration Properties related to a resource.
type ResourceConfiguration struct {

	// The version of the spark used while creating an Oracle Cloud Infrastructure Data Flow application.
	SparkVersion *string `mandatory:"true" json:"sparkVersion"`

	// The VM shape of the driver used while creating an Oracle Cloud Infrastructure Data Flow application. It sets the driver cores and memory.
	DriverShape *string `mandatory:"true" json:"driverShape"`

	// The shape of the executor used while creating an Oracle Cloud Infrastructure Data Flow application. It sets the executor cores and memory.
	ExecutorShape *string `mandatory:"true" json:"executorShape"`

	// Number of executor VMs requested while creating an Oracle Cloud Infrastructure Data Flow application.
	TotalExecutors *int `mandatory:"true" json:"totalExecutors"`
}

func (m ResourceConfiguration) String() string {
	return common.PointerString(m)
}
