// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TdeWalletBackupDestination Backup destination for the TDE wallet backups.
type TdeWalletBackupDestination struct {

	// Destination where TDE Wallet backups are to be placed.
	BackupDestinationType TdeWalletBackupDestinationBackupDestinationTypeEnum `mandatory:"true" json:"backupDestinationType"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
	BackupDestinationId *string `mandatory:"false" json:"backupDestinationId"`
}

func (m TdeWalletBackupDestination) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TdeWalletBackupDestination) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTdeWalletBackupDestinationBackupDestinationTypeEnum(string(m.BackupDestinationType)); !ok && m.BackupDestinationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupDestinationType: %s. Supported values are: %s.", m.BackupDestinationType, strings.Join(GetTdeWalletBackupDestinationBackupDestinationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TdeWalletBackupDestinationBackupDestinationTypeEnum Enum with underlying type: string
type TdeWalletBackupDestinationBackupDestinationTypeEnum string

// Set of constants representing the allowable values for TdeWalletBackupDestinationBackupDestinationTypeEnum
const (
	TdeWalletBackupDestinationBackupDestinationTypeOss TdeWalletBackupDestinationBackupDestinationTypeEnum = "OSS"
	TdeWalletBackupDestinationBackupDestinationTypeNfs TdeWalletBackupDestinationBackupDestinationTypeEnum = "NFS"
)

var mappingTdeWalletBackupDestinationBackupDestinationTypeEnum = map[string]TdeWalletBackupDestinationBackupDestinationTypeEnum{
	"OSS": TdeWalletBackupDestinationBackupDestinationTypeOss,
	"NFS": TdeWalletBackupDestinationBackupDestinationTypeNfs,
}

var mappingTdeWalletBackupDestinationBackupDestinationTypeEnumLowerCase = map[string]TdeWalletBackupDestinationBackupDestinationTypeEnum{
	"oss": TdeWalletBackupDestinationBackupDestinationTypeOss,
	"nfs": TdeWalletBackupDestinationBackupDestinationTypeNfs,
}

// GetTdeWalletBackupDestinationBackupDestinationTypeEnumValues Enumerates the set of values for TdeWalletBackupDestinationBackupDestinationTypeEnum
func GetTdeWalletBackupDestinationBackupDestinationTypeEnumValues() []TdeWalletBackupDestinationBackupDestinationTypeEnum {
	values := make([]TdeWalletBackupDestinationBackupDestinationTypeEnum, 0)
	for _, v := range mappingTdeWalletBackupDestinationBackupDestinationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTdeWalletBackupDestinationBackupDestinationTypeEnumStringValues Enumerates the set of values in String for TdeWalletBackupDestinationBackupDestinationTypeEnum
func GetTdeWalletBackupDestinationBackupDestinationTypeEnumStringValues() []string {
	return []string{
		"OSS",
		"NFS",
	}
}

// GetMappingTdeWalletBackupDestinationBackupDestinationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTdeWalletBackupDestinationBackupDestinationTypeEnum(val string) (TdeWalletBackupDestinationBackupDestinationTypeEnum, bool) {
	enum, ok := mappingTdeWalletBackupDestinationBackupDestinationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
