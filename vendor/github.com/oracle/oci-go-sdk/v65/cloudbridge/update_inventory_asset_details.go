// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateInventoryAssetDetails The information of a JSON formatted asset to be updated.
type UpdateInventoryAssetDetails struct {

	// The class name of the asset.
	AssetClassName *string `mandatory:"true" json:"assetClassName"`

	// The version of the asset class.
	AssetClassVersion *string `mandatory:"true" json:"assetClassVersion"`

	// The details of the asset.
	AssetDetails map[string]interface{} `mandatory:"true" json:"assetDetails"`

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

//GetDisplayName returns DisplayName
func (m UpdateInventoryAssetDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetAssetSourceIds returns AssetSourceIds
func (m UpdateInventoryAssetDetails) GetAssetSourceIds() []string {
	return m.AssetSourceIds
}

//GetFreeformTags returns FreeformTags
func (m UpdateInventoryAssetDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateInventoryAssetDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateInventoryAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInventoryAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateInventoryAssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateInventoryAssetDetails UpdateInventoryAssetDetails
	s := struct {
		DiscriminatorParam string `json:"assetType"`
		MarshalTypeUpdateInventoryAssetDetails
	}{
		"INVENTORY_ASSET",
		(MarshalTypeUpdateInventoryAssetDetails)(m),
	}

	return json.Marshal(&s)
}
