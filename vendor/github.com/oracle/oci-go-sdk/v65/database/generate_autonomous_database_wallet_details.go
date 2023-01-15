// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// GenerateAutonomousDatabaseWalletDetails Details to create and download an Oracle Autonomous Database wallet.
type GenerateAutonomousDatabaseWalletDetails struct {

	// The password to encrypt the keys inside the wallet. The password must be at least 8 characters long and must include at least 1 letter and either 1 numeric character or 1 special character.
	Password *string `mandatory:"true" json:"password"`

	// The type of wallet to generate.
	// **Shared Exadata infrastructure usage:**
	// * `SINGLE` - used to generate a wallet for a single database
	// * `ALL` - used to generate wallet for all databases in the region
	// **Dedicated Exadata infrastructure usage:** Value must be `NULL` if attribute is used.
	GenerateType GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum `mandatory:"false" json:"generateType,omitempty"`
}

func (m GenerateAutonomousDatabaseWalletDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateAutonomousDatabaseWalletDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum(string(m.GenerateType)); !ok && m.GenerateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GenerateType: %s. Supported values are: %s.", m.GenerateType, strings.Join(GetGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum Enum with underlying type: string
type GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum string

// Set of constants representing the allowable values for GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum
const (
	GenerateAutonomousDatabaseWalletDetailsGenerateTypeAll    GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum = "ALL"
	GenerateAutonomousDatabaseWalletDetailsGenerateTypeSingle GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum = "SINGLE"
)

var mappingGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum = map[string]GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum{
	"ALL":    GenerateAutonomousDatabaseWalletDetailsGenerateTypeAll,
	"SINGLE": GenerateAutonomousDatabaseWalletDetailsGenerateTypeSingle,
}

var mappingGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumLowerCase = map[string]GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum{
	"all":    GenerateAutonomousDatabaseWalletDetailsGenerateTypeAll,
	"single": GenerateAutonomousDatabaseWalletDetailsGenerateTypeSingle,
}

// GetGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumValues Enumerates the set of values for GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum
func GetGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumValues() []GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum {
	values := make([]GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum, 0)
	for _, v := range mappingGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumStringValues Enumerates the set of values in String for GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum
func GetGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"SINGLE",
	}
}

// GetMappingGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum(val string) (GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum, bool) {
	enum, ok := mappingGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
