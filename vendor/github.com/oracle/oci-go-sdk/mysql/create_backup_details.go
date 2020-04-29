// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateBackupDetails Complete information for a Backup.
type CreateBackupDetails struct {

	// The OCID of the DB System the Backup is associated with.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// A user-supplied display name for the backup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-supplied description for the backup.
	Description *string `mandatory:"false" json:"description"`

	// The type of backup.
	BackupType CreateBackupDetailsBackupTypeEnum `mandatory:"false" json:"backupType,omitempty"`

	// Number of days to retain this backup.
	RetentionInDays *int `mandatory:"false" json:"retentionInDays"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBackupDetails) String() string {
	return common.PointerString(m)
}

// CreateBackupDetailsBackupTypeEnum Enum with underlying type: string
type CreateBackupDetailsBackupTypeEnum string

// Set of constants representing the allowable values for CreateBackupDetailsBackupTypeEnum
const (
	CreateBackupDetailsBackupTypeFull        CreateBackupDetailsBackupTypeEnum = "FULL"
	CreateBackupDetailsBackupTypeIncremental CreateBackupDetailsBackupTypeEnum = "INCREMENTAL"
)

var mappingCreateBackupDetailsBackupType = map[string]CreateBackupDetailsBackupTypeEnum{
	"FULL":        CreateBackupDetailsBackupTypeFull,
	"INCREMENTAL": CreateBackupDetailsBackupTypeIncremental,
}

// GetCreateBackupDetailsBackupTypeEnumValues Enumerates the set of values for CreateBackupDetailsBackupTypeEnum
func GetCreateBackupDetailsBackupTypeEnumValues() []CreateBackupDetailsBackupTypeEnum {
	values := make([]CreateBackupDetailsBackupTypeEnum, 0)
	for _, v := range mappingCreateBackupDetailsBackupType {
		values = append(values, v)
	}
	return values
}
