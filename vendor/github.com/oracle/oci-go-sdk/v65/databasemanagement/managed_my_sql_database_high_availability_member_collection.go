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

// ManagedMySqlDatabaseHighAvailabilityMemberCollection Information pertaining to high availability of a MySQL server.
type ManagedMySqlDatabaseHighAvailabilityMemberCollection struct {

	// A list of MySqlHighAvailabilityMember records.
	Items []ManagedMySqlDatabaseHighAvailabilityMemberSummary `mandatory:"true" json:"items"`

	// The name of the group to which this server instance belongs.
	GroupName *string `mandatory:"false" json:"groupName"`

	// Indicates if the replication group is running in single-primary mode.
	SinglePrimaryMode *string `mandatory:"false" json:"singlePrimaryMode"`

	// The interval between successive values for auto-incremented columns for transactions that execute on this server instance.
	GroupAutoIncrement *int `mandatory:"false" json:"groupAutoIncrement"`

	// The mode used for flow control.
	FlowControl *string `mandatory:"false" json:"flowControl"`

	// The state of this server as a group replication member.
	MemberState *string `mandatory:"false" json:"memberState"`

	// The role of this server as a group replication member.
	MemberRole *string `mandatory:"false" json:"memberRole"`

	// The current view identifier for this group.
	ViewId *string `mandatory:"false" json:"viewId"`

	// The number of transactions that were replicated within the cluster.
	TransactionsInGtidExecuted *int64 `mandatory:"false" json:"transactionsInGtidExecuted"`

	StatusSummary *MySqlHighAvailabilityStatusSummary `mandatory:"false" json:"statusSummary"`
}

func (m ManagedMySqlDatabaseHighAvailabilityMemberCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedMySqlDatabaseHighAvailabilityMemberCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
