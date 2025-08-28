// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerSummary Summary representation of a container
type ContainerSummary struct {

	// The name of the container. This can be same for different tags
	ContainerName *string `mandatory:"true" json:"containerName"`

	// The latest tag of the container.
	IsLatest *bool `mandatory:"true" json:"isLatest"`

	// Container Tag.
	Tag *string `mandatory:"true" json:"tag"`

	// Container Version LifecycleState.
	LifecycleState ContainerVersionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The display name of the container.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The family name of the container.
	FamilyName *string `mandatory:"false" json:"familyName"`

	// Description of the container.
	Description *string `mandatory:"false" json:"description"`

	// The list of target workload. This Container can be used with given data science resources.
	TargetWorkloads []ContainerSummaryTargetWorkloadsEnum `mandatory:"false" json:"targetWorkloads,omitempty"`

	// The list of usages of this container. This Container can be used for given use-cases.
	Usages []ContainerSummaryUsagesEnum `mandatory:"false" json:"usages,omitempty"`

	// workload configuration of the container.
	WorkloadConfigurationDetailsList []WorkloadConfigurationDetails `mandatory:"false" json:"workloadConfigurationDetailsList"`

	// An array of defined metadata details for the model.
	TagConfigurationList []TagConfiguration `mandatory:"false" json:"tagConfigurationList"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ContainerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerVersionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetContainerVersionLifecycleStateEnumStringValues(), ",")))
	}

	for _, val := range m.TargetWorkloads {
		if _, ok := GetMappingContainerSummaryTargetWorkloadsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetWorkloads: %s. Supported values are: %s.", val, strings.Join(GetContainerSummaryTargetWorkloadsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.Usages {
		if _, ok := GetMappingContainerSummaryUsagesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Usages: %s. Supported values are: %s.", val, strings.Join(GetContainerSummaryUsagesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ContainerSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                      *string                               `json:"displayName"`
		FamilyName                       *string                               `json:"familyName"`
		Description                      *string                               `json:"description"`
		TargetWorkloads                  []ContainerSummaryTargetWorkloadsEnum `json:"targetWorkloads"`
		Usages                           []ContainerSummaryUsagesEnum          `json:"usages"`
		WorkloadConfigurationDetailsList []workloadconfigurationdetails        `json:"workloadConfigurationDetailsList"`
		TagConfigurationList             []TagConfiguration                    `json:"tagConfigurationList"`
		FreeformTags                     map[string]string                     `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}     `json:"definedTags"`
		ContainerName                    *string                               `json:"containerName"`
		IsLatest                         *bool                                 `json:"isLatest"`
		Tag                              *string                               `json:"tag"`
		LifecycleState                   ContainerVersionLifecycleStateEnum    `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FamilyName = model.FamilyName

	m.Description = model.Description

	m.TargetWorkloads = make([]ContainerSummaryTargetWorkloadsEnum, len(model.TargetWorkloads))
	copy(m.TargetWorkloads, model.TargetWorkloads)
	m.Usages = make([]ContainerSummaryUsagesEnum, len(model.Usages))
	copy(m.Usages, model.Usages)
	m.WorkloadConfigurationDetailsList = make([]WorkloadConfigurationDetails, len(model.WorkloadConfigurationDetailsList))
	for i, n := range model.WorkloadConfigurationDetailsList {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.WorkloadConfigurationDetailsList[i] = nn.(WorkloadConfigurationDetails)
		} else {
			m.WorkloadConfigurationDetailsList[i] = nil
		}
	}
	m.TagConfigurationList = make([]TagConfiguration, len(model.TagConfigurationList))
	copy(m.TagConfigurationList, model.TagConfigurationList)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ContainerName = model.ContainerName

	m.IsLatest = model.IsLatest

	m.Tag = model.Tag

	m.LifecycleState = model.LifecycleState

	return
}

// ContainerSummaryTargetWorkloadsEnum Enum with underlying type: string
type ContainerSummaryTargetWorkloadsEnum string

// Set of constants representing the allowable values for ContainerSummaryTargetWorkloadsEnum
const (
	ContainerSummaryTargetWorkloadsModelDeployment ContainerSummaryTargetWorkloadsEnum = "MODEL_DEPLOYMENT"
	ContainerSummaryTargetWorkloadsJobRun          ContainerSummaryTargetWorkloadsEnum = "JOB_RUN"
)

var mappingContainerSummaryTargetWorkloadsEnum = map[string]ContainerSummaryTargetWorkloadsEnum{
	"MODEL_DEPLOYMENT": ContainerSummaryTargetWorkloadsModelDeployment,
	"JOB_RUN":          ContainerSummaryTargetWorkloadsJobRun,
}

var mappingContainerSummaryTargetWorkloadsEnumLowerCase = map[string]ContainerSummaryTargetWorkloadsEnum{
	"model_deployment": ContainerSummaryTargetWorkloadsModelDeployment,
	"job_run":          ContainerSummaryTargetWorkloadsJobRun,
}

// GetContainerSummaryTargetWorkloadsEnumValues Enumerates the set of values for ContainerSummaryTargetWorkloadsEnum
func GetContainerSummaryTargetWorkloadsEnumValues() []ContainerSummaryTargetWorkloadsEnum {
	values := make([]ContainerSummaryTargetWorkloadsEnum, 0)
	for _, v := range mappingContainerSummaryTargetWorkloadsEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerSummaryTargetWorkloadsEnumStringValues Enumerates the set of values in String for ContainerSummaryTargetWorkloadsEnum
func GetContainerSummaryTargetWorkloadsEnumStringValues() []string {
	return []string{
		"MODEL_DEPLOYMENT",
		"JOB_RUN",
	}
}

// GetMappingContainerSummaryTargetWorkloadsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerSummaryTargetWorkloadsEnum(val string) (ContainerSummaryTargetWorkloadsEnum, bool) {
	enum, ok := mappingContainerSummaryTargetWorkloadsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ContainerSummaryUsagesEnum Enum with underlying type: string
type ContainerSummaryUsagesEnum string

// Set of constants representing the allowable values for ContainerSummaryUsagesEnum
const (
	ContainerSummaryUsagesInference      ContainerSummaryUsagesEnum = "INFERENCE"
	ContainerSummaryUsagesFineTune       ContainerSummaryUsagesEnum = "FINE_TUNE"
	ContainerSummaryUsagesEvaluation     ContainerSummaryUsagesEnum = "EVALUATION"
	ContainerSummaryUsagesBatchInference ContainerSummaryUsagesEnum = "BATCH_INFERENCE"
	ContainerSummaryUsagesOther          ContainerSummaryUsagesEnum = "OTHER"
)

var mappingContainerSummaryUsagesEnum = map[string]ContainerSummaryUsagesEnum{
	"INFERENCE":       ContainerSummaryUsagesInference,
	"FINE_TUNE":       ContainerSummaryUsagesFineTune,
	"EVALUATION":      ContainerSummaryUsagesEvaluation,
	"BATCH_INFERENCE": ContainerSummaryUsagesBatchInference,
	"OTHER":           ContainerSummaryUsagesOther,
}

var mappingContainerSummaryUsagesEnumLowerCase = map[string]ContainerSummaryUsagesEnum{
	"inference":       ContainerSummaryUsagesInference,
	"fine_tune":       ContainerSummaryUsagesFineTune,
	"evaluation":      ContainerSummaryUsagesEvaluation,
	"batch_inference": ContainerSummaryUsagesBatchInference,
	"other":           ContainerSummaryUsagesOther,
}

// GetContainerSummaryUsagesEnumValues Enumerates the set of values for ContainerSummaryUsagesEnum
func GetContainerSummaryUsagesEnumValues() []ContainerSummaryUsagesEnum {
	values := make([]ContainerSummaryUsagesEnum, 0)
	for _, v := range mappingContainerSummaryUsagesEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerSummaryUsagesEnumStringValues Enumerates the set of values in String for ContainerSummaryUsagesEnum
func GetContainerSummaryUsagesEnumStringValues() []string {
	return []string{
		"INFERENCE",
		"FINE_TUNE",
		"EVALUATION",
		"BATCH_INFERENCE",
		"OTHER",
	}
}

// GetMappingContainerSummaryUsagesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerSummaryUsagesEnum(val string) (ContainerSummaryUsagesEnum, bool) {
	enum, ok := mappingContainerSummaryUsagesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
