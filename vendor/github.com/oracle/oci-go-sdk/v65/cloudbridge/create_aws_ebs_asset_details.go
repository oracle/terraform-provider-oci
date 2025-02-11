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

// CreateAwsEbsAssetDetails Create AWS EBS type of asset.
type CreateAwsEbsAssetDetails struct {

	// Inventory ID to which an asset belongs.
	InventoryId *string `mandatory:"true" json:"inventoryId"`

	// The OCID of the compartment that the asset belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The source key to which the asset belongs.
	SourceKey *string `mandatory:"true" json:"sourceKey"`

	// The key of the asset from the external environment.
	ExternalAssetKey *string `mandatory:"true" json:"externalAssetKey"`

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
}

// GetDisplayName returns DisplayName
func (m CreateAwsEbsAssetDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetInventoryId returns InventoryId
func (m CreateAwsEbsAssetDetails) GetInventoryId() *string {
	return m.InventoryId
}

// GetCompartmentId returns CompartmentId
func (m CreateAwsEbsAssetDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceKey returns SourceKey
func (m CreateAwsEbsAssetDetails) GetSourceKey() *string {
	return m.SourceKey
}

// GetExternalAssetKey returns ExternalAssetKey
func (m CreateAwsEbsAssetDetails) GetExternalAssetKey() *string {
	return m.ExternalAssetKey
}

// GetAssetSourceIds returns AssetSourceIds
func (m CreateAwsEbsAssetDetails) GetAssetSourceIds() []string {
	return m.AssetSourceIds
}

// GetFreeformTags returns FreeformTags
func (m CreateAwsEbsAssetDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAwsEbsAssetDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateAwsEbsAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAwsEbsAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAwsEbsAssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAwsEbsAssetDetails CreateAwsEbsAssetDetails
	s := struct {
		DiscriminatorParam string `json:"assetType"`
		MarshalTypeCreateAwsEbsAssetDetails
	}{
		"AWS_EBS",
		(MarshalTypeCreateAwsEbsAssetDetails)(m),
	}

	return json.Marshal(&s)
}
