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

// FilesystemSnapshotPolicySummary Summary information for a file system snapshot policy.
type FilesystemSnapshotPolicySummary struct {

	// The availability domain that the file system snapshot policy is in.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the file system snapshot policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system snapshot policy.
	Id *string `mandatory:"true" json:"id"`

	// The current state of this file system snapshot policy.
	LifecycleState FilesystemSnapshotPolicySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time that the file system snapshot policy was created
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2020-02-04T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My Filesystem Snapshot Policy`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The prefix to apply to all snapshots created by this policy.
	// Example: `acme`
	PolicyPrefix *string `mandatory:"false" json:"policyPrefix"`

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

func (m FilesystemSnapshotPolicySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FilesystemSnapshotPolicySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFilesystemSnapshotPolicySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFilesystemSnapshotPolicySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FilesystemSnapshotPolicySummaryLifecycleStateEnum Enum with underlying type: string
type FilesystemSnapshotPolicySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for FilesystemSnapshotPolicySummaryLifecycleStateEnum
const (
	FilesystemSnapshotPolicySummaryLifecycleStateCreating FilesystemSnapshotPolicySummaryLifecycleStateEnum = "CREATING"
	FilesystemSnapshotPolicySummaryLifecycleStateActive   FilesystemSnapshotPolicySummaryLifecycleStateEnum = "ACTIVE"
	FilesystemSnapshotPolicySummaryLifecycleStateDeleting FilesystemSnapshotPolicySummaryLifecycleStateEnum = "DELETING"
	FilesystemSnapshotPolicySummaryLifecycleStateDeleted  FilesystemSnapshotPolicySummaryLifecycleStateEnum = "DELETED"
	FilesystemSnapshotPolicySummaryLifecycleStateInactive FilesystemSnapshotPolicySummaryLifecycleStateEnum = "INACTIVE"
	FilesystemSnapshotPolicySummaryLifecycleStateFailed   FilesystemSnapshotPolicySummaryLifecycleStateEnum = "FAILED"
)

var mappingFilesystemSnapshotPolicySummaryLifecycleStateEnum = map[string]FilesystemSnapshotPolicySummaryLifecycleStateEnum{
	"CREATING": FilesystemSnapshotPolicySummaryLifecycleStateCreating,
	"ACTIVE":   FilesystemSnapshotPolicySummaryLifecycleStateActive,
	"DELETING": FilesystemSnapshotPolicySummaryLifecycleStateDeleting,
	"DELETED":  FilesystemSnapshotPolicySummaryLifecycleStateDeleted,
	"INACTIVE": FilesystemSnapshotPolicySummaryLifecycleStateInactive,
	"FAILED":   FilesystemSnapshotPolicySummaryLifecycleStateFailed,
}

var mappingFilesystemSnapshotPolicySummaryLifecycleStateEnumLowerCase = map[string]FilesystemSnapshotPolicySummaryLifecycleStateEnum{
	"creating": FilesystemSnapshotPolicySummaryLifecycleStateCreating,
	"active":   FilesystemSnapshotPolicySummaryLifecycleStateActive,
	"deleting": FilesystemSnapshotPolicySummaryLifecycleStateDeleting,
	"deleted":  FilesystemSnapshotPolicySummaryLifecycleStateDeleted,
	"inactive": FilesystemSnapshotPolicySummaryLifecycleStateInactive,
	"failed":   FilesystemSnapshotPolicySummaryLifecycleStateFailed,
}

// GetFilesystemSnapshotPolicySummaryLifecycleStateEnumValues Enumerates the set of values for FilesystemSnapshotPolicySummaryLifecycleStateEnum
func GetFilesystemSnapshotPolicySummaryLifecycleStateEnumValues() []FilesystemSnapshotPolicySummaryLifecycleStateEnum {
	values := make([]FilesystemSnapshotPolicySummaryLifecycleStateEnum, 0)
	for _, v := range mappingFilesystemSnapshotPolicySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFilesystemSnapshotPolicySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for FilesystemSnapshotPolicySummaryLifecycleStateEnum
func GetFilesystemSnapshotPolicySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"INACTIVE",
		"FAILED",
	}
}

// GetMappingFilesystemSnapshotPolicySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFilesystemSnapshotPolicySummaryLifecycleStateEnum(val string) (FilesystemSnapshotPolicySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingFilesystemSnapshotPolicySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
