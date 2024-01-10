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

// ImportInventoryViaAssetsDetails Details for importing assets from a file.
type ImportInventoryViaAssetsDetails struct {

	// The OCID of the compartmentId that resources import.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The file body to be sent in the request.
	Data []byte `mandatory:"false" json:"data"`

	// The type of asset.
	AssetType AssetTypeEnum `mandatory:"false" json:"assetType,omitempty"`
}

// GetCompartmentId returns CompartmentId
func (m ImportInventoryViaAssetsDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m ImportInventoryViaAssetsDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m ImportInventoryViaAssetsDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m ImportInventoryViaAssetsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportInventoryViaAssetsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAssetTypeEnum(string(m.AssetType)); !ok && m.AssetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssetType: %s. Supported values are: %s.", m.AssetType, strings.Join(GetAssetTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ImportInventoryViaAssetsDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImportInventoryViaAssetsDetails ImportInventoryViaAssetsDetails
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeImportInventoryViaAssetsDetails
	}{
		"ASSET",
		(MarshalTypeImportInventoryViaAssetsDetails)(m),
	}

	return json.Marshal(&s)
}
