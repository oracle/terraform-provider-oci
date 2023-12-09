// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RestoreDbSystemDetails Backup details to restore the database system.
type RestoreDbSystemDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database system backup.
	BackupId *string `mandatory:"true" json:"backupId"`

	// The desired AD for regions with three ADs. This parameter is optional.
	// If not set, the AD is chosen based on the database system's current AD.
	Ad *string `mandatory:"false" json:"ad"`
}

func (m RestoreDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestoreDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
