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

// ManageDataSafeDetails Details for registering datasafe operation on a database and its pluggable databases
type ManageDataSafeDetails struct {

	// datasafe user account name
	Username *string `mandatory:"false" json:"username"`

	// Current password
	OldPassword *string `mandatory:"false" json:"oldPassword"`

	// New password
	NewPassword *string `mandatory:"false" json:"newPassword"`

	// Features of Data Safe customer has opted-in
	Features []string `mandatory:"false" json:"features"`

	// True if action applies to container database.
	ShouldApplyFeaturesToCdb *bool `mandatory:"false" json:"shouldApplyFeaturesToCdb"`

	// SHA-256 checksum of the one-off patch.
	Sha256Sum *string `mandatory:"false" json:"sha256Sum"`

	// pdbs to register for datasafe
	PluggableDatabases []PdbDataSafeDetail `mandatory:"false" json:"pluggableDatabases"`

	// Supported actions for Datasafe operations
	Action ManageDataSafeDetailsActionEnum `mandatory:"false" json:"action,omitempty"`
}

func (m ManageDataSafeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManageDataSafeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingManageDataSafeDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetManageDataSafeDetailsActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManageDataSafeDetailsActionEnum Enum with underlying type: string
type ManageDataSafeDetailsActionEnum string

// Set of constants representing the allowable values for ManageDataSafeDetailsActionEnum
const (
	ManageDataSafeDetailsActionRegister       ManageDataSafeDetailsActionEnum = "REGISTER"
	ManageDataSafeDetailsActionRotatePassword ManageDataSafeDetailsActionEnum = "ROTATE_PASSWORD"
	ManageDataSafeDetailsActionResetPassword  ManageDataSafeDetailsActionEnum = "RESET_PASSWORD"
	ManageDataSafeDetailsActionUpdateFeature  ManageDataSafeDetailsActionEnum = "UPDATE_FEATURE"
	ManageDataSafeDetailsActionDeregister     ManageDataSafeDetailsActionEnum = "DEREGISTER"
)

var mappingManageDataSafeDetailsActionEnum = map[string]ManageDataSafeDetailsActionEnum{
	"REGISTER":        ManageDataSafeDetailsActionRegister,
	"ROTATE_PASSWORD": ManageDataSafeDetailsActionRotatePassword,
	"RESET_PASSWORD":  ManageDataSafeDetailsActionResetPassword,
	"UPDATE_FEATURE":  ManageDataSafeDetailsActionUpdateFeature,
	"DEREGISTER":      ManageDataSafeDetailsActionDeregister,
}

var mappingManageDataSafeDetailsActionEnumLowerCase = map[string]ManageDataSafeDetailsActionEnum{
	"register":        ManageDataSafeDetailsActionRegister,
	"rotate_password": ManageDataSafeDetailsActionRotatePassword,
	"reset_password":  ManageDataSafeDetailsActionResetPassword,
	"update_feature":  ManageDataSafeDetailsActionUpdateFeature,
	"deregister":      ManageDataSafeDetailsActionDeregister,
}

// GetManageDataSafeDetailsActionEnumValues Enumerates the set of values for ManageDataSafeDetailsActionEnum
func GetManageDataSafeDetailsActionEnumValues() []ManageDataSafeDetailsActionEnum {
	values := make([]ManageDataSafeDetailsActionEnum, 0)
	for _, v := range mappingManageDataSafeDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetManageDataSafeDetailsActionEnumStringValues Enumerates the set of values in String for ManageDataSafeDetailsActionEnum
func GetManageDataSafeDetailsActionEnumStringValues() []string {
	return []string{
		"REGISTER",
		"ROTATE_PASSWORD",
		"RESET_PASSWORD",
		"UPDATE_FEATURE",
		"DEREGISTER",
	}
}

// GetMappingManageDataSafeDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManageDataSafeDetailsActionEnum(val string) (ManageDataSafeDetailsActionEnum, bool) {
	enum, ok := mappingManageDataSafeDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
