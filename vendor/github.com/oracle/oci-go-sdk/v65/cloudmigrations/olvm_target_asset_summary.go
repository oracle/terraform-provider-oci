// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// OlvmTargetAssetSummary Summary of the VM target asset.
type OlvmTargetAssetSummary struct {

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

	// Compartment identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Created resource identifier
	CreatedResourceId *string `mandatory:"false" json:"createdResourceId"`

	// Messages about compatibility issues.
	CompatibilityMessages []CompatibilityMessage `mandatory:"false" json:"compatibilityMessages"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	MigrationAsset *MigrationAssetSummary `mandatory:"false" json:"migrationAsset"`

	// Microsoft license for VM configuration.
	MsLicense *string `mandatory:"false" json:"msLicense"`

	// The current state of the target asset.
	LifecycleState TargetAssetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m OlvmTargetAssetSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m OlvmTargetAssetSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetLifecycleState returns LifecycleState
func (m OlvmTargetAssetSummary) GetLifecycleState() TargetAssetLifecycleStateEnum {
	return m.LifecycleState
}

// GetMigrationPlanId returns MigrationPlanId
func (m OlvmTargetAssetSummary) GetMigrationPlanId() *string {
	return m.MigrationPlanId
}

// GetCompartmentId returns CompartmentId
func (m OlvmTargetAssetSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetCreatedResourceId returns CreatedResourceId
func (m OlvmTargetAssetSummary) GetCreatedResourceId() *string {
	return m.CreatedResourceId
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m OlvmTargetAssetSummary) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

// GetCompatibilityMessages returns CompatibilityMessages
func (m OlvmTargetAssetSummary) GetCompatibilityMessages() []CompatibilityMessage {
	return m.CompatibilityMessages
}

// GetEstimatedCost returns EstimatedCost
func (m OlvmTargetAssetSummary) GetEstimatedCost() *CostEstimation {
	return m.EstimatedCost
}

// GetTimeCreated returns TimeCreated
func (m OlvmTargetAssetSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleDetails returns LifecycleDetails
func (m OlvmTargetAssetSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeUpdated returns TimeUpdated
func (m OlvmTargetAssetSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeAssessed returns TimeAssessed
func (m OlvmTargetAssetSummary) GetTimeAssessed() *common.SDKTime {
	return m.TimeAssessed
}

// GetMigrationAsset returns MigrationAsset
func (m OlvmTargetAssetSummary) GetMigrationAsset() *MigrationAssetSummary {
	return m.MigrationAsset
}

func (m OlvmTargetAssetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmTargetAssetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTargetAssetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetAssetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OlvmTargetAssetSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOlvmTargetAssetSummary OlvmTargetAssetSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOlvmTargetAssetSummary
	}{
		"OLVM_INSTANCE",
		(MarshalTypeOlvmTargetAssetSummary)(m),
	}

	return json.Marshal(&s)
}
