// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdwAdminPasswordDetails Details for the ADW Admin password.
// Password can be passed as `VaultSecretPasswordDetails` or `PlainTextPasswordDetails`.
// Example: `{"passwordType":"PLAIN_TEXT","password":"..."}`
// Example: `{"passwordType":"VAULT_SECRET","secretId":"ocid..."}`
type AdwAdminPasswordDetails interface {
}

type adwadminpassworddetails struct {
	JsonData     []byte
	PasswordType string `json:"passwordType"`
}

// UnmarshalJSON unmarshals json
func (m *adwadminpassworddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleradwadminpassworddetails adwadminpassworddetails
	s := struct {
		Model Unmarshaleradwadminpassworddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PasswordType = s.Model.PasswordType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *adwadminpassworddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

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
		common.Logf("Received unsupported enum value for AdwAdminPasswordDetails: %s.", m.PasswordType)
		return *m, nil
	}
}

func (m adwadminpassworddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m adwadminpassworddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AdwAdminPasswordDetailsPasswordTypeEnum Enum with underlying type: string
type AdwAdminPasswordDetailsPasswordTypeEnum string

// Set of constants representing the allowable values for AdwAdminPasswordDetailsPasswordTypeEnum
const (
	AdwAdminPasswordDetailsPasswordTypePlainText   AdwAdminPasswordDetailsPasswordTypeEnum = "PLAIN_TEXT"
	AdwAdminPasswordDetailsPasswordTypeVaultSecret AdwAdminPasswordDetailsPasswordTypeEnum = "VAULT_SECRET"
)

var mappingAdwAdminPasswordDetailsPasswordTypeEnum = map[string]AdwAdminPasswordDetailsPasswordTypeEnum{
	"PLAIN_TEXT":   AdwAdminPasswordDetailsPasswordTypePlainText,
	"VAULT_SECRET": AdwAdminPasswordDetailsPasswordTypeVaultSecret,
}

var mappingAdwAdminPasswordDetailsPasswordTypeEnumLowerCase = map[string]AdwAdminPasswordDetailsPasswordTypeEnum{
	"plain_text":   AdwAdminPasswordDetailsPasswordTypePlainText,
	"vault_secret": AdwAdminPasswordDetailsPasswordTypeVaultSecret,
}

// GetAdwAdminPasswordDetailsPasswordTypeEnumValues Enumerates the set of values for AdwAdminPasswordDetailsPasswordTypeEnum
func GetAdwAdminPasswordDetailsPasswordTypeEnumValues() []AdwAdminPasswordDetailsPasswordTypeEnum {
	values := make([]AdwAdminPasswordDetailsPasswordTypeEnum, 0)
	for _, v := range mappingAdwAdminPasswordDetailsPasswordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAdwAdminPasswordDetailsPasswordTypeEnumStringValues Enumerates the set of values in String for AdwAdminPasswordDetailsPasswordTypeEnum
func GetAdwAdminPasswordDetailsPasswordTypeEnumStringValues() []string {
	return []string{
		"PLAIN_TEXT",
		"VAULT_SECRET",
	}
}

// GetMappingAdwAdminPasswordDetailsPasswordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdwAdminPasswordDetailsPasswordTypeEnum(val string) (AdwAdminPasswordDetailsPasswordTypeEnum, bool) {
	enum, ok := mappingAdwAdminPasswordDetailsPasswordTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
