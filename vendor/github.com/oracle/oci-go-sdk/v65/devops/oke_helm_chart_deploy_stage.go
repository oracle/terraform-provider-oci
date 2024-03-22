// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OkeHelmChartDeployStage Specifies the OKE cluster deployment stage using helm charts.
type OkeHelmChartDeployStage struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Kubernetes cluster environment OCID for deployment.
	OkeClusterDeployEnvironmentId *string `mandatory:"true" json:"okeClusterDeployEnvironmentId"`

	// Helm chart artifact OCID.
	HelmChartDeployArtifactId *string `mandatory:"true" json:"helmChartDeployArtifactId"`

	// Release name of the Helm chart.
	ReleaseName *string `mandatory:"true" json:"releaseName"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// List of values.yaml file artifact OCIDs.
	ValuesArtifactIds []string `mandatory:"false" json:"valuesArtifactIds"`

	// Uninstall the Helm chart release on deleting the stage.
	IsUninstallOnStageDelete *bool `mandatory:"false" json:"isUninstallOnStageDelete"`

	// List of Helm command artifact OCIDs.
	HelmCommandArtifactIds []string `mandatory:"false" json:"helmCommandArtifactIds"`

	// Default namespace to be used for Kubernetes deployment when not specified in the manifest.
	Namespace *string `mandatory:"false" json:"namespace"`

	// Time to wait for execution of a helm stage. Defaults to 300 seconds.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	RollbackPolicy DeployStageRollbackPolicy `mandatory:"false" json:"rollbackPolicy"`

	SetValues *HelmSetValueCollection `mandatory:"false" json:"setValues"`

	SetString *HelmSetValueCollection `mandatory:"false" json:"setString"`

	// Disable pre/post upgrade hooks. Set to false by default.
	AreHooksEnabled *bool `mandatory:"false" json:"areHooksEnabled"`

	// During upgrade, reuse the values of the last release and merge overrides from the command line. Set to false by default.
	ShouldReuseValues *bool `mandatory:"false" json:"shouldReuseValues"`

	// During upgrade, reset the values to the ones built into the chart. It overrides shouldReuseValues. Set to false by default.
	ShouldResetValues *bool `mandatory:"false" json:"shouldResetValues"`

	// Force resource update through delete; or if required, recreate. Set to false by default.
	IsForceEnabled *bool `mandatory:"false" json:"isForceEnabled"`

	// Allow deletion of new resources created during when an upgrade fails. Set to false by default.
	ShouldCleanupOnFail *bool `mandatory:"false" json:"shouldCleanupOnFail"`

	// Limit the maximum number of revisions saved per release. Use 0 for no limit. Set to 10 by default
	MaxHistory *int `mandatory:"false" json:"maxHistory"`

	// If set, no CRDs are installed. By default, CRDs are installed only if they are not present already. Set to false by default.
	ShouldSkipCrds *bool `mandatory:"false" json:"shouldSkipCrds"`

	// If set, renders subchart notes along with the parent. Set to false by default.
	ShouldSkipRenderSubchartNotes *bool `mandatory:"false" json:"shouldSkipRenderSubchartNotes"`

	// Waits until all the resources are in a ready state to mark the release as successful. Set to false by default.
	ShouldNotWait *bool `mandatory:"false" json:"shouldNotWait"`

	// Enables helm --debug option to stream output to tf stdout. Set to false by default.
	IsDebugEnabled *bool `mandatory:"false" json:"isDebugEnabled"`

	// The purpose of running this Helm stage
	Purpose OkeHelmChartDeployStagePurposeEnum `mandatory:"false" json:"purpose,omitempty"`

	// The current state of the deployment stage.
	LifecycleState DeployStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m OkeHelmChartDeployStage) GetId() *string {
	return m.Id
}

// GetDescription returns Description
func (m OkeHelmChartDeployStage) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m OkeHelmChartDeployStage) GetDisplayName() *string {
	return m.DisplayName
}

// GetProjectId returns ProjectId
func (m OkeHelmChartDeployStage) GetProjectId() *string {
	return m.ProjectId
}

// GetDeployPipelineId returns DeployPipelineId
func (m OkeHelmChartDeployStage) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

// GetCompartmentId returns CompartmentId
func (m OkeHelmChartDeployStage) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m OkeHelmChartDeployStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OkeHelmChartDeployStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m OkeHelmChartDeployStage) GetLifecycleState() DeployStageLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m OkeHelmChartDeployStage) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m OkeHelmChartDeployStage) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m OkeHelmChartDeployStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OkeHelmChartDeployStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OkeHelmChartDeployStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m OkeHelmChartDeployStage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OkeHelmChartDeployStage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOkeHelmChartDeployStagePurposeEnum(string(m.Purpose)); !ok && m.Purpose != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Purpose: %s. Supported values are: %s.", m.Purpose, strings.Join(GetOkeHelmChartDeployStagePurposeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDeployStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeployStageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OkeHelmChartDeployStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOkeHelmChartDeployStage OkeHelmChartDeployStage
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeOkeHelmChartDeployStage
	}{
		"OKE_HELM_CHART_DEPLOYMENT",
		(MarshalTypeOkeHelmChartDeployStage)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OkeHelmChartDeployStage) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                            `json:"description"`
		DisplayName                      *string                            `json:"displayName"`
		TimeCreated                      *common.SDKTime                    `json:"timeCreated"`
		TimeUpdated                      *common.SDKTime                    `json:"timeUpdated"`
		LifecycleState                   DeployStageLifecycleStateEnum      `json:"lifecycleState"`
		LifecycleDetails                 *string                            `json:"lifecycleDetails"`
		DeployStagePredecessorCollection *DeployStagePredecessorCollection  `json:"deployStagePredecessorCollection"`
		FreeformTags                     map[string]string                  `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}  `json:"definedTags"`
		SystemTags                       map[string]map[string]interface{}  `json:"systemTags"`
		ValuesArtifactIds                []string                           `json:"valuesArtifactIds"`
		IsUninstallOnStageDelete         *bool                              `json:"isUninstallOnStageDelete"`
		HelmCommandArtifactIds           []string                           `json:"helmCommandArtifactIds"`
		Purpose                          OkeHelmChartDeployStagePurposeEnum `json:"purpose"`
		Namespace                        *string                            `json:"namespace"`
		TimeoutInSeconds                 *int                               `json:"timeoutInSeconds"`
		RollbackPolicy                   deploystagerollbackpolicy          `json:"rollbackPolicy"`
		SetValues                        *HelmSetValueCollection            `json:"setValues"`
		SetString                        *HelmSetValueCollection            `json:"setString"`
		AreHooksEnabled                  *bool                              `json:"areHooksEnabled"`
		ShouldReuseValues                *bool                              `json:"shouldReuseValues"`
		ShouldResetValues                *bool                              `json:"shouldResetValues"`
		IsForceEnabled                   *bool                              `json:"isForceEnabled"`
		ShouldCleanupOnFail              *bool                              `json:"shouldCleanupOnFail"`
		MaxHistory                       *int                               `json:"maxHistory"`
		ShouldSkipCrds                   *bool                              `json:"shouldSkipCrds"`
		ShouldSkipRenderSubchartNotes    *bool                              `json:"shouldSkipRenderSubchartNotes"`
		ShouldNotWait                    *bool                              `json:"shouldNotWait"`
		IsDebugEnabled                   *bool                              `json:"isDebugEnabled"`
		Id                               *string                            `json:"id"`
		ProjectId                        *string                            `json:"projectId"`
		DeployPipelineId                 *string                            `json:"deployPipelineId"`
		CompartmentId                    *string                            `json:"compartmentId"`
		OkeClusterDeployEnvironmentId    *string                            `json:"okeClusterDeployEnvironmentId"`
		HelmChartDeployArtifactId        *string                            `json:"helmChartDeployArtifactId"`
		ReleaseName                      *string                            `json:"releaseName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.ValuesArtifactIds = make([]string, len(model.ValuesArtifactIds))
	copy(m.ValuesArtifactIds, model.ValuesArtifactIds)
	m.IsUninstallOnStageDelete = model.IsUninstallOnStageDelete

	m.HelmCommandArtifactIds = make([]string, len(model.HelmCommandArtifactIds))
	copy(m.HelmCommandArtifactIds, model.HelmCommandArtifactIds)
	m.Purpose = model.Purpose

	m.Namespace = model.Namespace

	m.TimeoutInSeconds = model.TimeoutInSeconds

	nn, e = model.RollbackPolicy.UnmarshalPolymorphicJSON(model.RollbackPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RollbackPolicy = nn.(DeployStageRollbackPolicy)
	} else {
		m.RollbackPolicy = nil
	}

	m.SetValues = model.SetValues

	m.SetString = model.SetString

	m.AreHooksEnabled = model.AreHooksEnabled

	m.ShouldReuseValues = model.ShouldReuseValues

	m.ShouldResetValues = model.ShouldResetValues

	m.IsForceEnabled = model.IsForceEnabled

	m.ShouldCleanupOnFail = model.ShouldCleanupOnFail

	m.MaxHistory = model.MaxHistory

	m.ShouldSkipCrds = model.ShouldSkipCrds

	m.ShouldSkipRenderSubchartNotes = model.ShouldSkipRenderSubchartNotes

	m.ShouldNotWait = model.ShouldNotWait

	m.IsDebugEnabled = model.IsDebugEnabled

	m.Id = model.Id

	m.ProjectId = model.ProjectId

	m.DeployPipelineId = model.DeployPipelineId

	m.CompartmentId = model.CompartmentId

	m.OkeClusterDeployEnvironmentId = model.OkeClusterDeployEnvironmentId

	m.HelmChartDeployArtifactId = model.HelmChartDeployArtifactId

	m.ReleaseName = model.ReleaseName

	return
}

// OkeHelmChartDeployStagePurposeEnum Enum with underlying type: string
type OkeHelmChartDeployStagePurposeEnum string

// Set of constants representing the allowable values for OkeHelmChartDeployStagePurposeEnum
const (
	OkeHelmChartDeployStagePurposeUpgrade OkeHelmChartDeployStagePurposeEnum = "EXECUTE_HELM_UPGRADE"
	OkeHelmChartDeployStagePurposeCommand OkeHelmChartDeployStagePurposeEnum = "EXECUTE_HELM_COMMAND"
)

var mappingOkeHelmChartDeployStagePurposeEnum = map[string]OkeHelmChartDeployStagePurposeEnum{
	"EXECUTE_HELM_UPGRADE": OkeHelmChartDeployStagePurposeUpgrade,
	"EXECUTE_HELM_COMMAND": OkeHelmChartDeployStagePurposeCommand,
}

var mappingOkeHelmChartDeployStagePurposeEnumLowerCase = map[string]OkeHelmChartDeployStagePurposeEnum{
	"execute_helm_upgrade": OkeHelmChartDeployStagePurposeUpgrade,
	"execute_helm_command": OkeHelmChartDeployStagePurposeCommand,
}

// GetOkeHelmChartDeployStagePurposeEnumValues Enumerates the set of values for OkeHelmChartDeployStagePurposeEnum
func GetOkeHelmChartDeployStagePurposeEnumValues() []OkeHelmChartDeployStagePurposeEnum {
	values := make([]OkeHelmChartDeployStagePurposeEnum, 0)
	for _, v := range mappingOkeHelmChartDeployStagePurposeEnum {
		values = append(values, v)
	}
	return values
}

// GetOkeHelmChartDeployStagePurposeEnumStringValues Enumerates the set of values in String for OkeHelmChartDeployStagePurposeEnum
func GetOkeHelmChartDeployStagePurposeEnumStringValues() []string {
	return []string{
		"EXECUTE_HELM_UPGRADE",
		"EXECUTE_HELM_COMMAND",
	}
}

// GetMappingOkeHelmChartDeployStagePurposeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOkeHelmChartDeployStagePurposeEnum(val string) (OkeHelmChartDeployStagePurposeEnum, bool) {
	enum, ok := mappingOkeHelmChartDeployStagePurposeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
