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

// SyncJobSummary Summary information associated with sync jobs.
type SyncJobSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the sync job.
	Id *string `mandatory:"true" json:"id"`

	// The type of the sync job.
	JobType SyncJobSummaryJobTypeEnum `mandatory:"true" json:"jobType"`

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

	// Count of total files transferred successfully.
	ObjectsTransferred *int64 `mandatory:"true" json:"objectsTransferred"`

	// Bytes transferred during the sync. This value changes while sync is still in progress.
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

	// A message that describes the current state of the sync job in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the job finished, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2024-07-21T20:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m SyncJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SyncJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSyncJobSummaryJobTypeEnum(string(m.JobType)); !ok && m.JobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobType: %s. Supported values are: %s.", m.JobType, strings.Join(GetSyncJobSummaryJobTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSyncJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSyncJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SyncJobSummaryJobTypeEnum Enum with underlying type: string
type SyncJobSummaryJobTypeEnum string

// Set of constants representing the allowable values for SyncJobSummaryJobTypeEnum
const (
	SyncJobSummaryJobTypeImport SyncJobSummaryJobTypeEnum = "IMPORT"
	SyncJobSummaryJobTypeExport SyncJobSummaryJobTypeEnum = "EXPORT"
)

var mappingSyncJobSummaryJobTypeEnum = map[string]SyncJobSummaryJobTypeEnum{
	"IMPORT": SyncJobSummaryJobTypeImport,
	"EXPORT": SyncJobSummaryJobTypeExport,
}

var mappingSyncJobSummaryJobTypeEnumLowerCase = map[string]SyncJobSummaryJobTypeEnum{
	"import": SyncJobSummaryJobTypeImport,
	"export": SyncJobSummaryJobTypeExport,
}

// GetSyncJobSummaryJobTypeEnumValues Enumerates the set of values for SyncJobSummaryJobTypeEnum
func GetSyncJobSummaryJobTypeEnumValues() []SyncJobSummaryJobTypeEnum {
	values := make([]SyncJobSummaryJobTypeEnum, 0)
	for _, v := range mappingSyncJobSummaryJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSyncJobSummaryJobTypeEnumStringValues Enumerates the set of values in String for SyncJobSummaryJobTypeEnum
func GetSyncJobSummaryJobTypeEnumStringValues() []string {
	return []string{
		"IMPORT",
		"EXPORT",
	}
}

// GetMappingSyncJobSummaryJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSyncJobSummaryJobTypeEnum(val string) (SyncJobSummaryJobTypeEnum, bool) {
	enum, ok := mappingSyncJobSummaryJobTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
