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

// UpdateDataAssetFromAdwc Details for the Autonomous Data Warehouse data asset type.
type UpdateDataAssetFromAdwc struct {

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

	// The Autonomous Data Warehouse instance service name.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The Autonomous Data Warehouse driver class.
	DriverClass *string `mandatory:"false" json:"driverClass"`

	// The credential file content from a Autonomous Data Warehouse wallet.
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

	DefaultConnection *UpdateConnectionFromAdwc `mandatory:"false" json:"defaultConnection"`
}

// GetKey returns Key
func (m UpdateDataAssetFromAdwc) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m UpdateDataAssetFromAdwc) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m UpdateDataAssetFromAdwc) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m UpdateDataAssetFromAdwc) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m UpdateDataAssetFromAdwc) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetObjectVersion returns ObjectVersion
func (m UpdateDataAssetFromAdwc) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetIdentifier returns Identifier
func (m UpdateDataAssetFromAdwc) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m UpdateDataAssetFromAdwc) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m UpdateDataAssetFromAdwc) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m UpdateDataAssetFromAdwc) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateDataAssetFromAdwc) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDataAssetFromAdwc) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDataAssetFromAdwc) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDataAssetFromAdwc UpdateDataAssetFromAdwc
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateDataAssetFromAdwc
	}{
		"ORACLE_ADWC_DATA_ASSET",
		(MarshalTypeUpdateDataAssetFromAdwc)(m),
	}

	return json.Marshal(&s)
}
