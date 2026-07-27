// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpgradeDbSystemDetails Upgrade details for the database system.
type UpgradeDbSystemDetails struct {

	// The target major version of the database system software.
	// Example: "17"
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration for target database
	// version to be associated with database system.
	// Optional, if database system is currently running with default configuration object.
	// Mandatory, if database system is currently running with custom configuration object.
	ConfigId *string `mandatory:"false" json:"configId"`
}

func (m UpgradeDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpgradeDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
