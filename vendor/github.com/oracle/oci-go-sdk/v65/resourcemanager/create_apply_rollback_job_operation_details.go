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

// CreateApplyRollbackJobOperationDetails Job details that are specific to an apply rollback job. For more information about apply rollback jobs, see
// Creating an Apply Rollback Job (https://docs.oracle.com/iaas/Content/ResourceManager/Tasks/create-job-apply-rollback.htm).
type CreateApplyRollbackJobOperationDetails struct {

	// Specifies whether or not to upgrade provider versions.
	// Within the version constraints of your Terraform configuration, use the latest versions available from the source of Terraform providers.
	// For more information about this option, see Dependency Lock File (terraform.io) (https://www.terraform.io/language/files/dependency-lock).
	IsProviderUpgradeRequired *bool `mandatory:"false" json:"isProviderUpgradeRequired"`

	TerraformAdvancedOptions *TerraformAdvancedOptions `mandatory:"false" json:"terraformAdvancedOptions"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a plan rollback job, for use when specifying `"FROM_PLAN_ROLLBACK_JOB_ID"` as the `executionPlanRollbackStrategy`.
	ExecutionPlanRollbackJobId *string `mandatory:"false" json:"executionPlanRollbackJobId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a successful apply job, for use when specifying `"AUTO_APPROVED"` as the `executionPlanRollbackStrategy`.
	TargetRollbackJobId *string `mandatory:"false" json:"targetRollbackJobId"`

	// Specifies the source of the execution plan for rollback to apply.
	// Use `AUTO_APPROVED` to run the job without an execution plan for rollback job.
	ExecutionPlanRollbackStrategy ApplyRollbackJobOperationDetailsExecutionPlanRollbackStrategyEnum `mandatory:"true" json:"executionPlanRollbackStrategy"`
}

// GetIsProviderUpgradeRequired returns IsProviderUpgradeRequired
func (m CreateApplyRollbackJobOperationDetails) GetIsProviderUpgradeRequired() *bool {
	return m.IsProviderUpgradeRequired
}

func (m CreateApplyRollbackJobOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateApplyRollbackJobOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApplyRollbackJobOperationDetailsExecutionPlanRollbackStrategyEnum(string(m.ExecutionPlanRollbackStrategy)); !ok && m.ExecutionPlanRollbackStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionPlanRollbackStrategy: %s. Supported values are: %s.", m.ExecutionPlanRollbackStrategy, strings.Join(GetApplyRollbackJobOperationDetailsExecutionPlanRollbackStrategyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateApplyRollbackJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateApplyRollbackJobOperationDetails CreateApplyRollbackJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeCreateApplyRollbackJobOperationDetails
	}{
		"APPLY_ROLLBACK",
		(MarshalTypeCreateApplyRollbackJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}
