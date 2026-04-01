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

// CreateAwsEc2AssetDetails Create AWS EC2 VM type of asset.
type CreateAwsEc2AssetDetails struct {

	// Inventory ID to which an asset belongs.
	InventoryId *string `mandatory:"true" json:"inventoryId"`

	// The OCID of the compartment that the asset belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The source key to which the asset belongs.
	SourceKey *string `mandatory:"true" json:"sourceKey"`

	// The key of the asset from the external environment.
	ExternalAssetKey *string `mandatory:"true" json:"externalAssetKey"`

	Compute *ComputeProperties `mandatory:"true" json:"compute"`

	Vm *VmProperties `mandatory:"true" json:"vm"`

	AwsEc2 *AwsEc2Properties `mandatory:"true" json:"awsEc2"`

	// Asset display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// List of asset source OCID.
	AssetSourceIds []string `mandatory:"false" json:"assetSourceIds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	AwsEc2Cost *MonthlyCostSummary `mandatory:"false" json:"awsEc2Cost"`

	AttachedEbsVolumesCost *MonthlyCostSummary `mandatory:"false" json:"attachedEbsVolumesCost"`
}

// GetDisplayName returns DisplayName
func (m CreateAwsEc2AssetDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetInventoryId returns InventoryId
func (m CreateAwsEc2AssetDetails) GetInventoryId() *string {
	return m.InventoryId
}

// GetCompartmentId returns CompartmentId
func (m CreateAwsEc2AssetDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceKey returns SourceKey
func (m CreateAwsEc2AssetDetails) GetSourceKey() *string {
	return m.SourceKey
}

// GetExternalAssetKey returns ExternalAssetKey
func (m CreateAwsEc2AssetDetails) GetExternalAssetKey() *string {
	return m.ExternalAssetKey
}

// GetAssetSourceIds returns AssetSourceIds
func (m CreateAwsEc2AssetDetails) GetAssetSourceIds() []string {
	return m.AssetSourceIds
}

// GetFreeformTags returns FreeformTags
func (m CreateAwsEc2AssetDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAwsEc2AssetDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateAwsEc2AssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAwsEc2AssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAwsEc2AssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAwsEc2AssetDetails CreateAwsEc2AssetDetails
	s := struct {
		DiscriminatorParam string `json:"assetType"`
		MarshalTypeCreateAwsEc2AssetDetails
	}{
		"AWS_EC2",
		(MarshalTypeCreateAwsEc2AssetDetails)(m),
	}

	return json.Marshal(&s)
}
