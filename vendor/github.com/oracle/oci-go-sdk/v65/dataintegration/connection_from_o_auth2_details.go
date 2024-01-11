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

// ConnectionFromOAuth2Details The connection details for a OAuth connection.
type ConnectionFromOAuth2Details struct {

	// Specifies the endpoint used to exchange authentication credentials for access tokens
	AccessTokenUrl *string `mandatory:"true" json:"accessTokenUrl"`

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

	// Specifies the client ID key for specific application
	ClientId *string `mandatory:"false" json:"clientId"`

	ClientSecret *SensitiveAttribute `mandatory:"false" json:"clientSecret"`

	// Specifies the OAuth scopes that limit the permissions granted by an access token.
	Scope *string `mandatory:"false" json:"scope"`

	// Specifies the OAuth2 grant mechanism. Example CLIENT_CREDENTIALS, Implicit Flow etc.
	GrantType ConnectionFromOAuth2DetailsGrantTypeEnum `mandatory:"false" json:"grantType,omitempty"`
}

// GetKey returns Key
func (m ConnectionFromOAuth2Details) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m ConnectionFromOAuth2Details) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m ConnectionFromOAuth2Details) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m ConnectionFromOAuth2Details) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m ConnectionFromOAuth2Details) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m ConnectionFromOAuth2Details) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m ConnectionFromOAuth2Details) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m ConnectionFromOAuth2Details) GetIdentifier() *string {
	return m.Identifier
}

// GetPrimarySchema returns PrimarySchema
func (m ConnectionFromOAuth2Details) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

// GetConnectionProperties returns ConnectionProperties
func (m ConnectionFromOAuth2Details) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetIsDefault returns IsDefault
func (m ConnectionFromOAuth2Details) GetIsDefault() *bool {
	return m.IsDefault
}

// GetMetadata returns Metadata
func (m ConnectionFromOAuth2Details) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m ConnectionFromOAuth2Details) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectionFromOAuth2Details) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionFromOAuth2DetailsGrantTypeEnum(string(m.GrantType)); !ok && m.GrantType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantType: %s. Supported values are: %s.", m.GrantType, strings.Join(GetConnectionFromOAuth2DetailsGrantTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ConnectionFromOAuth2Details) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConnectionFromOAuth2Details ConnectionFromOAuth2Details
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConnectionFromOAuth2Details
	}{
		"OAUTH2_CONNECTION",
		(MarshalTypeConnectionFromOAuth2Details)(m),
	}

	return json.Marshal(&s)
}

// ConnectionFromOAuth2DetailsGrantTypeEnum Enum with underlying type: string
type ConnectionFromOAuth2DetailsGrantTypeEnum string

// Set of constants representing the allowable values for ConnectionFromOAuth2DetailsGrantTypeEnum
const (
	ConnectionFromOAuth2DetailsGrantTypeClientCredentials ConnectionFromOAuth2DetailsGrantTypeEnum = "CLIENT_CREDENTIALS"
)

var mappingConnectionFromOAuth2DetailsGrantTypeEnum = map[string]ConnectionFromOAuth2DetailsGrantTypeEnum{
	"CLIENT_CREDENTIALS": ConnectionFromOAuth2DetailsGrantTypeClientCredentials,
}

var mappingConnectionFromOAuth2DetailsGrantTypeEnumLowerCase = map[string]ConnectionFromOAuth2DetailsGrantTypeEnum{
	"client_credentials": ConnectionFromOAuth2DetailsGrantTypeClientCredentials,
}

// GetConnectionFromOAuth2DetailsGrantTypeEnumValues Enumerates the set of values for ConnectionFromOAuth2DetailsGrantTypeEnum
func GetConnectionFromOAuth2DetailsGrantTypeEnumValues() []ConnectionFromOAuth2DetailsGrantTypeEnum {
	values := make([]ConnectionFromOAuth2DetailsGrantTypeEnum, 0)
	for _, v := range mappingConnectionFromOAuth2DetailsGrantTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionFromOAuth2DetailsGrantTypeEnumStringValues Enumerates the set of values in String for ConnectionFromOAuth2DetailsGrantTypeEnum
func GetConnectionFromOAuth2DetailsGrantTypeEnumStringValues() []string {
	return []string{
		"CLIENT_CREDENTIALS",
	}
}

// GetMappingConnectionFromOAuth2DetailsGrantTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionFromOAuth2DetailsGrantTypeEnum(val string) (ConnectionFromOAuth2DetailsGrantTypeEnum, bool) {
	enum, ok := mappingConnectionFromOAuth2DetailsGrantTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
