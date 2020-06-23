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

// ConnectionSummaryFromOracle The Oracle connection details object.
type ConnectionSummaryFromOracle struct {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	PrimarySchema *Schema `mandatory:"false" json:"primarySchema"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	// The default property for the connection.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A map, if provided key is replaced with generated key, this structure provides mapping between user provided key and generated key
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	// The user name for the connection.
	Username *string `mandatory:"false" json:"username"`

	// The password for the connection.
	Password *string `mandatory:"false" json:"password"`
}

//GetKey returns Key
func (m ConnectionSummaryFromOracle) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ConnectionSummaryFromOracle) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ConnectionSummaryFromOracle) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m ConnectionSummaryFromOracle) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m ConnectionSummaryFromOracle) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m ConnectionSummaryFromOracle) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m ConnectionSummaryFromOracle) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m ConnectionSummaryFromOracle) GetIdentifier() *string {
	return m.Identifier
}

//GetPrimarySchema returns PrimarySchema
func (m ConnectionSummaryFromOracle) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

//GetConnectionProperties returns ConnectionProperties
func (m ConnectionSummaryFromOracle) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

//GetIsDefault returns IsDefault
func (m ConnectionSummaryFromOracle) GetIsDefault() *bool {
	return m.IsDefault
}

//GetMetadata returns Metadata
func (m ConnectionSummaryFromOracle) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

//GetKeyMap returns KeyMap
func (m ConnectionSummaryFromOracle) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m ConnectionSummaryFromOracle) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ConnectionSummaryFromOracle) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConnectionSummaryFromOracle ConnectionSummaryFromOracle
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConnectionSummaryFromOracle
	}{
		"ORACLEDB_CONNECTION",
		(MarshalTypeConnectionSummaryFromOracle)(m),
	}

	return json.Marshal(&s)
}
