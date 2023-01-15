// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KeyStoreTypeDetails Key store type details.
type KeyStoreTypeDetails interface {
}

type keystoretypedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *keystoretypedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerkeystoretypedetails keystoretypedetails
	s := struct {
		Model Unmarshalerkeystoretypedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *keystoretypedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_KEY_VAULT":
		mm := KeyStoreTypeFromOracleKeyVaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m keystoretypedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m keystoretypedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KeyStoreTypeDetailsTypeEnum Enum with underlying type: string
type KeyStoreTypeDetailsTypeEnum string

// Set of constants representing the allowable values for KeyStoreTypeDetailsTypeEnum
const (
	KeyStoreTypeDetailsTypeOracleKeyVault KeyStoreTypeDetailsTypeEnum = "ORACLE_KEY_VAULT"
)

var mappingKeyStoreTypeDetailsTypeEnum = map[string]KeyStoreTypeDetailsTypeEnum{
	"ORACLE_KEY_VAULT": KeyStoreTypeDetailsTypeOracleKeyVault,
}

var mappingKeyStoreTypeDetailsTypeEnumLowerCase = map[string]KeyStoreTypeDetailsTypeEnum{
	"oracle_key_vault": KeyStoreTypeDetailsTypeOracleKeyVault,
}

// GetKeyStoreTypeDetailsTypeEnumValues Enumerates the set of values for KeyStoreTypeDetailsTypeEnum
func GetKeyStoreTypeDetailsTypeEnumValues() []KeyStoreTypeDetailsTypeEnum {
	values := make([]KeyStoreTypeDetailsTypeEnum, 0)
	for _, v := range mappingKeyStoreTypeDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyStoreTypeDetailsTypeEnumStringValues Enumerates the set of values in String for KeyStoreTypeDetailsTypeEnum
func GetKeyStoreTypeDetailsTypeEnumStringValues() []string {
	return []string{
		"ORACLE_KEY_VAULT",
	}
}

// GetMappingKeyStoreTypeDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyStoreTypeDetailsTypeEnum(val string) (KeyStoreTypeDetailsTypeEnum, bool) {
	enum, ok := mappingKeyStoreTypeDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
