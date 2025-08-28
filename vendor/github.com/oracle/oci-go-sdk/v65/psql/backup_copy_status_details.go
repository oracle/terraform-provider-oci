// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// BackupCopyStatusDetails Backup Copy Status details
type BackupCopyStatusDetails struct {

	// Region name of the remote region
	Region *string `mandatory:"true" json:"region"`

	// Copy States
	State BackupCopyStatusDetailsStateEnum `mandatory:"false" json:"state,omitempty"`

	// A message describing the current state of copy in more detail
	StateDetails *string `mandatory:"false" json:"stateDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup in the source region
	BackupId *string `mandatory:"false" json:"backupId"`
}

func (m BackupCopyStatusDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupCopyStatusDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBackupCopyStatusDetailsStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetBackupCopyStatusDetailsStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupCopyStatusDetailsStateEnum Enum with underlying type: string
type BackupCopyStatusDetailsStateEnum string

// Set of constants representing the allowable values for BackupCopyStatusDetailsStateEnum
const (
	BackupCopyStatusDetailsStateNotStarted BackupCopyStatusDetailsStateEnum = "NOT_STARTED"
	BackupCopyStatusDetailsStateCopying    BackupCopyStatusDetailsStateEnum = "COPYING"
	BackupCopyStatusDetailsStateCopied     BackupCopyStatusDetailsStateEnum = "COPIED"
	BackupCopyStatusDetailsStateFailed     BackupCopyStatusDetailsStateEnum = "FAILED"
)

var mappingBackupCopyStatusDetailsStateEnum = map[string]BackupCopyStatusDetailsStateEnum{
	"NOT_STARTED": BackupCopyStatusDetailsStateNotStarted,
	"COPYING":     BackupCopyStatusDetailsStateCopying,
	"COPIED":      BackupCopyStatusDetailsStateCopied,
	"FAILED":      BackupCopyStatusDetailsStateFailed,
}

var mappingBackupCopyStatusDetailsStateEnumLowerCase = map[string]BackupCopyStatusDetailsStateEnum{
	"not_started": BackupCopyStatusDetailsStateNotStarted,
	"copying":     BackupCopyStatusDetailsStateCopying,
	"copied":      BackupCopyStatusDetailsStateCopied,
	"failed":      BackupCopyStatusDetailsStateFailed,
}

// GetBackupCopyStatusDetailsStateEnumValues Enumerates the set of values for BackupCopyStatusDetailsStateEnum
func GetBackupCopyStatusDetailsStateEnumValues() []BackupCopyStatusDetailsStateEnum {
	values := make([]BackupCopyStatusDetailsStateEnum, 0)
	for _, v := range mappingBackupCopyStatusDetailsStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupCopyStatusDetailsStateEnumStringValues Enumerates the set of values in String for BackupCopyStatusDetailsStateEnum
func GetBackupCopyStatusDetailsStateEnumStringValues() []string {
	return []string{
		"NOT_STARTED",
		"COPYING",
		"COPIED",
		"FAILED",
	}
}

// GetMappingBackupCopyStatusDetailsStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupCopyStatusDetailsStateEnum(val string) (BackupCopyStatusDetailsStateEnum, bool) {
	enum, ok := mappingBackupCopyStatusDetailsStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
