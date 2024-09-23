// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CredentialDetails Credential Details
type CredentialDetails interface {
}

type credentialdetails struct {
	JsonData       []byte
	CredentialType string `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *credentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercredentialdetails credentialdetails
	s := struct {
		Model Unmarshalercredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *credentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "PLAIN_TEXT":
		mm := PlainTextCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KEY_ENCRYPTION":
		mm := KeyEncryptionCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VAULT_SECRET":
		mm := VaultSecretCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CredentialDetails: %s.", m.CredentialType)
		return *m, nil
	}
}

func (m credentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m credentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CredentialDetailsCredentialTypeEnum Enum with underlying type: string
type CredentialDetailsCredentialTypeEnum string

// Set of constants representing the allowable values for CredentialDetailsCredentialTypeEnum
const (
	CredentialDetailsCredentialTypePlainText     CredentialDetailsCredentialTypeEnum = "PLAIN_TEXT"
	CredentialDetailsCredentialTypeVaultSecret   CredentialDetailsCredentialTypeEnum = "VAULT_SECRET"
	CredentialDetailsCredentialTypeKeyEncryption CredentialDetailsCredentialTypeEnum = "KEY_ENCRYPTION"
)

var mappingCredentialDetailsCredentialTypeEnum = map[string]CredentialDetailsCredentialTypeEnum{
	"PLAIN_TEXT":     CredentialDetailsCredentialTypePlainText,
	"VAULT_SECRET":   CredentialDetailsCredentialTypeVaultSecret,
	"KEY_ENCRYPTION": CredentialDetailsCredentialTypeKeyEncryption,
}

var mappingCredentialDetailsCredentialTypeEnumLowerCase = map[string]CredentialDetailsCredentialTypeEnum{
	"plain_text":     CredentialDetailsCredentialTypePlainText,
	"vault_secret":   CredentialDetailsCredentialTypeVaultSecret,
	"key_encryption": CredentialDetailsCredentialTypeKeyEncryption,
}

// GetCredentialDetailsCredentialTypeEnumValues Enumerates the set of values for CredentialDetailsCredentialTypeEnum
func GetCredentialDetailsCredentialTypeEnumValues() []CredentialDetailsCredentialTypeEnum {
	values := make([]CredentialDetailsCredentialTypeEnum, 0)
	for _, v := range mappingCredentialDetailsCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialDetailsCredentialTypeEnumStringValues Enumerates the set of values in String for CredentialDetailsCredentialTypeEnum
func GetCredentialDetailsCredentialTypeEnumStringValues() []string {
	return []string{
		"PLAIN_TEXT",
		"VAULT_SECRET",
		"KEY_ENCRYPTION",
	}
}

// GetMappingCredentialDetailsCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialDetailsCredentialTypeEnum(val string) (CredentialDetailsCredentialTypeEnum, bool) {
	enum, ok := mappingCredentialDetailsCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
