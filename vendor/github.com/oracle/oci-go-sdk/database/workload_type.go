// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
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
