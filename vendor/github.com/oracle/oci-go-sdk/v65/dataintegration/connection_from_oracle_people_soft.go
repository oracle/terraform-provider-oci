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

// ConnectionFromOraclePeopleSoft The connection details for an Oracle PeopleSoft data asset.
type ConnectionFromOraclePeopleSoft struct {

	// The user name for the connection.
	Username *string `mandatory:"true" json:"username"`

	// The password for the connection.
	Password *string `mandatory:"true" json:"password"`

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

	PasswordSecret *SensitiveAttribute `mandatory:"false" json:"passwordSecret"`
}

// GetKey returns Key
func (m ConnectionFromOraclePeopleSoft) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m ConnectionFromOraclePeopleSoft) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m ConnectionFromOraclePeopleSoft) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m ConnectionFromOraclePeopleSoft) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m ConnectionFromOraclePeopleSoft) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m ConnectionFromOraclePeopleSoft) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m ConnectionFromOraclePeopleSoft) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m ConnectionFromOraclePeopleSoft) GetIdentifier() *string {
	return m.Identifier
}

// GetPrimarySchema returns PrimarySchema
func (m ConnectionFromOraclePeopleSoft) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

// GetConnectionProperties returns ConnectionProperties
func (m ConnectionFromOraclePeopleSoft) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetIsDefault returns IsDefault
func (m ConnectionFromOraclePeopleSoft) GetIsDefault() *bool {
	return m.IsDefault
}

// GetMetadata returns Metadata
func (m ConnectionFromOraclePeopleSoft) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

// GetKeyMap returns KeyMap
func (m ConnectionFromOraclePeopleSoft) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m ConnectionFromOraclePeopleSoft) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectionFromOraclePeopleSoft) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ConnectionFromOraclePeopleSoft) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConnectionFromOraclePeopleSoft ConnectionFromOraclePeopleSoft
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConnectionFromOraclePeopleSoft
	}{
		"ORACLE_PEOPLESOFT_CONNECTION",
		(MarshalTypeConnectionFromOraclePeopleSoft)(m),
	}

	return json.Marshal(&s)
}
