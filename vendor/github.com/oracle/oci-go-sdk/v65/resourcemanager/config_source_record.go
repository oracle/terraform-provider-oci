// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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
	case "DEVOPS_CONFIG_SOURCE":
		mm := DevOpsConfigSourceRecord{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GIT_CONFIG_SOURCE":
		mm := GitConfigSourceRecord{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZIP_UPLOAD":
		mm := ZipUploadConfigSourceRecord{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_CLOUD_CONFIG_SOURCE":
		mm := BitbucketCloudConfigSourceRecord{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_SERVER_CONFIG_SOURCE":
		mm := BitbucketServerConfigSourceRecord{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_CONFIG_SOURCE":
		mm := ObjectStorageConfigSourceRecord{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConfigSourceRecord: %s.", m.ConfigSourceRecordType)
		return *m, nil
	}
}

func (m configsourcerecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configsourcerecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigSourceRecordConfigSourceRecordTypeEnum Enum with underlying type: string
type ConfigSourceRecordConfigSourceRecordTypeEnum string

// Set of constants representing the allowable values for ConfigSourceRecordConfigSourceRecordTypeEnum
const (
	ConfigSourceRecordConfigSourceRecordTypeBitbucketCloudConfigSource  ConfigSourceRecordConfigSourceRecordTypeEnum = "BITBUCKET_CLOUD_CONFIG_SOURCE"
	ConfigSourceRecordConfigSourceRecordTypeBitbucketServerConfigSource ConfigSourceRecordConfigSourceRecordTypeEnum = "BITBUCKET_SERVER_CONFIG_SOURCE"
	ConfigSourceRecordConfigSourceRecordTypeCompartmentConfigSource     ConfigSourceRecordConfigSourceRecordTypeEnum = "COMPARTMENT_CONFIG_SOURCE"
	ConfigSourceRecordConfigSourceRecordTypeDevopsConfigSource          ConfigSourceRecordConfigSourceRecordTypeEnum = "DEVOPS_CONFIG_SOURCE"
	ConfigSourceRecordConfigSourceRecordTypeGitConfigSource             ConfigSourceRecordConfigSourceRecordTypeEnum = "GIT_CONFIG_SOURCE"
	ConfigSourceRecordConfigSourceRecordTypeObjectStorageConfigSource   ConfigSourceRecordConfigSourceRecordTypeEnum = "OBJECT_STORAGE_CONFIG_SOURCE"
	ConfigSourceRecordConfigSourceRecordTypeZipUpload                   ConfigSourceRecordConfigSourceRecordTypeEnum = "ZIP_UPLOAD"
)

var mappingConfigSourceRecordConfigSourceRecordTypeEnum = map[string]ConfigSourceRecordConfigSourceRecordTypeEnum{
	"BITBUCKET_CLOUD_CONFIG_SOURCE":  ConfigSourceRecordConfigSourceRecordTypeBitbucketCloudConfigSource,
	"BITBUCKET_SERVER_CONFIG_SOURCE": ConfigSourceRecordConfigSourceRecordTypeBitbucketServerConfigSource,
	"COMPARTMENT_CONFIG_SOURCE":      ConfigSourceRecordConfigSourceRecordTypeCompartmentConfigSource,
	"DEVOPS_CONFIG_SOURCE":           ConfigSourceRecordConfigSourceRecordTypeDevopsConfigSource,
	"GIT_CONFIG_SOURCE":              ConfigSourceRecordConfigSourceRecordTypeGitConfigSource,
	"OBJECT_STORAGE_CONFIG_SOURCE":   ConfigSourceRecordConfigSourceRecordTypeObjectStorageConfigSource,
	"ZIP_UPLOAD":                     ConfigSourceRecordConfigSourceRecordTypeZipUpload,
}

var mappingConfigSourceRecordConfigSourceRecordTypeEnumLowerCase = map[string]ConfigSourceRecordConfigSourceRecordTypeEnum{
	"bitbucket_cloud_config_source":  ConfigSourceRecordConfigSourceRecordTypeBitbucketCloudConfigSource,
	"bitbucket_server_config_source": ConfigSourceRecordConfigSourceRecordTypeBitbucketServerConfigSource,
	"compartment_config_source":      ConfigSourceRecordConfigSourceRecordTypeCompartmentConfigSource,
	"devops_config_source":           ConfigSourceRecordConfigSourceRecordTypeDevopsConfigSource,
	"git_config_source":              ConfigSourceRecordConfigSourceRecordTypeGitConfigSource,
	"object_storage_config_source":   ConfigSourceRecordConfigSourceRecordTypeObjectStorageConfigSource,
	"zip_upload":                     ConfigSourceRecordConfigSourceRecordTypeZipUpload,
}

// GetConfigSourceRecordConfigSourceRecordTypeEnumValues Enumerates the set of values for ConfigSourceRecordConfigSourceRecordTypeEnum
func GetConfigSourceRecordConfigSourceRecordTypeEnumValues() []ConfigSourceRecordConfigSourceRecordTypeEnum {
	values := make([]ConfigSourceRecordConfigSourceRecordTypeEnum, 0)
	for _, v := range mappingConfigSourceRecordConfigSourceRecordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigSourceRecordConfigSourceRecordTypeEnumStringValues Enumerates the set of values in String for ConfigSourceRecordConfigSourceRecordTypeEnum
func GetConfigSourceRecordConfigSourceRecordTypeEnumStringValues() []string {
	return []string{
		"BITBUCKET_CLOUD_CONFIG_SOURCE",
		"BITBUCKET_SERVER_CONFIG_SOURCE",
		"COMPARTMENT_CONFIG_SOURCE",
		"DEVOPS_CONFIG_SOURCE",
		"GIT_CONFIG_SOURCE",
		"OBJECT_STORAGE_CONFIG_SOURCE",
		"ZIP_UPLOAD",
	}
}

// GetMappingConfigSourceRecordConfigSourceRecordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigSourceRecordConfigSourceRecordTypeEnum(val string) (ConfigSourceRecordConfigSourceRecordTypeEnum, bool) {
	enum, ok := mappingConfigSourceRecordConfigSourceRecordTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
