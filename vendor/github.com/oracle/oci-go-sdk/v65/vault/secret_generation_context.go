// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

	// Name of the predefined secret generation template.
	GetGenerationTemplate() SecretGenerationContextGenerationTemplateEnum
}

type secretgenerationcontext struct {
	JsonData           []byte
	GenerationTemplate SecretGenerationContextGenerationTemplateEnum `mandatory:"true" json:"generationTemplate"`
	GenerationType     string                                        `json:"generationType"`
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
	m.GenerationTemplate = s.Model.GenerationTemplate
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
	default:
		common.Logf("Recieved unsupported enum value for SecretGenerationContext: %s.", m.GenerationType)
		return *m, nil
	}
}

//GetGenerationTemplate returns GenerationTemplate
func (m secretgenerationcontext) GetGenerationTemplate() SecretGenerationContextGenerationTemplateEnum {
	return m.GenerationTemplate
}

func (m secretgenerationcontext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m secretgenerationcontext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecretGenerationContextGenerationTemplateEnum(string(m.GenerationTemplate)); !ok && m.GenerationTemplate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GenerationTemplate: %s. Supported values are: %s.", m.GenerationTemplate, strings.Join(GetSecretGenerationContextGenerationTemplateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecretGenerationContextGenerationTemplateEnum Enum with underlying type: string
type SecretGenerationContextGenerationTemplateEnum string

// Set of constants representing the allowable values for SecretGenerationContextGenerationTemplateEnum
const (
	SecretGenerationContextGenerationTemplateSecretsDefaultPassword SecretGenerationContextGenerationTemplateEnum = "SECRETS_DEFAULT_PASSWORD"
	SecretGenerationContextGenerationTemplateDbaasDefaultPassword   SecretGenerationContextGenerationTemplateEnum = "DBAAS_DEFAULT_PASSWORD"
)

var mappingSecretGenerationContextGenerationTemplateEnum = map[string]SecretGenerationContextGenerationTemplateEnum{
	"SECRETS_DEFAULT_PASSWORD": SecretGenerationContextGenerationTemplateSecretsDefaultPassword,
	"DBAAS_DEFAULT_PASSWORD":   SecretGenerationContextGenerationTemplateDbaasDefaultPassword,
}

var mappingSecretGenerationContextGenerationTemplateEnumLowerCase = map[string]SecretGenerationContextGenerationTemplateEnum{
	"secrets_default_password": SecretGenerationContextGenerationTemplateSecretsDefaultPassword,
	"dbaas_default_password":   SecretGenerationContextGenerationTemplateDbaasDefaultPassword,
}

// GetSecretGenerationContextGenerationTemplateEnumValues Enumerates the set of values for SecretGenerationContextGenerationTemplateEnum
func GetSecretGenerationContextGenerationTemplateEnumValues() []SecretGenerationContextGenerationTemplateEnum {
	values := make([]SecretGenerationContextGenerationTemplateEnum, 0)
	for _, v := range mappingSecretGenerationContextGenerationTemplateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretGenerationContextGenerationTemplateEnumStringValues Enumerates the set of values in String for SecretGenerationContextGenerationTemplateEnum
func GetSecretGenerationContextGenerationTemplateEnumStringValues() []string {
	return []string{
		"SECRETS_DEFAULT_PASSWORD",
		"DBAAS_DEFAULT_PASSWORD",
	}
}

// GetMappingSecretGenerationContextGenerationTemplateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretGenerationContextGenerationTemplateEnum(val string) (SecretGenerationContextGenerationTemplateEnum, bool) {
	enum, ok := mappingSecretGenerationContextGenerationTemplateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecretGenerationContextGenerationTypeEnum Enum with underlying type: string
type SecretGenerationContextGenerationTypeEnum string

// Set of constants representing the allowable values for SecretGenerationContextGenerationTypeEnum
const (
	SecretGenerationContextGenerationTypePassphrase SecretGenerationContextGenerationTypeEnum = "PASSPHRASE"
)

var mappingSecretGenerationContextGenerationTypeEnum = map[string]SecretGenerationContextGenerationTypeEnum{
	"PASSPHRASE": SecretGenerationContextGenerationTypePassphrase,
}

var mappingSecretGenerationContextGenerationTypeEnumLowerCase = map[string]SecretGenerationContextGenerationTypeEnum{
	"passphrase": SecretGenerationContextGenerationTypePassphrase,
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
	}
}

// GetMappingSecretGenerationContextGenerationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretGenerationContextGenerationTypeEnum(val string) (SecretGenerationContextGenerationTypeEnum, bool) {
	enum, ok := mappingSecretGenerationContextGenerationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
