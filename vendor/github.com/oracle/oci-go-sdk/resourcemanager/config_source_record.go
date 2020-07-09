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

// ConfigSourceRecord Information about the Terraform configuration.
type ConfigSourceRecord interface {
}

type configsourcerecord struct {
	JsonData               []byte
	ConfigSourceRecordType string `json:"configSourceRecordType"`
}

// UnmarshalJSON unmarshals json
func (m *configsourcerecord) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigsourcerecord configsourcerecord
	s := struct {
		Model Unmarshalerconfigsourcerecord
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConfigSourceRecordType = s.Model.ConfigSourceRecordType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configsourcerecord) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceRecordType {
	case "GIT_CONFIG_SOURCE":
		mm := GitConfigSourceRecord{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZIP_UPLOAD":
		mm := ZipUploadConfigSourceRecord{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m configsourcerecord) String() string {
	return common.PointerString(m)
}

// ConfigSourceRecordConfigSourceRecordTypeEnum Enum with underlying type: string
type ConfigSourceRecordConfigSourceRecordTypeEnum string

// Set of constants representing the allowable values for ConfigSourceRecordConfigSourceRecordTypeEnum
const (
	ConfigSourceRecordConfigSourceRecordTypeZipUpload       ConfigSourceRecordConfigSourceRecordTypeEnum = "ZIP_UPLOAD"
	ConfigSourceRecordConfigSourceRecordTypeGitConfigSource ConfigSourceRecordConfigSourceRecordTypeEnum = "GIT_CONFIG_SOURCE"
)

var mappingConfigSourceRecordConfigSourceRecordType = map[string]ConfigSourceRecordConfigSourceRecordTypeEnum{
	"ZIP_UPLOAD":        ConfigSourceRecordConfigSourceRecordTypeZipUpload,
	"GIT_CONFIG_SOURCE": ConfigSourceRecordConfigSourceRecordTypeGitConfigSource,
}

// GetConfigSourceRecordConfigSourceRecordTypeEnumValues Enumerates the set of values for ConfigSourceRecordConfigSourceRecordTypeEnum
func GetConfigSourceRecordConfigSourceRecordTypeEnumValues() []ConfigSourceRecordConfigSourceRecordTypeEnum {
	values := make([]ConfigSourceRecordConfigSourceRecordTypeEnum, 0)
	for _, v := range mappingConfigSourceRecordConfigSourceRecordType {
		values = append(values, v)
	}
	return values
}
