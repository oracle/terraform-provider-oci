// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkloadType The number of consumed OCPUs, by database workload type.
type WorkloadType struct {

	// The total number of OCPU cores in use for Autonomous Transaction Processing databases in the infrastructure instance.
	Atp *float32 `mandatory:"false" json:"atp"`

	// The total number of OCPU cores in use for Autonomous Data Warehouse databases in the infrastructure instance.
	Adw *float32 `mandatory:"false" json:"adw"`
}

func (m WorkloadType) String() string {
	return common.PointerString(m)
}
