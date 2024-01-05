// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VerificationKeySource The source of the verification material.
type VerificationKeySource interface {
}

type verificationkeysource struct {
	JsonData                  []byte
	VerificationKeySourceType string `json:"verificationKeySourceType"`
}

// UnmarshalJSON unmarshals json
func (m *verificationkeysource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerverificationkeysource verificationkeysource
	s := struct {
		Model Unmarshalerverificationkeysource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.VerificationKeySourceType = s.Model.VerificationKeySourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *verificationkeysource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.VerificationKeySourceType {
	case "INLINE_PUBLIC_KEY":
		mm := InlinePublicKeyVerificationKeySource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VAULT_SECRET":
		mm := VaultSecretVerificationKeySource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NoneVerificationKeySource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for VerificationKeySource: %s.", m.VerificationKeySourceType)
		return *m, nil
	}
}

func (m verificationkeysource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m verificationkeysource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VerificationKeySourceVerificationKeySourceTypeEnum Enum with underlying type: string
type VerificationKeySourceVerificationKeySourceTypeEnum string

// Set of constants representing the allowable values for VerificationKeySourceVerificationKeySourceTypeEnum
const (
	VerificationKeySourceVerificationKeySourceTypeVaultSecret     VerificationKeySourceVerificationKeySourceTypeEnum = "VAULT_SECRET"
	VerificationKeySourceVerificationKeySourceTypeInlinePublicKey VerificationKeySourceVerificationKeySourceTypeEnum = "INLINE_PUBLIC_KEY"
	VerificationKeySourceVerificationKeySourceTypeNone            VerificationKeySourceVerificationKeySourceTypeEnum = "NONE"
)

var mappingVerificationKeySourceVerificationKeySourceTypeEnum = map[string]VerificationKeySourceVerificationKeySourceTypeEnum{
	"VAULT_SECRET":      VerificationKeySourceVerificationKeySourceTypeVaultSecret,
	"INLINE_PUBLIC_KEY": VerificationKeySourceVerificationKeySourceTypeInlinePublicKey,
	"NONE":              VerificationKeySourceVerificationKeySourceTypeNone,
}

var mappingVerificationKeySourceVerificationKeySourceTypeEnumLowerCase = map[string]VerificationKeySourceVerificationKeySourceTypeEnum{
	"vault_secret":      VerificationKeySourceVerificationKeySourceTypeVaultSecret,
	"inline_public_key": VerificationKeySourceVerificationKeySourceTypeInlinePublicKey,
	"none":              VerificationKeySourceVerificationKeySourceTypeNone,
}

// GetVerificationKeySourceVerificationKeySourceTypeEnumValues Enumerates the set of values for VerificationKeySourceVerificationKeySourceTypeEnum
func GetVerificationKeySourceVerificationKeySourceTypeEnumValues() []VerificationKeySourceVerificationKeySourceTypeEnum {
	values := make([]VerificationKeySourceVerificationKeySourceTypeEnum, 0)
	for _, v := range mappingVerificationKeySourceVerificationKeySourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVerificationKeySourceVerificationKeySourceTypeEnumStringValues Enumerates the set of values in String for VerificationKeySourceVerificationKeySourceTypeEnum
func GetVerificationKeySourceVerificationKeySourceTypeEnumStringValues() []string {
	return []string{
		"VAULT_SECRET",
		"INLINE_PUBLIC_KEY",
		"NONE",
	}
}

// GetMappingVerificationKeySourceVerificationKeySourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVerificationKeySourceVerificationKeySourceTypeEnum(val string) (VerificationKeySourceVerificationKeySourceTypeEnum, bool) {
	enum, ok := mappingVerificationKeySourceVerificationKeySourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
