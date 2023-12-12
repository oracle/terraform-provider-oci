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

// ConnectionDetails Database system connection information. Used to connect to PostgreSQL instance(s).
type ConnectionDetails struct {

	// The CA certificate to be used by the PosgreSQL client to connect to the database.
	// The CA certificate is used to authenticate the server identity.
	// It is issued by PostgreSQL Service Private CA.
	CaCertificate *string `mandatory:"true" json:"caCertificate"`

	PrimaryDbEndpoint *Endpoint `mandatory:"true" json:"primaryDbEndpoint"`

	// The list of database instance node endpoints in the database system.
	InstanceEndpoints []DbInstanceEndpoint `mandatory:"true" json:"instanceEndpoints"`
}

func (m ConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
