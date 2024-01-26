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

// CreateConnectionFromOAuth2 The details to create a OAuth2 connection
type CreateConnectionFromOAuth2 struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Specifies the endpoint used to exchange authentication credentials for access tokens
	AccessTokenUrl *string `mandatory:"true" json:"accessTokenUrl"`

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// User-defined description for the connection.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// Specifies the client ID key for specific application
	ClientId *string `mandatory:"false" json:"clientId"`

	ClientSecret *SensitiveAttribute `mandatory:"false" json:"clientSecret"`

	// Specifies the OAuth scopes that limit the permissions granted by an access token.
	Scope *string `mandatory:"false" json:"scope"`

	// Specifies the OAuth2 grant mechanism. Example CLIENT_CREDENTIALS, Implicit Flow etc.
	GrantType CreateConnectionFromOAuth2GrantTypeEnum `mandatory:"false" json:"grantType,omitempty"`
}

// GetKey returns Key
func (m CreateConnectionFromOAuth2) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m CreateConnectionFromOAuth2) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m CreateConnectionFromOAuth2) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m CreateConnectionFromOAuth2) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateConnectionFromOAuth2) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m CreateConnectionFromOAuth2) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m CreateConnectionFromOAuth2) GetIdentifier() *string {
	return m.Identifier
}

// GetConnectionProperties returns ConnectionProperties
func (m CreateConnectionFromOAuth2) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m CreateConnectionFromOAuth2) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateConnectionFromOAuth2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateConnectionFromOAuth2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateConnectionFromOAuth2GrantTypeEnum(string(m.GrantType)); !ok && m.GrantType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantType: %s. Supported values are: %s.", m.GrantType, strings.Join(GetCreateConnectionFromOAuth2GrantTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateConnectionFromOAuth2) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateConnectionFromOAuth2 CreateConnectionFromOAuth2
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateConnectionFromOAuth2
	}{
		"OAUTH2_CONNECTION",
		(MarshalTypeCreateConnectionFromOAuth2)(m),
	}

	return json.Marshal(&s)
}

// CreateConnectionFromOAuth2GrantTypeEnum Enum with underlying type: string
type CreateConnectionFromOAuth2GrantTypeEnum string

// Set of constants representing the allowable values for CreateConnectionFromOAuth2GrantTypeEnum
const (
	CreateConnectionFromOAuth2GrantTypeClientCredentials CreateConnectionFromOAuth2GrantTypeEnum = "CLIENT_CREDENTIALS"
)

var mappingCreateConnectionFromOAuth2GrantTypeEnum = map[string]CreateConnectionFromOAuth2GrantTypeEnum{
	"CLIENT_CREDENTIALS": CreateConnectionFromOAuth2GrantTypeClientCredentials,
}

var mappingCreateConnectionFromOAuth2GrantTypeEnumLowerCase = map[string]CreateConnectionFromOAuth2GrantTypeEnum{
	"client_credentials": CreateConnectionFromOAuth2GrantTypeClientCredentials,
}

// GetCreateConnectionFromOAuth2GrantTypeEnumValues Enumerates the set of values for CreateConnectionFromOAuth2GrantTypeEnum
func GetCreateConnectionFromOAuth2GrantTypeEnumValues() []CreateConnectionFromOAuth2GrantTypeEnum {
	values := make([]CreateConnectionFromOAuth2GrantTypeEnum, 0)
	for _, v := range mappingCreateConnectionFromOAuth2GrantTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateConnectionFromOAuth2GrantTypeEnumStringValues Enumerates the set of values in String for CreateConnectionFromOAuth2GrantTypeEnum
func GetCreateConnectionFromOAuth2GrantTypeEnumStringValues() []string {
	return []string{
		"CLIENT_CREDENTIALS",
	}
}

// GetMappingCreateConnectionFromOAuth2GrantTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateConnectionFromOAuth2GrantTypeEnum(val string) (CreateConnectionFromOAuth2GrantTypeEnum, bool) {
	enum, ok := mappingCreateConnectionFromOAuth2GrantTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
