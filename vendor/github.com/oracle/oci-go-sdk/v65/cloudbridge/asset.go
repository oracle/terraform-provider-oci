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

// Asset Description of an asset.
type Asset interface {

	// Inventory ID to which an asset belongs to.
	GetInventoryId() *string

	// Asset OCID that is immutable on creation.
	GetId() *string

	// The OCID of the compartment to which an asset belongs to.
	GetCompartmentId() *string

	// The source key that the asset belongs to.
	GetSourceKey() *string

	// The key of the asset from the external environment.
	GetExternalAssetKey() *string

	// The time when the asset was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time when the asset was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// The current state of the asset.
	GetLifecycleState() AssetLifecycleStateEnum

	// Asset display name.
	GetDisplayName() *string

	// List of asset source OCID.
	GetAssetSourceIds() []string

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

type asset struct {
	JsonData         []byte
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	AssetSourceIds   []string                          `mandatory:"false" json:"assetSourceIds"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	InventoryId      *string                           `mandatory:"true" json:"inventoryId"`
	Id               *string                           `mandatory:"true" json:"id"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	SourceKey        *string                           `mandatory:"true" json:"sourceKey"`
	ExternalAssetKey *string                           `mandatory:"true" json:"externalAssetKey"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	LifecycleState   AssetLifecycleStateEnum           `mandatory:"true" json:"lifecycleState"`
	AssetType        string                            `json:"assetType"`
}

// UnmarshalJSON unmarshals json
func (m *asset) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerasset asset
	s := struct {
		Model Unmarshalerasset
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InventoryId = s.Model.InventoryId
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.SourceKey = s.Model.SourceKey
	m.ExternalAssetKey = s.Model.ExternalAssetKey
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.AssetSourceIds = s.Model.AssetSourceIds
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.AssetType = s.Model.AssetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *asset) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AssetType {
	case "VMWARE_VM":
		mm := VmwareVmAsset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VM":
		mm := VmAsset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Asset: %s.", m.AssetType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m asset) GetDisplayName() *string {
	return m.DisplayName
}

// GetAssetSourceIds returns AssetSourceIds
func (m asset) GetAssetSourceIds() []string {
	return m.AssetSourceIds
}

// GetFreeformTags returns FreeformTags
func (m asset) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m asset) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m asset) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetInventoryId returns InventoryId
func (m asset) GetInventoryId() *string {
	return m.InventoryId
}

// GetId returns Id
func (m asset) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m asset) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceKey returns SourceKey
func (m asset) GetSourceKey() *string {
	return m.SourceKey
}

// GetExternalAssetKey returns ExternalAssetKey
func (m asset) GetExternalAssetKey() *string {
	return m.ExternalAssetKey
}

// GetTimeCreated returns TimeCreated
func (m asset) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m asset) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m asset) GetLifecycleState() AssetLifecycleStateEnum {
	return m.LifecycleState
}

func (m asset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m asset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssetLifecycleStateEnum Enum with underlying type: string
type AssetLifecycleStateEnum string

// Set of constants representing the allowable values for AssetLifecycleStateEnum
const (
	AssetLifecycleStateActive  AssetLifecycleStateEnum = "ACTIVE"
	AssetLifecycleStateDeleted AssetLifecycleStateEnum = "DELETED"
)

var mappingAssetLifecycleStateEnum = map[string]AssetLifecycleStateEnum{
	"ACTIVE":  AssetLifecycleStateActive,
	"DELETED": AssetLifecycleStateDeleted,
}

var mappingAssetLifecycleStateEnumLowerCase = map[string]AssetLifecycleStateEnum{
	"active":  AssetLifecycleStateActive,
	"deleted": AssetLifecycleStateDeleted,
}

// GetAssetLifecycleStateEnumValues Enumerates the set of values for AssetLifecycleStateEnum
func GetAssetLifecycleStateEnumValues() []AssetLifecycleStateEnum {
	values := make([]AssetLifecycleStateEnum, 0)
	for _, v := range mappingAssetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetLifecycleStateEnumStringValues Enumerates the set of values in String for AssetLifecycleStateEnum
func GetAssetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingAssetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetLifecycleStateEnum(val string) (AssetLifecycleStateEnum, bool) {
	enum, ok := mappingAssetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
