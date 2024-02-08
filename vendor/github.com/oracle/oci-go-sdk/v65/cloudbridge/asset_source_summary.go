// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssetSourceSummary Summary of an asset source provided in the list.
type AssetSourceSummary interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment for the resource.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the environment.
	GetEnvironmentId() *string

	// A user-friendly name for the asset source. Does not have to be unique, and it's mutable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The current state of the asset source.
	GetLifecycleState() AssetSourceLifecycleStateEnum

	// The detailed state of the asset source.
	GetLifecycleDetails() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the inventory that will contain created assets.
	GetInventoryId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
	GetAssetsCompartmentId() *string

	// The time when the asset source was created in RFC3339 format.
	GetTimeCreated() *common.SDKTime

	// The point in time that the asset source was last updated in RFC3339 format.
	GetTimeUpdated() *common.SDKTime

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	GetSystemTags() map[string]map[string]interface{}
}

type assetsourcesummary struct {
	JsonData            []byte
	TimeCreated         *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated         *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                  *string                           `mandatory:"true" json:"id"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	EnvironmentId       *string                           `mandatory:"true" json:"environmentId"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	LifecycleState      AssetSourceLifecycleStateEnum     `mandatory:"true" json:"lifecycleState"`
	LifecycleDetails    *string                           `mandatory:"true" json:"lifecycleDetails"`
	InventoryId         *string                           `mandatory:"true" json:"inventoryId"`
	AssetsCompartmentId *string                           `mandatory:"true" json:"assetsCompartmentId"`
	Type                string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *assetsourcesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerassetsourcesummary assetsourcesummary
	s := struct {
		Model Unmarshalerassetsourcesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.EnvironmentId = s.Model.EnvironmentId
	m.DisplayName = s.Model.DisplayName
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.InventoryId = s.Model.InventoryId
	m.AssetsCompartmentId = s.Model.AssetsCompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *assetsourcesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VMWARE":
		mm := VmWareAssetSourceSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AssetSourceSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetTimeCreated returns TimeCreated
func (m assetsourcesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m assetsourcesummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m assetsourcesummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m assetsourcesummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m assetsourcesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m assetsourcesummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m assetsourcesummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetEnvironmentId returns EnvironmentId
func (m assetsourcesummary) GetEnvironmentId() *string {
	return m.EnvironmentId
}

// GetDisplayName returns DisplayName
func (m assetsourcesummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetLifecycleState returns LifecycleState
func (m assetsourcesummary) GetLifecycleState() AssetSourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m assetsourcesummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetInventoryId returns InventoryId
func (m assetsourcesummary) GetInventoryId() *string {
	return m.InventoryId
}

// GetAssetsCompartmentId returns AssetsCompartmentId
func (m assetsourcesummary) GetAssetsCompartmentId() *string {
	return m.AssetsCompartmentId
}

func (m assetsourcesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m assetsourcesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssetSourceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
