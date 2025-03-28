// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDataAssetFromFusionApp Details for the FUSION_APP data asset type.
type CreateDataAssetFromFusionApp struct {

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
	ServiceUrl *string `mandatory:"false" json:"serviceUrl"`

	DefaultConnection CreateConnectionDetails `mandatory:"false" json:"defaultConnection"`

	StagingDataAsset *DataAssetSummaryFromObjectStorage `mandatory:"false" json:"stagingDataAsset"`

	StagingConnection *ConnectionSummaryFromObjectStorage `mandatory:"false" json:"stagingConnection"`

	BucketSchema *Schema `mandatory:"false" json:"bucketSchema"`
}

// GetKey returns Key
func (m CreateDataAssetFromFusionApp) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m CreateDataAssetFromFusionApp) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m CreateDataAssetFromFusionApp) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateDataAssetFromFusionApp) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m CreateDataAssetFromFusionApp) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m CreateDataAssetFromFusionApp) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m CreateDataAssetFromFusionApp) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m CreateDataAssetFromFusionApp) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m CreateDataAssetFromFusionApp) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateDataAssetFromFusionApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataAssetFromFusionApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDataAssetFromFusionApp) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataAssetFromFusionApp CreateDataAssetFromFusionApp
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateDataAssetFromFusionApp
	}{
		"FUSION_APP_DATA_ASSET",
		(MarshalTypeCreateDataAssetFromFusionApp)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDataAssetFromFusionApp) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key               *string                             `json:"key"`
		ModelVersion      *string                             `json:"modelVersion"`
		Description       *string                             `json:"description"`
		ObjectStatus      *int                                `json:"objectStatus"`
		ExternalKey       *string                             `json:"externalKey"`
		AssetProperties   map[string]string                   `json:"assetProperties"`
		RegistryMetadata  *RegistryMetadata                   `json:"registryMetadata"`
		ServiceUrl        *string                             `json:"serviceUrl"`
		DefaultConnection createconnectiondetails             `json:"defaultConnection"`
		StagingDataAsset  *DataAssetSummaryFromObjectStorage  `json:"stagingDataAsset"`
		StagingConnection *ConnectionSummaryFromObjectStorage `json:"stagingConnection"`
		BucketSchema      *Schema                             `json:"bucketSchema"`
		Name              *string                             `json:"name"`
		Identifier        *string                             `json:"identifier"`
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

	m.ServiceUrl = model.ServiceUrl

	nn, e = model.DefaultConnection.UnmarshalPolymorphicJSON(model.DefaultConnection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DefaultConnection = nn.(CreateConnectionDetails)
	} else {
		m.DefaultConnection = nil
	}

	m.StagingDataAsset = model.StagingDataAsset

	m.StagingConnection = model.StagingConnection

	m.BucketSchema = model.BucketSchema

	m.Name = model.Name

	m.Identifier = model.Identifier

	return
}
