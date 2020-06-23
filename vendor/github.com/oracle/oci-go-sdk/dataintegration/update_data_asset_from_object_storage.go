// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateDataAssetFromObjectStorage The Oracle data asset details.
type UpdateDataAssetFromObjectStorage struct {

	// Generated key that can be used in API calls to identify data asset.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The external key for the object
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// assetProperties
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// url
	Url *string `mandatory:"false" json:"url"`

	// The OCI tenancy OCID.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
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
