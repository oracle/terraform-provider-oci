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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FailureDetails The representation of FailureDetails
type FailureDetails struct {

	// Job failure reason.
	Code FailureDetailsCodeEnum `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`
}

func (m FailureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FailureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFailureDetailsCodeEnum(string(m.Code)); !ok && m.Code != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Code: %s. Supported values are: %s.", m.Code, strings.Join(GetFailureDetailsCodeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FailureDetailsCodeEnum Enum with underlying type: string
type FailureDetailsCodeEnum string

// Set of constants representing the allowable values for FailureDetailsCodeEnum
const (
	FailureDetailsCodeInternalServiceError                                    FailureDetailsCodeEnum = "INTERNAL_SERVICE_ERROR"
	FailureDetailsCodeTerraformExecutionError                                 FailureDetailsCodeEnum = "TERRAFORM_EXECUTION_ERROR"
	FailureDetailsCodeTerraformConfigUnzipFailed                              FailureDetailsCodeEnum = "TERRAFORM_CONFIG_UNZIP_FAILED"
	FailureDetailsCodeInvalidWorkingDirectory                                 FailureDetailsCodeEnum = "INVALID_WORKING_DIRECTORY"
	FailureDetailsCodeJobTimeout                                              FailureDetailsCodeEnum = "JOB_TIMEOUT"
	FailureDetailsCodeTerraformConfigVirusFound                               FailureDetailsCodeEnum = "TERRAFORM_CONFIG_VIRUS_FOUND"
	FailureDetailsCodeTerraformGitCloneFailure                                FailureDetailsCodeEnum = "TERRAFORM_GIT_CLONE_FAILURE"
	FailureDetailsCodeTerraformGitCheckoutFailure                             FailureDetailsCodeEnum = "TERRAFORM_GIT_CHECKOUT_FAILURE"
	FailureDetailsCodeTerraformObjectStorageConfigSourceEmptyBucket           FailureDetailsCodeEnum = "TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_EMPTY_BUCKET"
	FailureDetailsCodeTerraformObjectStorageConfigSourceNoTfFilePresent       FailureDetailsCodeEnum = "TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_NO_TF_FILE_PRESENT"
	FailureDetailsCodeTerraformObjectStorageConfigSourceUnsupportedObjectSize FailureDetailsCodeEnum = "TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_UNSUPPORTED_OBJECT_SIZE"
	FailureDetailsCodeCustomTerraformProviderBucketNotFound                   FailureDetailsCodeEnum = "CUSTOM_TERRAFORM_PROVIDER_BUCKET_NOT_FOUND"
	FailureDetailsCodeCustomTerraformProviderUnsupportedObjectSize            FailureDetailsCodeEnum = "CUSTOM_TERRAFORM_PROVIDER_UNSUPPORTED_OBJECT_SIZE"
)

var mappingFailureDetailsCodeEnum = map[string]FailureDetailsCodeEnum{
	"INTERNAL_SERVICE_ERROR":                                         FailureDetailsCodeInternalServiceError,
	"TERRAFORM_EXECUTION_ERROR":                                      FailureDetailsCodeTerraformExecutionError,
	"TERRAFORM_CONFIG_UNZIP_FAILED":                                  FailureDetailsCodeTerraformConfigUnzipFailed,
	"INVALID_WORKING_DIRECTORY":                                      FailureDetailsCodeInvalidWorkingDirectory,
	"JOB_TIMEOUT":                                                    FailureDetailsCodeJobTimeout,
	"TERRAFORM_CONFIG_VIRUS_FOUND":                                   FailureDetailsCodeTerraformConfigVirusFound,
	"TERRAFORM_GIT_CLONE_FAILURE":                                    FailureDetailsCodeTerraformGitCloneFailure,
	"TERRAFORM_GIT_CHECKOUT_FAILURE":                                 FailureDetailsCodeTerraformGitCheckoutFailure,
	"TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_EMPTY_BUCKET":            FailureDetailsCodeTerraformObjectStorageConfigSourceEmptyBucket,
	"TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_NO_TF_FILE_PRESENT":      FailureDetailsCodeTerraformObjectStorageConfigSourceNoTfFilePresent,
	"TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_UNSUPPORTED_OBJECT_SIZE": FailureDetailsCodeTerraformObjectStorageConfigSourceUnsupportedObjectSize,
	"CUSTOM_TERRAFORM_PROVIDER_BUCKET_NOT_FOUND":                     FailureDetailsCodeCustomTerraformProviderBucketNotFound,
	"CUSTOM_TERRAFORM_PROVIDER_UNSUPPORTED_OBJECT_SIZE":              FailureDetailsCodeCustomTerraformProviderUnsupportedObjectSize,
}

var mappingFailureDetailsCodeEnumLowerCase = map[string]FailureDetailsCodeEnum{
	"internal_service_error":                                         FailureDetailsCodeInternalServiceError,
	"terraform_execution_error":                                      FailureDetailsCodeTerraformExecutionError,
	"terraform_config_unzip_failed":                                  FailureDetailsCodeTerraformConfigUnzipFailed,
	"invalid_working_directory":                                      FailureDetailsCodeInvalidWorkingDirectory,
	"job_timeout":                                                    FailureDetailsCodeJobTimeout,
	"terraform_config_virus_found":                                   FailureDetailsCodeTerraformConfigVirusFound,
	"terraform_git_clone_failure":                                    FailureDetailsCodeTerraformGitCloneFailure,
	"terraform_git_checkout_failure":                                 FailureDetailsCodeTerraformGitCheckoutFailure,
	"terraform_object_storage_config_source_empty_bucket":            FailureDetailsCodeTerraformObjectStorageConfigSourceEmptyBucket,
	"terraform_object_storage_config_source_no_tf_file_present":      FailureDetailsCodeTerraformObjectStorageConfigSourceNoTfFilePresent,
	"terraform_object_storage_config_source_unsupported_object_size": FailureDetailsCodeTerraformObjectStorageConfigSourceUnsupportedObjectSize,
	"custom_terraform_provider_bucket_not_found":                     FailureDetailsCodeCustomTerraformProviderBucketNotFound,
	"custom_terraform_provider_unsupported_object_size":              FailureDetailsCodeCustomTerraformProviderUnsupportedObjectSize,
}

// GetFailureDetailsCodeEnumValues Enumerates the set of values for FailureDetailsCodeEnum
func GetFailureDetailsCodeEnumValues() []FailureDetailsCodeEnum {
	values := make([]FailureDetailsCodeEnum, 0)
	for _, v := range mappingFailureDetailsCodeEnum {
		values = append(values, v)
	}
	return values
}

// GetFailureDetailsCodeEnumStringValues Enumerates the set of values in String for FailureDetailsCodeEnum
func GetFailureDetailsCodeEnumStringValues() []string {
	return []string{
		"INTERNAL_SERVICE_ERROR",
		"TERRAFORM_EXECUTION_ERROR",
		"TERRAFORM_CONFIG_UNZIP_FAILED",
		"INVALID_WORKING_DIRECTORY",
		"JOB_TIMEOUT",
		"TERRAFORM_CONFIG_VIRUS_FOUND",
		"TERRAFORM_GIT_CLONE_FAILURE",
		"TERRAFORM_GIT_CHECKOUT_FAILURE",
		"TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_EMPTY_BUCKET",
		"TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_NO_TF_FILE_PRESENT",
		"TERRAFORM_OBJECT_STORAGE_CONFIG_SOURCE_UNSUPPORTED_OBJECT_SIZE",
		"CUSTOM_TERRAFORM_PROVIDER_BUCKET_NOT_FOUND",
		"CUSTOM_TERRAFORM_PROVIDER_UNSUPPORTED_OBJECT_SIZE",
	}
}

// GetMappingFailureDetailsCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFailureDetailsCodeEnum(val string) (FailureDetailsCodeEnum, bool) {
	enum, ok := mappingFailureDetailsCodeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
