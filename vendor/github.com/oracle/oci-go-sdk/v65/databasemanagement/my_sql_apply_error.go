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

// MySqlApplyError Error from the apply operation of a MySQL server replication channel.
type MySqlApplyError struct {

	// The error number of the most recent error that caused the SQL or coordinator thread to stop.
	LastErrorNumber *int `mandatory:"false" json:"lastErrorNumber"`

	// The error message of the most recent error that caused the SQL or coordinator thread to stop.
	LastErrorMessage *string `mandatory:"false" json:"lastErrorMessage"`

	// The timestamp when the most recent SQL or coordinator error occurred.
	TimeLastError *common.SDKTime `mandatory:"false" json:"timeLastError"`

	// A list of MySqlApplyErrorWorker records.
	WorkerErrors []MySqlApplyErrorWorker `mandatory:"false" json:"workerErrors"`
}

func (m MySqlApplyError) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlApplyError) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
