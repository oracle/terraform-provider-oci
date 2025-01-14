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

// ZeroEtlPipelineSummary Summary of the ZeroETL pipeline.
type ZeroEtlPipelineSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the pipeline. This option applies when retrieving a pipeline.
	Id *string `mandatory:"true" json:"id"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SourceConnectionDetails *SourcePipelineConnectionDetails `mandatory:"true" json:"sourceConnectionDetails"`

	TargetConnectionDetails *TargetPipelineConnectionDetails `mandatory:"true" json:"targetConnectionDetails"`

	// The Minimum number of OCPUs to be made available for this Deployment.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	ProcessOptions *ProcessOptions `mandatory:"true" json:"processOptions"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// When the resource was last updated. This option applies when retrieving a pipeline. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-07-25T21:10:29.600Z`.
	TimeLastRecorded *common.SDKTime `mandatory:"false" json:"timeLastRecorded"`

	// Lifecycle state for the pipeline summary.
	LifecycleState PipelineLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"true" json:"licenseModel"`

	// Possible lifecycle substates when retrieving a pipeline.
	LifecycleSubState PipelineLifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`
}

// GetId returns Id
func (m ZeroEtlPipelineSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ZeroEtlPipelineSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m ZeroEtlPipelineSummary) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m ZeroEtlPipelineSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceConnectionDetails returns SourceConnectionDetails
func (m ZeroEtlPipelineSummary) GetSourceConnectionDetails() *SourcePipelineConnectionDetails {
	return m.SourceConnectionDetails
}

// GetTargetConnectionDetails returns TargetConnectionDetails
func (m ZeroEtlPipelineSummary) GetTargetConnectionDetails() *TargetPipelineConnectionDetails {
	return m.TargetConnectionDetails
}

// GetFreeformTags returns FreeformTags
func (m ZeroEtlPipelineSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLicenseModel returns LicenseModel
func (m ZeroEtlPipelineSummary) GetLicenseModel() LicenseModelEnum {
	return m.LicenseModel
}

// GetCpuCoreCount returns CpuCoreCount
func (m ZeroEtlPipelineSummary) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

// GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m ZeroEtlPipelineSummary) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

// GetDefinedTags returns DefinedTags
func (m ZeroEtlPipelineSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m ZeroEtlPipelineSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m ZeroEtlPipelineSummary) GetLocks() []ResourceLock {
	return m.Locks
}

// GetLifecycleState returns LifecycleState
func (m ZeroEtlPipelineSummary) GetLifecycleState() PipelineLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleSubState returns LifecycleSubState
func (m ZeroEtlPipelineSummary) GetLifecycleSubState() PipelineLifecycleSubStateEnum {
	return m.LifecycleSubState
}

// GetLifecycleDetails returns LifecycleDetails
func (m ZeroEtlPipelineSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m ZeroEtlPipelineSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ZeroEtlPipelineSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m ZeroEtlPipelineSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZeroEtlPipelineSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPipelineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPipelineLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPipelineLifecycleSubStateEnum(string(m.LifecycleSubState)); !ok && m.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", m.LifecycleSubState, strings.Join(GetPipelineLifecycleSubStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ZeroEtlPipelineSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeZeroEtlPipelineSummary ZeroEtlPipelineSummary
	s := struct {
		DiscriminatorParam string `json:"recipeType"`
		MarshalTypeZeroEtlPipelineSummary
	}{
		"ZERO_ETL",
		(MarshalTypeZeroEtlPipelineSummary)(m),
	}

	return json.Marshal(&s)
}
