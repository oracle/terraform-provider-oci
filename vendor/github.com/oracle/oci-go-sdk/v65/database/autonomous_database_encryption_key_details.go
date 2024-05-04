// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDatabaseEncryptionKeyDetails Details of the Autonomous Database encryption key.
type AutonomousDatabaseEncryptionKeyDetails interface {
}

type autonomousdatabaseencryptionkeydetails struct {
	JsonData []byte
	Provider string `json:"provider"`
}

// UnmarshalJSON unmarshals json
func (m *autonomousdatabaseencryptionkeydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerautonomousdatabaseencryptionkeydetails autonomousdatabaseencryptionkeydetails
	s := struct {
		Model Unmarshalerautonomousdatabaseencryptionkeydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Provider = s.Model.Provider

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *autonomousdatabaseencryptionkeydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Provider {
	case "OKV":
		mm := OkvKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE":
		mm := AzureKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWS":
		mm := AwsKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI":
		mm := OciKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_MANAGED":
		mm := OracleManagedKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AutonomousDatabaseEncryptionKeyDetails: %s.", m.Provider)
		return *m, nil
	}
}

func (m autonomousdatabaseencryptionkeydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m autonomousdatabaseencryptionkeydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseEncryptionKeyDetailsProviderEnum Enum with underlying type: string
type AutonomousDatabaseEncryptionKeyDetailsProviderEnum string

// Set of constants representing the allowable values for AutonomousDatabaseEncryptionKeyDetailsProviderEnum
const (
	AutonomousDatabaseEncryptionKeyDetailsProviderAws           AutonomousDatabaseEncryptionKeyDetailsProviderEnum = "AWS"
	AutonomousDatabaseEncryptionKeyDetailsProviderAzure         AutonomousDatabaseEncryptionKeyDetailsProviderEnum = "AZURE"
	AutonomousDatabaseEncryptionKeyDetailsProviderOci           AutonomousDatabaseEncryptionKeyDetailsProviderEnum = "OCI"
	AutonomousDatabaseEncryptionKeyDetailsProviderOracleManaged AutonomousDatabaseEncryptionKeyDetailsProviderEnum = "ORACLE_MANAGED"
	AutonomousDatabaseEncryptionKeyDetailsProviderOkv           AutonomousDatabaseEncryptionKeyDetailsProviderEnum = "OKV"
)

var mappingAutonomousDatabaseEncryptionKeyDetailsProviderEnum = map[string]AutonomousDatabaseEncryptionKeyDetailsProviderEnum{
	"AWS":            AutonomousDatabaseEncryptionKeyDetailsProviderAws,
	"AZURE":          AutonomousDatabaseEncryptionKeyDetailsProviderAzure,
	"OCI":            AutonomousDatabaseEncryptionKeyDetailsProviderOci,
	"ORACLE_MANAGED": AutonomousDatabaseEncryptionKeyDetailsProviderOracleManaged,
	"OKV":            AutonomousDatabaseEncryptionKeyDetailsProviderOkv,
}

var mappingAutonomousDatabaseEncryptionKeyDetailsProviderEnumLowerCase = map[string]AutonomousDatabaseEncryptionKeyDetailsProviderEnum{
	"aws":            AutonomousDatabaseEncryptionKeyDetailsProviderAws,
	"azure":          AutonomousDatabaseEncryptionKeyDetailsProviderAzure,
	"oci":            AutonomousDatabaseEncryptionKeyDetailsProviderOci,
	"oracle_managed": AutonomousDatabaseEncryptionKeyDetailsProviderOracleManaged,
	"okv":            AutonomousDatabaseEncryptionKeyDetailsProviderOkv,
}

// GetAutonomousDatabaseEncryptionKeyDetailsProviderEnumValues Enumerates the set of values for AutonomousDatabaseEncryptionKeyDetailsProviderEnum
func GetAutonomousDatabaseEncryptionKeyDetailsProviderEnumValues() []AutonomousDatabaseEncryptionKeyDetailsProviderEnum {
	values := make([]AutonomousDatabaseEncryptionKeyDetailsProviderEnum, 0)
	for _, v := range mappingAutonomousDatabaseEncryptionKeyDetailsProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseEncryptionKeyDetailsProviderEnumStringValues Enumerates the set of values in String for AutonomousDatabaseEncryptionKeyDetailsProviderEnum
func GetAutonomousDatabaseEncryptionKeyDetailsProviderEnumStringValues() []string {
	return []string{
		"AWS",
		"AZURE",
		"OCI",
		"ORACLE_MANAGED",
		"OKV",
	}
}

// GetMappingAutonomousDatabaseEncryptionKeyDetailsProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseEncryptionKeyDetailsProviderEnum(val string) (AutonomousDatabaseEncryptionKeyDetailsProviderEnum, bool) {
	enum, ok := mappingAutonomousDatabaseEncryptionKeyDetailsProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
