// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DbSystemSnapshotSummary A summary of snapshot of the DB system details at the time of the backup.
type DbSystemSnapshotSummary struct {

	// The user-friendly name for the DB system. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the DB System.
	Id *string `mandatory:"true" json:"id"`

	// The region identifier of the region where the DB system exists.
	// For more information, please see Regions and Availability Domains (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
	Region *string `mandatory:"true" json:"region"`
}

func (m DbSystemSnapshotSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemSnapshotSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
