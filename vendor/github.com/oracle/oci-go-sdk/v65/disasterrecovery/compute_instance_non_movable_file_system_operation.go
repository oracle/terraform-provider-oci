// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeInstanceNonMovableFileSystemOperation The details of operations performed on a file system.
type ComputeInstanceNonMovableFileSystemOperation struct {

	// The export path of the file system.
	// Example: `/fs-export-path`
	ExportPath *string `mandatory:"true" json:"exportPath"`

	// The physical mount point of the file system on a host.
	// Example: `/mnt/yourmountpoint`
	MountPoint *string `mandatory:"true" json:"mountPoint"`

	// The OCID of mount target.
	// Example: `ocid1.mounttarget.oc1..uniqueID`
	MountTargetId *string `mandatory:"true" json:"mountTargetId"`
}

func (m ComputeInstanceNonMovableFileSystemOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeInstanceNonMovableFileSystemOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
