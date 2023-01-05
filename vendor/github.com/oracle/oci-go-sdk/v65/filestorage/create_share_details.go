// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CreateShareDetails Details for creating the share.
type CreateShareDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the share set of this share is in.
	ShareSetId *string `mandatory:"true" json:"shareSetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this share's file system.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// Share name used to access the associated file system.
	// Avoid entering confidential information.
	// Example: `mediafiles`
	ShareName *string `mandatory:"true" json:"shareName"`

	// Share options for the new share. If left unspecified,
	// defaults to:
	//        [
	//          {
	//             "source" : "0.0.0.0/0",
	//             "access" : "READ_WRITE",
	//           }
	//        ]
	//   **Note:** Mount targets do not have Internet-routable IP
	//   addresses.  Therefore they will not be reachable from the
	//   Internet, even if an associated `ClientShareOptions` item has
	//   a source of `0.0.0.0/0`.
	//   **If set to the empty array then the share will not be
	//   visible to any clients.**
	//   The share's `shareOptions` can be changed after creation
	//   using the `UpdateShare` operation.
	ShareOptions []ClientShareOptions `mandatory:"false" json:"shareOptions"`

	// A short comment description of the Share.
	// Avoid entering confidential information.
	// Example: `accounting`
	Comment *string `mandatory:"false" json:"comment"`

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

func (m CreateShareDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateShareDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
