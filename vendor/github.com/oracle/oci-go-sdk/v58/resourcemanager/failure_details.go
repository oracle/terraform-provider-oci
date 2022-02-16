// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
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
	}
}

// GetMappingFailureDetailsCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFailureDetailsCodeEnum(val string) (FailureDetailsCodeEnum, bool) {
	mappingFailureDetailsCodeEnumIgnoreCase := make(map[string]FailureDetailsCodeEnum)
	for k, v := range mappingFailureDetailsCodeEnum {
		mappingFailureDetailsCodeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingFailureDetailsCodeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
