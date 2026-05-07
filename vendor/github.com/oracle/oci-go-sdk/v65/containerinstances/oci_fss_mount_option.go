// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciFssMountOption A mount option to be used in the mount command while mounting the OCI File Storage Service (FSS) File System to Containers.
// The mount option will look like option=value in the mount command.
type OciFssMountOption struct {

	// A generic (https://man7.org/linux/man-pages/man8/mount.8.html) or nfs (https://man7.org/linux/man-pages/man5/nfs.5.html) mount option.
	Option *string `mandatory:"false" json:"option"`

	// The value of the mount option.
	Value *string `mandatory:"false" json:"value"`
}

func (m OciFssMountOption) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciFssMountOption) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
