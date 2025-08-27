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

// ManagedMySqlDatabaseHighAvailabilityMemberSummary Information about a member of a MySQL server group replication for high availability.
type ManagedMySqlDatabaseHighAvailabilityMemberSummary struct {

	// The host name of the group member that clients use to connect to it.
	MemberHost *string `mandatory:"true" json:"memberHost"`

	// The port number of the group member that clients use to connect to it.
	MemberPort *int `mandatory:"true" json:"memberPort"`

	// The Universally Unique Identifier (UUID) of the member server.
	MemberUuid *string `mandatory:"true" json:"memberUuid"`

	// The current state of the group member.
	MemberState *string `mandatory:"false" json:"memberState"`

	// The current role of the group member in the group.
	MemberRole *string `mandatory:"false" json:"memberRole"`
}

func (m ManagedMySqlDatabaseHighAvailabilityMemberSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedMySqlDatabaseHighAvailabilityMemberSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
