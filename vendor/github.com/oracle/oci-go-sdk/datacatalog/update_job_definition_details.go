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

// UpdateJobDefinitionDetails Update information for a job definition resource.
type UpdateJobDefinitionDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Specifies if the job definition is incremental or full.
	IsIncremental *bool `mandatory:"false" json:"isIncremental"`

	// The key of the data asset for which the job is defined.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Detailed description of the job definition.
	Description *string `mandatory:"false" json:"description"`

	// The key of the connection resource to be used for harvest, sampling, profiling jobs.
	ConnectionKey *string `mandatory:"false" json:"connectionKey"`

	// Specify if sample data to be extracted as part of this harvest.
	IsSampleDataExtracted *bool `mandatory:"false" json:"isSampleDataExtracted"`

	// Specify the sample data size in MB, specified as number of rows, for this metadata harvest.
	SampleDataSizeInMBs *int `mandatory:"false" json:"sampleDataSizeInMBs"`

	// A map of maps that contains the properties which are specific to the job type. Each job type
	// definition may define it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// job definitions have required properties within the "default" category.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m UpdateJobDefinitionDetails) String() string {
	return common.PointerString(m)
}
