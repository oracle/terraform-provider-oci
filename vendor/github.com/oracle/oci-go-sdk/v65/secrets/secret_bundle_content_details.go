// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Retrieval API
//
// Use the Secret Retrieval API to retrieve secrets and secret versions from vaults. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package secrets

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecretBundleContentDetails The contents of the secret.
type SecretBundleContentDetails interface {
}

type secretbundlecontentdetails struct {
	JsonData    []byte
	ContentType string `json:"contentType"`
}

// UnmarshalJSON unmarshals json
func (m *secretbundlecontentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecretbundlecontentdetails secretbundlecontentdetails
	s := struct {
		Model Unmarshalersecretbundlecontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ContentType = s.Model.ContentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *secretbundlecontentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ContentType {
	case "BASE64":
		mm := Base64SecretBundleContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m secretbundlecontentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m secretbundlecontentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecretBundleContentDetailsContentTypeEnum Enum with underlying type: string
type SecretBundleContentDetailsContentTypeEnum string

// Set of constants representing the allowable values for SecretBundleContentDetailsContentTypeEnum
const (
	SecretBundleContentDetailsContentTypeBase64 SecretBundleContentDetailsContentTypeEnum = "BASE64"
)

var mappingSecretBundleContentDetailsContentTypeEnum = map[string]SecretBundleContentDetailsContentTypeEnum{
	"BASE64": SecretBundleContentDetailsContentTypeBase64,
}

var mappingSecretBundleContentDetailsContentTypeEnumLowerCase = map[string]SecretBundleContentDetailsContentTypeEnum{
	"base64": SecretBundleContentDetailsContentTypeBase64,
}

// GetSecretBundleContentDetailsContentTypeEnumValues Enumerates the set of values for SecretBundleContentDetailsContentTypeEnum
func GetSecretBundleContentDetailsContentTypeEnumValues() []SecretBundleContentDetailsContentTypeEnum {
	values := make([]SecretBundleContentDetailsContentTypeEnum, 0)
	for _, v := range mappingSecretBundleContentDetailsContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretBundleContentDetailsContentTypeEnumStringValues Enumerates the set of values in String for SecretBundleContentDetailsContentTypeEnum
func GetSecretBundleContentDetailsContentTypeEnumStringValues() []string {
	return []string{
		"BASE64",
	}
}

// GetMappingSecretBundleContentDetailsContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretBundleContentDetailsContentTypeEnum(val string) (SecretBundleContentDetailsContentTypeEnum, bool) {
	enum, ok := mappingSecretBundleContentDetailsContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
