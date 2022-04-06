// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalBackupJob Provides all the details that apply to an external backup job.
type ExternalBackupJob struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated backup resource.
	BackupId *string `mandatory:"true" json:"backupId"`

	// An indicator for the provisioning state of the resource. If `TRUE`, the resource is still being provisioned.
	Provisioning *bool `mandatory:"true" json:"provisioning"`

	// The Swift path to use as a destination for the standalone backup.
	SwiftPath *string `mandatory:"true" json:"swiftPath"`

	// The name of the Swift compartment bucket where the backup should be stored.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The tag for RMAN to apply to the backup.
	Tag *string `mandatory:"true" json:"tag"`

	// The Swift user name to use for transferring the standalone backup to the designated Swift compartment bucket.
	UserName *string `mandatory:"true" json:"userName"`

	// The auth token to use for access to the Swift compartment bucket that will store the standalone backup.
	// For information about auth tokens, see Working with Auth Tokens (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingcredentials.htm#two).
	SwiftPassword *string `mandatory:"false" json:"swiftPassword"`
}

func (m ExternalBackupJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalBackupJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
