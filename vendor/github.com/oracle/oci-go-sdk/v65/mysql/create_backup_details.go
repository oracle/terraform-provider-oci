// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateBackupDetailsBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetCreateBackupDetailsBackupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateBackupDetailsBackupTypeEnum Enum with underlying type: string
type CreateBackupDetailsBackupTypeEnum string

// Set of constants representing the allowable values for CreateBackupDetailsBackupTypeEnum
const (
	CreateBackupDetailsBackupTypeFull        CreateBackupDetailsBackupTypeEnum = "FULL"
	CreateBackupDetailsBackupTypeIncremental CreateBackupDetailsBackupTypeEnum = "INCREMENTAL"
)

var mappingCreateBackupDetailsBackupTypeEnum = map[string]CreateBackupDetailsBackupTypeEnum{
	"FULL":        CreateBackupDetailsBackupTypeFull,
	"INCREMENTAL": CreateBackupDetailsBackupTypeIncremental,
}

var mappingCreateBackupDetailsBackupTypeEnumLowerCase = map[string]CreateBackupDetailsBackupTypeEnum{
	"full":        CreateBackupDetailsBackupTypeFull,
	"incremental": CreateBackupDetailsBackupTypeIncremental,
}

// GetCreateBackupDetailsBackupTypeEnumValues Enumerates the set of values for CreateBackupDetailsBackupTypeEnum
func GetCreateBackupDetailsBackupTypeEnumValues() []CreateBackupDetailsBackupTypeEnum {
	values := make([]CreateBackupDetailsBackupTypeEnum, 0)
	for _, v := range mappingCreateBackupDetailsBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBackupDetailsBackupTypeEnumStringValues Enumerates the set of values in String for CreateBackupDetailsBackupTypeEnum
func GetCreateBackupDetailsBackupTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"INCREMENTAL",
	}
}

// GetMappingCreateBackupDetailsBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBackupDetailsBackupTypeEnum(val string) (CreateBackupDetailsBackupTypeEnum, bool) {
	enum, ok := mappingCreateBackupDetailsBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
