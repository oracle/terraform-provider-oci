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

// MySqlQuerySampleDetails The details of a query sample including the query text, execution time and other details.
type MySqlQuerySampleDetails struct {

	// The query sample mapped by MySQL to a given normalized query.
	QuerySampleText *string `mandatory:"true" json:"querySampleText"`

	// The date and time the query sample was last seen.
	TimeQuerySampleSeen *common.SDKTime `mandatory:"true" json:"timeQuerySampleSeen"`

	// The total amount of time that has been spent executing the query sample.
	ExecutionTime *int64 `mandatory:"true" json:"executionTime"`

	// The thread id of the connection.
	ThreadId *int `mandatory:"true" json:"threadId"`

	// The user who ran the query sample.
	User *string `mandatory:"true" json:"user"`

	// The host from which the query sample was run.
	Host *string `mandatory:"true" json:"host"`

	// The MySQL instance against which the query sample was run.
	MysqlInstance *string `mandatory:"true" json:"mysqlInstance"`
}

func (m MySqlQuerySampleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlQuerySampleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
