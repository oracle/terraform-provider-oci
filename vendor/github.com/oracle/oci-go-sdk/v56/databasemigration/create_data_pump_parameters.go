// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateDataPumpParameters Optional parameters for Data Pump Export and Import. Refer to Configuring Optional Initial Load Advanced Settings (https://docs.us.oracle.com/en/cloud/paas/database-migration/dmsus/working-migration-resources.html#GUID-24BD3054-FDF8-48FF-8492-636C1D4B71ED)
type CreateDataPumpParameters struct {

	// Set to false to force Data Pump worker process to run on one instance.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// Estimate size of dumps that will be generated.
	Estimate DataPumpEstimateEnum `mandatory:"false" json:"estimate,omitempty"`

	// IMPORT: Specifies the action to be performed when data is loaded into a preexisting table.
	TableExistsAction DataPumpTableExistsActionEnum `mandatory:"false" json:"tableExistsAction,omitempty"`

	// Exclude paratemers for Export and Import.
	ExcludeParameters []DataPumpExcludeParametersEnum `mandatory:"false" json:"excludeParameters"`

	// Maximum number of worker processes that can be used for a Data Pump Import job.
	// For an Autonomous Database, ODMS will automatically query its CPU core count and set this property.
	ImportParallelismDegree *int `mandatory:"false" json:"importParallelismDegree"`

	// Maximum number of worker processes that can be used for a Data Pump Export job.
	ExportParallelismDegree *int `mandatory:"false" json:"exportParallelismDegree"`
}

func (m CreateDataPumpParameters) String() string {
	return common.PointerString(m)
}
