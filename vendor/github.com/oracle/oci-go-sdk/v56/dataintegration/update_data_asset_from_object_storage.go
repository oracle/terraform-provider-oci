// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateDataAssetFromObjectStorage Details for the Oracle Object storage data asset type.
type UpdateDataAssetFromObjectStorage struct {

	// Generated key that can be used in API calls to identify data asset.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The user-defined description of the data asset.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Additional properties for the data asset.
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// The Oracle Object storage Region ie. us-ashburn-1
	OciRegion *string `mandatory:"false" json:"ociRegion"`

	// The Oracle Object storage URL.
	Url *string `mandatory:"false" json:"url"`

	// The OCI tenancy OCID.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// The namespace for the specified Oracle Object storage resource. You can find the namespace under Object Storage Settings in the Console.
	Namespace *string `mandatory:"false" json:"namespace"`

	DefaultConnection *UpdateConnectionFromObjectStorage `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m UpdateDataAssetFromObjectStorage) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateDataAssetFromObjectStorage) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m UpdateDataAssetFromObjectStorage) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateDataAssetFromObjectStorage) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateDataAssetFromObjectStorage) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateDataAssetFromObjectStorage) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateDataAssetFromObjectStorage) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m UpdateDataAssetFromObjectStorage) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m UpdateDataAssetFromObjectStorage) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateDataAssetFromObjectStorage) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateDataAssetFromObjectStorage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateDataAssetFromObjectStorage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDataAssetFromObjectStorage UpdateDataAssetFromObjectStorage
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateDataAssetFromObjectStorage
	}{
		"ORACLE_OBJECT_STORAGE_DATA_ASSET",
		(MarshalTypeUpdateDataAssetFromObjectStorage)(m),
	}

	return json.Marshal(&s)
}
