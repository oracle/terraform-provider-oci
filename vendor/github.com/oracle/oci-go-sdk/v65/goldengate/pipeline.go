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

// Pipeline Represents the metadata details of a pipeline in the same compartment.
type Pipeline interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline. This option applies when retrieving a pipeline.
	GetId() *string

	// An object's Display Name.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	GetCompartmentId() *string

	// The Oracle license model that applies to a Deployment.
	GetLicenseModel() LicenseModelEnum

	// The Minimum number of OCPUs to be made available for this Deployment.
	GetCpuCoreCount() *int

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	GetIsAutoScalingEnabled() *bool

	GetSourceConnectionDetails() *SourcePipelineConnectionDetails

	GetTargetConnectionDetails() *TargetPipelineConnectionDetails

	// Lifecycle state of the pipeline.
	GetLifecycleState() PipelineLifecycleStateEnum

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	GetTimeCreated() *common.SDKTime

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	GetTimeUpdated() *common.SDKTime

	// Metadata about this specific object.
	GetDescription() *string

	GetPipelineDiagnosticData() *PipelineDiagnosticData

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

type pipeline struct {
	JsonData                []byte
	Description             *string                           `mandatory:"false" json:"description"`
	PipelineDiagnosticData  *PipelineDiagnosticData           `mandatory:"false" json:"pipelineDiagnosticData"`
	FreeformTags            map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags             map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags              map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Locks                   []ResourceLock                    `mandatory:"false" json:"locks"`
	LifecycleSubState       PipelineLifecycleSubStateEnum     `mandatory:"false" json:"lifecycleSubState,omitempty"`
	LifecycleDetails        *string                           `mandatory:"false" json:"lifecycleDetails"`
	Id                      *string                           `mandatory:"true" json:"id"`
	DisplayName             *string                           `mandatory:"true" json:"displayName"`
	CompartmentId           *string                           `mandatory:"true" json:"compartmentId"`
	LicenseModel            LicenseModelEnum                  `mandatory:"true" json:"licenseModel"`
	CpuCoreCount            *int                              `mandatory:"true" json:"cpuCoreCount"`
	IsAutoScalingEnabled    *bool                             `mandatory:"true" json:"isAutoScalingEnabled"`
	SourceConnectionDetails *SourcePipelineConnectionDetails  `mandatory:"true" json:"sourceConnectionDetails"`
	TargetConnectionDetails *TargetPipelineConnectionDetails  `mandatory:"true" json:"targetConnectionDetails"`
	LifecycleState          PipelineLifecycleStateEnum        `mandatory:"true" json:"lifecycleState"`
	TimeCreated             *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated             *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	RecipeType              string                            `json:"recipeType"`
}

// UnmarshalJSON unmarshals json
func (m *pipeline) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipeline pipeline
	s := struct {
		Model Unmarshalerpipeline
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.LicenseModel = s.Model.LicenseModel
	m.CpuCoreCount = s.Model.CpuCoreCount
	m.IsAutoScalingEnabled = s.Model.IsAutoScalingEnabled
	m.SourceConnectionDetails = s.Model.SourceConnectionDetails
	m.TargetConnectionDetails = s.Model.TargetConnectionDetails
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Description = s.Model.Description
	m.PipelineDiagnosticData = s.Model.PipelineDiagnosticData
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
func (m *pipeline) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RecipeType {
	case "ZERO_ETL":
		mm := ZeroEtlPipeline{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Pipeline: %s.", m.RecipeType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m pipeline) GetDescription() *string {
	return m.Description
}

// GetPipelineDiagnosticData returns PipelineDiagnosticData
func (m pipeline) GetPipelineDiagnosticData() *PipelineDiagnosticData {
	return m.PipelineDiagnosticData
}

// GetFreeformTags returns FreeformTags
func (m pipeline) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m pipeline) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m pipeline) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m pipeline) GetLocks() []ResourceLock {
	return m.Locks
}

// GetLifecycleSubState returns LifecycleSubState
func (m pipeline) GetLifecycleSubState() PipelineLifecycleSubStateEnum {
	return m.LifecycleSubState
}

// GetLifecycleDetails returns LifecycleDetails
func (m pipeline) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m pipeline) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m pipeline) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m pipeline) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLicenseModel returns LicenseModel
func (m pipeline) GetLicenseModel() LicenseModelEnum {
	return m.LicenseModel
}

// GetCpuCoreCount returns CpuCoreCount
func (m pipeline) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

// GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m pipeline) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

// GetSourceConnectionDetails returns SourceConnectionDetails
func (m pipeline) GetSourceConnectionDetails() *SourcePipelineConnectionDetails {
	return m.SourceConnectionDetails
}

// GetTargetConnectionDetails returns TargetConnectionDetails
func (m pipeline) GetTargetConnectionDetails() *TargetPipelineConnectionDetails {
	return m.TargetConnectionDetails
}

// GetLifecycleState returns LifecycleState
func (m pipeline) GetLifecycleState() PipelineLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m pipeline) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m pipeline) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m pipeline) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipeline) ValidateEnumValue() (bool, error) {
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

// PipelineLifecycleStateEnum Enum with underlying type: string
type PipelineLifecycleStateEnum string

// Set of constants representing the allowable values for PipelineLifecycleStateEnum
const (
	PipelineLifecycleStateCreating       PipelineLifecycleStateEnum = "CREATING"
	PipelineLifecycleStateUpdating       PipelineLifecycleStateEnum = "UPDATING"
	PipelineLifecycleStateActive         PipelineLifecycleStateEnum = "ACTIVE"
	PipelineLifecycleStateNeedsAttention PipelineLifecycleStateEnum = "NEEDS_ATTENTION"
	PipelineLifecycleStateDeleting       PipelineLifecycleStateEnum = "DELETING"
	PipelineLifecycleStateDeleted        PipelineLifecycleStateEnum = "DELETED"
	PipelineLifecycleStateFailed         PipelineLifecycleStateEnum = "FAILED"
)

var mappingPipelineLifecycleStateEnum = map[string]PipelineLifecycleStateEnum{
	"CREATING":        PipelineLifecycleStateCreating,
	"UPDATING":        PipelineLifecycleStateUpdating,
	"ACTIVE":          PipelineLifecycleStateActive,
	"NEEDS_ATTENTION": PipelineLifecycleStateNeedsAttention,
	"DELETING":        PipelineLifecycleStateDeleting,
	"DELETED":         PipelineLifecycleStateDeleted,
	"FAILED":          PipelineLifecycleStateFailed,
}

var mappingPipelineLifecycleStateEnumLowerCase = map[string]PipelineLifecycleStateEnum{
	"creating":        PipelineLifecycleStateCreating,
	"updating":        PipelineLifecycleStateUpdating,
	"active":          PipelineLifecycleStateActive,
	"needs_attention": PipelineLifecycleStateNeedsAttention,
	"deleting":        PipelineLifecycleStateDeleting,
	"deleted":         PipelineLifecycleStateDeleted,
	"failed":          PipelineLifecycleStateFailed,
}

// GetPipelineLifecycleStateEnumValues Enumerates the set of values for PipelineLifecycleStateEnum
func GetPipelineLifecycleStateEnumValues() []PipelineLifecycleStateEnum {
	values := make([]PipelineLifecycleStateEnum, 0)
	for _, v := range mappingPipelineLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineLifecycleStateEnumStringValues Enumerates the set of values in String for PipelineLifecycleStateEnum
func GetPipelineLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPipelineLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineLifecycleStateEnum(val string) (PipelineLifecycleStateEnum, bool) {
	enum, ok := mappingPipelineLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
