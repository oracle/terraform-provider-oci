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

// CreateDataAssetFromAtp Details for the Autonomous Transaction Processing data asset type.
type CreateDataAssetFromAtp struct {

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

	// The Autonomous Transaction Processing instance service name.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The Autonomous Transaction Processing driver class.
	DriverClass *string `mandatory:"false" json:"driverClass"`

	// The credential file content from an Autonomous Transaction Processing wallet.
	CredentialFileContent *string `mandatory:"false" json:"credentialFileContent"`

	WalletSecret *SensitiveAttribute `mandatory:"false" json:"walletSecret"`

	WalletPasswordSecret *SensitiveAttribute `mandatory:"false" json:"walletPasswordSecret"`

	// The Autonomous Data Warehouse instance region Id.
	RegionId *string `mandatory:"false" json:"regionId"`

	// The Autonomous Data Warehouse instance tenancy Id.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// The Autonomous Data Warehouse instance compartment Id.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Tha Autonomous Database Id
	AutonomousDbId *string `mandatory:"false" json:"autonomousDbId"`

	DefaultConnection *CreateConnectionFromAtp `mandatory:"false" json:"defaultConnection"`

	StagingDataAsset DataAsset `mandatory:"false" json:"stagingDataAsset"`

	StagingConnection Connection `mandatory:"false" json:"stagingConnection"`

	BucketSchema *Schema `mandatory:"false" json:"bucketSchema"`
}

// GetKey returns Key
func (m CreateDataAssetFromAtp) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m CreateDataAssetFromAtp) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m CreateDataAssetFromAtp) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateDataAssetFromAtp) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m CreateDataAssetFromAtp) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m CreateDataAssetFromAtp) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m CreateDataAssetFromAtp) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m CreateDataAssetFromAtp) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m CreateDataAssetFromAtp) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateDataAssetFromAtp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataAssetFromAtp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

// UnmarshalJSON unmarshals from json
func (m *CreateDataAssetFromAtp) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                   *string                  `json:"key"`
		ModelVersion          *string                  `json:"modelVersion"`
		Description           *string                  `json:"description"`
		ObjectStatus          *int                     `json:"objectStatus"`
		ExternalKey           *string                  `json:"externalKey"`
		AssetProperties       map[string]string        `json:"assetProperties"`
		RegistryMetadata      *RegistryMetadata        `json:"registryMetadata"`
		ServiceName           *string                  `json:"serviceName"`
		DriverClass           *string                  `json:"driverClass"`
		CredentialFileContent *string                  `json:"credentialFileContent"`
		WalletSecret          *SensitiveAttribute      `json:"walletSecret"`
		WalletPasswordSecret  *SensitiveAttribute      `json:"walletPasswordSecret"`
		RegionId              *string                  `json:"regionId"`
		TenancyId             *string                  `json:"tenancyId"`
		CompartmentId         *string                  `json:"compartmentId"`
		AutonomousDbId        *string                  `json:"autonomousDbId"`
		DefaultConnection     *CreateConnectionFromAtp `json:"defaultConnection"`
		StagingDataAsset      dataasset                `json:"stagingDataAsset"`
		StagingConnection     connection               `json:"stagingConnection"`
		BucketSchema          *Schema                  `json:"bucketSchema"`
		Name                  *string                  `json:"name"`
		Identifier            *string                  `json:"identifier"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.Description = model.Description

	m.ObjectStatus = model.ObjectStatus

	m.ExternalKey = model.ExternalKey

	m.AssetProperties = model.AssetProperties

	m.RegistryMetadata = model.RegistryMetadata

	m.ServiceName = model.ServiceName

	m.DriverClass = model.DriverClass

	m.CredentialFileContent = model.CredentialFileContent

	m.WalletSecret = model.WalletSecret

	m.WalletPasswordSecret = model.WalletPasswordSecret

	m.RegionId = model.RegionId

	m.TenancyId = model.TenancyId

	m.CompartmentId = model.CompartmentId

	m.AutonomousDbId = model.AutonomousDbId

	m.DefaultConnection = model.DefaultConnection

	nn, e = model.StagingDataAsset.UnmarshalPolymorphicJSON(model.StagingDataAsset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StagingDataAsset = nn.(DataAsset)
	} else {
		m.StagingDataAsset = nil
	}

	nn, e = model.StagingConnection.UnmarshalPolymorphicJSON(model.StagingConnection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StagingConnection = nn.(Connection)
	} else {
		m.StagingConnection = nil
	}

	m.BucketSchema = model.BucketSchema

	m.Name = model.Name

	m.Identifier = model.Identifier

	return
}
