// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceProviderCredentialDetails Contains credentials associated with the source provider.
type SourceProviderCredentialDetails interface {
}

type sourceprovidercredentialdetails struct {
	JsonData       []byte
	CredentialType string `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *sourceprovidercredentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourceprovidercredentialdetails sourceprovidercredentialdetails
	s := struct {
		Model Unmarshalersourceprovidercredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourceprovidercredentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "KEY_ENCRYPTION":
		mm := SourceProviderKeyEncryptionCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VAULT_SECRET":
		mm := SourceProviderVaultSecretCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SourceProviderCredentialDetails: %s.", m.CredentialType)
		return *m, nil
	}
}

func (m sourceprovidercredentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sourceprovidercredentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceProviderCredentialDetailsCredentialTypeEnum Enum with underlying type: string
type SourceProviderCredentialDetailsCredentialTypeEnum string

// Set of constants representing the allowable values for SourceProviderCredentialDetailsCredentialTypeEnum
const (
	SourceProviderCredentialDetailsCredentialTypeVaultSecret   SourceProviderCredentialDetailsCredentialTypeEnum = "VAULT_SECRET"
	SourceProviderCredentialDetailsCredentialTypeKeyEncryption SourceProviderCredentialDetailsCredentialTypeEnum = "KEY_ENCRYPTION"
)

var mappingSourceProviderCredentialDetailsCredentialTypeEnum = map[string]SourceProviderCredentialDetailsCredentialTypeEnum{
	"VAULT_SECRET":   SourceProviderCredentialDetailsCredentialTypeVaultSecret,
	"KEY_ENCRYPTION": SourceProviderCredentialDetailsCredentialTypeKeyEncryption,
}

var mappingSourceProviderCredentialDetailsCredentialTypeEnumLowerCase = map[string]SourceProviderCredentialDetailsCredentialTypeEnum{
	"vault_secret":   SourceProviderCredentialDetailsCredentialTypeVaultSecret,
	"key_encryption": SourceProviderCredentialDetailsCredentialTypeKeyEncryption,
}

// GetSourceProviderCredentialDetailsCredentialTypeEnumValues Enumerates the set of values for SourceProviderCredentialDetailsCredentialTypeEnum
func GetSourceProviderCredentialDetailsCredentialTypeEnumValues() []SourceProviderCredentialDetailsCredentialTypeEnum {
	values := make([]SourceProviderCredentialDetailsCredentialTypeEnum, 0)
	for _, v := range mappingSourceProviderCredentialDetailsCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceProviderCredentialDetailsCredentialTypeEnumStringValues Enumerates the set of values in String for SourceProviderCredentialDetailsCredentialTypeEnum
func GetSourceProviderCredentialDetailsCredentialTypeEnumStringValues() []string {
	return []string{
		"VAULT_SECRET",
		"KEY_ENCRYPTION",
	}
}

// GetMappingSourceProviderCredentialDetailsCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceProviderCredentialDetailsCredentialTypeEnum(val string) (SourceProviderCredentialDetailsCredentialTypeEnum, bool) {
	enum, ok := mappingSourceProviderCredentialDetailsCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
