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

// DataAssetSummaryFromAtp The Oracle data asset details.
type DataAssetSummaryFromAtp struct {

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

	// The service name for the data asset.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Array of service names that are available for selection in the serviceName property.
	ServiceNames []string `mandatory:"false" json:"serviceNames"`

	// The driver class for the data asset.
	DriverClass *string `mandatory:"false" json:"driverClass"`

	DefaultConnection *ConnectionSummaryFromAtp `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m DataAssetSummaryFromAtp) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DataAssetSummaryFromAtp) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m DataAssetSummaryFromAtp) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DataAssetSummaryFromAtp) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m DataAssetSummaryFromAtp) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DataAssetSummaryFromAtp) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m DataAssetSummaryFromAtp) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m DataAssetSummaryFromAtp) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m DataAssetSummaryFromAtp) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m DataAssetSummaryFromAtp) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m DataAssetSummaryFromAtp) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetMetadata returns Metadata
func (m DataAssetSummaryFromAtp) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DataAssetSummaryFromAtp) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataAssetSummaryFromAtp) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetSummaryFromAtp DataAssetSummaryFromAtp
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetSummaryFromAtp
	}{
		"ORACLE_ATP_DATA_ASSET",
		(MarshalTypeDataAssetSummaryFromAtp)(m),
	}

	return json.Marshal(&s)
}
