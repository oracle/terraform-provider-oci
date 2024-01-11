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

// UpdateConnectionFromOAuth2 The details to update a OAuth connection.
type UpdateConnectionFromOAuth2 struct {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// Specifies the endpoint used to exchange authentication credentials for access tokens
	AccessTokenUrl *string `mandatory:"true" json:"accessTokenUrl"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description for the connection.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// Specifies the client ID key for specific application
	ClientId *string `mandatory:"false" json:"clientId"`

	ClientSecret *SensitiveAttribute `mandatory:"false" json:"clientSecret"`

	// Specifies the OAuth scopes that limit the permissions granted by an access token.
	Scope *string `mandatory:"false" json:"scope"`

	// Specifies the OAuth2 grant mechanism. Example CLIENT_CREDENTIALS, Implicit Flow etc.
	GrantType UpdateConnectionFromOAuth2GrantTypeEnum `mandatory:"false" json:"grantType,omitempty"`
}

// GetKey returns Key
func (m UpdateConnectionFromOAuth2) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m UpdateConnectionFromOAuth2) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m UpdateConnectionFromOAuth2) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m UpdateConnectionFromOAuth2) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m UpdateConnectionFromOAuth2) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m UpdateConnectionFromOAuth2) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetObjectVersion returns ObjectVersion
func (m UpdateConnectionFromOAuth2) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetIdentifier returns Identifier
func (m UpdateConnectionFromOAuth2) GetIdentifier() *string {
	return m.Identifier
}

// GetConnectionProperties returns ConnectionProperties
func (m UpdateConnectionFromOAuth2) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m UpdateConnectionFromOAuth2) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateConnectionFromOAuth2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateConnectionFromOAuth2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateConnectionFromOAuth2GrantTypeEnum(string(m.GrantType)); !ok && m.GrantType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantType: %s. Supported values are: %s.", m.GrantType, strings.Join(GetUpdateConnectionFromOAuth2GrantTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateConnectionFromOAuth2) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateConnectionFromOAuth2 UpdateConnectionFromOAuth2
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateConnectionFromOAuth2
	}{
		"OAUTH2_CONNECTION",
		(MarshalTypeUpdateConnectionFromOAuth2)(m),
	}

	return json.Marshal(&s)
}

// UpdateConnectionFromOAuth2GrantTypeEnum Enum with underlying type: string
type UpdateConnectionFromOAuth2GrantTypeEnum string

// Set of constants representing the allowable values for UpdateConnectionFromOAuth2GrantTypeEnum
const (
	UpdateConnectionFromOAuth2GrantTypeClientCredentials UpdateConnectionFromOAuth2GrantTypeEnum = "CLIENT_CREDENTIALS"
)

var mappingUpdateConnectionFromOAuth2GrantTypeEnum = map[string]UpdateConnectionFromOAuth2GrantTypeEnum{
	"CLIENT_CREDENTIALS": UpdateConnectionFromOAuth2GrantTypeClientCredentials,
}

var mappingUpdateConnectionFromOAuth2GrantTypeEnumLowerCase = map[string]UpdateConnectionFromOAuth2GrantTypeEnum{
	"client_credentials": UpdateConnectionFromOAuth2GrantTypeClientCredentials,
}

// GetUpdateConnectionFromOAuth2GrantTypeEnumValues Enumerates the set of values for UpdateConnectionFromOAuth2GrantTypeEnum
func GetUpdateConnectionFromOAuth2GrantTypeEnumValues() []UpdateConnectionFromOAuth2GrantTypeEnum {
	values := make([]UpdateConnectionFromOAuth2GrantTypeEnum, 0)
	for _, v := range mappingUpdateConnectionFromOAuth2GrantTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateConnectionFromOAuth2GrantTypeEnumStringValues Enumerates the set of values in String for UpdateConnectionFromOAuth2GrantTypeEnum
func GetUpdateConnectionFromOAuth2GrantTypeEnumStringValues() []string {
	return []string{
		"CLIENT_CREDENTIALS",
	}
}

// GetMappingUpdateConnectionFromOAuth2GrantTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateConnectionFromOAuth2GrantTypeEnum(val string) (UpdateConnectionFromOAuth2GrantTypeEnum, bool) {
	enum, ok := mappingUpdateConnectionFromOAuth2GrantTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
