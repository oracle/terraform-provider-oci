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

// CreateDataAssetFromAtp The ATP data asset details.
type CreateDataAssetFromAtp struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Currently not used on data asset creation. Reserved for future.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The external key for the object
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// assetProperties
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// The service name for the data asset.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The driver class for the data asset.
	DriverClass *string `mandatory:"false" json:"driverClass"`

	// The credential file content from a wallet for the data asset.
	CredentialFileContent *string `mandatory:"false" json:"credentialFileContent"`

	DefaultConnection *CreateConnectionFromAtp `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m CreateDataAssetFromAtp) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CreateDataAssetFromAtp) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m CreateDataAssetFromAtp) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateDataAssetFromAtp) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m CreateDataAssetFromAtp) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m CreateDataAssetFromAtp) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m CreateDataAssetFromAtp) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m CreateDataAssetFromAtp) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m CreateDataAssetFromAtp) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateDataAssetFromAtp) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDataAssetFromAtp) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataAssetFromAtp CreateDataAssetFromAtp
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateDataAssetFromAtp
	}{
		"ORACLE_ATP_DATA_ASSET",
		(MarshalTypeCreateDataAssetFromAtp)(m),
	}

	return json.Marshal(&s)
}
