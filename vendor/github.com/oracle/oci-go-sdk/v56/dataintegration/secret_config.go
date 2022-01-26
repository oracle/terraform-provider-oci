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
	"github.com/oracle/oci-go-sdk/v56/common"
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

// SecretConfigModelTypeEnum Enum with underlying type: string
type SecretConfigModelTypeEnum string

// Set of constants representing the allowable values for SecretConfigModelTypeEnum
const (
	SecretConfigModelTypeOciVaultSecretConfig SecretConfigModelTypeEnum = "OCI_VAULT_SECRET_CONFIG"
)

var mappingSecretConfigModelType = map[string]SecretConfigModelTypeEnum{
	"OCI_VAULT_SECRET_CONFIG": SecretConfigModelTypeOciVaultSecretConfig,
}

// GetSecretConfigModelTypeEnumValues Enumerates the set of values for SecretConfigModelTypeEnum
func GetSecretConfigModelTypeEnumValues() []SecretConfigModelTypeEnum {
	values := make([]SecretConfigModelTypeEnum, 0)
	for _, v := range mappingSecretConfigModelType {
		values = append(values, v)
	}
	return values
}
