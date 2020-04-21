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

// OcpUs The details of the available and consumed CPU cores of the Autonomous Exadata Infrastructure instance, including consumption by database workload type.
type OcpUs struct {

	// The total number of OCPUs in the Autonomous Exadata Infrastructure instance.
	TotalCpu *float32 `mandatory:"false" json:"totalCpu"`

	// The total number of consumed OCPUs in the Autonomous Exadata Infrastructure instance.
	ConsumedCpu *float32 `mandatory:"false" json:"consumedCpu"`

	ByWorkloadType *WorkloadType `mandatory:"false" json:"byWorkloadType"`
}

func (m OcpUs) String() string {
	return common.PointerString(m)
}
