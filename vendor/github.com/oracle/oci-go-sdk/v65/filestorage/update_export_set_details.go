// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateExportSetDetails Details for updating the export set.
type UpdateExportSetDetails struct {

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My export set`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Controls the maximum `tbytes`, `fbytes`, and `abytes`
	// values reported by `NFS FSSTAT` calls through any associated
	// mount targets. This is an advanced feature. For most
	// applications, use the default value. The
	// `tbytes` value reported by `FSSTAT` will be
	// `maxFsStatBytes`. The value of `fbytes` and `abytes` will be
	// `maxFsStatBytes` minus the metered size of the file
	// system. If the metered size is larger than `maxFsStatBytes`,
	// then `fbytes` and `abytes` will both be '0'.
	MaxFsStatBytes *int64 `mandatory:"false" json:"maxFsStatBytes"`

	// Controls the maximum `ffiles`, `ffiles`, and `afiles`
	// values reported by `NFS FSSTAT` calls through any associated
	// mount targets. This is an advanced feature. For most
	// applications, use the default value. The
	// `tfiles` value reported by `FSSTAT` will be
	// `maxFsStatFiles`. The value of `ffiles` and `afiles` will be
	// `maxFsStatFiles` minus the metered size of the file
	// system. If the metered size is larger than `maxFsStatFiles`,
	// then `ffiles` and `afiles` will both be '0'.
	MaxFsStatFiles *int64 `mandatory:"false" json:"maxFsStatFiles"`
}

func (m UpdateExportSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateExportSetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
