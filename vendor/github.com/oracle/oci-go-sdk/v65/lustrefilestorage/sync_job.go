// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SyncJob Details associated with sync job runs.
type SyncJob struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the sync job.
	Id *string `mandatory:"true" json:"id"`

	// The type of the sync job.
	JobType SyncJobJobTypeEnum `mandatory:"true" json:"jobType"`

	// The current state of the sync job.
	LifecycleState SyncJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The flag is an identifier to tell whether this specific job run has overwrite enabled.
	// If `isOverwrite` is false, the file to be imported or exported will be skipped if it already exists.
	// If `isOverwrite` is true, the file to be imported or exported will be overwritten if it already exists.
	IsOverwrite *bool `mandatory:"true" json:"isOverwrite"`

	// Total object count for scanned files for import or export as part of this sync job.
	TotalObjectsScanned *int64 `mandatory:"true" json:"totalObjectsScanned"`

	// Count of total files that transferred successfully.
	ObjectsTransferred *int64 `mandatory:"true" json:"objectsTransferred"`

	// Bytes transferred during the sync. This value changes while the sync is still in progress.
	BytesTransferred *int64 `mandatory:"true" json:"bytesTransferred"`

	// Count of files or objects that failed to export or import due to errors.
	SkippedErrorCount *int64 `mandatory:"true" json:"skippedErrorCount"`

	// The date and time the job was started, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2020-07-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The path in the Lustre file system used for this Object Storage link.
	// Example: `myFileSystem/mount/myDirectory`
	LustreFileSystemPath *string `mandatory:"true" json:"lustreFileSystemPath"`

	// The Object Storage namespace and bucket name, including optional object prefix string, to use as the source for imports or destination for exports.
	// Example: `objectStorageNamespace:/bucketName/optionalFolder/optionalPrefix`
	ObjectStoragePath *string `mandatory:"true" json:"objectStoragePath"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Object Storage link.
	ParentId *string `mandatory:"false" json:"parentId"`

	// A message that describes the current state of the sync job in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the job finished, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2020-07-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m SyncJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SyncJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSyncJobJobTypeEnum(string(m.JobType)); !ok && m.JobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobType: %s. Supported values are: %s.", m.JobType, strings.Join(GetSyncJobJobTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSyncJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSyncJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SyncJobJobTypeEnum Enum with underlying type: string
type SyncJobJobTypeEnum string

// Set of constants representing the allowable values for SyncJobJobTypeEnum
const (
	SyncJobJobTypeImport SyncJobJobTypeEnum = "IMPORT"
	SyncJobJobTypeExport SyncJobJobTypeEnum = "EXPORT"
)

var mappingSyncJobJobTypeEnum = map[string]SyncJobJobTypeEnum{
	"IMPORT": SyncJobJobTypeImport,
	"EXPORT": SyncJobJobTypeExport,
}

var mappingSyncJobJobTypeEnumLowerCase = map[string]SyncJobJobTypeEnum{
	"import": SyncJobJobTypeImport,
	"export": SyncJobJobTypeExport,
}

// GetSyncJobJobTypeEnumValues Enumerates the set of values for SyncJobJobTypeEnum
func GetSyncJobJobTypeEnumValues() []SyncJobJobTypeEnum {
	values := make([]SyncJobJobTypeEnum, 0)
	for _, v := range mappingSyncJobJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSyncJobJobTypeEnumStringValues Enumerates the set of values in String for SyncJobJobTypeEnum
func GetSyncJobJobTypeEnumStringValues() []string {
	return []string{
		"IMPORT",
		"EXPORT",
	}
}

// GetMappingSyncJobJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSyncJobJobTypeEnum(val string) (SyncJobJobTypeEnum, bool) {
	enum, ok := mappingSyncJobJobTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SyncJobLifecycleStateEnum Enum with underlying type: string
type SyncJobLifecycleStateEnum string

// Set of constants representing the allowable values for SyncJobLifecycleStateEnum
const (
	SyncJobLifecycleStateInProgress SyncJobLifecycleStateEnum = "IN_PROGRESS"
	SyncJobLifecycleStateSucceeded  SyncJobLifecycleStateEnum = "SUCCEEDED"
	SyncJobLifecycleStateCanceling  SyncJobLifecycleStateEnum = "CANCELING"
	SyncJobLifecycleStateCanceled   SyncJobLifecycleStateEnum = "CANCELED"
	SyncJobLifecycleStateFailed     SyncJobLifecycleStateEnum = "FAILED"
	SyncJobLifecycleStateFailing    SyncJobLifecycleStateEnum = "FAILING"
)

var mappingSyncJobLifecycleStateEnum = map[string]SyncJobLifecycleStateEnum{
	"IN_PROGRESS": SyncJobLifecycleStateInProgress,
	"SUCCEEDED":   SyncJobLifecycleStateSucceeded,
	"CANCELING":   SyncJobLifecycleStateCanceling,
	"CANCELED":    SyncJobLifecycleStateCanceled,
	"FAILED":      SyncJobLifecycleStateFailed,
	"FAILING":     SyncJobLifecycleStateFailing,
}

var mappingSyncJobLifecycleStateEnumLowerCase = map[string]SyncJobLifecycleStateEnum{
	"in_progress": SyncJobLifecycleStateInProgress,
	"succeeded":   SyncJobLifecycleStateSucceeded,
	"canceling":   SyncJobLifecycleStateCanceling,
	"canceled":    SyncJobLifecycleStateCanceled,
	"failed":      SyncJobLifecycleStateFailed,
	"failing":     SyncJobLifecycleStateFailing,
}

// GetSyncJobLifecycleStateEnumValues Enumerates the set of values for SyncJobLifecycleStateEnum
func GetSyncJobLifecycleStateEnumValues() []SyncJobLifecycleStateEnum {
	values := make([]SyncJobLifecycleStateEnum, 0)
	for _, v := range mappingSyncJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSyncJobLifecycleStateEnumStringValues Enumerates the set of values in String for SyncJobLifecycleStateEnum
func GetSyncJobLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"FAILED",
		"FAILING",
	}
}

// GetMappingSyncJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSyncJobLifecycleStateEnum(val string) (SyncJobLifecycleStateEnum, bool) {
	enum, ok := mappingSyncJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
