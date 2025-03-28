// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PipelineSummary Summary details of the pipeline.
type PipelineSummary interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline. This option applies when retrieving a pipeline.
	GetId() *string

	// An object's Display Name.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	GetCompartmentId() *string

	GetSourceConnectionDetails() *SourcePipelineConnectionDetails

	GetTargetConnectionDetails() *TargetPipelineConnectionDetails

	// The Oracle license model that applies to a Deployment.
	GetLicenseModel() LicenseModelEnum

	// The Minimum number of OCPUs to be made available for this Deployment.
	GetCpuCoreCount() *int

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	GetIsAutoScalingEnabled() *bool

	// Lifecycle state for the pipeline summary.
	GetLifecycleState() PipelineLifecycleStateEnum

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	GetTimeCreated() *common.SDKTime

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	GetTimeUpdated() *common.SDKTime

	// Metadata about this specific object.
	GetDescription() *string

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	GetSystemTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []ResourceLock

	// Possible lifecycle substates when retrieving a pipeline.
	GetLifecycleSubState() PipelineLifecycleSubStateEnum

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	GetLifecycleDetails() *string
}

type pipelinesummary struct {
	JsonData                []byte
	Description             *string                           `mandatory:"false" json:"description"`
	FreeformTags            map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags             map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags              map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Locks                   []ResourceLock                    `mandatory:"false" json:"locks"`
	LifecycleSubState       PipelineLifecycleSubStateEnum     `mandatory:"false" json:"lifecycleSubState,omitempty"`
	LifecycleDetails        *string                           `mandatory:"false" json:"lifecycleDetails"`
	Id                      *string                           `mandatory:"true" json:"id"`
	DisplayName             *string                           `mandatory:"true" json:"displayName"`
	CompartmentId           *string                           `mandatory:"true" json:"compartmentId"`
	SourceConnectionDetails *SourcePipelineConnectionDetails  `mandatory:"true" json:"sourceConnectionDetails"`
	TargetConnectionDetails *TargetPipelineConnectionDetails  `mandatory:"true" json:"targetConnectionDetails"`
	LicenseModel            LicenseModelEnum                  `mandatory:"true" json:"licenseModel"`
	CpuCoreCount            *int                              `mandatory:"true" json:"cpuCoreCount"`
	IsAutoScalingEnabled    *bool                             `mandatory:"true" json:"isAutoScalingEnabled"`
	LifecycleState          PipelineLifecycleStateEnum        `mandatory:"true" json:"lifecycleState"`
	TimeCreated             *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated             *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	RecipeType              string                            `json:"recipeType"`
}

// UnmarshalJSON unmarshals json
func (m *pipelinesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipelinesummary pipelinesummary
	s := struct {
		Model Unmarshalerpipelinesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.SourceConnectionDetails = s.Model.SourceConnectionDetails
	m.TargetConnectionDetails = s.Model.TargetConnectionDetails
	m.LicenseModel = s.Model.LicenseModel
	m.CpuCoreCount = s.Model.CpuCoreCount
	m.IsAutoScalingEnabled = s.Model.IsAutoScalingEnabled
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Locks = s.Model.Locks
	m.LifecycleSubState = s.Model.LifecycleSubState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.RecipeType = s.Model.RecipeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pipelinesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RecipeType {
	case "ZERO_ETL":
		mm := ZeroEtlPipelineSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PipelineSummary: %s.", m.RecipeType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m pipelinesummary) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m pipelinesummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m pipelinesummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m pipelinesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m pipelinesummary) GetLocks() []ResourceLock {
	return m.Locks
}

// GetLifecycleSubState returns LifecycleSubState
func (m pipelinesummary) GetLifecycleSubState() PipelineLifecycleSubStateEnum {
	return m.LifecycleSubState
}

// GetLifecycleDetails returns LifecycleDetails
func (m pipelinesummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m pipelinesummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m pipelinesummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m pipelinesummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceConnectionDetails returns SourceConnectionDetails
func (m pipelinesummary) GetSourceConnectionDetails() *SourcePipelineConnectionDetails {
	return m.SourceConnectionDetails
}

// GetTargetConnectionDetails returns TargetConnectionDetails
func (m pipelinesummary) GetTargetConnectionDetails() *TargetPipelineConnectionDetails {
	return m.TargetConnectionDetails
}

// GetLicenseModel returns LicenseModel
func (m pipelinesummary) GetLicenseModel() LicenseModelEnum {
	return m.LicenseModel
}

// GetCpuCoreCount returns CpuCoreCount
func (m pipelinesummary) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

// GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m pipelinesummary) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

// GetLifecycleState returns LifecycleState
func (m pipelinesummary) GetLifecycleState() PipelineLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m pipelinesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m pipelinesummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m pipelinesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipelinesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPipelineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPipelineLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPipelineLifecycleSubStateEnum(string(m.LifecycleSubState)); !ok && m.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", m.LifecycleSubState, strings.Join(GetPipelineLifecycleSubStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
