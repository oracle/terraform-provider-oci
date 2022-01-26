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

// CreateDataAssetFromJdbc Details for the generic JDBC data asset type.
type CreateDataAssetFromJdbc struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

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

	// The generic JDBC host name.
	Host *string `mandatory:"false" json:"host"`

	// The generic JDBC port number.
	Port *string `mandatory:"false" json:"port"`

	// The data asset type for the generic JDBC data asset.
	DataAssetType *string `mandatory:"false" json:"dataAssetType"`

	DefaultConnection *CreateConnectionFromJdbc `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m CreateDataAssetFromJdbc) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CreateDataAssetFromJdbc) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m CreateDataAssetFromJdbc) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateDataAssetFromJdbc) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m CreateDataAssetFromJdbc) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m CreateDataAssetFromJdbc) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m CreateDataAssetFromJdbc) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m CreateDataAssetFromJdbc) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m CreateDataAssetFromJdbc) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateDataAssetFromJdbc) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDataAssetFromJdbc) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataAssetFromJdbc CreateDataAssetFromJdbc
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateDataAssetFromJdbc
	}{
		"GENERIC_JDBC_DATA_ASSET",
		(MarshalTypeCreateDataAssetFromJdbc)(m),
	}

	return json.Marshal(&s)
}
