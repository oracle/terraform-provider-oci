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

// CreateObjectStorageLinkDetails The details to create a Object Storage link.
type CreateObjectStorageLinkDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Object Storage link.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain that the Lustre file system is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

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

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My Object Storage Link`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateObjectStorageLinkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateObjectStorageLinkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
