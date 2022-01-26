// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ImportDataAssetJobResult Information about a data asset import operation.
type ImportDataAssetJobResult struct {

	// The unique key of the data asset on which import is triggered.
	DataAssetKey *string `mandatory:"true" json:"dataAssetKey"`

	// The unique key of the job definition resource that is used for the import.
	ImportJobDefinitionKey *string `mandatory:"false" json:"importJobDefinitionKey"`

	// The unique key of the job policy for the import.
	ImportJobKey *string `mandatory:"false" json:"importJobKey"`

	// The unique key of the parent job execution for which the log resource is created.
	ImportJobExecutionKey *string `mandatory:"false" json:"importJobExecutionKey"`

	// The status of the import job execution.
	ImportJobExecutionStatus JobExecutionStateEnum `mandatory:"false" json:"importJobExecutionStatus,omitempty"`
}

func (m ImportDataAssetJobResult) String() string {
	return common.PointerString(m)
}
