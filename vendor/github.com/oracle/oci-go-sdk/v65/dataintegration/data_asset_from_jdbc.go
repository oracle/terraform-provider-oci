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

// DataAssetFromJdbc Details for the generic JDBC data asset type.
type DataAssetFromJdbc struct {

	// Generated key that can be used in API calls to identify data asset.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description of the data asset.
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

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	// The generic JDBC host name.
	Host *string `mandatory:"false" json:"host"`

	// The generic JDBC port number.
	Port *string `mandatory:"false" json:"port"`

	// The data asset type for the generic JDBC data asset.
	DataAssetType *string `mandatory:"false" json:"dataAssetType"`

	DefaultConnection *ConnectionFromJdbcDetails `mandatory:"false" json:"defaultConnection"`
}

// GetKey returns Key
func (m DataAssetFromJdbc) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m DataAssetFromJdbc) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m DataAssetFromJdbc) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m DataAssetFromJdbc) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m DataAssetFromJdbc) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m DataAssetFromJdbc) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m DataAssetFromJdbc) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m DataAssetFromJdbc) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetNativeTypeSystem returns NativeTypeSystem
func (m DataAssetFromJdbc) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

// GetObjectVersion returns ObjectVersion
func (m DataAssetFromJdbc) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetParentRef returns ParentRef
func (m DataAssetFromJdbc) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetMetadata returns Metadata
func (m DataAssetFromJdbc) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

// GetKeyMap returns KeyMap
func (m DataAssetFromJdbc) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m DataAssetFromJdbc) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataAssetFromJdbc) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataAssetFromJdbc) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetFromJdbc DataAssetFromJdbc
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetFromJdbc
	}{
		"GENERIC_JDBC_DATA_ASSET",
		(MarshalTypeDataAssetFromJdbc)(m),
	}

	return json.Marshal(&s)
}
