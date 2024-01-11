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

// ConnectionSummaryFromOAuth2 The connection summary for a OAuth connection.
type ConnectionSummaryFromOAuth2 struct {

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

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	PrimarySchema *Schema `mandatory:"false" json:"primarySchema"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	// The default property for the connection.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	// Specifies the client ID key for specific application
	ClientId *string `mandatory:"false" json:"clientId"`

	ClientSecret *SensitiveAttribute `mandatory:"false" json:"clientSecret"`

	// Specifies the OAuth scopes that limit the permissions granted by an access token.
	Scope *string `mandatory:"false" json:"scope"`

	// Specifies the OAuth2 grant mechanism. Example CLIENT_CREDENTIALS, Implicit Flow etc.
	GrantType ConnectionSummaryFromOAuth2GrantTypeEnum `mandatory:"false" json:"grantType,omitempty"`
}

// GetKey returns Key
func (m ConnectionSummaryFromOAuth2) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m ConnectionSummaryFromOAuth2) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m ConnectionSummaryFromOAuth2) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m ConnectionSummaryFromOAuth2) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m ConnectionSummaryFromOAuth2) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m ConnectionSummaryFromOAuth2) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m ConnectionSummaryFromOAuth2) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m ConnectionSummaryFromOAuth2) GetIdentifier() *string {
	return m.Identifier
}

// GetPrimarySchema returns PrimarySchema
func (m ConnectionSummaryFromOAuth2) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

// GetConnectionProperties returns ConnectionProperties
func (m ConnectionSummaryFromOAuth2) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetIsDefault returns IsDefault
func (m ConnectionSummaryFromOAuth2) GetIsDefault() *bool {
	return m.IsDefault
}

// GetMetadata returns Metadata
func (m ConnectionSummaryFromOAuth2) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

// GetKeyMap returns KeyMap
func (m ConnectionSummaryFromOAuth2) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m ConnectionSummaryFromOAuth2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectionSummaryFromOAuth2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionSummaryFromOAuth2GrantTypeEnum(string(m.GrantType)); !ok && m.GrantType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantType: %s. Supported values are: %s.", m.GrantType, strings.Join(GetConnectionSummaryFromOAuth2GrantTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ConnectionSummaryFromOAuth2) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConnectionSummaryFromOAuth2 ConnectionSummaryFromOAuth2
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConnectionSummaryFromOAuth2
	}{
		"OAUTH2_CONNECTION",
		(MarshalTypeConnectionSummaryFromOAuth2)(m),
	}

	return json.Marshal(&s)
}

// ConnectionSummaryFromOAuth2GrantTypeEnum Enum with underlying type: string
type ConnectionSummaryFromOAuth2GrantTypeEnum string

// Set of constants representing the allowable values for ConnectionSummaryFromOAuth2GrantTypeEnum
const (
	ConnectionSummaryFromOAuth2GrantTypeClientCredentials ConnectionSummaryFromOAuth2GrantTypeEnum = "CLIENT_CREDENTIALS"
)

var mappingConnectionSummaryFromOAuth2GrantTypeEnum = map[string]ConnectionSummaryFromOAuth2GrantTypeEnum{
	"CLIENT_CREDENTIALS": ConnectionSummaryFromOAuth2GrantTypeClientCredentials,
}

var mappingConnectionSummaryFromOAuth2GrantTypeEnumLowerCase = map[string]ConnectionSummaryFromOAuth2GrantTypeEnum{
	"client_credentials": ConnectionSummaryFromOAuth2GrantTypeClientCredentials,
}

// GetConnectionSummaryFromOAuth2GrantTypeEnumValues Enumerates the set of values for ConnectionSummaryFromOAuth2GrantTypeEnum
func GetConnectionSummaryFromOAuth2GrantTypeEnumValues() []ConnectionSummaryFromOAuth2GrantTypeEnum {
	values := make([]ConnectionSummaryFromOAuth2GrantTypeEnum, 0)
	for _, v := range mappingConnectionSummaryFromOAuth2GrantTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionSummaryFromOAuth2GrantTypeEnumStringValues Enumerates the set of values in String for ConnectionSummaryFromOAuth2GrantTypeEnum
func GetConnectionSummaryFromOAuth2GrantTypeEnumStringValues() []string {
	return []string{
		"CLIENT_CREDENTIALS",
	}
}

// GetMappingConnectionSummaryFromOAuth2GrantTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionSummaryFromOAuth2GrantTypeEnum(val string) (ConnectionSummaryFromOAuth2GrantTypeEnum, bool) {
	enum, ok := mappingConnectionSummaryFromOAuth2GrantTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
