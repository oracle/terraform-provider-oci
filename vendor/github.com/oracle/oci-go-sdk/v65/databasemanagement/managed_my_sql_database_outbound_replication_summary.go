// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedMySqlDatabaseOutboundReplicationSummary An outbound replication record of a MySQL server.
type ManagedMySqlDatabaseOutboundReplicationSummary struct {

	// The Universally Unique Identifier (UUID) value of the replica server.
	ReplicaUuid *string `mandatory:"true" json:"replicaUuid"`

	// The server ID value of the replica.
	ReplicaServerId *int64 `mandatory:"true" json:"replicaServerId"`

	// The host name of the replica server, as specified on the replica with the --report-host option. This can differ from the machine name as configured in the operating system.
	ReplicaHost *string `mandatory:"false" json:"replicaHost"`

	// The port on the replica server, as specified on the replica with the --report-port option. A zero in this column means that the replica port (--report-port) was not set.
	ReplicaPort *int `mandatory:"false" json:"replicaPort"`
}

func (m ManagedMySqlDatabaseOutboundReplicationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedMySqlDatabaseOutboundReplicationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
