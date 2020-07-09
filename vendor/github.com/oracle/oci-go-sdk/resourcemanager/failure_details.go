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
	"github.com/oracle/oci-go-sdk/common"
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

// FailureDetailsCodeEnum Enum with underlying type: string
type FailureDetailsCodeEnum string

// Set of constants representing the allowable values for FailureDetailsCodeEnum
const (
	FailureDetailsCodeInternalServiceError        FailureDetailsCodeEnum = "INTERNAL_SERVICE_ERROR"
	FailureDetailsCodeTerraformExecutionError     FailureDetailsCodeEnum = "TERRAFORM_EXECUTION_ERROR"
	FailureDetailsCodeTerraformConfigUnzipFailed  FailureDetailsCodeEnum = "TERRAFORM_CONFIG_UNZIP_FAILED"
	FailureDetailsCodeInvalidWorkingDirectory     FailureDetailsCodeEnum = "INVALID_WORKING_DIRECTORY"
	FailureDetailsCodeJobTimeout                  FailureDetailsCodeEnum = "JOB_TIMEOUT"
	FailureDetailsCodeTerraformConfigVirusFound   FailureDetailsCodeEnum = "TERRAFORM_CONFIG_VIRUS_FOUND"
	FailureDetailsCodeTerraformGitCloneFailure    FailureDetailsCodeEnum = "TERRAFORM_GIT_CLONE_FAILURE"
	FailureDetailsCodeTerraformGitCheckoutFailure FailureDetailsCodeEnum = "TERRAFORM_GIT_CHECKOUT_FAILURE"
)

var mappingFailureDetailsCode = map[string]FailureDetailsCodeEnum{
	"INTERNAL_SERVICE_ERROR":         FailureDetailsCodeInternalServiceError,
	"TERRAFORM_EXECUTION_ERROR":      FailureDetailsCodeTerraformExecutionError,
	"TERRAFORM_CONFIG_UNZIP_FAILED":  FailureDetailsCodeTerraformConfigUnzipFailed,
	"INVALID_WORKING_DIRECTORY":      FailureDetailsCodeInvalidWorkingDirectory,
	"JOB_TIMEOUT":                    FailureDetailsCodeJobTimeout,
	"TERRAFORM_CONFIG_VIRUS_FOUND":   FailureDetailsCodeTerraformConfigVirusFound,
	"TERRAFORM_GIT_CLONE_FAILURE":    FailureDetailsCodeTerraformGitCloneFailure,
	"TERRAFORM_GIT_CHECKOUT_FAILURE": FailureDetailsCodeTerraformGitCheckoutFailure,
}

// GetFailureDetailsCodeEnumValues Enumerates the set of values for FailureDetailsCodeEnum
func GetFailureDetailsCodeEnumValues() []FailureDetailsCodeEnum {
	values := make([]FailureDetailsCodeEnum, 0)
	for _, v := range mappingFailureDetailsCode {
		values = append(values, v)
	}
	return values
}
