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

// ConfigSource Information about the Terraform configuration.
type ConfigSource interface {

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// Required when using a zip Terraform configuration (`configSourceType` value of `ZIP_UPLOAD`) that contains folders.
	// Ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	// For more information about required and recommended file structure, see
	// File Structure (Terraform Configurations for Resource Manager) (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/terraformconfigresourcemanager.htm#filestructure).
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
	case "DEVOPS_CONFIG_SOURCE":
		mm := DevOpsConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_CONFIG_SOURCE":
		mm := ObjectStorageConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_SERVER_CONFIG_SOURCE":
		mm := BitbucketServerConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_CLOUD_CONFIG_SOURCE":
		mm := BitbucketCloudConfigSource{}
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
		common.Logf("Recieved unsupported enum value for ConfigSource: %s.", m.ConfigSourceType)
		return *m, nil
	}
}

// GetWorkingDirectory returns WorkingDirectory
func (m configsource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m configsource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigSourceConfigSourceTypeEnum Enum with underlying type: string
type ConfigSourceConfigSourceTypeEnum string

// Set of constants representing the allowable values for ConfigSourceConfigSourceTypeEnum
const (
	ConfigSourceConfigSourceTypeBitbucketCloudConfigSource  ConfigSourceConfigSourceTypeEnum = "BITBUCKET_CLOUD_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeBitbucketServerConfigSource ConfigSourceConfigSourceTypeEnum = "BITBUCKET_SERVER_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeCompartmentConfigSource     ConfigSourceConfigSourceTypeEnum = "COMPARTMENT_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeDevopsConfigSource          ConfigSourceConfigSourceTypeEnum = "DEVOPS_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeGitConfigSource             ConfigSourceConfigSourceTypeEnum = "GIT_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeObjectStorageConfigSource   ConfigSourceConfigSourceTypeEnum = "OBJECT_STORAGE_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeZipUpload                   ConfigSourceConfigSourceTypeEnum = "ZIP_UPLOAD"
)

var mappingConfigSourceConfigSourceTypeEnum = map[string]ConfigSourceConfigSourceTypeEnum{
	"BITBUCKET_CLOUD_CONFIG_SOURCE":  ConfigSourceConfigSourceTypeBitbucketCloudConfigSource,
	"BITBUCKET_SERVER_CONFIG_SOURCE": ConfigSourceConfigSourceTypeBitbucketServerConfigSource,
	"COMPARTMENT_CONFIG_SOURCE":      ConfigSourceConfigSourceTypeCompartmentConfigSource,
	"DEVOPS_CONFIG_SOURCE":           ConfigSourceConfigSourceTypeDevopsConfigSource,
	"GIT_CONFIG_SOURCE":              ConfigSourceConfigSourceTypeGitConfigSource,
	"OBJECT_STORAGE_CONFIG_SOURCE":   ConfigSourceConfigSourceTypeObjectStorageConfigSource,
	"ZIP_UPLOAD":                     ConfigSourceConfigSourceTypeZipUpload,
}

var mappingConfigSourceConfigSourceTypeEnumLowerCase = map[string]ConfigSourceConfigSourceTypeEnum{
	"bitbucket_cloud_config_source":  ConfigSourceConfigSourceTypeBitbucketCloudConfigSource,
	"bitbucket_server_config_source": ConfigSourceConfigSourceTypeBitbucketServerConfigSource,
	"compartment_config_source":      ConfigSourceConfigSourceTypeCompartmentConfigSource,
	"devops_config_source":           ConfigSourceConfigSourceTypeDevopsConfigSource,
	"git_config_source":              ConfigSourceConfigSourceTypeGitConfigSource,
	"object_storage_config_source":   ConfigSourceConfigSourceTypeObjectStorageConfigSource,
	"zip_upload":                     ConfigSourceConfigSourceTypeZipUpload,
}

// GetConfigSourceConfigSourceTypeEnumValues Enumerates the set of values for ConfigSourceConfigSourceTypeEnum
func GetConfigSourceConfigSourceTypeEnumValues() []ConfigSourceConfigSourceTypeEnum {
	values := make([]ConfigSourceConfigSourceTypeEnum, 0)
	for _, v := range mappingConfigSourceConfigSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigSourceConfigSourceTypeEnumStringValues Enumerates the set of values in String for ConfigSourceConfigSourceTypeEnum
func GetConfigSourceConfigSourceTypeEnumStringValues() []string {
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

// GetMappingConfigSourceConfigSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigSourceConfigSourceTypeEnum(val string) (ConfigSourceConfigSourceTypeEnum, bool) {
	enum, ok := mappingConfigSourceConfigSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
