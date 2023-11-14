// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// A description of the PGSQL Control Plane API
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PasswordDetails Details for the DbSystem password.
// Password can be passed as `VaultSecretPasswordDetails`(Vault) or `PlainTextPasswordDetails`.
type PasswordDetails interface {
}

type passworddetails struct {
	JsonData     []byte
	PasswordType string `json:"passwordType"`
}

// UnmarshalJSON unmarshals json
func (m *passworddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpassworddetails passworddetails
	s := struct {
		Model Unmarshalerpassworddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PasswordType = s.Model.PasswordType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *passworddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PasswordType {
	case "PLAIN_TEXT":
		mm := PlainTextPasswordDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VAULT_SECRET":
		mm := VaultSecretPasswordDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PasswordDetails: %s.", m.PasswordType)
		return *m, nil
	}
}

func (m passworddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m passworddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PasswordDetailsPasswordTypeEnum Enum with underlying type: string
type PasswordDetailsPasswordTypeEnum string

// Set of constants representing the allowable values for PasswordDetailsPasswordTypeEnum
const (
	PasswordDetailsPasswordTypePlainText   PasswordDetailsPasswordTypeEnum = "PLAIN_TEXT"
	PasswordDetailsPasswordTypeVaultSecret PasswordDetailsPasswordTypeEnum = "VAULT_SECRET"
)

var mappingPasswordDetailsPasswordTypeEnum = map[string]PasswordDetailsPasswordTypeEnum{
	"PLAIN_TEXT":   PasswordDetailsPasswordTypePlainText,
	"VAULT_SECRET": PasswordDetailsPasswordTypeVaultSecret,
}

var mappingPasswordDetailsPasswordTypeEnumLowerCase = map[string]PasswordDetailsPasswordTypeEnum{
	"plain_text":   PasswordDetailsPasswordTypePlainText,
	"vault_secret": PasswordDetailsPasswordTypeVaultSecret,
}

// GetPasswordDetailsPasswordTypeEnumValues Enumerates the set of values for PasswordDetailsPasswordTypeEnum
func GetPasswordDetailsPasswordTypeEnumValues() []PasswordDetailsPasswordTypeEnum {
	values := make([]PasswordDetailsPasswordTypeEnum, 0)
	for _, v := range mappingPasswordDetailsPasswordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPasswordDetailsPasswordTypeEnumStringValues Enumerates the set of values in String for PasswordDetailsPasswordTypeEnum
func GetPasswordDetailsPasswordTypeEnumStringValues() []string {
	return []string{
		"PLAIN_TEXT",
		"VAULT_SECRET",
	}
}

// GetMappingPasswordDetailsPasswordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPasswordDetailsPasswordTypeEnum(val string) (PasswordDetailsPasswordTypeEnum, bool) {
	enum, ok := mappingPasswordDetailsPasswordTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
