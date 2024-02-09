// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseKmsKeyManagementDetails Details for performing key updates.
type DatabaseKmsKeyManagementDetails struct {

	// The specified key update action.
	Action DatabaseKmsKeyManagementDetailsActionEnum `mandatory:"true" json:"action"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`
}

func (m DatabaseKmsKeyManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseKmsKeyManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseKmsKeyManagementDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDatabaseKmsKeyManagementDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseKmsKeyManagementDetailsActionEnum Enum with underlying type: string
type DatabaseKmsKeyManagementDetailsActionEnum string

// Set of constants representing the allowable values for DatabaseKmsKeyManagementDetailsActionEnum
const (
	DatabaseKmsKeyManagementDetailsActionSetKeyVersion    DatabaseKmsKeyManagementDetailsActionEnum = "SET_KEY_VERSION"
	DatabaseKmsKeyManagementDetailsActionRotateKey        DatabaseKmsKeyManagementDetailsActionEnum = "ROTATE_KEY"
	DatabaseKmsKeyManagementDetailsActionMigrateFileToHsm DatabaseKmsKeyManagementDetailsActionEnum = "MIGRATE_FILE_TO_HSM"
)

var mappingDatabaseKmsKeyManagementDetailsActionEnum = map[string]DatabaseKmsKeyManagementDetailsActionEnum{
	"SET_KEY_VERSION":     DatabaseKmsKeyManagementDetailsActionSetKeyVersion,
	"ROTATE_KEY":          DatabaseKmsKeyManagementDetailsActionRotateKey,
	"MIGRATE_FILE_TO_HSM": DatabaseKmsKeyManagementDetailsActionMigrateFileToHsm,
}

var mappingDatabaseKmsKeyManagementDetailsActionEnumLowerCase = map[string]DatabaseKmsKeyManagementDetailsActionEnum{
	"set_key_version":     DatabaseKmsKeyManagementDetailsActionSetKeyVersion,
	"rotate_key":          DatabaseKmsKeyManagementDetailsActionRotateKey,
	"migrate_file_to_hsm": DatabaseKmsKeyManagementDetailsActionMigrateFileToHsm,
}

// GetDatabaseKmsKeyManagementDetailsActionEnumValues Enumerates the set of values for DatabaseKmsKeyManagementDetailsActionEnum
func GetDatabaseKmsKeyManagementDetailsActionEnumValues() []DatabaseKmsKeyManagementDetailsActionEnum {
	values := make([]DatabaseKmsKeyManagementDetailsActionEnum, 0)
	for _, v := range mappingDatabaseKmsKeyManagementDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseKmsKeyManagementDetailsActionEnumStringValues Enumerates the set of values in String for DatabaseKmsKeyManagementDetailsActionEnum
func GetDatabaseKmsKeyManagementDetailsActionEnumStringValues() []string {
	return []string{
		"SET_KEY_VERSION",
		"ROTATE_KEY",
		"MIGRATE_FILE_TO_HSM",
	}
}

// GetMappingDatabaseKmsKeyManagementDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseKmsKeyManagementDetailsActionEnum(val string) (DatabaseKmsKeyManagementDetailsActionEnum, bool) {
	enum, ok := mappingDatabaseKmsKeyManagementDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
