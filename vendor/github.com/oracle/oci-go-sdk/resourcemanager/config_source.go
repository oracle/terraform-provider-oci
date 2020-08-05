// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ConfigSource Information about the Terraform configuration.
type ConfigSource interface {

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	GetWorkingDirectory() *string
}

type configsource struct {
	JsonData         []byte
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
	ConfigSourceType string  `json:"configSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *configsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigsource configsource
	s := struct {
		Model Unmarshalerconfigsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WorkingDirectory = s.Model.WorkingDirectory
	m.ConfigSourceType = s.Model.ConfigSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceType {
	case "GIT_CONFIG_SOURCE":
		mm := GitConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPARTMENT_CONFIG_SOURCE":
		mm := CompartmentConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZIP_UPLOAD":
		mm := ZipUploadConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetWorkingDirectory returns WorkingDirectory
func (m configsource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m configsource) String() string {
	return common.PointerString(m)
}

// ConfigSourceConfigSourceTypeEnum Enum with underlying type: string
type ConfigSourceConfigSourceTypeEnum string

// Set of constants representing the allowable values for ConfigSourceConfigSourceTypeEnum
const (
	ConfigSourceConfigSourceTypeZipUpload               ConfigSourceConfigSourceTypeEnum = "ZIP_UPLOAD"
	ConfigSourceConfigSourceTypeGitConfigSource         ConfigSourceConfigSourceTypeEnum = "GIT_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeCompartmentConfigSource ConfigSourceConfigSourceTypeEnum = "COMPARTMENT_CONFIG_SOURCE"
)

var mappingConfigSourceConfigSourceType = map[string]ConfigSourceConfigSourceTypeEnum{
	"ZIP_UPLOAD":                ConfigSourceConfigSourceTypeZipUpload,
	"GIT_CONFIG_SOURCE":         ConfigSourceConfigSourceTypeGitConfigSource,
	"COMPARTMENT_CONFIG_SOURCE": ConfigSourceConfigSourceTypeCompartmentConfigSource,
}

// GetConfigSourceConfigSourceTypeEnumValues Enumerates the set of values for ConfigSourceConfigSourceTypeEnum
func GetConfigSourceConfigSourceTypeEnumValues() []ConfigSourceConfigSourceTypeEnum {
	values := make([]ConfigSourceConfigSourceTypeEnum, 0)
	for _, v := range mappingConfigSourceConfigSourceType {
		values = append(values, v)
	}
	return values
}
