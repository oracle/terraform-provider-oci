// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmTargetAsset Description of the VM target asset.
type VmTargetAsset struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the associated migration plan.
	MigrationPlanId *string `mandatory:"true" json:"migrationPlanId"`

	// A boolean indicating whether the asset should be migrated.
	IsExcludedFromExecution *bool `mandatory:"true" json:"isExcludedFromExecution"`

	EstimatedCost *CostEstimation `mandatory:"true" json:"estimatedCost"`

	// The time when the target asset was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the target asset was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time when the assessment was done. An RFC3339 formatted datetime string.
	TimeAssessed *common.SDKTime `mandatory:"true" json:"timeAssessed"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Compartment identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Created resource identifier
	CreatedResourceId *string `mandatory:"false" json:"createdResourceId"`

	// Messages about the compatibility issues.
	CompatibilityMessages []CompatibilityMessage `mandatory:"false" json:"compatibilityMessages"`

	MigrationAsset *MigrationAsset `mandatory:"false" json:"migrationAsset"`

	TestSpec *LaunchInstanceDetails `mandatory:"false" json:"testSpec"`

	// Performance of the block volumes.
	BlockVolumesPerformance *int `mandatory:"false" json:"blockVolumesPerformance"`

	// Microsoft license for VM configuration.
	MsLicense *string `mandatory:"false" json:"msLicense"`

	UserSpec *LaunchInstanceDetails `mandatory:"false" json:"userSpec"`

	RecommendedSpec *LaunchInstanceDetails `mandatory:"false" json:"recommendedSpec"`

	// Preferred VM shape type that you provide.
	PreferredShapeType VmTargetAssetPreferredShapeTypeEnum `mandatory:"true" json:"preferredShapeType"`

	// The current state of the target asset.
	LifecycleState TargetAssetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m VmTargetAsset) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m VmTargetAsset) GetDisplayName() *string {
	return m.DisplayName
}

// GetLifecycleState returns LifecycleState
func (m VmTargetAsset) GetLifecycleState() TargetAssetLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m VmTargetAsset) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetMigrationPlanId returns MigrationPlanId
func (m VmTargetAsset) GetMigrationPlanId() *string {
	return m.MigrationPlanId
}

// GetCompartmentId returns CompartmentId
func (m VmTargetAsset) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetCreatedResourceId returns CreatedResourceId
func (m VmTargetAsset) GetCreatedResourceId() *string {
	return m.CreatedResourceId
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m VmTargetAsset) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

// GetCompatibilityMessages returns CompatibilityMessages
func (m VmTargetAsset) GetCompatibilityMessages() []CompatibilityMessage {
	return m.CompatibilityMessages
}

// GetEstimatedCost returns EstimatedCost
func (m VmTargetAsset) GetEstimatedCost() *CostEstimation {
	return m.EstimatedCost
}

// GetTimeCreated returns TimeCreated
func (m VmTargetAsset) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m VmTargetAsset) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeAssessed returns TimeAssessed
func (m VmTargetAsset) GetTimeAssessed() *common.SDKTime {
	return m.TimeAssessed
}

// GetMigrationAsset returns MigrationAsset
func (m VmTargetAsset) GetMigrationAsset() *MigrationAsset {
	return m.MigrationAsset
}

func (m VmTargetAsset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmTargetAsset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmTargetAssetPreferredShapeTypeEnum(string(m.PreferredShapeType)); !ok && m.PreferredShapeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferredShapeType: %s. Supported values are: %s.", m.PreferredShapeType, strings.Join(GetVmTargetAssetPreferredShapeTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTargetAssetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetAssetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VmTargetAsset) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVmTargetAsset VmTargetAsset
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVmTargetAsset
	}{
		"INSTANCE",
		(MarshalTypeVmTargetAsset)(m),
	}

	return json.Marshal(&s)
}

// VmTargetAssetPreferredShapeTypeEnum Enum with underlying type: string
type VmTargetAssetPreferredShapeTypeEnum string

// Set of constants representing the allowable values for VmTargetAssetPreferredShapeTypeEnum
const (
	VmTargetAssetPreferredShapeTypeVm               VmTargetAssetPreferredShapeTypeEnum = "VM"
	VmTargetAssetPreferredShapeTypeVmIntel          VmTargetAssetPreferredShapeTypeEnum = "VM_INTEL"
	VmTargetAssetPreferredShapeTypeVmIntelStandard  VmTargetAssetPreferredShapeTypeEnum = "VM_INTEL_Standard"
	VmTargetAssetPreferredShapeTypeVmIntelDensio    VmTargetAssetPreferredShapeTypeEnum = "VM_INTEL_DensIO"
	VmTargetAssetPreferredShapeTypeVmIntelGpu       VmTargetAssetPreferredShapeTypeEnum = "VM_INTEL_GPU"
	VmTargetAssetPreferredShapeTypeVmIntelOptimized VmTargetAssetPreferredShapeTypeEnum = "VM_INTEL_Optimized"
	VmTargetAssetPreferredShapeTypeVmAmd            VmTargetAssetPreferredShapeTypeEnum = "VM_AMD"
	VmTargetAssetPreferredShapeTypeVmAmdStandard    VmTargetAssetPreferredShapeTypeEnum = "VM_AMD_Standard"
)

var mappingVmTargetAssetPreferredShapeTypeEnum = map[string]VmTargetAssetPreferredShapeTypeEnum{
	"VM":                 VmTargetAssetPreferredShapeTypeVm,
	"VM_INTEL":           VmTargetAssetPreferredShapeTypeVmIntel,
	"VM_INTEL_Standard":  VmTargetAssetPreferredShapeTypeVmIntelStandard,
	"VM_INTEL_DensIO":    VmTargetAssetPreferredShapeTypeVmIntelDensio,
	"VM_INTEL_GPU":       VmTargetAssetPreferredShapeTypeVmIntelGpu,
	"VM_INTEL_Optimized": VmTargetAssetPreferredShapeTypeVmIntelOptimized,
	"VM_AMD":             VmTargetAssetPreferredShapeTypeVmAmd,
	"VM_AMD_Standard":    VmTargetAssetPreferredShapeTypeVmAmdStandard,
}

var mappingVmTargetAssetPreferredShapeTypeEnumLowerCase = map[string]VmTargetAssetPreferredShapeTypeEnum{
	"vm":                 VmTargetAssetPreferredShapeTypeVm,
	"vm_intel":           VmTargetAssetPreferredShapeTypeVmIntel,
	"vm_intel_standard":  VmTargetAssetPreferredShapeTypeVmIntelStandard,
	"vm_intel_densio":    VmTargetAssetPreferredShapeTypeVmIntelDensio,
	"vm_intel_gpu":       VmTargetAssetPreferredShapeTypeVmIntelGpu,
	"vm_intel_optimized": VmTargetAssetPreferredShapeTypeVmIntelOptimized,
	"vm_amd":             VmTargetAssetPreferredShapeTypeVmAmd,
	"vm_amd_standard":    VmTargetAssetPreferredShapeTypeVmAmdStandard,
}

// GetVmTargetAssetPreferredShapeTypeEnumValues Enumerates the set of values for VmTargetAssetPreferredShapeTypeEnum
func GetVmTargetAssetPreferredShapeTypeEnumValues() []VmTargetAssetPreferredShapeTypeEnum {
	values := make([]VmTargetAssetPreferredShapeTypeEnum, 0)
	for _, v := range mappingVmTargetAssetPreferredShapeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmTargetAssetPreferredShapeTypeEnumStringValues Enumerates the set of values in String for VmTargetAssetPreferredShapeTypeEnum
func GetVmTargetAssetPreferredShapeTypeEnumStringValues() []string {
	return []string{
		"VM",
		"VM_INTEL",
		"VM_INTEL_Standard",
		"VM_INTEL_DensIO",
		"VM_INTEL_GPU",
		"VM_INTEL_Optimized",
		"VM_AMD",
		"VM_AMD_Standard",
	}
}

// GetMappingVmTargetAssetPreferredShapeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmTargetAssetPreferredShapeTypeEnum(val string) (VmTargetAssetPreferredShapeTypeEnum, bool) {
	enum, ok := mappingVmTargetAssetPreferredShapeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
