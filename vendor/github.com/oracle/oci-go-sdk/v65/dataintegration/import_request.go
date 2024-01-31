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

// ImportRequest Import metadata object response.
type ImportRequest struct {

	// Import object request key
	Key *string `mandatory:"false" json:"key"`

	// The name of the Object Storage Bucket where the objects will be imported from
	BucketName *string `mandatory:"false" json:"bucketName"`

	// Name of the zip file from which objects will be imported.
	FileName *string `mandatory:"false" json:"fileName"`

	// Optional parameter to point to object storage tenancy (if using Object Storage of different tenancy)
	ObjectStorageTenancyId *string `mandatory:"false" json:"objectStorageTenancyId"`

	// Region of the object storage (if using object storage of different region)
	ObjectStorageRegion *string `mandatory:"false" json:"objectStorageRegion"`

	// Key of the object inside which all the objects will be imported
	ObjectKeyForImport *string `mandatory:"false" json:"objectKeyForImport"`

	// This field controls if the data asset references will be included during import.
	AreDataAssetReferencesIncluded *bool `mandatory:"false" json:"areDataAssetReferencesIncluded"`

	ImportConflictResolution *ImportConflictResolution `mandatory:"false" json:"importConflictResolution"`

	// Import Objects request status.
	Status ImportRequestStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Name of the user who initiated import request.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Number of objects that are imported.
	TotalImportedObjectCount *int `mandatory:"false" json:"totalImportedObjectCount"`

	// Time at which the request started getting processed.
	TimeStartedInMillis *int64 `mandatory:"false" json:"timeStartedInMillis"`

	// Time at which the request was completely processed.
	TimeEndedInMillis *int64 `mandatory:"false" json:"timeEndedInMillis"`

	// Contains key of the error
	ErrorMessages map[string]string `mandatory:"false" json:"errorMessages"`

	// The array of imported object details.
	ImportedObjects []ImportObjectMetadataSummary `mandatory:"false" json:"importedObjects"`

	// Name of the import request.
	Name *string `mandatory:"false" json:"name"`
}

func (m ImportRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImportRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetImportRequestStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportRequestStatusEnum Enum with underlying type: string
type ImportRequestStatusEnum string

// Set of constants representing the allowable values for ImportRequestStatusEnum
const (
	ImportRequestStatusSuccessful  ImportRequestStatusEnum = "SUCCESSFUL"
	ImportRequestStatusFailed      ImportRequestStatusEnum = "FAILED"
	ImportRequestStatusInProgress  ImportRequestStatusEnum = "IN_PROGRESS"
	ImportRequestStatusTerminating ImportRequestStatusEnum = "TERMINATING"
	ImportRequestStatusTerminated  ImportRequestStatusEnum = "TERMINATED"
	ImportRequestStatusQueued      ImportRequestStatusEnum = "QUEUED"
)

var mappingImportRequestStatusEnum = map[string]ImportRequestStatusEnum{
	"SUCCESSFUL":  ImportRequestStatusSuccessful,
	"FAILED":      ImportRequestStatusFailed,
	"IN_PROGRESS": ImportRequestStatusInProgress,
	"TERMINATING": ImportRequestStatusTerminating,
	"TERMINATED":  ImportRequestStatusTerminated,
	"QUEUED":      ImportRequestStatusQueued,
}

var mappingImportRequestStatusEnumLowerCase = map[string]ImportRequestStatusEnum{
	"successful":  ImportRequestStatusSuccessful,
	"failed":      ImportRequestStatusFailed,
	"in_progress": ImportRequestStatusInProgress,
	"terminating": ImportRequestStatusTerminating,
	"terminated":  ImportRequestStatusTerminated,
	"queued":      ImportRequestStatusQueued,
}

// GetImportRequestStatusEnumValues Enumerates the set of values for ImportRequestStatusEnum
func GetImportRequestStatusEnumValues() []ImportRequestStatusEnum {
	values := make([]ImportRequestStatusEnum, 0)
	for _, v := range mappingImportRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetImportRequestStatusEnumStringValues Enumerates the set of values in String for ImportRequestStatusEnum
func GetImportRequestStatusEnumStringValues() []string {
	return []string{
		"SUCCESSFUL",
		"FAILED",
		"IN_PROGRESS",
		"TERMINATING",
		"TERMINATED",
		"QUEUED",
	}
}

// GetMappingImportRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportRequestStatusEnum(val string) (ImportRequestStatusEnum, bool) {
	enum, ok := mappingImportRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
