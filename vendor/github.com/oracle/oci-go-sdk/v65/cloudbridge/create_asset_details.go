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

// CreateAssetDetails The information about the new asset.
type CreateAssetDetails interface {

	// Inventory ID to which an asset belongs.
	GetInventoryId() *string

	// The OCID of the compartment that the asset belongs to.
	GetCompartmentId() *string

	// The source key to which the asset belongs.
	GetSourceKey() *string

	// The key of the asset from the external environment.
	GetExternalAssetKey() *string

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
}

type createassetdetails struct {
	JsonData         []byte
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	AssetSourceIds   []string                          `mandatory:"false" json:"assetSourceIds"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	InventoryId      *string                           `mandatory:"true" json:"inventoryId"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	SourceKey        *string                           `mandatory:"true" json:"sourceKey"`
	ExternalAssetKey *string                           `mandatory:"true" json:"externalAssetKey"`
	AssetType        string                            `json:"assetType"`
}

// UnmarshalJSON unmarshals json
func (m *createassetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateassetdetails createassetdetails
	s := struct {
		Model Unmarshalercreateassetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InventoryId = s.Model.InventoryId
	m.CompartmentId = s.Model.CompartmentId
	m.SourceKey = s.Model.SourceKey
	m.ExternalAssetKey = s.Model.ExternalAssetKey
	m.DisplayName = s.Model.DisplayName
	m.AssetSourceIds = s.Model.AssetSourceIds
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.AssetType = s.Model.AssetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createassetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AssetType {
	case "VMWARE_VM":
		mm := CreateVmwareVmAssetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateAssetDetails: %s.", m.AssetType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createassetdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetAssetSourceIds returns AssetSourceIds
func (m createassetdetails) GetAssetSourceIds() []string {
	return m.AssetSourceIds
}

// GetFreeformTags returns FreeformTags
func (m createassetdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createassetdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetInventoryId returns InventoryId
func (m createassetdetails) GetInventoryId() *string {
	return m.InventoryId
}

// GetCompartmentId returns CompartmentId
func (m createassetdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceKey returns SourceKey
func (m createassetdetails) GetSourceKey() *string {
	return m.SourceKey
}

// GetExternalAssetKey returns ExternalAssetKey
func (m createassetdetails) GetExternalAssetKey() *string {
	return m.ExternalAssetKey
}

func (m createassetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createassetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
