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

// TargetAsset Description of the target asset.
type TargetAsset interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The current state of the target asset.
	GetLifecycleState() TargetAssetLifecycleStateEnum

	// OCID of the associated migration plan.
	GetMigrationPlanId() *string

	// A boolean indicating whether the asset should be migrated.
	GetIsExcludedFromExecution() *bool

	GetEstimatedCost() *CostEstimation

	// The time when the target asset was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time when the target asset was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// The time when the assessment was done. An RFC3339 formatted datetime string.
	GetTimeAssessed() *common.SDKTime

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Compartment identifier
	GetCompartmentId() *string

	// Created resource identifier
	GetCreatedResourceId() *string

	// Messages about the compatibility issues.
	GetCompatibilityMessages() []CompatibilityMessage

	GetMigrationAsset() *MigrationAsset
}

type targetasset struct {
	JsonData                []byte
	DisplayName             *string                       `mandatory:"false" json:"displayName"`
	LifecycleDetails        *string                       `mandatory:"false" json:"lifecycleDetails"`
	CompartmentId           *string                       `mandatory:"false" json:"compartmentId"`
	CreatedResourceId       *string                       `mandatory:"false" json:"createdResourceId"`
	CompatibilityMessages   []CompatibilityMessage        `mandatory:"false" json:"compatibilityMessages"`
	MigrationAsset          *MigrationAsset               `mandatory:"false" json:"migrationAsset"`
	Id                      *string                       `mandatory:"true" json:"id"`
	LifecycleState          TargetAssetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	MigrationPlanId         *string                       `mandatory:"true" json:"migrationPlanId"`
	IsExcludedFromExecution *bool                         `mandatory:"true" json:"isExcludedFromExecution"`
	EstimatedCost           *CostEstimation               `mandatory:"true" json:"estimatedCost"`
	TimeCreated             *common.SDKTime               `mandatory:"true" json:"timeCreated"`
	TimeUpdated             *common.SDKTime               `mandatory:"true" json:"timeUpdated"`
	TimeAssessed            *common.SDKTime               `mandatory:"true" json:"timeAssessed"`
	Type                    string                        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *targetasset) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetasset targetasset
	s := struct {
		Model Unmarshalertargetasset
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.LifecycleState = s.Model.LifecycleState
	m.MigrationPlanId = s.Model.MigrationPlanId
	m.IsExcludedFromExecution = s.Model.IsExcludedFromExecution
	m.EstimatedCost = s.Model.EstimatedCost
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.TimeAssessed = s.Model.TimeAssessed
	m.DisplayName = s.Model.DisplayName
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.CompartmentId = s.Model.CompartmentId
	m.CreatedResourceId = s.Model.CreatedResourceId
	m.CompatibilityMessages = s.Model.CompatibilityMessages
	m.MigrationAsset = s.Model.MigrationAsset
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetasset) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "INSTANCE":
		mm := VmTargetAsset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TargetAsset: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m targetasset) GetDisplayName() *string {
	return m.DisplayName
}

// GetLifecycleDetails returns LifecycleDetails
func (m targetasset) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetCompartmentId returns CompartmentId
func (m targetasset) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetCreatedResourceId returns CreatedResourceId
func (m targetasset) GetCreatedResourceId() *string {
	return m.CreatedResourceId
}

// GetCompatibilityMessages returns CompatibilityMessages
func (m targetasset) GetCompatibilityMessages() []CompatibilityMessage {
	return m.CompatibilityMessages
}

// GetMigrationAsset returns MigrationAsset
func (m targetasset) GetMigrationAsset() *MigrationAsset {
	return m.MigrationAsset
}

// GetId returns Id
func (m targetasset) GetId() *string {
	return m.Id
}

// GetLifecycleState returns LifecycleState
func (m targetasset) GetLifecycleState() TargetAssetLifecycleStateEnum {
	return m.LifecycleState
}

// GetMigrationPlanId returns MigrationPlanId
func (m targetasset) GetMigrationPlanId() *string {
	return m.MigrationPlanId
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m targetasset) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

// GetEstimatedCost returns EstimatedCost
func (m targetasset) GetEstimatedCost() *CostEstimation {
	return m.EstimatedCost
}

// GetTimeCreated returns TimeCreated
func (m targetasset) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m targetasset) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeAssessed returns TimeAssessed
func (m targetasset) GetTimeAssessed() *common.SDKTime {
	return m.TimeAssessed
}

func (m targetasset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targetasset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetAssetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetAssetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetAssetLifecycleStateEnum Enum with underlying type: string
type TargetAssetLifecycleStateEnum string

// Set of constants representing the allowable values for TargetAssetLifecycleStateEnum
const (
	TargetAssetLifecycleStateCreating       TargetAssetLifecycleStateEnum = "CREATING"
	TargetAssetLifecycleStateUpdating       TargetAssetLifecycleStateEnum = "UPDATING"
	TargetAssetLifecycleStateNeedsAttention TargetAssetLifecycleStateEnum = "NEEDS_ATTENTION"
	TargetAssetLifecycleStateActive         TargetAssetLifecycleStateEnum = "ACTIVE"
	TargetAssetLifecycleStateDeleting       TargetAssetLifecycleStateEnum = "DELETING"
	TargetAssetLifecycleStateDeleted        TargetAssetLifecycleStateEnum = "DELETED"
	TargetAssetLifecycleStateFailed         TargetAssetLifecycleStateEnum = "FAILED"
)

var mappingTargetAssetLifecycleStateEnum = map[string]TargetAssetLifecycleStateEnum{
	"CREATING":        TargetAssetLifecycleStateCreating,
	"UPDATING":        TargetAssetLifecycleStateUpdating,
	"NEEDS_ATTENTION": TargetAssetLifecycleStateNeedsAttention,
	"ACTIVE":          TargetAssetLifecycleStateActive,
	"DELETING":        TargetAssetLifecycleStateDeleting,
	"DELETED":         TargetAssetLifecycleStateDeleted,
	"FAILED":          TargetAssetLifecycleStateFailed,
}

var mappingTargetAssetLifecycleStateEnumLowerCase = map[string]TargetAssetLifecycleStateEnum{
	"creating":        TargetAssetLifecycleStateCreating,
	"updating":        TargetAssetLifecycleStateUpdating,
	"needs_attention": TargetAssetLifecycleStateNeedsAttention,
	"active":          TargetAssetLifecycleStateActive,
	"deleting":        TargetAssetLifecycleStateDeleting,
	"deleted":         TargetAssetLifecycleStateDeleted,
	"failed":          TargetAssetLifecycleStateFailed,
}

// GetTargetAssetLifecycleStateEnumValues Enumerates the set of values for TargetAssetLifecycleStateEnum
func GetTargetAssetLifecycleStateEnumValues() []TargetAssetLifecycleStateEnum {
	values := make([]TargetAssetLifecycleStateEnum, 0)
	for _, v := range mappingTargetAssetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetAssetLifecycleStateEnumStringValues Enumerates the set of values in String for TargetAssetLifecycleStateEnum
func GetTargetAssetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"NEEDS_ATTENTION",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingTargetAssetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetAssetLifecycleStateEnum(val string) (TargetAssetLifecycleStateEnum, bool) {
	enum, ok := mappingTargetAssetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TargetAssetTypeEnum Enum with underlying type: string
type TargetAssetTypeEnum string

// Set of constants representing the allowable values for TargetAssetTypeEnum
const (
	TargetAssetTypeInstance TargetAssetTypeEnum = "INSTANCE"
)

var mappingTargetAssetTypeEnum = map[string]TargetAssetTypeEnum{
	"INSTANCE": TargetAssetTypeInstance,
}

var mappingTargetAssetTypeEnumLowerCase = map[string]TargetAssetTypeEnum{
	"instance": TargetAssetTypeInstance,
}

// GetTargetAssetTypeEnumValues Enumerates the set of values for TargetAssetTypeEnum
func GetTargetAssetTypeEnumValues() []TargetAssetTypeEnum {
	values := make([]TargetAssetTypeEnum, 0)
	for _, v := range mappingTargetAssetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetAssetTypeEnumStringValues Enumerates the set of values in String for TargetAssetTypeEnum
func GetTargetAssetTypeEnumStringValues() []string {
	return []string{
		"INSTANCE",
	}
}

// GetMappingTargetAssetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetAssetTypeEnum(val string) (TargetAssetTypeEnum, bool) {
	enum, ok := mappingTargetAssetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
