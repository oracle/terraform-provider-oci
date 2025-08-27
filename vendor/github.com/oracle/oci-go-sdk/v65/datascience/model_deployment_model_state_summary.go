// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelDeploymentModelStateSummary Status of the model in a model group deployment.
type ModelDeploymentModelStateSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployed model in model deployment.
	ModelId *string `mandatory:"true" json:"modelId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// SaaS friendly name for the model OCID.
	InferenceKey *string `mandatory:"true" json:"inferenceKey"`

	// The state of the deployed model in model deployment.
	State ModelDeploymentModelStateSummaryStateEnum `mandatory:"true" json:"state"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ModelDeploymentModelStateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelDeploymentModelStateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelDeploymentModelStateSummaryStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetModelDeploymentModelStateSummaryStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelDeploymentModelStateSummaryStateEnum Enum with underlying type: string
type ModelDeploymentModelStateSummaryStateEnum string

// Set of constants representing the allowable values for ModelDeploymentModelStateSummaryStateEnum
const (
	ModelDeploymentModelStateSummaryStateSuccess      ModelDeploymentModelStateSummaryStateEnum = "SUCCESS"
	ModelDeploymentModelStateSummaryStateFailed       ModelDeploymentModelStateSummaryStateEnum = "FAILED"
	ModelDeploymentModelStateSummaryStateInconsistent ModelDeploymentModelStateSummaryStateEnum = "INCONSISTENT"
)

var mappingModelDeploymentModelStateSummaryStateEnum = map[string]ModelDeploymentModelStateSummaryStateEnum{
	"SUCCESS":      ModelDeploymentModelStateSummaryStateSuccess,
	"FAILED":       ModelDeploymentModelStateSummaryStateFailed,
	"INCONSISTENT": ModelDeploymentModelStateSummaryStateInconsistent,
}

var mappingModelDeploymentModelStateSummaryStateEnumLowerCase = map[string]ModelDeploymentModelStateSummaryStateEnum{
	"success":      ModelDeploymentModelStateSummaryStateSuccess,
	"failed":       ModelDeploymentModelStateSummaryStateFailed,
	"inconsistent": ModelDeploymentModelStateSummaryStateInconsistent,
}

// GetModelDeploymentModelStateSummaryStateEnumValues Enumerates the set of values for ModelDeploymentModelStateSummaryStateEnum
func GetModelDeploymentModelStateSummaryStateEnumValues() []ModelDeploymentModelStateSummaryStateEnum {
	values := make([]ModelDeploymentModelStateSummaryStateEnum, 0)
	for _, v := range mappingModelDeploymentModelStateSummaryStateEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDeploymentModelStateSummaryStateEnumStringValues Enumerates the set of values in String for ModelDeploymentModelStateSummaryStateEnum
func GetModelDeploymentModelStateSummaryStateEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILED",
		"INCONSISTENT",
	}
}

// GetMappingModelDeploymentModelStateSummaryStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDeploymentModelStateSummaryStateEnum(val string) (ModelDeploymentModelStateSummaryStateEnum, bool) {
	enum, ok := mappingModelDeploymentModelStateSummaryStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
