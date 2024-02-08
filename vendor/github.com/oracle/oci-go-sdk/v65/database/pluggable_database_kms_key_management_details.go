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

// PluggableDatabaseKmsKeyManagementDetails Details for performing the pluggable key updates.
type PluggableDatabaseKmsKeyManagementDetails struct {

	// The specified key update action.
	Action PluggableDatabaseKmsKeyManagementDetailsActionEnum `mandatory:"true" json:"action"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`
}

func (m PluggableDatabaseKmsKeyManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PluggableDatabaseKmsKeyManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPluggableDatabaseKmsKeyManagementDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetPluggableDatabaseKmsKeyManagementDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PluggableDatabaseKmsKeyManagementDetailsActionEnum Enum with underlying type: string
type PluggableDatabaseKmsKeyManagementDetailsActionEnum string

// Set of constants representing the allowable values for PluggableDatabaseKmsKeyManagementDetailsActionEnum
const (
	PluggableDatabaseKmsKeyManagementDetailsActionRotateKey     PluggableDatabaseKmsKeyManagementDetailsActionEnum = "ROTATE_KEY"
	PluggableDatabaseKmsKeyManagementDetailsActionSetKeyVersion PluggableDatabaseKmsKeyManagementDetailsActionEnum = "SET_KEY_VERSION"
)

var mappingPluggableDatabaseKmsKeyManagementDetailsActionEnum = map[string]PluggableDatabaseKmsKeyManagementDetailsActionEnum{
	"ROTATE_KEY":      PluggableDatabaseKmsKeyManagementDetailsActionRotateKey,
	"SET_KEY_VERSION": PluggableDatabaseKmsKeyManagementDetailsActionSetKeyVersion,
}

var mappingPluggableDatabaseKmsKeyManagementDetailsActionEnumLowerCase = map[string]PluggableDatabaseKmsKeyManagementDetailsActionEnum{
	"rotate_key":      PluggableDatabaseKmsKeyManagementDetailsActionRotateKey,
	"set_key_version": PluggableDatabaseKmsKeyManagementDetailsActionSetKeyVersion,
}

// GetPluggableDatabaseKmsKeyManagementDetailsActionEnumValues Enumerates the set of values for PluggableDatabaseKmsKeyManagementDetailsActionEnum
func GetPluggableDatabaseKmsKeyManagementDetailsActionEnumValues() []PluggableDatabaseKmsKeyManagementDetailsActionEnum {
	values := make([]PluggableDatabaseKmsKeyManagementDetailsActionEnum, 0)
	for _, v := range mappingPluggableDatabaseKmsKeyManagementDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetPluggableDatabaseKmsKeyManagementDetailsActionEnumStringValues Enumerates the set of values in String for PluggableDatabaseKmsKeyManagementDetailsActionEnum
func GetPluggableDatabaseKmsKeyManagementDetailsActionEnumStringValues() []string {
	return []string{
		"ROTATE_KEY",
		"SET_KEY_VERSION",
	}
}

// GetMappingPluggableDatabaseKmsKeyManagementDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluggableDatabaseKmsKeyManagementDetailsActionEnum(val string) (PluggableDatabaseKmsKeyManagementDetailsActionEnum, bool) {
	enum, ok := mappingPluggableDatabaseKmsKeyManagementDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
