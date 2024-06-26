// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMySqlInitialLoadSettings Optional dump settings
type CreateMySqlInitialLoadSettings struct {

	// MySql Job Mode
	JobMode JobModeMySqlEnum `mandatory:"true" json:"jobMode"`

	// Enable (true) or disable (false) consistent data dumps by locking the instance for backup during the dump.
	IsConsistent *bool `mandatory:"false" json:"isConsistent"`

	// Include a statement at the start of the dump to set the time zone to UTC.
	IsTzUtc *bool `mandatory:"false" json:"isTzUtc"`

	// Apply the specified requirements for compatibility with MySQL Database Service for all tables in the dump
	// output, altering the dump files as necessary.
	Compatibility []CompatibilityOptionEnum `mandatory:"false" json:"compatibility"`

	// Primary key compatibility option
	PrimaryKeyCompatibility PrimaryKeyCompatibilityEnum `mandatory:"false" json:"primaryKeyCompatibility,omitempty"`

	// Import the dump even if it contains objects that already exist in the target schema in the MySQL instance.
	IsIgnoreExistingObjects *bool `mandatory:"false" json:"isIgnoreExistingObjects"`

	// The action taken in the event of errors related to GRANT or REVOKE errors.
	HandleGrantErrors HandleGrantErrorsEnum `mandatory:"false" json:"handleGrantErrors,omitempty"`
}

func (m CreateMySqlInitialLoadSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMySqlInitialLoadSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobModeMySqlEnum(string(m.JobMode)); !ok && m.JobMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobMode: %s. Supported values are: %s.", m.JobMode, strings.Join(GetJobModeMySqlEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPrimaryKeyCompatibilityEnum(string(m.PrimaryKeyCompatibility)); !ok && m.PrimaryKeyCompatibility != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrimaryKeyCompatibility: %s. Supported values are: %s.", m.PrimaryKeyCompatibility, strings.Join(GetPrimaryKeyCompatibilityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHandleGrantErrorsEnum(string(m.HandleGrantErrors)); !ok && m.HandleGrantErrors != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HandleGrantErrors: %s. Supported values are: %s.", m.HandleGrantErrors, strings.Join(GetHandleGrantErrorsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
