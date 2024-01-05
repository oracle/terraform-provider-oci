// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDataAssetFromLake Details for the Lake data asset type.
type CreateDataAssetFromLake struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The Lake Ocid.
	LakeId *string `mandatory:"true" json:"lakeId"`

	DefaultConnection *CreateConnectionFromLake `mandatory:"true" json:"defaultConnection"`

	// Currently not used on data asset creation. Reserved for future.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// User-defined description of the data asset.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Additional properties for the data asset.
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// The metastoreId for the specified Lake Resource.
	MetastoreId *string `mandatory:"false" json:"metastoreId"`

	// The lakeProxyEndpoint for the specified Lake Resource.
	LakeProxyEndpoint *string `mandatory:"false" json:"lakeProxyEndpoint"`
}

// GetKey returns Key
func (m CreateDataAssetFromLake) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m CreateDataAssetFromLake) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m CreateDataAssetFromLake) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateDataAssetFromLake) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m CreateDataAssetFromLake) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m CreateDataAssetFromLake) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m CreateDataAssetFromLake) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m CreateDataAssetFromLake) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m CreateDataAssetFromLake) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateDataAssetFromLake) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataAssetFromLake) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDataAssetFromLake) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataAssetFromLake CreateDataAssetFromLake
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateDataAssetFromLake
	}{
		"LAKE_DATA_ASSET",
		(MarshalTypeCreateDataAssetFromLake)(m),
	}

	return json.Marshal(&s)
}
