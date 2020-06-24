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

// UpdateConnectionFromObjectStorage The Object Storage connection details.
type UpdateConnectionFromObjectStorage struct {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// The credential file content from a wallet for the data asset.
	CredentialFileContent *string `mandatory:"false" json:"credentialFileContent"`

	// The OCI user OCID for the user to connect to.
	UserId *string `mandatory:"false" json:"userId"`

	// The fingeprint for the user.
	FingerPrint *string `mandatory:"false" json:"fingerPrint"`

	// The pass phrase for the connection.
	PassPhrase *string `mandatory:"false" json:"passPhrase"`
}

//GetKey returns Key
func (m UpdateConnectionFromObjectStorage) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateConnectionFromObjectStorage) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateConnectionFromObjectStorage) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateConnectionFromObjectStorage) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateConnectionFromObjectStorage) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateConnectionFromObjectStorage) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateConnectionFromObjectStorage) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateConnectionFromObjectStorage) GetIdentifier() *string {
	return m.Identifier
}

//GetConnectionProperties returns ConnectionProperties
func (m UpdateConnectionFromObjectStorage) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateConnectionFromObjectStorage) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateConnectionFromObjectStorage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateConnectionFromObjectStorage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateConnectionFromObjectStorage UpdateConnectionFromObjectStorage
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateConnectionFromObjectStorage
	}{
		"ORACLE_OBJECT_STORAGE_CONNECTION",
		(MarshalTypeUpdateConnectionFromObjectStorage)(m),
	}

	return json.Marshal(&s)
}
