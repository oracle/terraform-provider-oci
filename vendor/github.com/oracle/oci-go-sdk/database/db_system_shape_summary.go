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

// DbSystemShapeSummary The shape of the DB system. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator.
// If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type DbSystemShapeSummary struct {

	// The name of the shape used for the DB system.
	Name *string `mandatory:"true" json:"name"`

	// The maximum number of CPU cores that can be enabled on the DB system for this shape.
	AvailableCoreCount *int `mandatory:"true" json:"availableCoreCount"`

	// The family of the shape used for the DB system.
	ShapeFamily *string `mandatory:"false" json:"shapeFamily"`

	// Deprecated. Use `name` instead of `shape`.
	Shape *string `mandatory:"false" json:"shape"`

	// The minimum number of CPU cores that can be enabled on the DB system for this shape.
	MinimumCoreCount *int `mandatory:"false" json:"minimumCoreCount"`

	// The discrete number by which the CPU core count for this shape can be increased or decreased.
	CoreCountIncrement *int `mandatory:"false" json:"coreCountIncrement"`

	// The minimum number of CPU cores that can be enabled per node for this shape.
	MinCoreCountPerNode *int `mandatory:"false" json:"minCoreCountPerNode"`

	// The maximum memory that can be enabled for this shape.
	AvailableMemoryInGBs *int `mandatory:"false" json:"availableMemoryInGBs"`

	// The minimum memory that need be allocated per node for this shape.
	MinMemoryPerNodeInGBs *int `mandatory:"false" json:"minMemoryPerNodeInGBs"`

	// The maximum Db Node storage that can be enabled for this shape.
	AvailableDbNodeStorageInGBs *int `mandatory:"false" json:"availableDbNodeStorageInGBs"`

	// The minimum Db Node storage that need be allocated per node for this shape.
	MinDbNodeStoragePerNodeInGBs *int `mandatory:"false" json:"minDbNodeStoragePerNodeInGBs"`

	// The maximum DATA storage that can be enabled for this shape.
	AvailableDataStorageInTBs *int `mandatory:"false" json:"availableDataStorageInTBs"`

	// The minimum data storage that need be allocated for this shape.
	MinDataStorageInTBs *int `mandatory:"false" json:"minDataStorageInTBs"`

	// The minimum number of database nodes available for this shape.
	MinimumNodeCount *int `mandatory:"false" json:"minimumNodeCount"`

	// The maximum number of database nodes available for this shape.
	MaximumNodeCount *int `mandatory:"false" json:"maximumNodeCount"`
}

func (m DbSystemShapeSummary) String() string {
	return common.PointerString(m)
}
