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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ApplyJobOperationDetailsSummary Job details that are specific to apply operations.
type ApplyJobOperationDetailsSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the plan job that contains the execution plan used for this job,
	// or `null` if no execution plan was used.
	ExecutionPlanJobId *string `mandatory:"false" json:"executionPlanJobId"`

	// Specifies the source of the execution plan to apply.
	// Use `AUTO_APPROVED` to run the job without an execution plan.
	ExecutionPlanStrategy ApplyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"true" json:"executionPlanStrategy"`
}

func (m ApplyJobOperationDetailsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplyJobOperationDetailsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApplyJobOperationDetailsExecutionPlanStrategyEnum(string(m.ExecutionPlanStrategy)); !ok && m.ExecutionPlanStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionPlanStrategy: %s. Supported values are: %s.", m.ExecutionPlanStrategy, strings.Join(GetApplyJobOperationDetailsExecutionPlanStrategyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ApplyJobOperationDetailsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApplyJobOperationDetailsSummary ApplyJobOperationDetailsSummary
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeApplyJobOperationDetailsSummary
	}{
		"APPLY",
		(MarshalTypeApplyJobOperationDetailsSummary)(m),
	}

	return json.Marshal(&s)
}
