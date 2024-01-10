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

// AssetSource Asset source.
type AssetSource interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment for the resource.
	GetCompartmentId() *string

	// A user-friendly name for the asset source. Does not have to be unique, and it's mutable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the environment.
	GetEnvironmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the inventory that will contain created assets.
	GetInventoryId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
	GetAssetsCompartmentId() *string

	// The current state of the asset source.
	GetLifecycleState() AssetSourceLifecycleStateEnum

	// The detailed state of the asset source.
	GetLifecycleDetails() *string

	// The time when the asset source was created in the RFC3339 format.
	GetTimeCreated() *common.SDKTime

	// The point in time that the asset source was last updated in the RFC3339 format.
	GetTimeUpdated() *common.SDKTime

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of an attached discovery schedule.
	GetDiscoveryScheduleId() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type assetsource struct {
	JsonData            []byte
	DiscoveryScheduleId *string                           `mandatory:"false" json:"discoveryScheduleId"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                  *string                           `mandatory:"true" json:"id"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	EnvironmentId       *string                           `mandatory:"true" json:"environmentId"`
	InventoryId         *string                           `mandatory:"true" json:"inventoryId"`
	AssetsCompartmentId *string                           `mandatory:"true" json:"assetsCompartmentId"`
	LifecycleState      AssetSourceLifecycleStateEnum     `mandatory:"true" json:"lifecycleState"`
	LifecycleDetails    *string                           `mandatory:"true" json:"lifecycleDetails"`
	TimeCreated         *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated         *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	Type                string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *assetsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerassetsource assetsource
	s := struct {
		Model Unmarshalerassetsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.EnvironmentId = s.Model.EnvironmentId
	m.InventoryId = s.Model.InventoryId
	m.AssetsCompartmentId = s.Model.AssetsCompartmentId
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.DiscoveryScheduleId = s.Model.DiscoveryScheduleId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *assetsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VMWARE":
		mm := VmWareAssetSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AssetSource: %s.", m.Type)
		return *m, nil
	}
}

// GetDiscoveryScheduleId returns DiscoveryScheduleId
func (m assetsource) GetDiscoveryScheduleId() *string {
	return m.DiscoveryScheduleId
}

// GetFreeformTags returns FreeformTags
func (m assetsource) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m assetsource) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m assetsource) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m assetsource) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m assetsource) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m assetsource) GetDisplayName() *string {
	return m.DisplayName
}

// GetEnvironmentId returns EnvironmentId
func (m assetsource) GetEnvironmentId() *string {
	return m.EnvironmentId
}

// GetInventoryId returns InventoryId
func (m assetsource) GetInventoryId() *string {
	return m.InventoryId
}

// GetAssetsCompartmentId returns AssetsCompartmentId
func (m assetsource) GetAssetsCompartmentId() *string {
	return m.AssetsCompartmentId
}

// GetLifecycleState returns LifecycleState
func (m assetsource) GetLifecycleState() AssetSourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m assetsource) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m assetsource) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m assetsource) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m assetsource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m assetsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssetSourceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
