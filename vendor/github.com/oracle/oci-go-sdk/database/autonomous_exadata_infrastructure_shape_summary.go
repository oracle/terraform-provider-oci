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

// AutonomousExadataInfrastructureShapeSummary The shape of the Autonomous Exadata Infrastructure. The shape determines resources to allocate to the Autonomous Exadata Infrastructure (CPU cores, memory and storage).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator.
// If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type AutonomousExadataInfrastructureShapeSummary struct {

	// The name of the shape used for the Autonomous Exadata Infrastructure.
	Name *string `mandatory:"true" json:"name"`

	// The maximum number of CPU cores that can be enabled on the Autonomous Exadata Infrastructure.
	AvailableCoreCount *int `mandatory:"true" json:"availableCoreCount"`

	// The minimum number of CPU cores that can be enabled on the Autonomous Exadata Infrastructure.
	MinimumCoreCount *int `mandatory:"false" json:"minimumCoreCount"`

	// The increment in which core count can be increased or decreased.
	CoreCountIncrement *int `mandatory:"false" json:"coreCountIncrement"`

	// The minimum number of nodes available for the shape.
	MinimumNodeCount *int `mandatory:"false" json:"minimumNodeCount"`

	// The maximum number of nodes available for the shape.
	MaximumNodeCount *int `mandatory:"false" json:"maximumNodeCount"`
}

func (m AutonomousExadataInfrastructureShapeSummary) String() string {
	return common.PointerString(m)
}
