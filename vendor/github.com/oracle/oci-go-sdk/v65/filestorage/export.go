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

// Export A file system and the path that you can use to mount it. Each export
// resource belongs to exactly one export set.
// The export's path attribute is not a path in the
// referenced file system, but the value used by clients for the path
// component of the remotetarget argument when mounting the file
// system.
// The path must start with a slash (/) followed by a sequence of zero or more
// slash-separated path elements. For any two export resources associated with
// the same export set, except those in a 'DELETED' state, the path element
// sequence for the first export resource can't contain the
// complete path element sequence of the second export resource.
//
// For example, the following are acceptable:
//   - /example and /path
//   - /example1 and /example2
//   - /example and /example1
//
// The following examples are not acceptable:
//   - /example and /example/path
//   - / and /example
//
// Paths may not end in a slash (/). No path element can be a period (.)
// or two periods in sequence (..). All path elements must be 255 bytes or less.
// No two non-'DELETED' export resources in the same export set can
// reference the same file system.
// Use `exportOptions` to control access to an export. For more information, see
// Export Options (https://docs.cloud.oracle.com/Content/File/Tasks/exportoptions.htm).
type Export struct {

	// Policies that apply to NFS requests made through this
	// export. `exportOptions` contains a sequential list of
	// `ClientOptions`. Each `ClientOptions` item defines the
	// export options that are applied to a specified
	// set of clients.
	// For each NFS request, the first `ClientOptions` option
	// in the list whose `source` attribute matches the source
	// IP address of the request is applied.
	// If a client source IP address does not match the `source`
	// property of any `ClientOptions` in the list, then the
	// export will be invisible to that client. This export will
	// not be returned by `MOUNTPROC_EXPORT` calls made by the client
	// and any attempt to mount or access the file system through
	// this export will result in an error.
	// **Exports without defined `ClientOptions` are invisible to all clients.**
	// If one export is invisible to a particular client, associated file
	// systems may still be accessible through other exports on the same
	// or different mount targets.
	// To completely deny client access to a file system, be sure that the client
	// source IP address is not included in any export for any mount target
	// associated with the file system.
	ExportOptions []ClientOptions `mandatory:"true" json:"exportOptions"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export's export set.
	ExportSetId *string `mandatory:"true" json:"exportSetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export's file system.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export.
	Id *string `mandatory:"true" json:"id"`

	// The current state of this export.
	LifecycleState ExportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Path used to access the associated file system.
	// Avoid entering confidential information.
	// Example: `/accounting`
	Path *string `mandatory:"true" json:"path"`

	// The date and time the export was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Whether or not the export should use ID mapping for Unix groups rather than the group list provided within an NFS request's RPC header. When this flag is true the Unix UID from the RPC header is used to retrieve the list of secondary groups from a the ID mapping subsystem. The primary GID is always taken from the RPC header. If ID mapping is not configured, incorrectly configured, unavailable, or cannot be used to determine a list of secondary groups then an empty secondary group list is used for authorization. If the number of groups exceeds the limit of 256 groups, the list retrieved from LDAP is truncated to the first 256 groups read.
	IsIdmapGroupsForSysAuth *bool `mandatory:"false" json:"isIdmapGroupsForSysAuth"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m Export) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Export) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExportLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportLifecycleStateEnum Enum with underlying type: string
type ExportLifecycleStateEnum string

// Set of constants representing the allowable values for ExportLifecycleStateEnum
const (
	ExportLifecycleStateCreating ExportLifecycleStateEnum = "CREATING"
	ExportLifecycleStateActive   ExportLifecycleStateEnum = "ACTIVE"
	ExportLifecycleStateDeleting ExportLifecycleStateEnum = "DELETING"
	ExportLifecycleStateDeleted  ExportLifecycleStateEnum = "DELETED"
)

var mappingExportLifecycleStateEnum = map[string]ExportLifecycleStateEnum{
	"CREATING": ExportLifecycleStateCreating,
	"ACTIVE":   ExportLifecycleStateActive,
	"DELETING": ExportLifecycleStateDeleting,
	"DELETED":  ExportLifecycleStateDeleted,
}

var mappingExportLifecycleStateEnumLowerCase = map[string]ExportLifecycleStateEnum{
	"creating": ExportLifecycleStateCreating,
	"active":   ExportLifecycleStateActive,
	"deleting": ExportLifecycleStateDeleting,
	"deleted":  ExportLifecycleStateDeleted,
}

// GetExportLifecycleStateEnumValues Enumerates the set of values for ExportLifecycleStateEnum
func GetExportLifecycleStateEnumValues() []ExportLifecycleStateEnum {
	values := make([]ExportLifecycleStateEnum, 0)
	for _, v := range mappingExportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExportLifecycleStateEnumStringValues Enumerates the set of values in String for ExportLifecycleStateEnum
func GetExportLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingExportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportLifecycleStateEnum(val string) (ExportLifecycleStateEnum, bool) {
	enum, ok := mappingExportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
