// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportBackupDetails The parameters required to export a DB system backup.
type ExportBackupDetails struct {

	// The Object Storage bucket name.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage namespace.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	ExportOptions *ExportOptions `mandatory:"false" json:"exportOptions"`

	// The OCID of the subnet which the temporary MySQL instance is associated with.
	// If not specified, the subnet OCID of the DB system from which the backup is taken will be used by default.
	// To export a backup copied from another region, it is mandatory to specify a valid subnet OCID in the current region.
	SubnetId *string `mandatory:"false" json:"subnetId"`
}

func (m ExportBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
