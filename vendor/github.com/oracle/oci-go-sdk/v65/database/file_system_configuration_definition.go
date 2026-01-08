// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FileSystemConfigurationDefinition Details of the file system definition of the Exadata VM cluster.
type FileSystemConfigurationDefinition struct {

	// The mount point of file system.
	MountPoint *string `mandatory:"false" json:"mountPoint"`

	// The minimum size of file system.
	MinSizeGBs *int `mandatory:"false" json:"minSizeGBs"`

	// The maximum size of file system.
	MaxSizeGBs *int `mandatory:"false" json:"maxSizeGBs"`

	// If true, the file system resize is allowed for the Exadata VM cluster. If false, the file system resize is not allowed.
	IsResizable *bool `mandatory:"false" json:"isResizable"`

	// If true, the file system is used to create a backup prior to Exadata VM OS update.
	IsBackupPartition *bool `mandatory:"false" json:"isBackupPartition"`
}

func (m FileSystemConfigurationDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileSystemConfigurationDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
