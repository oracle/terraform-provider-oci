// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault The content of a Database Tools database API gateway config pool sub resource.
type DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault struct {

	// A system generated string that uniquely identifies a pool sub resource.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The pool route value provided in requests to target this pool.
	PoolRouteValue *string `mandatory:"false" json:"poolRouteValue"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection. Specifies the OCI database tools connection ocid to build the connection pool from.
	DatabaseToolsConnectionId *string `mandatory:"false" json:"databaseToolsConnectionId"`

	// Specifies the maximum number of database connections allowed for the pool.
	MaxPoolSize *int `mandatory:"false" json:"maxPoolSize"`

	// Specifies the minimum number of database connections allowed for the pool.
	MinPoolSize *int `mandatory:"false" json:"minPoolSize"`

	// Specifies the initial size for the number of database connections that will be created for the pool.
	InitialPoolSize *int `mandatory:"false" json:"initialPoolSize"`

	// Specifies the URL of the JSON Web Key (JWK) that is used to verify the signature of the JWT token.
	JwtProfileJwkUrl *string `mandatory:"false" json:"jwtProfileJwkUrl"`

	// Specifies the issuer of the JWT token. This value is used to validate the iss claim in the JWT token.
	JwtProfileIssuer *string `mandatory:"false" json:"jwtProfileIssuer"`

	// Specifies the expected audience for the JWT token. This value is used to validate the aud claim in the JWT token.
	JwtProfileAudience *string `mandatory:"false" json:"jwtProfileAudience"`

	// Specifies the JSON pointer to the claim in the JWT token that contains the roles of the users.
	JwtProfileRoleClaimName *string `mandatory:"false" json:"jwtProfileRoleClaimName"`

	// Advanced pool properties.
	AdvancedProperties map[string]string `mandatory:"false" json:"advancedProperties"`

	// Array of Database Tools database API gateway config API spec items.
	ApiSpecs []DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec `mandatory:"false" json:"apiSpecs"`

	// Array of Database Tools database API gateway config auto API spec items.
	AutoApiSpecs []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec `mandatory:"false" json:"autoApiSpecs"`

	// The time the resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Specifies to enable the Database Actions feature.
	DatabaseActionsStatus DatabaseToolsDatabaseApiGatewayConfigPoolContentDatabaseActionsStatusEnum `mandatory:"false" json:"databaseActionsStatus,omitempty"`

	// Specifies whether the REST-Enabled SQL service is active.
	RestEnabledSqlStatus DatabaseToolsDatabaseApiGatewayConfigPoolContentRestEnabledSqlStatusEnum `mandatory:"false" json:"restEnabledSqlStatus,omitempty"`
}

// GetKey returns Key
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetKey() *string {
	return m.Key
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetDisplayName() *string {
	return m.DisplayName
}

// GetPoolRouteValue returns PoolRouteValue
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetPoolRouteValue() *string {
	return m.PoolRouteValue
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetMaxPoolSize returns MaxPoolSize
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetMaxPoolSize() *int {
	return m.MaxPoolSize
}

// GetMinPoolSize returns MinPoolSize
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetMinPoolSize() *int {
	return m.MinPoolSize
}

// GetInitialPoolSize returns InitialPoolSize
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetInitialPoolSize() *int {
	return m.InitialPoolSize
}

// GetJwtProfileJwkUrl returns JwtProfileJwkUrl
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetJwtProfileJwkUrl() *string {
	return m.JwtProfileJwkUrl
}

// GetJwtProfileIssuer returns JwtProfileIssuer
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetJwtProfileIssuer() *string {
	return m.JwtProfileIssuer
}

// GetJwtProfileAudience returns JwtProfileAudience
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetJwtProfileAudience() *string {
	return m.JwtProfileAudience
}

// GetJwtProfileRoleClaimName returns JwtProfileRoleClaimName
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetJwtProfileRoleClaimName() *string {
	return m.JwtProfileRoleClaimName
}

// GetDatabaseActionsStatus returns DatabaseActionsStatus
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetDatabaseActionsStatus() DatabaseToolsDatabaseApiGatewayConfigPoolContentDatabaseActionsStatusEnum {
	return m.DatabaseActionsStatus
}

// GetRestEnabledSqlStatus returns RestEnabledSqlStatus
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetRestEnabledSqlStatus() DatabaseToolsDatabaseApiGatewayConfigPoolContentRestEnabledSqlStatusEnum {
	return m.RestEnabledSqlStatus
}

// GetAdvancedProperties returns AdvancedProperties
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

// GetApiSpecs returns ApiSpecs
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetApiSpecs() []DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec {
	return m.ApiSpecs
}

// GetAutoApiSpecs returns AutoApiSpecs
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetAutoApiSpecs() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec {
	return m.AutoApiSpecs
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolContentDatabaseActionsStatusEnum(string(m.DatabaseActionsStatus)); !ok && m.DatabaseActionsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseActionsStatus: %s. Supported values are: %s.", m.DatabaseActionsStatus, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolContentDatabaseActionsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolContentRestEnabledSqlStatusEnum(string(m.RestEnabledSqlStatus)); !ok && m.RestEnabledSqlStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RestEnabledSqlStatus: %s. Supported values are: %s.", m.RestEnabledSqlStatus, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolContentRestEnabledSqlStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolContentDefault DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolContentDefault
	}{
		"DEFAULT",
		(MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolContentDefault)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsDatabaseApiGatewayConfigPoolContentDefault) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName               *string                                                                   `json:"displayName"`
		PoolRouteValue            *string                                                                   `json:"poolRouteValue"`
		DatabaseToolsConnectionId *string                                                                   `json:"databaseToolsConnectionId"`
		MaxPoolSize               *int                                                                      `json:"maxPoolSize"`
		MinPoolSize               *int                                                                      `json:"minPoolSize"`
		InitialPoolSize           *int                                                                      `json:"initialPoolSize"`
		JwtProfileJwkUrl          *string                                                                   `json:"jwtProfileJwkUrl"`
		JwtProfileIssuer          *string                                                                   `json:"jwtProfileIssuer"`
		JwtProfileAudience        *string                                                                   `json:"jwtProfileAudience"`
		JwtProfileRoleClaimName   *string                                                                   `json:"jwtProfileRoleClaimName"`
		DatabaseActionsStatus     DatabaseToolsDatabaseApiGatewayConfigPoolContentDatabaseActionsStatusEnum `json:"databaseActionsStatus"`
		RestEnabledSqlStatus      DatabaseToolsDatabaseApiGatewayConfigPoolContentRestEnabledSqlStatusEnum  `json:"restEnabledSqlStatus"`
		AdvancedProperties        map[string]string                                                         `json:"advancedProperties"`
		ApiSpecs                  []databasetoolsdatabaseapigatewayconfigpoolapispec                        `json:"apiSpecs"`
		AutoApiSpecs              []databasetoolsdatabaseapigatewayconfigpoolautoapispec                    `json:"autoApiSpecs"`
		TimeCreated               *common.SDKTime                                                           `json:"timeCreated"`
		TimeUpdated               *common.SDKTime                                                           `json:"timeUpdated"`
		Key                       *string                                                                   `json:"key"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.PoolRouteValue = model.PoolRouteValue

	m.DatabaseToolsConnectionId = model.DatabaseToolsConnectionId

	m.MaxPoolSize = model.MaxPoolSize

	m.MinPoolSize = model.MinPoolSize

	m.InitialPoolSize = model.InitialPoolSize

	m.JwtProfileJwkUrl = model.JwtProfileJwkUrl

	m.JwtProfileIssuer = model.JwtProfileIssuer

	m.JwtProfileAudience = model.JwtProfileAudience

	m.JwtProfileRoleClaimName = model.JwtProfileRoleClaimName

	m.DatabaseActionsStatus = model.DatabaseActionsStatus

	m.RestEnabledSqlStatus = model.RestEnabledSqlStatus

	m.AdvancedProperties = model.AdvancedProperties

	m.ApiSpecs = make([]DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec, len(model.ApiSpecs))
	for i, n := range model.ApiSpecs {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ApiSpecs[i] = nn.(DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec)
		} else {
			m.ApiSpecs[i] = nil
		}
	}
	m.AutoApiSpecs = make([]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec, len(model.AutoApiSpecs))
	for i, n := range model.AutoApiSpecs {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.AutoApiSpecs[i] = nn.(DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec)
		} else {
			m.AutoApiSpecs[i] = nil
		}
	}
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.Key = model.Key

	return
}
