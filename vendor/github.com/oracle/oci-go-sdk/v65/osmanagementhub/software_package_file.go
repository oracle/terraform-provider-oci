// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwarePackageFile A file associated with a package.
type SoftwarePackageFile struct {

	// File path.
	Path *string `mandatory:"false" json:"path"`

	// Type of the file.
	Type *string `mandatory:"false" json:"type"`

	// The date and time of the last modification to this file, as described
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`

	// Checksum of the file.
	Checksum *string `mandatory:"false" json:"checksum"`

	// Type of the checksum.
	ChecksumType *string `mandatory:"false" json:"checksumType"`

	// Size of the file in bytes.
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`
}

func (m SoftwarePackageFile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwarePackageFile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
