// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KeyManagementSystemDetails Autonomous VM Cluster Key Management System details.
type KeyManagementSystemDetails interface {

	// If true, ACDs within this Cluster cannot use a different key management system than what is configured in AVM Cluster.
	GetIsExclusive() *bool
}

type keymanagementsystemdetails struct {
	JsonData            []byte
	IsExclusive         *bool  `mandatory:"false" json:"isExclusive"`
	KeyManagementSystem string `json:"keyManagementSystem"`
}

// UnmarshalJSON unmarshals json
func (m *keymanagementsystemdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerkeymanagementsystemdetails keymanagementsystemdetails
	s := struct {
		Model Unmarshalerkeymanagementsystemdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsExclusive = s.Model.IsExclusive
	m.KeyManagementSystem = s.Model.KeyManagementSystem

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *keymanagementsystemdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.KeyManagementSystem {
	case "ORACLE_MANAGED_KEYS":
		mm := OracleManagedKeyManagementSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "THALES":
		mm := ThalesKeyManagementSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_VAULT":
		mm := OciVaultKeyManagementSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_KEY_VAULT":
		mm := OracleKeyVaultKeyManagementSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for KeyManagementSystemDetails: %s.", m.KeyManagementSystem)
		return *m, nil
	}
}

// GetIsExclusive returns IsExclusive
func (m keymanagementsystemdetails) GetIsExclusive() *bool {
	return m.IsExclusive
}

func (m keymanagementsystemdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m keymanagementsystemdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KeyManagementSystemDetailsKeyManagementSystemEnum Enum with underlying type: string
type KeyManagementSystemDetailsKeyManagementSystemEnum string

// Set of constants representing the allowable values for KeyManagementSystemDetailsKeyManagementSystemEnum
const (
	KeyManagementSystemDetailsKeyManagementSystemThales            KeyManagementSystemDetailsKeyManagementSystemEnum = "THALES"
	KeyManagementSystemDetailsKeyManagementSystemOracleKeyVault    KeyManagementSystemDetailsKeyManagementSystemEnum = "ORACLE_KEY_VAULT"
	KeyManagementSystemDetailsKeyManagementSystemOciVault          KeyManagementSystemDetailsKeyManagementSystemEnum = "OCI_VAULT"
	KeyManagementSystemDetailsKeyManagementSystemOracleManagedKeys KeyManagementSystemDetailsKeyManagementSystemEnum = "ORACLE_MANAGED_KEYS"
)

var mappingKeyManagementSystemDetailsKeyManagementSystemEnum = map[string]KeyManagementSystemDetailsKeyManagementSystemEnum{
	"THALES":              KeyManagementSystemDetailsKeyManagementSystemThales,
	"ORACLE_KEY_VAULT":    KeyManagementSystemDetailsKeyManagementSystemOracleKeyVault,
	"OCI_VAULT":           KeyManagementSystemDetailsKeyManagementSystemOciVault,
	"ORACLE_MANAGED_KEYS": KeyManagementSystemDetailsKeyManagementSystemOracleManagedKeys,
}

var mappingKeyManagementSystemDetailsKeyManagementSystemEnumLowerCase = map[string]KeyManagementSystemDetailsKeyManagementSystemEnum{
	"thales":              KeyManagementSystemDetailsKeyManagementSystemThales,
	"oracle_key_vault":    KeyManagementSystemDetailsKeyManagementSystemOracleKeyVault,
	"oci_vault":           KeyManagementSystemDetailsKeyManagementSystemOciVault,
	"oracle_managed_keys": KeyManagementSystemDetailsKeyManagementSystemOracleManagedKeys,
}

// GetKeyManagementSystemDetailsKeyManagementSystemEnumValues Enumerates the set of values for KeyManagementSystemDetailsKeyManagementSystemEnum
func GetKeyManagementSystemDetailsKeyManagementSystemEnumValues() []KeyManagementSystemDetailsKeyManagementSystemEnum {
	values := make([]KeyManagementSystemDetailsKeyManagementSystemEnum, 0)
	for _, v := range mappingKeyManagementSystemDetailsKeyManagementSystemEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyManagementSystemDetailsKeyManagementSystemEnumStringValues Enumerates the set of values in String for KeyManagementSystemDetailsKeyManagementSystemEnum
func GetKeyManagementSystemDetailsKeyManagementSystemEnumStringValues() []string {
	return []string{
		"THALES",
		"ORACLE_KEY_VAULT",
		"OCI_VAULT",
		"ORACLE_MANAGED_KEYS",
	}
}

// GetMappingKeyManagementSystemDetailsKeyManagementSystemEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyManagementSystemDetailsKeyManagementSystemEnum(val string) (KeyManagementSystemDetailsKeyManagementSystemEnum, bool) {
	enum, ok := mappingKeyManagementSystemDetailsKeyManagementSystemEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
