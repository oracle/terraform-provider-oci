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

// Job The properties of a job.
// A job performs the actions that are defined in your Terraform configuration.
// For instructions on managing jobs, see
// Managing Jobs (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/jobs.htm).
// For more information about jobs, see
// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__jobdefinition).
type Job struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack that is associated with the job.
	StackId *string `mandatory:"false" json:"stackId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the job's associated stack resides.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The job's display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of job executing.
	Operation JobOperationEnum `mandatory:"false" json:"operation,omitempty"`

	// When `true`, the stack sources third-party Terraform providers from
	// Terraform Registry (https://registry.terraform.io/browse/providers) and allows
	// CustomTerraformProvider.
	// For more information about stack sourcing of third-party Terraform providers, see
	// Third-party Provider Configuration (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/terraformconfigresourcemanager.htm#third-party-providers).
	IsThirdPartyProviderExperienceEnabled *bool `mandatory:"false" json:"isThirdPartyProviderExperienceEnabled"`

	// Specifies whether or not to upgrade provider versions.
	// Within the version constraints of your Terraform configuration, use the latest versions available from the source of Terraform providers.
	// For more information about this option, see Dependency Lock File (terraform.io) (https://www.terraform.io/language/files/dependency-lock).
	IsProviderUpgradeRequired *bool `mandatory:"false" json:"isProviderUpgradeRequired"`

	JobOperationDetails JobOperationDetails `mandatory:"false" json:"jobOperationDetails"`

	ApplyJobPlanResolution *ApplyJobPlanResolution `mandatory:"false" json:"applyJobPlanResolution"`

	// Deprecated. Use the property `executionPlanJobId` in `jobOperationDetails` instead.
	// The plan job OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that was used (if this was an apply job and was not auto-approved).
	ResolvedPlanJobId *string `mandatory:"false" json:"resolvedPlanJobId"`

	// The date and time when the job was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the job stopped running, irrespective of whether the job ran successfully.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Current state of the specified job.
	// For more information about job lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__JobStates).
	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	FailureDetails *FailureDetails `mandatory:"false" json:"failureDetails"`

	CancellationDetails *CancellationDetails `mandatory:"false" json:"cancellationDetails"`

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// Required when using a zip Terraform configuration (`configSourceType` value of `ZIP_UPLOAD`) that contains folders.
	// Ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	// For more information about required and recommended file structure, see
	// File Structure (Terraform Configurations for Resource Manager) (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/terraformconfigresourcemanager.htm#filestructure).
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// Terraform variables associated with this resource.
	// Maximum number of variables supported is 250.
	// The maximum size of each variable, including both name and value, is 8192 bytes.
	// Example: `{"CompartmentId": "compartment-id-value"}`
	Variables map[string]string `mandatory:"false" json:"variables"`

	ConfigSource ConfigSourceRecord `mandatory:"false" json:"configSource"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Job) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Job) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetJobOperationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Job) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                                    *string                           `json:"id"`
		StackId                               *string                           `json:"stackId"`
		CompartmentId                         *string                           `json:"compartmentId"`
		DisplayName                           *string                           `json:"displayName"`
		Operation                             JobOperationEnum                  `json:"operation"`
		IsThirdPartyProviderExperienceEnabled *bool                             `json:"isThirdPartyProviderExperienceEnabled"`
		IsProviderUpgradeRequired             *bool                             `json:"isProviderUpgradeRequired"`
		JobOperationDetails                   joboperationdetails               `json:"jobOperationDetails"`
		ApplyJobPlanResolution                *ApplyJobPlanResolution           `json:"applyJobPlanResolution"`
		ResolvedPlanJobId                     *string                           `json:"resolvedPlanJobId"`
		TimeCreated                           *common.SDKTime                   `json:"timeCreated"`
		TimeFinished                          *common.SDKTime                   `json:"timeFinished"`
		LifecycleState                        JobLifecycleStateEnum             `json:"lifecycleState"`
		FailureDetails                        *FailureDetails                   `json:"failureDetails"`
		CancellationDetails                   *CancellationDetails              `json:"cancellationDetails"`
		WorkingDirectory                      *string                           `json:"workingDirectory"`
		Variables                             map[string]string                 `json:"variables"`
		ConfigSource                          configsourcerecord                `json:"configSource"`
		FreeformTags                          map[string]string                 `json:"freeformTags"`
		DefinedTags                           map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.StackId = model.StackId

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Operation = model.Operation

	m.IsThirdPartyProviderExperienceEnabled = model.IsThirdPartyProviderExperienceEnabled

	m.IsProviderUpgradeRequired = model.IsProviderUpgradeRequired

	nn, e = model.JobOperationDetails.UnmarshalPolymorphicJSON(model.JobOperationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobOperationDetails = nn.(JobOperationDetails)
	} else {
		m.JobOperationDetails = nil
	}

	m.ApplyJobPlanResolution = model.ApplyJobPlanResolution

	m.ResolvedPlanJobId = model.ResolvedPlanJobId

	m.TimeCreated = model.TimeCreated

	m.TimeFinished = model.TimeFinished

	m.LifecycleState = model.LifecycleState

	m.FailureDetails = model.FailureDetails

	m.CancellationDetails = model.CancellationDetails

	m.WorkingDirectory = model.WorkingDirectory

	m.Variables = model.Variables

	nn, e = model.ConfigSource.UnmarshalPolymorphicJSON(model.ConfigSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigSource = nn.(ConfigSourceRecord)
	} else {
		m.ConfigSource = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// JobOperationEnum Enum with underlying type: string
type JobOperationEnum string

// Set of constants representing the allowable values for JobOperationEnum
const (
	JobOperationPlan          JobOperationEnum = "PLAN"
	JobOperationApply         JobOperationEnum = "APPLY"
	JobOperationDestroy       JobOperationEnum = "DESTROY"
	JobOperationImportTfState JobOperationEnum = "IMPORT_TF_STATE"
	JobOperationPlanRollback  JobOperationEnum = "PLAN_ROLLBACK"
	JobOperationApplyRollback JobOperationEnum = "APPLY_ROLLBACK"
)

var mappingJobOperationEnum = map[string]JobOperationEnum{
	"PLAN":            JobOperationPlan,
	"APPLY":           JobOperationApply,
	"DESTROY":         JobOperationDestroy,
	"IMPORT_TF_STATE": JobOperationImportTfState,
	"PLAN_ROLLBACK":   JobOperationPlanRollback,
	"APPLY_ROLLBACK":  JobOperationApplyRollback,
}

var mappingJobOperationEnumLowerCase = map[string]JobOperationEnum{
	"plan":            JobOperationPlan,
	"apply":           JobOperationApply,
	"destroy":         JobOperationDestroy,
	"import_tf_state": JobOperationImportTfState,
	"plan_rollback":   JobOperationPlanRollback,
	"apply_rollback":  JobOperationApplyRollback,
}

// GetJobOperationEnumValues Enumerates the set of values for JobOperationEnum
func GetJobOperationEnumValues() []JobOperationEnum {
	values := make([]JobOperationEnum, 0)
	for _, v := range mappingJobOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetJobOperationEnumStringValues Enumerates the set of values in String for JobOperationEnum
func GetJobOperationEnumStringValues() []string {
	return []string{
		"PLAN",
		"APPLY",
		"DESTROY",
		"IMPORT_TF_STATE",
		"PLAN_ROLLBACK",
		"APPLY_ROLLBACK",
	}
}

// GetMappingJobOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobOperationEnum(val string) (JobOperationEnum, bool) {
	enum, ok := mappingJobOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// JobLifecycleStateEnum Enum with underlying type: string
type JobLifecycleStateEnum string

// Set of constants representing the allowable values for JobLifecycleStateEnum
const (
	JobLifecycleStateAccepted   JobLifecycleStateEnum = "ACCEPTED"
	JobLifecycleStateInProgress JobLifecycleStateEnum = "IN_PROGRESS"
	JobLifecycleStateFailed     JobLifecycleStateEnum = "FAILED"
	JobLifecycleStateSucceeded  JobLifecycleStateEnum = "SUCCEEDED"
	JobLifecycleStateCanceling  JobLifecycleStateEnum = "CANCELING"
	JobLifecycleStateCanceled   JobLifecycleStateEnum = "CANCELED"
)

var mappingJobLifecycleStateEnum = map[string]JobLifecycleStateEnum{
	"ACCEPTED":    JobLifecycleStateAccepted,
	"IN_PROGRESS": JobLifecycleStateInProgress,
	"FAILED":      JobLifecycleStateFailed,
	"SUCCEEDED":   JobLifecycleStateSucceeded,
	"CANCELING":   JobLifecycleStateCanceling,
	"CANCELED":    JobLifecycleStateCanceled,
}

var mappingJobLifecycleStateEnumLowerCase = map[string]JobLifecycleStateEnum{
	"accepted":    JobLifecycleStateAccepted,
	"in_progress": JobLifecycleStateInProgress,
	"failed":      JobLifecycleStateFailed,
	"succeeded":   JobLifecycleStateSucceeded,
	"canceling":   JobLifecycleStateCanceling,
	"canceled":    JobLifecycleStateCanceled,
}

// GetJobLifecycleStateEnumValues Enumerates the set of values for JobLifecycleStateEnum
func GetJobLifecycleStateEnumValues() []JobLifecycleStateEnum {
	values := make([]JobLifecycleStateEnum, 0)
	for _, v := range mappingJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetJobLifecycleStateEnumStringValues Enumerates the set of values in String for JobLifecycleStateEnum
func GetJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobLifecycleStateEnum(val string) (JobLifecycleStateEnum, bool) {
	enum, ok := mappingJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
