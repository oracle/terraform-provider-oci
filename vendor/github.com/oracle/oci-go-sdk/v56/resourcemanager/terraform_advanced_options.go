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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TerraformAdvancedOptions Specifies advanced options for Terraform commands. These options are not necessary for normal usage of Terraform.
type TerraformAdvancedOptions struct {

	// Specifies whether to refresh the state for each resource before running the job (operation).
	// Refreshing the state can affect performance. Consider setting to `false` if the configuration includes several resources.
	// Used with the following operations: `PLAN`, `APPLY`, `DESTROY`.
	IsRefreshRequired *bool `mandatory:"false" json:"isRefreshRequired"`

	// Limits the number of concurrent Terraform operations when walking the graph (https://www.terraform.io/docs/internals/graph.html#walking-the-graph).
	// Use this parameter to help debug Terraform issues or to accomplish certain special use cases.
	// A higher value might cause resources to be throttled.
	// Used with the following operations: `PLAN`, `APPLY`, `DESTROY`.
	Parallelism *int `mandatory:"false" json:"parallelism"`

	// Enables detailed logs at the specified verbosity for running the job (operation).
	// Used with the following operations: `PLAN`, `APPLY`, `DESTROY`.
	DetailedLogLevel TerraformAdvancedOptionsDetailedLogLevelEnum `mandatory:"false" json:"detailedLogLevel,omitempty"`
}

func (m TerraformAdvancedOptions) String() string {
	return common.PointerString(m)
}

// TerraformAdvancedOptionsDetailedLogLevelEnum Enum with underlying type: string
type TerraformAdvancedOptionsDetailedLogLevelEnum string

// Set of constants representing the allowable values for TerraformAdvancedOptionsDetailedLogLevelEnum
const (
	TerraformAdvancedOptionsDetailedLogLevelError TerraformAdvancedOptionsDetailedLogLevelEnum = "ERROR"
	TerraformAdvancedOptionsDetailedLogLevelWarn  TerraformAdvancedOptionsDetailedLogLevelEnum = "WARN"
	TerraformAdvancedOptionsDetailedLogLevelInfo  TerraformAdvancedOptionsDetailedLogLevelEnum = "INFO"
	TerraformAdvancedOptionsDetailedLogLevelDebug TerraformAdvancedOptionsDetailedLogLevelEnum = "DEBUG"
	TerraformAdvancedOptionsDetailedLogLevelTrace TerraformAdvancedOptionsDetailedLogLevelEnum = "TRACE"
)

var mappingTerraformAdvancedOptionsDetailedLogLevel = map[string]TerraformAdvancedOptionsDetailedLogLevelEnum{
	"ERROR": TerraformAdvancedOptionsDetailedLogLevelError,
	"WARN":  TerraformAdvancedOptionsDetailedLogLevelWarn,
	"INFO":  TerraformAdvancedOptionsDetailedLogLevelInfo,
	"DEBUG": TerraformAdvancedOptionsDetailedLogLevelDebug,
	"TRACE": TerraformAdvancedOptionsDetailedLogLevelTrace,
}

// GetTerraformAdvancedOptionsDetailedLogLevelEnumValues Enumerates the set of values for TerraformAdvancedOptionsDetailedLogLevelEnum
func GetTerraformAdvancedOptionsDetailedLogLevelEnumValues() []TerraformAdvancedOptionsDetailedLogLevelEnum {
	values := make([]TerraformAdvancedOptionsDetailedLogLevelEnum, 0)
	for _, v := range mappingTerraformAdvancedOptionsDetailedLogLevel {
		values = append(values, v)
	}
	return values
}
