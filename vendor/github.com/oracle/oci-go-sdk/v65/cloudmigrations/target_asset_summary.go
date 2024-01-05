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

// TargetAssetSummary Summary of the target asset.
type TargetAssetSummary interface {

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

	// Compartment identifier
	GetCompartmentId() *string

	// Created resource identifier
	GetCreatedResourceId() *string

	// Messages about compatibility issues.
	GetCompatibilityMessages() []CompatibilityMessage

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	GetMigrationAsset() *MigrationAssetSummary
}

type targetassetsummary struct {
	JsonData                []byte
	DisplayName             *string                       `mandatory:"false" json:"displayName"`
	CompartmentId           *string                       `mandatory:"false" json:"compartmentId"`
	CreatedResourceId       *string                       `mandatory:"false" json:"createdResourceId"`
	CompatibilityMessages   []CompatibilityMessage        `mandatory:"false" json:"compatibilityMessages"`
	LifecycleDetails        *string                       `mandatory:"false" json:"lifecycleDetails"`
	MigrationAsset          *MigrationAssetSummary        `mandatory:"false" json:"migrationAsset"`
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
func (m *targetassetsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetassetsummary targetassetsummary
	s := struct {
		Model Unmarshalertargetassetsummary
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
	m.CompartmentId = s.Model.CompartmentId
	m.CreatedResourceId = s.Model.CreatedResourceId
	m.CompatibilityMessages = s.Model.CompatibilityMessages
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.MigrationAsset = s.Model.MigrationAsset
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetassetsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "INSTANCE":
		mm := VmTargetAssetSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TargetAssetSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m targetassetsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m targetassetsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetCreatedResourceId returns CreatedResourceId
func (m targetassetsummary) GetCreatedResourceId() *string {
	return m.CreatedResourceId
}

// GetCompatibilityMessages returns CompatibilityMessages
func (m targetassetsummary) GetCompatibilityMessages() []CompatibilityMessage {
	return m.CompatibilityMessages
}

// GetLifecycleDetails returns LifecycleDetails
func (m targetassetsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetMigrationAsset returns MigrationAsset
func (m targetassetsummary) GetMigrationAsset() *MigrationAssetSummary {
	return m.MigrationAsset
}

// GetId returns Id
func (m targetassetsummary) GetId() *string {
	return m.Id
}

// GetLifecycleState returns LifecycleState
func (m targetassetsummary) GetLifecycleState() TargetAssetLifecycleStateEnum {
	return m.LifecycleState
}

// GetMigrationPlanId returns MigrationPlanId
func (m targetassetsummary) GetMigrationPlanId() *string {
	return m.MigrationPlanId
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m targetassetsummary) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

// GetEstimatedCost returns EstimatedCost
func (m targetassetsummary) GetEstimatedCost() *CostEstimation {
	return m.EstimatedCost
}

// GetTimeCreated returns TimeCreated
func (m targetassetsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m targetassetsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeAssessed returns TimeAssessed
func (m targetassetsummary) GetTimeAssessed() *common.SDKTime {
	return m.TimeAssessed
}

func (m targetassetsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targetassetsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetAssetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetAssetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetAssetSummaryTypeEnum Enum with underlying type: string
type TargetAssetSummaryTypeEnum string

// Set of constants representing the allowable values for TargetAssetSummaryTypeEnum
const (
	TargetAssetSummaryTypeInstance TargetAssetSummaryTypeEnum = "INSTANCE"
)

var mappingTargetAssetSummaryTypeEnum = map[string]TargetAssetSummaryTypeEnum{
	"INSTANCE": TargetAssetSummaryTypeInstance,
}

var mappingTargetAssetSummaryTypeEnumLowerCase = map[string]TargetAssetSummaryTypeEnum{
	"instance": TargetAssetSummaryTypeInstance,
}

// GetTargetAssetSummaryTypeEnumValues Enumerates the set of values for TargetAssetSummaryTypeEnum
func GetTargetAssetSummaryTypeEnumValues() []TargetAssetSummaryTypeEnum {
	values := make([]TargetAssetSummaryTypeEnum, 0)
	for _, v := range mappingTargetAssetSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetAssetSummaryTypeEnumStringValues Enumerates the set of values in String for TargetAssetSummaryTypeEnum
func GetTargetAssetSummaryTypeEnumStringValues() []string {
	return []string{
		"INSTANCE",
	}
}

// GetMappingTargetAssetSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetAssetSummaryTypeEnum(val string) (TargetAssetSummaryTypeEnum, bool) {
	enum, ok := mappingTargetAssetSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
