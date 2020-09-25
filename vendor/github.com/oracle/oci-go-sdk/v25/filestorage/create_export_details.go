// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreateExportDetails Details for creating the export.
type CreateExportDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export's export set.
	ExportSetId *string `mandatory:"true" json:"exportSetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this export's file system.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// Path used to access the associated file system.
	// Avoid entering confidential information.
	// Example: `/mediafiles`
	Path *string `mandatory:"true" json:"path"`

	// Export options for the new export. If left unspecified,
	// defaults to:
	//        [
	//          {
	//             "source" : "0.0.0.0/0",
	//             "requirePrivilegedSourcePort" : false,
	//             "access" : "READ_WRITE",
	//             "identitySquash" : "NONE"
	//           }
	//        ]
	//   **Note:** Mount targets do not have Internet-routable IP
	//   addresses.  Therefore they will not be reachable from the
	//   Internet, even if an associated `ClientOptions` item has
	//   a source of `0.0.0.0/0`.
	//   **If set to the empty array then the export will not be
	//   visible to any clients.**
	//   The export's `exportOptions` can be changed after creation
	//   using the `UpdateExport` operation.
	ExportOptions []ClientOptions `mandatory:"false" json:"exportOptions"`

	// The export is modified to include a boolean to use ID mapping for Unix Groups rather than the group list provided within an NFS Request's RPC header. When this flag is true the Unix UID from the RPC header is used to retrieve the list of secondary groups from a the ID mapping subsystem. The primary GID is always taken from the RPC header. If ID mapping is not configured, incorrectly configured, unavailable or cannot be used to determine a list of secondary groups then the data path uses an empty secondary group list for authorization. If the number of groups exceeds the current limit of 256 groups the list retrieved from LDAP is truncated to the first 256 groups read.
	IsIdmapGroupsForSysAuth *bool `mandatory:"false" json:"isIdmapGroupsForSysAuth"`
}

func (m CreateExportDetails) String() string {
	return common.PointerString(m)
}
