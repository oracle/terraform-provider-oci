// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExportRequestDetails Details of export request. Export is supported using three ways.
// First, when objectKeys are provided, export of those objects take place.
// Second, when filter are provided, all the objects based on the filter provided are exported.
// Third, when neither objectKeys nor filters are provided, we export all the design objects for the workspace.
type CreateExportRequestDetails struct {

	// Name of the Object Storage bucket where the object will be exported.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Name of the exported zip file.
	FileName *string `mandatory:"false" json:"fileName"`

	// Optional parameter to point to object storage tenancy (if using Object Storage of different tenancy)
	ObjectStorageTenancyId *string `mandatory:"false" json:"objectStorageTenancyId"`

	// Region of the object storage (if using object storage of different region)
	ObjectStorageRegion *string `mandatory:"false" json:"objectStorageRegion"`

	// Flag to control whether to overwrite the object if it is already present at the provided object storage location.
	IsObjectOverwriteEnabled *bool `mandatory:"false" json:"isObjectOverwriteEnabled"`

	// Field is used to specify which object keys to export
	ObjectKeys []string `mandatory:"false" json:"objectKeys"`

	// This field controls if the references will be exported along with the objects
	AreReferencesIncluded *bool `mandatory:"false" json:"areReferencesIncluded"`

	// Filters for exported objects
	Filters []string `mandatory:"false" json:"filters"`
}

func (m CreateExportRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExportRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
