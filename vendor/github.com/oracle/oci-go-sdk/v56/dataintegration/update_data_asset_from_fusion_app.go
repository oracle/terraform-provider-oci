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

// UpdateDataAssetFromFusionApp Details for the Autonomous Transaction Processing data asset type.
type UpdateDataAssetFromFusionApp struct {

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

	// The service url of the BI Server.
	ServiceUrl *string `mandatory:"false" json:"serviceUrl"`
}

//GetKey returns Key
func (m UpdateDataAssetFromFusionApp) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateDataAssetFromFusionApp) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m UpdateDataAssetFromFusionApp) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateDataAssetFromFusionApp) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateDataAssetFromFusionApp) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateDataAssetFromFusionApp) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateDataAssetFromFusionApp) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m UpdateDataAssetFromFusionApp) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m UpdateDataAssetFromFusionApp) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateDataAssetFromFusionApp) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateDataAssetFromFusionApp) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateDataAssetFromFusionApp) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDataAssetFromFusionApp UpdateDataAssetFromFusionApp
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateDataAssetFromFusionApp
	}{
		"FUSION_APP_DATA_ASSET",
		(MarshalTypeUpdateDataAssetFromFusionApp)(m),
	}

	return json.Marshal(&s)
}
