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

// SecretContentDetails The content of the secret and metadata to help identify it.
type SecretContentDetails interface {

	// Names should be unique within a secret. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	GetName() *string

	// The rotation state of the secret content. The default is `CURRENT`, meaning that the secret is currently in use. A secret version
	// that you mark as `PENDING` is staged and available for use, but you don't yet want to rotate it into current, active use. For example,
	// you might create or update a secret and mark its rotation state as `PENDING` if you haven't yet updated the secret on the target system.
	// When creating a secret, only the value `CURRENT` is applicable, although the value `LATEST` is also automatically applied. When updating
	// a secret, you can specify a version's rotation state as either `CURRENT` or `PENDING`.
	GetStage() SecretContentDetailsStageEnum
}

type secretcontentdetails struct {
	JsonData    []byte
	Name        *string                       `mandatory:"false" json:"name"`
	Stage       SecretContentDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`
	ContentType string                        `json:"contentType"`
}

// UnmarshalJSON unmarshals json
func (m *secretcontentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecretcontentdetails secretcontentdetails
	s := struct {
		Model Unmarshalersecretcontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Stage = s.Model.Stage
	m.ContentType = s.Model.ContentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *secretcontentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ContentType {
	case "BASE64":
		mm := Base64SecretContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SecretContentDetails: %s.", m.ContentType)
		return *m, nil
	}
}

// GetName returns Name
func (m secretcontentdetails) GetName() *string {
	return m.Name
}

// GetStage returns Stage
func (m secretcontentdetails) GetStage() SecretContentDetailsStageEnum {
	return m.Stage
}

func (m secretcontentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m secretcontentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecretContentDetailsStageEnum(string(m.Stage)); !ok && m.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", m.Stage, strings.Join(GetSecretContentDetailsStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecretContentDetailsStageEnum Enum with underlying type: string
type SecretContentDetailsStageEnum string

// Set of constants representing the allowable values for SecretContentDetailsStageEnum
const (
	SecretContentDetailsStageCurrent SecretContentDetailsStageEnum = "CURRENT"
	SecretContentDetailsStagePending SecretContentDetailsStageEnum = "PENDING"
)

var mappingSecretContentDetailsStageEnum = map[string]SecretContentDetailsStageEnum{
	"CURRENT": SecretContentDetailsStageCurrent,
	"PENDING": SecretContentDetailsStagePending,
}

var mappingSecretContentDetailsStageEnumLowerCase = map[string]SecretContentDetailsStageEnum{
	"current": SecretContentDetailsStageCurrent,
	"pending": SecretContentDetailsStagePending,
}

// GetSecretContentDetailsStageEnumValues Enumerates the set of values for SecretContentDetailsStageEnum
func GetSecretContentDetailsStageEnumValues() []SecretContentDetailsStageEnum {
	values := make([]SecretContentDetailsStageEnum, 0)
	for _, v := range mappingSecretContentDetailsStageEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretContentDetailsStageEnumStringValues Enumerates the set of values in String for SecretContentDetailsStageEnum
func GetSecretContentDetailsStageEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
	}
}

// GetMappingSecretContentDetailsStageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretContentDetailsStageEnum(val string) (SecretContentDetailsStageEnum, bool) {
	enum, ok := mappingSecretContentDetailsStageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecretContentDetailsContentTypeEnum Enum with underlying type: string
type SecretContentDetailsContentTypeEnum string

// Set of constants representing the allowable values for SecretContentDetailsContentTypeEnum
const (
	SecretContentDetailsContentTypeBase64 SecretContentDetailsContentTypeEnum = "BASE64"
)

var mappingSecretContentDetailsContentTypeEnum = map[string]SecretContentDetailsContentTypeEnum{
	"BASE64": SecretContentDetailsContentTypeBase64,
}

var mappingSecretContentDetailsContentTypeEnumLowerCase = map[string]SecretContentDetailsContentTypeEnum{
	"base64": SecretContentDetailsContentTypeBase64,
}

// GetSecretContentDetailsContentTypeEnumValues Enumerates the set of values for SecretContentDetailsContentTypeEnum
func GetSecretContentDetailsContentTypeEnumValues() []SecretContentDetailsContentTypeEnum {
	values := make([]SecretContentDetailsContentTypeEnum, 0)
	for _, v := range mappingSecretContentDetailsContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretContentDetailsContentTypeEnumStringValues Enumerates the set of values in String for SecretContentDetailsContentTypeEnum
func GetSecretContentDetailsContentTypeEnumStringValues() []string {
	return []string{
		"BASE64",
	}
}

// GetMappingSecretContentDetailsContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretContentDetailsContentTypeEnum(val string) (SecretContentDetailsContentTypeEnum, bool) {
	enum, ok := mappingSecretContentDetailsContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
