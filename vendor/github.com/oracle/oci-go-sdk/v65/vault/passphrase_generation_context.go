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

// PassphraseGenerationContext Generates Passphrase type secrets. By default, passphrase type secrets have no structure. The generated content is stored in Base64 format.
// The SecretTemplate must have the %GENERATED_PASSPHRASE% keyword, which is later replaced with the generated content, if provided.
type PassphraseGenerationContext struct {

	// SecretTemplate captures structure in which customer wants to store secrets. This is optional and a default structure is available for each secret type.
	// The template can have any structure with static values that are not generated. Within the template, you can insert predefined placeholders to store secrets.
	// These placeholders are later replaced with the generated content and saved as a Base64 encoded content.
	SecretTemplate *string `mandatory:"false" json:"secretTemplate"`

	// Length of the passphrase to be generated
	PassphraseLength *int `mandatory:"false" json:"passphraseLength"`

	// Name of passphrase generation template to generate passphrase type secret.
	GenerationTemplate PassphraseGenerationContextGenerationTemplateEnum `mandatory:"true" json:"generationTemplate"`
}

// GetSecretTemplate returns SecretTemplate
func (m PassphraseGenerationContext) GetSecretTemplate() *string {
	return m.SecretTemplate
}

func (m PassphraseGenerationContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PassphraseGenerationContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPassphraseGenerationContextGenerationTemplateEnum(string(m.GenerationTemplate)); !ok && m.GenerationTemplate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GenerationTemplate: %s. Supported values are: %s.", m.GenerationTemplate, strings.Join(GetPassphraseGenerationContextGenerationTemplateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PassphraseGenerationContext) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePassphraseGenerationContext PassphraseGenerationContext
	s := struct {
		DiscriminatorParam string `json:"generationType"`
		MarshalTypePassphraseGenerationContext
	}{
		"PASSPHRASE",
		(MarshalTypePassphraseGenerationContext)(m),
	}

	return json.Marshal(&s)
}

// PassphraseGenerationContextGenerationTemplateEnum Enum with underlying type: string
type PassphraseGenerationContextGenerationTemplateEnum string

// Set of constants representing the allowable values for PassphraseGenerationContextGenerationTemplateEnum
const (
	PassphraseGenerationContextGenerationTemplateSecretsDefaultPassword PassphraseGenerationContextGenerationTemplateEnum = "SECRETS_DEFAULT_PASSWORD"
	PassphraseGenerationContextGenerationTemplateDbaasDefaultPassword   PassphraseGenerationContextGenerationTemplateEnum = "DBAAS_DEFAULT_PASSWORD"
)

var mappingPassphraseGenerationContextGenerationTemplateEnum = map[string]PassphraseGenerationContextGenerationTemplateEnum{
	"SECRETS_DEFAULT_PASSWORD": PassphraseGenerationContextGenerationTemplateSecretsDefaultPassword,
	"DBAAS_DEFAULT_PASSWORD":   PassphraseGenerationContextGenerationTemplateDbaasDefaultPassword,
}

var mappingPassphraseGenerationContextGenerationTemplateEnumLowerCase = map[string]PassphraseGenerationContextGenerationTemplateEnum{
	"secrets_default_password": PassphraseGenerationContextGenerationTemplateSecretsDefaultPassword,
	"dbaas_default_password":   PassphraseGenerationContextGenerationTemplateDbaasDefaultPassword,
}

// GetPassphraseGenerationContextGenerationTemplateEnumValues Enumerates the set of values for PassphraseGenerationContextGenerationTemplateEnum
func GetPassphraseGenerationContextGenerationTemplateEnumValues() []PassphraseGenerationContextGenerationTemplateEnum {
	values := make([]PassphraseGenerationContextGenerationTemplateEnum, 0)
	for _, v := range mappingPassphraseGenerationContextGenerationTemplateEnum {
		values = append(values, v)
	}
	return values
}

// GetPassphraseGenerationContextGenerationTemplateEnumStringValues Enumerates the set of values in String for PassphraseGenerationContextGenerationTemplateEnum
func GetPassphraseGenerationContextGenerationTemplateEnumStringValues() []string {
	return []string{
		"SECRETS_DEFAULT_PASSWORD",
		"DBAAS_DEFAULT_PASSWORD",
	}
}

// GetMappingPassphraseGenerationContextGenerationTemplateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPassphraseGenerationContextGenerationTemplateEnum(val string) (PassphraseGenerationContextGenerationTemplateEnum, bool) {
	enum, ok := mappingPassphraseGenerationContextGenerationTemplateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
