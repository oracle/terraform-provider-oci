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

// SaasAdminUserStatus SaaS administrative user status.
type SaasAdminUserStatus struct {

	// Indicates if the SaaS administrative user is enabled for the Autonomous Database.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The access type for the SaaS administrative user. If no access type is specified, the READ_ONLY access type is used.
	AccessType SaasAdminUserStatusAccessTypeEnum `mandatory:"false" json:"accessType,omitempty"`

	// The date and time the SaaS administrative user was enabled at, for the Autonomous Database.
	TimeSaasAdminUserEnabled *common.SDKTime `mandatory:"false" json:"timeSaasAdminUserEnabled"`
}

func (m SaasAdminUserStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SaasAdminUserStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSaasAdminUserStatusAccessTypeEnum(string(m.AccessType)); !ok && m.AccessType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessType: %s. Supported values are: %s.", m.AccessType, strings.Join(GetSaasAdminUserStatusAccessTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SaasAdminUserStatusAccessTypeEnum Enum with underlying type: string
type SaasAdminUserStatusAccessTypeEnum string

// Set of constants representing the allowable values for SaasAdminUserStatusAccessTypeEnum
const (
	SaasAdminUserStatusAccessTypeReadOnly  SaasAdminUserStatusAccessTypeEnum = "READ_ONLY"
	SaasAdminUserStatusAccessTypeReadWrite SaasAdminUserStatusAccessTypeEnum = "READ_WRITE"
	SaasAdminUserStatusAccessTypeAdmin     SaasAdminUserStatusAccessTypeEnum = "ADMIN"
)

var mappingSaasAdminUserStatusAccessTypeEnum = map[string]SaasAdminUserStatusAccessTypeEnum{
	"READ_ONLY":  SaasAdminUserStatusAccessTypeReadOnly,
	"READ_WRITE": SaasAdminUserStatusAccessTypeReadWrite,
	"ADMIN":      SaasAdminUserStatusAccessTypeAdmin,
}

var mappingSaasAdminUserStatusAccessTypeEnumLowerCase = map[string]SaasAdminUserStatusAccessTypeEnum{
	"read_only":  SaasAdminUserStatusAccessTypeReadOnly,
	"read_write": SaasAdminUserStatusAccessTypeReadWrite,
	"admin":      SaasAdminUserStatusAccessTypeAdmin,
}

// GetSaasAdminUserStatusAccessTypeEnumValues Enumerates the set of values for SaasAdminUserStatusAccessTypeEnum
func GetSaasAdminUserStatusAccessTypeEnumValues() []SaasAdminUserStatusAccessTypeEnum {
	values := make([]SaasAdminUserStatusAccessTypeEnum, 0)
	for _, v := range mappingSaasAdminUserStatusAccessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSaasAdminUserStatusAccessTypeEnumStringValues Enumerates the set of values in String for SaasAdminUserStatusAccessTypeEnum
func GetSaasAdminUserStatusAccessTypeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
		"ADMIN",
	}
}

// GetMappingSaasAdminUserStatusAccessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSaasAdminUserStatusAccessTypeEnum(val string) (SaasAdminUserStatusAccessTypeEnum, bool) {
	enum, ok := mappingSaasAdminUserStatusAccessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
