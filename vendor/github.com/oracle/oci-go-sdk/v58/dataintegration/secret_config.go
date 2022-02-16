// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SecretConfig Secret configuration if used for storing sensitive info
type SecretConfig interface {
}

type secretconfig struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *secretconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecretconfig secretconfig
	s := struct {
		Model Unmarshalersecretconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *secretconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "OCI_VAULT_SECRET_CONFIG":
		mm := OciVaultSecretConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m secretconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m secretconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecretConfigModelTypeEnum Enum with underlying type: string
type SecretConfigModelTypeEnum string

// Set of constants representing the allowable values for SecretConfigModelTypeEnum
const (
	SecretConfigModelTypeOciVaultSecretConfig SecretConfigModelTypeEnum = "OCI_VAULT_SECRET_CONFIG"
)

var mappingSecretConfigModelTypeEnum = map[string]SecretConfigModelTypeEnum{
	"OCI_VAULT_SECRET_CONFIG": SecretConfigModelTypeOciVaultSecretConfig,
}

// GetSecretConfigModelTypeEnumValues Enumerates the set of values for SecretConfigModelTypeEnum
func GetSecretConfigModelTypeEnumValues() []SecretConfigModelTypeEnum {
	values := make([]SecretConfigModelTypeEnum, 0)
	for _, v := range mappingSecretConfigModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretConfigModelTypeEnumStringValues Enumerates the set of values in String for SecretConfigModelTypeEnum
func GetSecretConfigModelTypeEnumStringValues() []string {
	return []string{
		"OCI_VAULT_SECRET_CONFIG",
	}
}

// GetMappingSecretConfigModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretConfigModelTypeEnum(val string) (SecretConfigModelTypeEnum, bool) {
	mappingSecretConfigModelTypeEnumIgnoreCase := make(map[string]SecretConfigModelTypeEnum)
	for k, v := range mappingSecretConfigModelTypeEnum {
		mappingSecretConfigModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSecretConfigModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
