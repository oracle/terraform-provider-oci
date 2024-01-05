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

// CopyObjectRequestSummary Copy metadata object response summary.
type CopyObjectRequestSummary struct {

	// Copy object request key.
	Key *string `mandatory:"false" json:"key"`

	// The workspace id of the source from where we need to copy object.
	SourceWorkspaceId *string `mandatory:"false" json:"sourceWorkspaceId"`

	// The list of the objects to be copied.
	ObjectKeys []string `mandatory:"false" json:"objectKeys"`

	CopyConflictResolution *CopyConflictResolution `mandatory:"false" json:"copyConflictResolution"`

	// Copy Object request status.
	CopyMetadataObjectRequestStatus CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum `mandatory:"false" json:"copyMetadataObjectRequestStatus,omitempty"`

	// OCID of the user who initiated copy request.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Name of the user who created the copy object request.
	CreatedByName *string `mandatory:"false" json:"createdByName"`

	// Number of source objects to be copied.
	TotalSourceObjectCount *int `mandatory:"false" json:"totalSourceObjectCount"`

	// Number of objects copied into the target.
	TotalObjectsCopiedIntoTarget *int `mandatory:"false" json:"totalObjectsCopiedIntoTarget"`

	// Time at which the request started getting processed.
	TimeStartedInMillis *int64 `mandatory:"false" json:"timeStartedInMillis"`

	// Time at which the request was completely processed.
	TimeEndedInMillis *int64 `mandatory:"false" json:"timeEndedInMillis"`

	// The array of copy object details.
	CopiedItems []CopyObjectMetadataSummary `mandatory:"false" json:"copiedItems"`

	// The array of copied referenced objects.
	ReferencedItems []CopyObjectMetadataSummary `mandatory:"false" json:"referencedItems"`

	// Name of the copy object request.
	Name *string `mandatory:"false" json:"name"`
}

func (m CopyObjectRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CopyObjectRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum(string(m.CopyMetadataObjectRequestStatus)); !ok && m.CopyMetadataObjectRequestStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CopyMetadataObjectRequestStatus: %s. Supported values are: %s.", m.CopyMetadataObjectRequestStatus, strings.Join(GetCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum Enum with underlying type: string
type CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum string

// Set of constants representing the allowable values for CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum
const (
	CopyObjectRequestSummaryCopyMetadataObjectRequestStatusSuccessful  CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum = "SUCCESSFUL"
	CopyObjectRequestSummaryCopyMetadataObjectRequestStatusFailed      CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum = "FAILED"
	CopyObjectRequestSummaryCopyMetadataObjectRequestStatusInProgress  CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum = "IN_PROGRESS"
	CopyObjectRequestSummaryCopyMetadataObjectRequestStatusQueued      CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum = "QUEUED"
	CopyObjectRequestSummaryCopyMetadataObjectRequestStatusTerminating CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum = "TERMINATING"
	CopyObjectRequestSummaryCopyMetadataObjectRequestStatusTerminated  CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum = "TERMINATED"
)

var mappingCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum = map[string]CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum{
	"SUCCESSFUL":  CopyObjectRequestSummaryCopyMetadataObjectRequestStatusSuccessful,
	"FAILED":      CopyObjectRequestSummaryCopyMetadataObjectRequestStatusFailed,
	"IN_PROGRESS": CopyObjectRequestSummaryCopyMetadataObjectRequestStatusInProgress,
	"QUEUED":      CopyObjectRequestSummaryCopyMetadataObjectRequestStatusQueued,
	"TERMINATING": CopyObjectRequestSummaryCopyMetadataObjectRequestStatusTerminating,
	"TERMINATED":  CopyObjectRequestSummaryCopyMetadataObjectRequestStatusTerminated,
}

var mappingCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnumLowerCase = map[string]CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum{
	"successful":  CopyObjectRequestSummaryCopyMetadataObjectRequestStatusSuccessful,
	"failed":      CopyObjectRequestSummaryCopyMetadataObjectRequestStatusFailed,
	"in_progress": CopyObjectRequestSummaryCopyMetadataObjectRequestStatusInProgress,
	"queued":      CopyObjectRequestSummaryCopyMetadataObjectRequestStatusQueued,
	"terminating": CopyObjectRequestSummaryCopyMetadataObjectRequestStatusTerminating,
	"terminated":  CopyObjectRequestSummaryCopyMetadataObjectRequestStatusTerminated,
}

// GetCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnumValues Enumerates the set of values for CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum
func GetCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnumValues() []CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum {
	values := make([]CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum, 0)
	for _, v := range mappingCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnumStringValues Enumerates the set of values in String for CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum
func GetCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnumStringValues() []string {
	return []string{
		"SUCCESSFUL",
		"FAILED",
		"IN_PROGRESS",
		"QUEUED",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum(val string) (CopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnum, bool) {
	enum, ok := mappingCopyObjectRequestSummaryCopyMetadataObjectRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
