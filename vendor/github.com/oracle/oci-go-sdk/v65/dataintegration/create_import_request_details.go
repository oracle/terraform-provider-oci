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

// CreateImportRequestDetails Details of import object.
type CreateImportRequestDetails struct {

	// Name of the Object Storage bucket where the object will be imported from.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Name of the zip file to be imported.
	FileName *string `mandatory:"true" json:"fileName"`

	// Optional parameter to point to object storage tenancy (if using Object Storage of different tenancy)
	ObjectStorageTenancyId *string `mandatory:"false" json:"objectStorageTenancyId"`

	// Region of the object storage (if using object storage of different region)
	ObjectStorageRegion *string `mandatory:"false" json:"objectStorageRegion"`

	// Key of the object inside which all the objects will be imported
	ObjectKeyForImport *string `mandatory:"false" json:"objectKeyForImport"`

	// This field controls if the data asset references will be included during import.
	AreDataAssetReferencesIncluded *bool `mandatory:"false" json:"areDataAssetReferencesIncluded"`

	ImportConflictResolution *ImportConflictResolution `mandatory:"false" json:"importConflictResolution"`
}

func (m CreateImportRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateImportRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
