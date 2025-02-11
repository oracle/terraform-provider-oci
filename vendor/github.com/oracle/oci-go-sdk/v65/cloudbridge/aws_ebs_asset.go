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

// AwsEbsAsset AWS EBS type of asset.
type AwsEbsAsset struct {

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

	AwsEbs *AwsEbsProperties `mandatory:"true" json:"awsEbs"`

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

	// The current state of the asset.
	LifecycleState AssetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetDisplayName returns DisplayName
func (m AwsEbsAsset) GetDisplayName() *string {
	return m.DisplayName
}

// GetInventoryId returns InventoryId
func (m AwsEbsAsset) GetInventoryId() *string {
	return m.InventoryId
}

// GetId returns Id
func (m AwsEbsAsset) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m AwsEbsAsset) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceKey returns SourceKey
func (m AwsEbsAsset) GetSourceKey() *string {
	return m.SourceKey
}

// GetExternalAssetKey returns ExternalAssetKey
func (m AwsEbsAsset) GetExternalAssetKey() *string {
	return m.ExternalAssetKey
}

// GetTimeCreated returns TimeCreated
func (m AwsEbsAsset) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m AwsEbsAsset) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetAssetSourceIds returns AssetSourceIds
func (m AwsEbsAsset) GetAssetSourceIds() []string {
	return m.AssetSourceIds
}

// GetLifecycleState returns LifecycleState
func (m AwsEbsAsset) GetLifecycleState() AssetLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m AwsEbsAsset) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m AwsEbsAsset) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m AwsEbsAsset) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m AwsEbsAsset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwsEbsAsset) ValidateEnumValue() (bool, error) {
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
func (m AwsEbsAsset) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwsEbsAsset AwsEbsAsset
	s := struct {
		DiscriminatorParam string `json:"assetType"`
		MarshalTypeAwsEbsAsset
	}{
		"AWS_EBS",
		(MarshalTypeAwsEbsAsset)(m),
	}

	return json.Marshal(&s)
}
