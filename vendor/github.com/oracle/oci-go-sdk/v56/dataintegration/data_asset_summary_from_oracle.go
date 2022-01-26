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

// DataAssetSummaryFromOracle Summary details for the Oracle Database data asset type.
type DataAssetSummaryFromOracle struct {

	// Generated key that can be used in API calls to identify data asset.
	Key *string `mandatory:"false" json:"key"`

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

	NativeTypeSystem *TypeSystem `mandatory:"false" json:"nativeTypeSystem"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The Oracle Database hostname.
	Host *string `mandatory:"false" json:"host"`

	// The Oracle Database port.
	Port *string `mandatory:"false" json:"port"`

	// The Oracle Database service name.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The Oracle Database driver class.
	DriverClass *string `mandatory:"false" json:"driverClass"`

	// The Oracle Database SID.
	Sid *string `mandatory:"false" json:"sid"`

	// The credential file content from a wallet for the data asset.
	CredentialFileContent *string `mandatory:"false" json:"credentialFileContent"`

	WalletSecret *SensitiveAttribute `mandatory:"false" json:"walletSecret"`

	WalletPasswordSecret *SensitiveAttribute `mandatory:"false" json:"walletPasswordSecret"`

	DefaultConnection *ConnectionSummaryFromOracle `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m DataAssetSummaryFromOracle) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DataAssetSummaryFromOracle) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m DataAssetSummaryFromOracle) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DataAssetSummaryFromOracle) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m DataAssetSummaryFromOracle) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DataAssetSummaryFromOracle) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m DataAssetSummaryFromOracle) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m DataAssetSummaryFromOracle) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m DataAssetSummaryFromOracle) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m DataAssetSummaryFromOracle) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m DataAssetSummaryFromOracle) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetMetadata returns Metadata
func (m DataAssetSummaryFromOracle) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DataAssetSummaryFromOracle) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataAssetSummaryFromOracle) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetSummaryFromOracle DataAssetSummaryFromOracle
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetSummaryFromOracle
	}{
		"ORACLE_DATA_ASSET",
		(MarshalTypeDataAssetSummaryFromOracle)(m),
	}

	return json.Marshal(&s)
}
