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

// ExportRequestSummary Export metadata object response summary.
type ExportRequestSummary struct {

	// Export object request key
	Key *string `mandatory:"false" json:"key"`

	// The list of the objects to be exported
	ObjectKeys []string `mandatory:"false" json:"objectKeys"`

	// The name of the Object Storage Bucket where the objects will be exported to
	BucketName *string `mandatory:"false" json:"bucketName"`

	// Name of the exported zip file.
	FileName *string `mandatory:"false" json:"fileName"`

	// Optional parameter to point to object storage tenancy (if using Object Storage of different tenancy)
	ObjectStorageTenancyId *string `mandatory:"false" json:"objectStorageTenancyId"`

	// Region of the object storage (if using object storage of different region)
	ObjectStorageRegion *string `mandatory:"false" json:"objectStorageRegion"`

	// Controls if the references will be exported along with the objects
	AreReferencesIncluded *bool `mandatory:"false" json:"areReferencesIncluded"`

	// Flag to control whether to overwrite the object if it is already present at the provided object storage location.
	IsObjectOverwriteEnabled *bool `mandatory:"false" json:"isObjectOverwriteEnabled"`

	// Export multiple objects based on filters.
	Filters []string `mandatory:"false" json:"filters"`

	// Export Objects request status.
	Status ExportRequestSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Name of the user who initiated export request.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Number of objects that are exported.
	TotalExportedObjectCount *int `mandatory:"false" json:"totalExportedObjectCount"`

	// Time at which the request started getting processed.
	TimeStartedInMillis *int64 `mandatory:"false" json:"timeStartedInMillis"`

	// Time at which the request was completely processed.
	TimeEndedInMillis *int64 `mandatory:"false" json:"timeEndedInMillis"`

	// Contains key of the error
	ErrorMessages map[string]string `mandatory:"false" json:"errorMessages"`

	// The array of exported object details.
	ExportedItems []ExportObjectMetadataSummary `mandatory:"false" json:"exportedItems"`

	// The array of exported referenced objects.
	ReferencedItems *string `mandatory:"false" json:"referencedItems"`

	// Name of the export request.
	Name *string `mandatory:"false" json:"name"`
}

func (m ExportRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExportRequestSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetExportRequestSummaryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportRequestSummaryStatusEnum Enum with underlying type: string
type ExportRequestSummaryStatusEnum string

// Set of constants representing the allowable values for ExportRequestSummaryStatusEnum
const (
	ExportRequestSummaryStatusSuccessful  ExportRequestSummaryStatusEnum = "SUCCESSFUL"
	ExportRequestSummaryStatusFailed      ExportRequestSummaryStatusEnum = "FAILED"
	ExportRequestSummaryStatusInProgress  ExportRequestSummaryStatusEnum = "IN_PROGRESS"
	ExportRequestSummaryStatusTerminating ExportRequestSummaryStatusEnum = "TERMINATING"
	ExportRequestSummaryStatusTerminated  ExportRequestSummaryStatusEnum = "TERMINATED"
	ExportRequestSummaryStatusQueued      ExportRequestSummaryStatusEnum = "QUEUED"
)

var mappingExportRequestSummaryStatusEnum = map[string]ExportRequestSummaryStatusEnum{
	"SUCCESSFUL":  ExportRequestSummaryStatusSuccessful,
	"FAILED":      ExportRequestSummaryStatusFailed,
	"IN_PROGRESS": ExportRequestSummaryStatusInProgress,
	"TERMINATING": ExportRequestSummaryStatusTerminating,
	"TERMINATED":  ExportRequestSummaryStatusTerminated,
	"QUEUED":      ExportRequestSummaryStatusQueued,
}

var mappingExportRequestSummaryStatusEnumLowerCase = map[string]ExportRequestSummaryStatusEnum{
	"successful":  ExportRequestSummaryStatusSuccessful,
	"failed":      ExportRequestSummaryStatusFailed,
	"in_progress": ExportRequestSummaryStatusInProgress,
	"terminating": ExportRequestSummaryStatusTerminating,
	"terminated":  ExportRequestSummaryStatusTerminated,
	"queued":      ExportRequestSummaryStatusQueued,
}

// GetExportRequestSummaryStatusEnumValues Enumerates the set of values for ExportRequestSummaryStatusEnum
func GetExportRequestSummaryStatusEnumValues() []ExportRequestSummaryStatusEnum {
	values := make([]ExportRequestSummaryStatusEnum, 0)
	for _, v := range mappingExportRequestSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExportRequestSummaryStatusEnumStringValues Enumerates the set of values in String for ExportRequestSummaryStatusEnum
func GetExportRequestSummaryStatusEnumStringValues() []string {
	return []string{
		"SUCCESSFUL",
		"FAILED",
		"IN_PROGRESS",
		"TERMINATING",
		"TERMINATED",
		"QUEUED",
	}
}

// GetMappingExportRequestSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportRequestSummaryStatusEnum(val string) (ExportRequestSummaryStatusEnum, bool) {
	enum, ok := mappingExportRequestSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
