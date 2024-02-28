// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecretGenerationContext Captures a configurable set of secret generation rules such as length, base characters, additional characters, and so on.
type SecretGenerationContext interface {

	// SecretTemplate captures structure in which customer wants to store secrets. This is optional and a default structure is available for each secret type.
	// The template can have any structure with static values that are not generated. Within the template, you can insert predefined placeholders to store secrets.
	// These placeholders are later replaced with the generated content and saved as a Base64 encoded content.
	GetSecretTemplate() *string
}

type secretgenerationcontext struct {
	JsonData       []byte
	SecretTemplate *string `mandatory:"false" json:"secretTemplate"`
	GenerationType string  `json:"generationType"`
}

// UnmarshalJSON unmarshals json
func (m *secretgenerationcontext) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecretgenerationcontext secretgenerationcontext
	s := struct {
		Model Unmarshalersecretgenerationcontext
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SecretTemplate = s.Model.SecretTemplate
	m.GenerationType = s.Model.GenerationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *secretgenerationcontext) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.GenerationType {
	case "PASSPHRASE":
		mm := PassphraseGenerationContext{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SSH_KEY":
		mm := SshKeyGenerationContext{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BYTES":
		mm := BytesGenerationContext{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SecretGenerationContext: %s.", m.GenerationType)
		return *m, nil
	}
}

// GetSecretTemplate returns SecretTemplate
func (m secretgenerationcontext) GetSecretTemplate() *string {
	return m.SecretTemplate
}

func (m secretgenerationcontext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m secretgenerationcontext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecretGenerationContextGenerationTypeEnum Enum with underlying type: string
type SecretGenerationContextGenerationTypeEnum string

// Set of constants representing the allowable values for SecretGenerationContextGenerationTypeEnum
const (
	SecretGenerationContextGenerationTypePassphrase SecretGenerationContextGenerationTypeEnum = "PASSPHRASE"
	SecretGenerationContextGenerationTypeSshKey     SecretGenerationContextGenerationTypeEnum = "SSH_KEY"
	SecretGenerationContextGenerationTypeBytes      SecretGenerationContextGenerationTypeEnum = "BYTES"
)

var mappingSecretGenerationContextGenerationTypeEnum = map[string]SecretGenerationContextGenerationTypeEnum{
	"PASSPHRASE": SecretGenerationContextGenerationTypePassphrase,
	"SSH_KEY":    SecretGenerationContextGenerationTypeSshKey,
	"BYTES":      SecretGenerationContextGenerationTypeBytes,
}

var mappingSecretGenerationContextGenerationTypeEnumLowerCase = map[string]SecretGenerationContextGenerationTypeEnum{
	"passphrase": SecretGenerationContextGenerationTypePassphrase,
	"ssh_key":    SecretGenerationContextGenerationTypeSshKey,
	"bytes":      SecretGenerationContextGenerationTypeBytes,
}

// GetSecretGenerationContextGenerationTypeEnumValues Enumerates the set of values for SecretGenerationContextGenerationTypeEnum
func GetSecretGenerationContextGenerationTypeEnumValues() []SecretGenerationContextGenerationTypeEnum {
	values := make([]SecretGenerationContextGenerationTypeEnum, 0)
	for _, v := range mappingSecretGenerationContextGenerationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretGenerationContextGenerationTypeEnumStringValues Enumerates the set of values in String for SecretGenerationContextGenerationTypeEnum
func GetSecretGenerationContextGenerationTypeEnumStringValues() []string {
	return []string{
		"PASSPHRASE",
		"SSH_KEY",
		"BYTES",
	}
}

// GetMappingSecretGenerationContextGenerationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretGenerationContextGenerationTypeEnum(val string) (SecretGenerationContextGenerationTypeEnum, bool) {
	enum, ok := mappingSecretGenerationContextGenerationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
