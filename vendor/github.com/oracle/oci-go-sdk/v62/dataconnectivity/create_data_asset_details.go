// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// CreateDataAssetDetails Properties used in data asset Create operations.
type CreateDataAssetDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// All the properties for the data asset in a key-value map format.
	Properties map[string]interface{} `mandatory:"true" json:"properties"`

	// Specific DataAsset Type
	Type *string `mandatory:"true" json:"type"`

	// Generated key that can be used in API calls to identify dataasset.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// User-defined description of the data asset.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Additional properties for the data asset.
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	NativeTypeSystem *TypeSystem `mandatory:"false" json:"nativeTypeSystem"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	DefaultConnection *Connection `mandatory:"false" json:"defaultConnection"`

	// The list of endpoints with which this data asset is associated.
	EndPoints []DpEndpoint `mandatory:"false" json:"endPoints"`
}

func (m CreateDataAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDataAssetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key               *string                `json:"key"`
		ModelVersion      *string                `json:"modelVersion"`
		ModelType         *string                `json:"modelType"`
		Description       *string                `json:"description"`
		ObjectStatus      *int                   `json:"objectStatus"`
		ObjectVersion     *int                   `json:"objectVersion"`
		ExternalKey       *string                `json:"externalKey"`
		AssetProperties   map[string]string      `json:"assetProperties"`
		NativeTypeSystem  *TypeSystem            `json:"nativeTypeSystem"`
		RegistryMetadata  *RegistryMetadata      `json:"registryMetadata"`
		Metadata          *ObjectMetadata        `json:"metadata"`
		DefaultConnection *Connection            `json:"defaultConnection"`
		EndPoints         []dpendpoint           `json:"endPoints"`
		Name              *string                `json:"name"`
		Identifier        *string                `json:"identifier"`
		Properties        map[string]interface{} `json:"properties"`
		Type              *string                `json:"type"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ModelType = model.ModelType

	m.Description = model.Description

	m.ObjectStatus = model.ObjectStatus

	m.ObjectVersion = model.ObjectVersion

	m.ExternalKey = model.ExternalKey

	m.AssetProperties = model.AssetProperties

	m.NativeTypeSystem = model.NativeTypeSystem

	m.RegistryMetadata = model.RegistryMetadata

	m.Metadata = model.Metadata

	m.DefaultConnection = model.DefaultConnection

	m.EndPoints = make([]DpEndpoint, len(model.EndPoints))
	for i, n := range model.EndPoints {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.EndPoints[i] = nn.(DpEndpoint)
		} else {
			m.EndPoints[i] = nil
		}
	}

	m.Name = model.Name

	m.Identifier = model.Identifier

	m.Properties = model.Properties

	m.Type = model.Type

	return
}
