// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Disk The assets disk.
type Disk struct {

	// Disk name.
	Name *string `mandatory:"false" json:"name"`

	// Order of boot volumes.
	BootOrder *int `mandatory:"false" json:"bootOrder"`

	// Disk UUID for the virtual disk, if available.
	Uuid *string `mandatory:"false" json:"uuid"`

	// Disk UUID LUN for the virtual disk, if available.
	UuidLun *string `mandatory:"false" json:"uuidLun"`

	// The size of the volume in MBs.
	SizeInMBs *int64 `mandatory:"false" json:"sizeInMBs"`

	// Location of the boot/data volume.
	Location *string `mandatory:"false" json:"location"`

	// The disk persistent mode.
	PersistentMode *string `mandatory:"false" json:"persistentMode"`
}

func (m Disk) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Disk) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
