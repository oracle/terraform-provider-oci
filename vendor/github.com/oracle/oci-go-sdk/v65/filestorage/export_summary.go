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

// ExportSummary Summary information for an export.
type ExportSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export's export set.
	ExportSetId *string `mandatory:"true" json:"exportSetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export's file system.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export.
	Id *string `mandatory:"true" json:"id"`

	// The current state of this export.
	LifecycleState ExportSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Path used to access the associated file system.
	// Avoid entering confidential information.
	// Example: `/mediafiles`
	Path *string `mandatory:"true" json:"path"`

	// The date and time the export was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Whether or not the export should use ID mapping for Unix groups rather than the group list provided within an NFS request's RPC header. When this flag is true the Unix UID from the RPC header is used to retrieve the list of secondary groups from a the ID mapping subsystem. The primary GID is always taken from the RPC header. If ID mapping is not configured, incorrectly configured, unavailable, or cannot be used to determine a list of secondary groups then an empty secondary group list is used for authorization. If the number of groups exceeds the limit of 256 groups, the list retrieved from LDAP is truncated to the first 256 groups read.
	IsIdmapGroupsForSysAuth *bool `mandatory:"false" json:"isIdmapGroupsForSysAuth"`
}

func (m ExportSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExportSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExportSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportSummaryLifecycleStateEnum Enum with underlying type: string
type ExportSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExportSummaryLifecycleStateEnum
const (
	ExportSummaryLifecycleStateCreating ExportSummaryLifecycleStateEnum = "CREATING"
	ExportSummaryLifecycleStateActive   ExportSummaryLifecycleStateEnum = "ACTIVE"
	ExportSummaryLifecycleStateDeleting ExportSummaryLifecycleStateEnum = "DELETING"
	ExportSummaryLifecycleStateDeleted  ExportSummaryLifecycleStateEnum = "DELETED"
)

var mappingExportSummaryLifecycleStateEnum = map[string]ExportSummaryLifecycleStateEnum{
	"CREATING": ExportSummaryLifecycleStateCreating,
	"ACTIVE":   ExportSummaryLifecycleStateActive,
	"DELETING": ExportSummaryLifecycleStateDeleting,
	"DELETED":  ExportSummaryLifecycleStateDeleted,
}

var mappingExportSummaryLifecycleStateEnumLowerCase = map[string]ExportSummaryLifecycleStateEnum{
	"creating": ExportSummaryLifecycleStateCreating,
	"active":   ExportSummaryLifecycleStateActive,
	"deleting": ExportSummaryLifecycleStateDeleting,
	"deleted":  ExportSummaryLifecycleStateDeleted,
}

// GetExportSummaryLifecycleStateEnumValues Enumerates the set of values for ExportSummaryLifecycleStateEnum
func GetExportSummaryLifecycleStateEnumValues() []ExportSummaryLifecycleStateEnum {
	values := make([]ExportSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExportSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExportSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExportSummaryLifecycleStateEnum
func GetExportSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingExportSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportSummaryLifecycleStateEnum(val string) (ExportSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExportSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
