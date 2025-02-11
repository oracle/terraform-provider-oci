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

// AwsEc2Asset AWS EC2 type of asset.
type AwsEc2Asset struct {

	// Inventory ID to which an asset belongs to.
	InventoryId *string `mandatory:"true" json:"inventoryId"`

	// Asset OCID that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment to which an asset belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The source key that the asset belongs to.
	SourceKey *string `mandatory:"true" json:"sourceKey"`

	// The key of the asset from the external environment.
	ExternalAssetKey *string `mandatory:"true" json:"externalAssetKey"`

	// The time when the asset was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the asset was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Compute *ComputeProperties `mandatory:"true" json:"compute"`

	Vm *VmProperties `mandatory:"true" json:"vm"`

	AwsEc2 *AwsEc2Properties `mandatory:"true" json:"awsEc2"`

	// Asset display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// List of asset source OCID.
	AssetSourceIds []string `mandatory:"false" json:"assetSourceIds"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	AwsEc2Cost *MonthlyCostSummary `mandatory:"false" json:"awsEc2Cost"`

	AttachedEbsVolumesCost *MonthlyCostSummary `mandatory:"false" json:"attachedEbsVolumesCost"`

	// The current state of the asset.
	LifecycleState AssetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetDisplayName returns DisplayName
func (m AwsEc2Asset) GetDisplayName() *string {
	return m.DisplayName
}

// GetInventoryId returns InventoryId
func (m AwsEc2Asset) GetInventoryId() *string {
	return m.InventoryId
}

// GetId returns Id
func (m AwsEc2Asset) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m AwsEc2Asset) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceKey returns SourceKey
func (m AwsEc2Asset) GetSourceKey() *string {
	return m.SourceKey
}

// GetExternalAssetKey returns ExternalAssetKey
func (m AwsEc2Asset) GetExternalAssetKey() *string {
	return m.ExternalAssetKey
}

// GetTimeCreated returns TimeCreated
func (m AwsEc2Asset) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m AwsEc2Asset) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetAssetSourceIds returns AssetSourceIds
func (m AwsEc2Asset) GetAssetSourceIds() []string {
	return m.AssetSourceIds
}

// GetLifecycleState returns LifecycleState
func (m AwsEc2Asset) GetLifecycleState() AssetLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m AwsEc2Asset) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m AwsEc2Asset) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m AwsEc2Asset) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m AwsEc2Asset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwsEc2Asset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAssetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AwsEc2Asset) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwsEc2Asset AwsEc2Asset
	s := struct {
		DiscriminatorParam string `json:"assetType"`
		MarshalTypeAwsEc2Asset
	}{
		"AWS_EC2",
		(MarshalTypeAwsEc2Asset)(m),
	}

	return json.Marshal(&s)
}
