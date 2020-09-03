// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
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

// GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum Enum with underlying type: string
type GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum string

// Set of constants representing the allowable values for GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum
const (
	GenerateAutonomousDatabaseWalletDetailsGenerateTypeAll    GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum = "ALL"
	GenerateAutonomousDatabaseWalletDetailsGenerateTypeSingle GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum = "SINGLE"
)

var mappingGenerateAutonomousDatabaseWalletDetailsGenerateType = map[string]GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum{
	"ALL":    GenerateAutonomousDatabaseWalletDetailsGenerateTypeAll,
	"SINGLE": GenerateAutonomousDatabaseWalletDetailsGenerateTypeSingle,
}

// GetGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumValues Enumerates the set of values for GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum
func GetGenerateAutonomousDatabaseWalletDetailsGenerateTypeEnumValues() []GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum {
	values := make([]GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum, 0)
	for _, v := range mappingGenerateAutonomousDatabaseWalletDetailsGenerateType {
		values = append(values, v)
	}
	return values
}
