// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
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

// KeyStoreTypeDetailsTypeEnum Enum with underlying type: string
type KeyStoreTypeDetailsTypeEnum string

// Set of constants representing the allowable values for KeyStoreTypeDetailsTypeEnum
const (
	KeyStoreTypeDetailsTypeOracleKeyVault KeyStoreTypeDetailsTypeEnum = "ORACLE_KEY_VAULT"
)

var mappingKeyStoreTypeDetailsType = map[string]KeyStoreTypeDetailsTypeEnum{
	"ORACLE_KEY_VAULT": KeyStoreTypeDetailsTypeOracleKeyVault,
}

// GetKeyStoreTypeDetailsTypeEnumValues Enumerates the set of values for KeyStoreTypeDetailsTypeEnum
func GetKeyStoreTypeDetailsTypeEnumValues() []KeyStoreTypeDetailsTypeEnum {
	values := make([]KeyStoreTypeDetailsTypeEnum, 0)
	for _, v := range mappingKeyStoreTypeDetailsType {
		values = append(values, v)
	}
	return values
}
