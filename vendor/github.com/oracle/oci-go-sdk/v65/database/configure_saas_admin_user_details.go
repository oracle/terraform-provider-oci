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

// ConfigureSaasAdminUserDetails Details to update SaaS administrative user configuration.
type ConfigureSaasAdminUserDetails struct {

	// Indicates if the SaaS administrative user is enabled for the Autonomous Database.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The access type for the SaaS administrative user. If no access type is specified, the READ_ONLY access type is used.
	AccessType ConfigureSaasAdminUserDetailsAccessTypeEnum `mandatory:"false" json:"accessType,omitempty"`

	// The date and time the SaaS administrative user was enabled at, for the Autonomous Database.
	TimeSaasAdminUserEnabled *common.SDKTime `mandatory:"false" json:"timeSaasAdminUserEnabled"`

	// A strong password for SaaS administrative user. The password must be a minimum of nine (9) characters and contain a minimum of two (2) uppercase, two (2) lowercase, two (2) numbers, and two (2) special characters from _ (underscore), \# (hashtag), or - (dash).
	Password *string `mandatory:"false" json:"password"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure secret (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	SecretId *string `mandatory:"false" json:"secretId"`

	// The version of the vault secret. If no version is specified, the latest version will be used.
	SecretVersionNumber *int `mandatory:"false" json:"secretVersionNumber"`

	// How long, in hours, the SaaS administrative user will stay enabled. If no duration is specified, the default value 1 will be used.
	Duration *int `mandatory:"false" json:"duration"`
}

func (m ConfigureSaasAdminUserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigureSaasAdminUserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigureSaasAdminUserDetailsAccessTypeEnum(string(m.AccessType)); !ok && m.AccessType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessType: %s. Supported values are: %s.", m.AccessType, strings.Join(GetConfigureSaasAdminUserDetailsAccessTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigureSaasAdminUserDetailsAccessTypeEnum Enum with underlying type: string
type ConfigureSaasAdminUserDetailsAccessTypeEnum string

// Set of constants representing the allowable values for ConfigureSaasAdminUserDetailsAccessTypeEnum
const (
	ConfigureSaasAdminUserDetailsAccessTypeReadOnly  ConfigureSaasAdminUserDetailsAccessTypeEnum = "READ_ONLY"
	ConfigureSaasAdminUserDetailsAccessTypeReadWrite ConfigureSaasAdminUserDetailsAccessTypeEnum = "READ_WRITE"
	ConfigureSaasAdminUserDetailsAccessTypeAdmin     ConfigureSaasAdminUserDetailsAccessTypeEnum = "ADMIN"
)

var mappingConfigureSaasAdminUserDetailsAccessTypeEnum = map[string]ConfigureSaasAdminUserDetailsAccessTypeEnum{
	"READ_ONLY":  ConfigureSaasAdminUserDetailsAccessTypeReadOnly,
	"READ_WRITE": ConfigureSaasAdminUserDetailsAccessTypeReadWrite,
	"ADMIN":      ConfigureSaasAdminUserDetailsAccessTypeAdmin,
}

var mappingConfigureSaasAdminUserDetailsAccessTypeEnumLowerCase = map[string]ConfigureSaasAdminUserDetailsAccessTypeEnum{
	"read_only":  ConfigureSaasAdminUserDetailsAccessTypeReadOnly,
	"read_write": ConfigureSaasAdminUserDetailsAccessTypeReadWrite,
	"admin":      ConfigureSaasAdminUserDetailsAccessTypeAdmin,
}

// GetConfigureSaasAdminUserDetailsAccessTypeEnumValues Enumerates the set of values for ConfigureSaasAdminUserDetailsAccessTypeEnum
func GetConfigureSaasAdminUserDetailsAccessTypeEnumValues() []ConfigureSaasAdminUserDetailsAccessTypeEnum {
	values := make([]ConfigureSaasAdminUserDetailsAccessTypeEnum, 0)
	for _, v := range mappingConfigureSaasAdminUserDetailsAccessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigureSaasAdminUserDetailsAccessTypeEnumStringValues Enumerates the set of values in String for ConfigureSaasAdminUserDetailsAccessTypeEnum
func GetConfigureSaasAdminUserDetailsAccessTypeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
		"ADMIN",
	}
}

// GetMappingConfigureSaasAdminUserDetailsAccessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigureSaasAdminUserDetailsAccessTypeEnum(val string) (ConfigureSaasAdminUserDetailsAccessTypeEnum, bool) {
	enum, ok := mappingConfigureSaasAdminUserDetailsAccessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
