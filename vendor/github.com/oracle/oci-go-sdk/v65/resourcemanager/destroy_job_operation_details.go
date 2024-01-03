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

// DestroyJobOperationDetails Job details that are specific to destroy operations.
type DestroyJobOperationDetails struct {
	TerraformAdvancedOptions *TerraformAdvancedOptions `mandatory:"false" json:"terraformAdvancedOptions"`

	// Specifies the source of the execution plan to apply.
	// Currently, only `AUTO_APPROVED` is allowed, which indicates that the job
	// will be run without an execution plan.
	ExecutionPlanStrategy DestroyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"true" json:"executionPlanStrategy"`
}

func (m DestroyJobOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DestroyJobOperationDetails) ValidateEnumValue() (bool, error) {
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
func (m DestroyJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDestroyJobOperationDetails DestroyJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeDestroyJobOperationDetails
	}{
		"DESTROY",
		(MarshalTypeDestroyJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}

// DestroyJobOperationDetailsExecutionPlanStrategyEnum Enum with underlying type: string
type DestroyJobOperationDetailsExecutionPlanStrategyEnum string

// Set of constants representing the allowable values for DestroyJobOperationDetailsExecutionPlanStrategyEnum
const (
	DestroyJobOperationDetailsExecutionPlanStrategyAutoApproved DestroyJobOperationDetailsExecutionPlanStrategyEnum = "AUTO_APPROVED"
)

var mappingDestroyJobOperationDetailsExecutionPlanStrategyEnum = map[string]DestroyJobOperationDetailsExecutionPlanStrategyEnum{
	"AUTO_APPROVED": DestroyJobOperationDetailsExecutionPlanStrategyAutoApproved,
}

var mappingDestroyJobOperationDetailsExecutionPlanStrategyEnumLowerCase = map[string]DestroyJobOperationDetailsExecutionPlanStrategyEnum{
	"auto_approved": DestroyJobOperationDetailsExecutionPlanStrategyAutoApproved,
}

// GetDestroyJobOperationDetailsExecutionPlanStrategyEnumValues Enumerates the set of values for DestroyJobOperationDetailsExecutionPlanStrategyEnum
func GetDestroyJobOperationDetailsExecutionPlanStrategyEnumValues() []DestroyJobOperationDetailsExecutionPlanStrategyEnum {
	values := make([]DestroyJobOperationDetailsExecutionPlanStrategyEnum, 0)
	for _, v := range mappingDestroyJobOperationDetailsExecutionPlanStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetDestroyJobOperationDetailsExecutionPlanStrategyEnumStringValues Enumerates the set of values in String for DestroyJobOperationDetailsExecutionPlanStrategyEnum
func GetDestroyJobOperationDetailsExecutionPlanStrategyEnumStringValues() []string {
	return []string{
		"AUTO_APPROVED",
	}
}

// GetMappingDestroyJobOperationDetailsExecutionPlanStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDestroyJobOperationDetailsExecutionPlanStrategyEnum(val string) (DestroyJobOperationDetailsExecutionPlanStrategyEnum, bool) {
	enum, ok := mappingDestroyJobOperationDetailsExecutionPlanStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
