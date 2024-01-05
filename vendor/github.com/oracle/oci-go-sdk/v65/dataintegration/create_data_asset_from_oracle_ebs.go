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

// CreateDataAssetFromOracleEbs Details for the E-Business Suite data asset type.
type CreateDataAssetFromOracleEbs struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The Oracle EBS hostname.
	Host *string `mandatory:"true" json:"host"`

	// The Oracle EBS port.
	Port *string `mandatory:"true" json:"port"`

	DefaultConnection *CreateConnectionFromOracleEbs `mandatory:"true" json:"defaultConnection"`

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

	// The service name for the data asset.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The Oracle EBS driver class.
	DriverClass *string `mandatory:"false" json:"driverClass"`

	// The Oracle EBS SID.
	Sid *string `mandatory:"false" json:"sid"`

	WalletSecret *SensitiveAttribute `mandatory:"false" json:"walletSecret"`

	WalletPasswordSecret *SensitiveAttribute `mandatory:"false" json:"walletPasswordSecret"`
}

// GetKey returns Key
func (m CreateDataAssetFromOracleEbs) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m CreateDataAssetFromOracleEbs) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m CreateDataAssetFromOracleEbs) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateDataAssetFromOracleEbs) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m CreateDataAssetFromOracleEbs) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m CreateDataAssetFromOracleEbs) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m CreateDataAssetFromOracleEbs) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m CreateDataAssetFromOracleEbs) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m CreateDataAssetFromOracleEbs) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateDataAssetFromOracleEbs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataAssetFromOracleEbs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDataAssetFromOracleEbs) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataAssetFromOracleEbs CreateDataAssetFromOracleEbs
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateDataAssetFromOracleEbs
	}{
		"ORACLE_EBS_DATA_ASSET",
		(MarshalTypeCreateDataAssetFromOracleEbs)(m),
	}

	return json.Marshal(&s)
}
