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

// ManagedMySqlDatabaseBinaryLogInformation Information pertaining to the binary logs of a MySQL server.
type ManagedMySqlDatabaseBinaryLogInformation struct {

	// The status of binary logging on the MySQL server.
	BinaryLogging *string `mandatory:"true" json:"binaryLogging"`

	// The binary logging format used by the MySQL server.
	BinaryLogFormat *string `mandatory:"false" json:"binaryLogFormat"`

	// Indicates whether compression is enabled for transactions written to binary log files on the MySQL server.
	BinaryLogCompression *string `mandatory:"false" json:"binaryLogCompression"`

	// The compression ratio for the binary log, expressed as a percentage.
	BinaryLogCompressionPercent *int `mandatory:"false" json:"binaryLogCompressionPercent"`

	// The name of the binary log file.
	BinaryLogName *string `mandatory:"false" json:"binaryLogName"`

	// The position within the binary log file.
	BinaryLogPosition *int64 `mandatory:"false" json:"binaryLogPosition"`
}

func (m ManagedMySqlDatabaseBinaryLogInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedMySqlDatabaseBinaryLogInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
