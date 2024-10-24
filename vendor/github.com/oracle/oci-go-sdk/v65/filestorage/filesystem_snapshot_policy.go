// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FilesystemSnapshotPolicy A file system snapshot policy is used to automate snapshot creation and deletion.
// It contains a list of snapshot schedules that define the frequency of
// snapshot creation for the associated file systems and the retention period of snapshots taken on schedule.
// For more information, see Snapshot Scheduling (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/snapshot-policies-and-schedules.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type FilesystemSnapshotPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the file system snapshot policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain that the file system snapshot policy is in. May be unset using a blank or NULL value.
	// Example: `Uocm:PHX-AD-2`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system snapshot policy.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the file system snapshot policy.
	LifecycleState FilesystemSnapshotPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the file system snapshot policy was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `policy1`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The prefix to apply to all snapshots created by this policy.
	// Example: `acme`
	PolicyPrefix *string `mandatory:"false" json:"policyPrefix"`

	// The list of associated snapshot schedules. A maximum of 10 schedules can be associated with a policy.
	Schedules []SnapshotSchedule `mandatory:"false" json:"schedules"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m FilesystemSnapshotPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FilesystemSnapshotPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFilesystemSnapshotPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFilesystemSnapshotPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FilesystemSnapshotPolicyLifecycleStateEnum Enum with underlying type: string
type FilesystemSnapshotPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for FilesystemSnapshotPolicyLifecycleStateEnum
const (
	FilesystemSnapshotPolicyLifecycleStateCreating FilesystemSnapshotPolicyLifecycleStateEnum = "CREATING"
	FilesystemSnapshotPolicyLifecycleStateActive   FilesystemSnapshotPolicyLifecycleStateEnum = "ACTIVE"
	FilesystemSnapshotPolicyLifecycleStateDeleting FilesystemSnapshotPolicyLifecycleStateEnum = "DELETING"
	FilesystemSnapshotPolicyLifecycleStateDeleted  FilesystemSnapshotPolicyLifecycleStateEnum = "DELETED"
	FilesystemSnapshotPolicyLifecycleStateInactive FilesystemSnapshotPolicyLifecycleStateEnum = "INACTIVE"
	FilesystemSnapshotPolicyLifecycleStateFailed   FilesystemSnapshotPolicyLifecycleStateEnum = "FAILED"
)

var mappingFilesystemSnapshotPolicyLifecycleStateEnum = map[string]FilesystemSnapshotPolicyLifecycleStateEnum{
	"CREATING": FilesystemSnapshotPolicyLifecycleStateCreating,
	"ACTIVE":   FilesystemSnapshotPolicyLifecycleStateActive,
	"DELETING": FilesystemSnapshotPolicyLifecycleStateDeleting,
	"DELETED":  FilesystemSnapshotPolicyLifecycleStateDeleted,
	"INACTIVE": FilesystemSnapshotPolicyLifecycleStateInactive,
	"FAILED":   FilesystemSnapshotPolicyLifecycleStateFailed,
}

var mappingFilesystemSnapshotPolicyLifecycleStateEnumLowerCase = map[string]FilesystemSnapshotPolicyLifecycleStateEnum{
	"creating": FilesystemSnapshotPolicyLifecycleStateCreating,
	"active":   FilesystemSnapshotPolicyLifecycleStateActive,
	"deleting": FilesystemSnapshotPolicyLifecycleStateDeleting,
	"deleted":  FilesystemSnapshotPolicyLifecycleStateDeleted,
	"inactive": FilesystemSnapshotPolicyLifecycleStateInactive,
	"failed":   FilesystemSnapshotPolicyLifecycleStateFailed,
}

// GetFilesystemSnapshotPolicyLifecycleStateEnumValues Enumerates the set of values for FilesystemSnapshotPolicyLifecycleStateEnum
func GetFilesystemSnapshotPolicyLifecycleStateEnumValues() []FilesystemSnapshotPolicyLifecycleStateEnum {
	values := make([]FilesystemSnapshotPolicyLifecycleStateEnum, 0)
	for _, v := range mappingFilesystemSnapshotPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFilesystemSnapshotPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for FilesystemSnapshotPolicyLifecycleStateEnum
func GetFilesystemSnapshotPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"INACTIVE",
		"FAILED",
	}
}

// GetMappingFilesystemSnapshotPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFilesystemSnapshotPolicyLifecycleStateEnum(val string) (FilesystemSnapshotPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingFilesystemSnapshotPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
