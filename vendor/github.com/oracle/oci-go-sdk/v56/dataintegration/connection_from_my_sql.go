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

// ConnectionFromMySql The connection details for a MYSQL data asset.
type ConnectionFromMySql struct {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description for the connection.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	PrimarySchema *Schema `mandatory:"false" json:"primarySchema"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	// The default property for the connection.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	// The user name for the connection.
	Username *string `mandatory:"false" json:"username"`
}

//GetKey returns Key
func (m ConnectionFromMySql) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ConnectionFromMySql) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ConnectionFromMySql) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m ConnectionFromMySql) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m ConnectionFromMySql) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m ConnectionFromMySql) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m ConnectionFromMySql) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m ConnectionFromMySql) GetIdentifier() *string {
	return m.Identifier
}

//GetPrimarySchema returns PrimarySchema
func (m ConnectionFromMySql) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

//GetConnectionProperties returns ConnectionProperties
func (m ConnectionFromMySql) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

//GetIsDefault returns IsDefault
func (m ConnectionFromMySql) GetIsDefault() *bool {
	return m.IsDefault
}

//GetMetadata returns Metadata
func (m ConnectionFromMySql) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

//GetKeyMap returns KeyMap
func (m ConnectionFromMySql) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m ConnectionFromMySql) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ConnectionFromMySql) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConnectionFromMySql ConnectionFromMySql
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConnectionFromMySql
	}{
		"MYSQL_CONNECTION",
		(MarshalTypeConnectionFromMySql)(m),
	}

	return json.Marshal(&s)
}
