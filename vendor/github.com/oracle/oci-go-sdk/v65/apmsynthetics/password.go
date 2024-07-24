// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Password Password.
type Password interface {
}

type password struct {
	JsonData     []byte
	PasswordType string `json:"passwordType"`
}

// UnmarshalJSON unmarshals json
func (m *password) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpassword password
	s := struct {
		Model Unmarshalerpassword
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PasswordType = s.Model.PasswordType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *password) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PasswordType {
	case "IN_TEXT":
		mm := PasswordInText{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VAULT_SECRET_ID":
		mm := PasswordInVault{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Password: %s.", m.PasswordType)
		return *m, nil
	}
}

func (m password) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m password) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PasswordPasswordTypeEnum Enum with underlying type: string
type PasswordPasswordTypeEnum string

// Set of constants representing the allowable values for PasswordPasswordTypeEnum
const (
	PasswordPasswordTypeInText        PasswordPasswordTypeEnum = "IN_TEXT"
	PasswordPasswordTypeVaultSecretId PasswordPasswordTypeEnum = "VAULT_SECRET_ID"
)

var mappingPasswordPasswordTypeEnum = map[string]PasswordPasswordTypeEnum{
	"IN_TEXT":         PasswordPasswordTypeInText,
	"VAULT_SECRET_ID": PasswordPasswordTypeVaultSecretId,
}

var mappingPasswordPasswordTypeEnumLowerCase = map[string]PasswordPasswordTypeEnum{
	"in_text":         PasswordPasswordTypeInText,
	"vault_secret_id": PasswordPasswordTypeVaultSecretId,
}

// GetPasswordPasswordTypeEnumValues Enumerates the set of values for PasswordPasswordTypeEnum
func GetPasswordPasswordTypeEnumValues() []PasswordPasswordTypeEnum {
	values := make([]PasswordPasswordTypeEnum, 0)
	for _, v := range mappingPasswordPasswordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPasswordPasswordTypeEnumStringValues Enumerates the set of values in String for PasswordPasswordTypeEnum
func GetPasswordPasswordTypeEnumStringValues() []string {
	return []string{
		"IN_TEXT",
		"VAULT_SECRET_ID",
	}
}

// GetMappingPasswordPasswordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPasswordPasswordTypeEnum(val string) (PasswordPasswordTypeEnum, bool) {
	enum, ok := mappingPasswordPasswordTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
