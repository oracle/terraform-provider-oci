// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SoftwarePackageFile A file associated with a package
type SoftwarePackageFile struct {

	// file path
	Path *string `mandatory:"false" json:"path"`

	// type of the file
	Type *string `mandatory:"false" json:"type"`

	// The date and time of the last modification to this file, as described
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`

	// checksum of the file
	Checksum *string `mandatory:"false" json:"checksum"`

	// type of the checksum
	ChecksumType *string `mandatory:"false" json:"checksumType"`

	// size of the file in bytes
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`
}

func (m SoftwarePackageFile) String() string {
	return common.PointerString(m)
}
