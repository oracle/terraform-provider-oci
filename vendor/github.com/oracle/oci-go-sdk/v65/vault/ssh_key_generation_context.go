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

// SshKeyGenerationContext Generates ssh keys. By default, the keys are stored as - {"publicKey": "%GENERATED_PUBLIC_KEY%", "privateKey": "%GENERATED_PRIVATE_KEY%"} in PEM format.
// The SecretTemplate must have both %GENERATED_PUBLIC_KEY% and %GENERATED_PRIVATE_KEY% keywords which is later replaced with the respective keys, if provided.
type SshKeyGenerationContext struct {

	// SecretTemplate captures structure in which customer wants to store secrets. This is optional and a default structure is available for each secret type.
	// The template can have any structure with static values that are not generated. Within the template, you can insert predefined placeholders to store secrets.
	// These placeholders are later replaced with the generated content and saved as a Base64 encoded content.
	SecretTemplate *string `mandatory:"false" json:"secretTemplate"`

	// Name of SSH key generation template to generate SSH key type secret.
	GenerationTemplate SshKeyGenerationContextGenerationTemplateEnum `mandatory:"true" json:"generationTemplate"`
}

// GetSecretTemplate returns SecretTemplate
func (m SshKeyGenerationContext) GetSecretTemplate() *string {
	return m.SecretTemplate
}

func (m SshKeyGenerationContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SshKeyGenerationContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSshKeyGenerationContextGenerationTemplateEnum(string(m.GenerationTemplate)); !ok && m.GenerationTemplate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GenerationTemplate: %s. Supported values are: %s.", m.GenerationTemplate, strings.Join(GetSshKeyGenerationContextGenerationTemplateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SshKeyGenerationContext) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSshKeyGenerationContext SshKeyGenerationContext
	s := struct {
		DiscriminatorParam string `json:"generationType"`
		MarshalTypeSshKeyGenerationContext
	}{
		"SSH_KEY",
		(MarshalTypeSshKeyGenerationContext)(m),
	}

	return json.Marshal(&s)
}

// SshKeyGenerationContextGenerationTemplateEnum Enum with underlying type: string
type SshKeyGenerationContextGenerationTemplateEnum string

// Set of constants representing the allowable values for SshKeyGenerationContextGenerationTemplateEnum
const (
	SshKeyGenerationContextGenerationTemplate2048 SshKeyGenerationContextGenerationTemplateEnum = "RSA_2048"
	SshKeyGenerationContextGenerationTemplate3072 SshKeyGenerationContextGenerationTemplateEnum = "RSA_3072"
	SshKeyGenerationContextGenerationTemplate4096 SshKeyGenerationContextGenerationTemplateEnum = "RSA_4096"
)

var mappingSshKeyGenerationContextGenerationTemplateEnum = map[string]SshKeyGenerationContextGenerationTemplateEnum{
	"RSA_2048": SshKeyGenerationContextGenerationTemplate2048,
	"RSA_3072": SshKeyGenerationContextGenerationTemplate3072,
	"RSA_4096": SshKeyGenerationContextGenerationTemplate4096,
}

var mappingSshKeyGenerationContextGenerationTemplateEnumLowerCase = map[string]SshKeyGenerationContextGenerationTemplateEnum{
	"rsa_2048": SshKeyGenerationContextGenerationTemplate2048,
	"rsa_3072": SshKeyGenerationContextGenerationTemplate3072,
	"rsa_4096": SshKeyGenerationContextGenerationTemplate4096,
}

// GetSshKeyGenerationContextGenerationTemplateEnumValues Enumerates the set of values for SshKeyGenerationContextGenerationTemplateEnum
func GetSshKeyGenerationContextGenerationTemplateEnumValues() []SshKeyGenerationContextGenerationTemplateEnum {
	values := make([]SshKeyGenerationContextGenerationTemplateEnum, 0)
	for _, v := range mappingSshKeyGenerationContextGenerationTemplateEnum {
		values = append(values, v)
	}
	return values
}

// GetSshKeyGenerationContextGenerationTemplateEnumStringValues Enumerates the set of values in String for SshKeyGenerationContextGenerationTemplateEnum
func GetSshKeyGenerationContextGenerationTemplateEnumStringValues() []string {
	return []string{
		"RSA_2048",
		"RSA_3072",
		"RSA_4096",
	}
}

// GetMappingSshKeyGenerationContextGenerationTemplateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSshKeyGenerationContextGenerationTemplateEnum(val string) (SshKeyGenerationContextGenerationTemplateEnum, bool) {
	enum, ok := mappingSshKeyGenerationContextGenerationTemplateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
