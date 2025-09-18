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

// ObjectStorageLink Object Storage links create the relationship between a directory in an File Storage with Lustre file system and a path within an Object Storage bucket.
// For more information, see Syncing Lustre file systems with Object Storage (https://docs.oracle.com/iaas/Content/lustre/object-storage-sync.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type ObjectStorageLink struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ObjectStorageLink.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Lustre file system.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain the file system is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My Object Storage Link`
	//
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the Lustre file system was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2024-04-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Object Storage link was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2024-04-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Object Storage link.
	LifecycleState ObjectStorageLinkLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated Lustre file system.
	LustreFileSystemId *string `mandatory:"true" json:"lustreFileSystemId"`

	// The path in the Lustre file system used for this Object Storage link.
	// Example: `myFileSystem/mount/myDirectory`
	FileSystemPath *string `mandatory:"true" json:"fileSystemPath"`

	// The Object Storage namespace and bucket name, including optional object prefix string, to use as the source for imports or destination for exports.
	// Example: `objectStorageNamespace:/bucketName/optionalFolder/optionalPrefix`
	ObjectStoragePrefix *string `mandatory:"true" json:"objectStoragePrefix"`

	// The flag is an identifier to tell whether the job run has overwrite enabled.
	// If `isOverwrite` is false, the file to be imported or exported will be skipped if it already exists.
	// If `isOverwrite` is true, the file to be imported or exported will be overwritten if it already exists.
	IsOverwrite *bool `mandatory:"true" json:"isOverwrite"`

	// A message that describes the current state of the Object Storage link in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of currently running sync job. If no sync job is running, then this will be empty.
	CurrentJobId *string `mandatory:"false" json:"currentJobId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of last succeeded sync job. If no sync job has previously run, then this will be empty.
	LastJobId *string `mandatory:"false" json:"lastJobId"`
}

func (m ObjectStorageLink) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageLink) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingObjectStorageLinkLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetObjectStorageLinkLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ObjectStorageLinkLifecycleStateEnum Enum with underlying type: string
type ObjectStorageLinkLifecycleStateEnum string

// Set of constants representing the allowable values for ObjectStorageLinkLifecycleStateEnum
const (
	ObjectStorageLinkLifecycleStateCreating ObjectStorageLinkLifecycleStateEnum = "CREATING"
	ObjectStorageLinkLifecycleStateActive   ObjectStorageLinkLifecycleStateEnum = "ACTIVE"
	ObjectStorageLinkLifecycleStateDeleting ObjectStorageLinkLifecycleStateEnum = "DELETING"
	ObjectStorageLinkLifecycleStateDeleted  ObjectStorageLinkLifecycleStateEnum = "DELETED"
	ObjectStorageLinkLifecycleStateFailed   ObjectStorageLinkLifecycleStateEnum = "FAILED"
)

var mappingObjectStorageLinkLifecycleStateEnum = map[string]ObjectStorageLinkLifecycleStateEnum{
	"CREATING": ObjectStorageLinkLifecycleStateCreating,
	"ACTIVE":   ObjectStorageLinkLifecycleStateActive,
	"DELETING": ObjectStorageLinkLifecycleStateDeleting,
	"DELETED":  ObjectStorageLinkLifecycleStateDeleted,
	"FAILED":   ObjectStorageLinkLifecycleStateFailed,
}

var mappingObjectStorageLinkLifecycleStateEnumLowerCase = map[string]ObjectStorageLinkLifecycleStateEnum{
	"creating": ObjectStorageLinkLifecycleStateCreating,
	"active":   ObjectStorageLinkLifecycleStateActive,
	"deleting": ObjectStorageLinkLifecycleStateDeleting,
	"deleted":  ObjectStorageLinkLifecycleStateDeleted,
	"failed":   ObjectStorageLinkLifecycleStateFailed,
}

// GetObjectStorageLinkLifecycleStateEnumValues Enumerates the set of values for ObjectStorageLinkLifecycleStateEnum
func GetObjectStorageLinkLifecycleStateEnumValues() []ObjectStorageLinkLifecycleStateEnum {
	values := make([]ObjectStorageLinkLifecycleStateEnum, 0)
	for _, v := range mappingObjectStorageLinkLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectStorageLinkLifecycleStateEnumStringValues Enumerates the set of values in String for ObjectStorageLinkLifecycleStateEnum
func GetObjectStorageLinkLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingObjectStorageLinkLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectStorageLinkLifecycleStateEnum(val string) (ObjectStorageLinkLifecycleStateEnum, bool) {
	enum, ok := mappingObjectStorageLinkLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
