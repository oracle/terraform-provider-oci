// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// FileSystemExportMappingDetails The mapping between a file system export in the primary region and a mount target in the standby region.
type FileSystemExportMappingDetails struct {

	// The OCID of the export path in the primary region used to mount or unmount the file system.
	// Example: `ocid1.export.oc1..uniqueID`
	ExportId *string `mandatory:"true" json:"exportId"`

	// The OCID of the destination mount target in the destination region which is used to export the file system.
	// Example: `ocid1.mounttarget.oc1..uniqueID`
	DestinationMountTargetId *string `mandatory:"true" json:"destinationMountTargetId"`
}

func (m FileSystemExportMappingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileSystemExportMappingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
