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

// DataAssetSummaryFromAmazonS3 Summary details for the Amazon s3 data asset type.
type DataAssetSummaryFromAmazonS3 struct {

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

	// The region for Amazon s3
	Region *string `mandatory:"false" json:"region"`

	DefaultConnection *ConnectionSummaryFromAmazonS3 `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m DataAssetSummaryFromAmazonS3) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DataAssetSummaryFromAmazonS3) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m DataAssetSummaryFromAmazonS3) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DataAssetSummaryFromAmazonS3) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m DataAssetSummaryFromAmazonS3) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DataAssetSummaryFromAmazonS3) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m DataAssetSummaryFromAmazonS3) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m DataAssetSummaryFromAmazonS3) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m DataAssetSummaryFromAmazonS3) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m DataAssetSummaryFromAmazonS3) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m DataAssetSummaryFromAmazonS3) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetMetadata returns Metadata
func (m DataAssetSummaryFromAmazonS3) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DataAssetSummaryFromAmazonS3) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataAssetSummaryFromAmazonS3) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetSummaryFromAmazonS3 DataAssetSummaryFromAmazonS3
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetSummaryFromAmazonS3
	}{
		"AMAZON_S3_DATA_ASSET",
		(MarshalTypeDataAssetSummaryFromAmazonS3)(m),
	}

	return json.Marshal(&s)
}
