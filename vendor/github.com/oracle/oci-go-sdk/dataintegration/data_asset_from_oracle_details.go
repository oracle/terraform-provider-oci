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

// DataAssetFromOracleDetails The Oracle data asset details.
type DataAssetFromOracleDetails struct {

	// Generated key that can be used in API calls to identify data asset.
	Key *string `mandatory:"false" json:"key"`

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

	NativeTypeSystem *TypeSystem `mandatory:"false" json:"nativeTypeSystem"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A map, if provided key is replaced with generated key, this structure provides mapping between user provided key and generated key
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	// The host details for the data asset.
	Host *string `mandatory:"false" json:"host"`

	// The port details for the data asset.
	Port *string `mandatory:"false" json:"port"`

	// The service name for the data asset.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The driver class for the data asset.
	DriverClass *string `mandatory:"false" json:"driverClass"`

	// sid
	Sid *string `mandatory:"false" json:"sid"`

	// The credential file content from a wallet for the data asset.
	CredentialFileContent *string `mandatory:"false" json:"credentialFileContent"`

	DefaultConnection *ConnectionFromOracleDetails `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m DataAssetFromOracleDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DataAssetFromOracleDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m DataAssetFromOracleDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DataAssetFromOracleDetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m DataAssetFromOracleDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DataAssetFromOracleDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m DataAssetFromOracleDetails) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m DataAssetFromOracleDetails) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m DataAssetFromOracleDetails) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m DataAssetFromOracleDetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m DataAssetFromOracleDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetMetadata returns Metadata
func (m DataAssetFromOracleDetails) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

//GetKeyMap returns KeyMap
func (m DataAssetFromOracleDetails) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m DataAssetFromOracleDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataAssetFromOracleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetFromOracleDetails DataAssetFromOracleDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetFromOracleDetails
	}{
		"ORACLE_DATA_ASSET",
		(MarshalTypeDataAssetFromOracleDetails)(m),
	}

	return json.Marshal(&s)
}
