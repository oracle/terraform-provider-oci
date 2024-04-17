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

// CopyBackupDetails Details required to copy a DB system backup from its source region to a destination region.
type CopyBackupDetails struct {

	// The OCID of the compartment the DB system backup is to be copied to.
	// **Note:** The compartment must be the same as the compartment of the DB system backup in the source region.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of DB system backup to be copied.
	SourceBackupId *string `mandatory:"true" json:"sourceBackupId"`

	// The region identifier of the source region where the DB system backup exists.
	// For more information, please see Regions and Availability Domains (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
	SourceRegion *string `mandatory:"true" json:"sourceRegion"`

	// A user-supplied description for the DB system backup.
	// By default, the source backup description will be used.
	Description *string `mandatory:"false" json:"description"`

	// A user-supplied display name for the DB system backup.
	// By default, the source backup display name will be used.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CopyBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CopyBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
