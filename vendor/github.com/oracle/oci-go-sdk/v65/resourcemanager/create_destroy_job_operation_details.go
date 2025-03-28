// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDestroyJobOperationDetails Job details that are specific to destroy operations.
type CreateDestroyJobOperationDetails struct {

	// Specifies whether or not to upgrade provider versions.
	// Within the version constraints of your Terraform configuration, use the latest versions available from the source of Terraform providers.
	// For more information about this option, see Dependency Lock File (terraform.io) (https://www.terraform.io/language/files/dependency-lock).
	IsProviderUpgradeRequired *bool `mandatory:"false" json:"isProviderUpgradeRequired"`

	TerraformAdvancedOptions *TerraformAdvancedOptions `mandatory:"false" json:"terraformAdvancedOptions"`

	// Specifies the source of the execution plan to apply.
	// Currently, only `AUTO_APPROVED` is allowed, which indicates that the job
	// will be run without an execution plan.
	ExecutionPlanStrategy DestroyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"true" json:"executionPlanStrategy"`
}

// GetIsProviderUpgradeRequired returns IsProviderUpgradeRequired
func (m CreateDestroyJobOperationDetails) GetIsProviderUpgradeRequired() *bool {
	return m.IsProviderUpgradeRequired
}

func (m CreateDestroyJobOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDestroyJobOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDestroyJobOperationDetailsExecutionPlanStrategyEnum(string(m.ExecutionPlanStrategy)); !ok && m.ExecutionPlanStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionPlanStrategy: %s. Supported values are: %s.", m.ExecutionPlanStrategy, strings.Join(GetDestroyJobOperationDetailsExecutionPlanStrategyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDestroyJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDestroyJobOperationDetails CreateDestroyJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeCreateDestroyJobOperationDetails
	}{
		"DESTROY",
		(MarshalTypeCreateDestroyJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}
