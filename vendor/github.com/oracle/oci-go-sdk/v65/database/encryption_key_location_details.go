// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EncryptionKeyLocationDetails Types of providers supported for managing database encryption keys
type EncryptionKeyLocationDetails interface {
}

type encryptionkeylocationdetails struct {
	JsonData     []byte
	ProviderType string `json:"providerType"`
}

// UnmarshalJSON unmarshals json
func (m *encryptionkeylocationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerencryptionkeylocationdetails encryptionkeylocationdetails
	s := struct {
		Model Unmarshalerencryptionkeylocationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ProviderType = s.Model.ProviderType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *encryptionkeylocationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ProviderType {
	case "EXTERNAL":
		mm := ExternalHsmEncryptionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GCP":
		mm := GoogleCloudProviderEncryptionKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE":
		mm := AzureEncryptionKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for EncryptionKeyLocationDetails: %s.", m.ProviderType)
		return *m, nil
	}
}

func (m encryptionkeylocationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m encryptionkeylocationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EncryptionKeyLocationDetailsProviderTypeEnum Enum with underlying type: string
type EncryptionKeyLocationDetailsProviderTypeEnum string

// Set of constants representing the allowable values for EncryptionKeyLocationDetailsProviderTypeEnum
const (
	EncryptionKeyLocationDetailsProviderTypeExternal EncryptionKeyLocationDetailsProviderTypeEnum = "EXTERNAL"
	EncryptionKeyLocationDetailsProviderTypeAzure    EncryptionKeyLocationDetailsProviderTypeEnum = "AZURE"
	EncryptionKeyLocationDetailsProviderTypeGcp      EncryptionKeyLocationDetailsProviderTypeEnum = "GCP"
)

var mappingEncryptionKeyLocationDetailsProviderTypeEnum = map[string]EncryptionKeyLocationDetailsProviderTypeEnum{
	"EXTERNAL": EncryptionKeyLocationDetailsProviderTypeExternal,
	"AZURE":    EncryptionKeyLocationDetailsProviderTypeAzure,
	"GCP":      EncryptionKeyLocationDetailsProviderTypeGcp,
}

var mappingEncryptionKeyLocationDetailsProviderTypeEnumLowerCase = map[string]EncryptionKeyLocationDetailsProviderTypeEnum{
	"external": EncryptionKeyLocationDetailsProviderTypeExternal,
	"azure":    EncryptionKeyLocationDetailsProviderTypeAzure,
	"gcp":      EncryptionKeyLocationDetailsProviderTypeGcp,
}

// GetEncryptionKeyLocationDetailsProviderTypeEnumValues Enumerates the set of values for EncryptionKeyLocationDetailsProviderTypeEnum
func GetEncryptionKeyLocationDetailsProviderTypeEnumValues() []EncryptionKeyLocationDetailsProviderTypeEnum {
	values := make([]EncryptionKeyLocationDetailsProviderTypeEnum, 0)
	for _, v := range mappingEncryptionKeyLocationDetailsProviderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEncryptionKeyLocationDetailsProviderTypeEnumStringValues Enumerates the set of values in String for EncryptionKeyLocationDetailsProviderTypeEnum
func GetEncryptionKeyLocationDetailsProviderTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL",
		"AZURE",
		"GCP",
	}
}

// GetMappingEncryptionKeyLocationDetailsProviderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEncryptionKeyLocationDetailsProviderTypeEnum(val string) (EncryptionKeyLocationDetailsProviderTypeEnum, bool) {
	enum, ok := mappingEncryptionKeyLocationDetailsProviderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
