// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// JobDefinitionScope Defines the rules or criteria based on which the scope for job definition is circumscribed.
type JobDefinitionScope struct {

	// Name of the folder or schema for this metadata harvest.
	FolderName *string `mandatory:"false" json:"folderName"`

	// Name of the entity for this metadata harvest.
	EntityName *string `mandatory:"false" json:"entityName"`

	// Filter rules with regular expression to specify folder names for this metadata harvest.
	FolderNameFilter *string `mandatory:"false" json:"folderNameFilter"`

	// Filter rules with regular expression to specify entity names for this metadata harvest.
	EntityNameFilter *string `mandatory:"false" json:"entityNameFilter"`

	// Specify if sample data to be extracted as part of this harvest.
	IsSampleDataExtracted *bool `mandatory:"false" json:"isSampleDataExtracted"`

	// Specify the sample data size in MB, specified as number of rows, for this metadata harvest.
	SampleDataSizeInMBs *int `mandatory:"false" json:"sampleDataSizeInMBs"`
}

func (m JobDefinitionScope) String() string {
	return common.PointerString(m)
}
