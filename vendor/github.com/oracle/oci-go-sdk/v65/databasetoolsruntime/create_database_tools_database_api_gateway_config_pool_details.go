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

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetails The content of a Database Tools database API gateway config pool sub resource to be created.
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetails interface {

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The pool route value provided in requests to target this pool.
	GetPoolRouteValue() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection. Specifies the OCI database tools connection ocid to build the connection pool from.
	GetDatabaseToolsConnectionId() *string

	// Specifies the maximum number of database connections allowed for the pool.
	GetMaxPoolSize() *int

	// Specifies the minimum number of database connections allowed for the pool.
	GetMinPoolSize() *int

	// Specifies the initial size for the number of database connections that will be created for the pool.
	GetInitialPoolSize() *int

	// Specifies the URL of the JSON Web Key (JWK) that is used to verify the signature of the JWT token.
	GetJwtProfileJwkUrl() *string

	// Specifies the issuer of the JWT token. This value is used to validate the iss claim in the JWT token.
	GetJwtProfileIssuer() *string

	// Specifies the expected audience for the JWT token. This value is used to validate the aud claim in the JWT token.
	GetJwtProfileAudience() *string

	// Specifies the JSON pointer to the claim in the JWT token that contains the roles of the users.
	GetJwtProfileRoleClaimName() *string

	// Specifies to enable the Database Actions feature.
	GetDatabaseActionsStatus() CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum

	// Specifies whether the REST-Enabled SQL service is active.
	GetRestEnabledSqlStatus() CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum

	// Advanced pool properties.
	GetAdvancedProperties() map[string]string
}

type createdatabasetoolsdatabaseapigatewayconfigpooldetails struct {
	JsonData                  []byte
	MaxPoolSize               *int                                                                            `mandatory:"false" json:"maxPoolSize"`
	MinPoolSize               *int                                                                            `mandatory:"false" json:"minPoolSize"`
	InitialPoolSize           *int                                                                            `mandatory:"false" json:"initialPoolSize"`
	JwtProfileJwkUrl          *string                                                                         `mandatory:"false" json:"jwtProfileJwkUrl"`
	JwtProfileIssuer          *string                                                                         `mandatory:"false" json:"jwtProfileIssuer"`
	JwtProfileAudience        *string                                                                         `mandatory:"false" json:"jwtProfileAudience"`
	JwtProfileRoleClaimName   *string                                                                         `mandatory:"false" json:"jwtProfileRoleClaimName"`
	DatabaseActionsStatus     CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum `mandatory:"false" json:"databaseActionsStatus,omitempty"`
	RestEnabledSqlStatus      CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum  `mandatory:"false" json:"restEnabledSqlStatus,omitempty"`
	AdvancedProperties        map[string]string                                                               `mandatory:"false" json:"advancedProperties"`
	DisplayName               *string                                                                         `mandatory:"true" json:"displayName"`
	PoolRouteValue            *string                                                                         `mandatory:"true" json:"poolRouteValue"`
	DatabaseToolsConnectionId *string                                                                         `mandatory:"true" json:"databaseToolsConnectionId"`
	Type                      string                                                                          `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsdatabaseapigatewayconfigpooldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsdatabaseapigatewayconfigpooldetails createdatabasetoolsdatabaseapigatewayconfigpooldetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsdatabaseapigatewayconfigpooldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.PoolRouteValue = s.Model.PoolRouteValue
	m.DatabaseToolsConnectionId = s.Model.DatabaseToolsConnectionId
	m.MaxPoolSize = s.Model.MaxPoolSize
	m.MinPoolSize = s.Model.MinPoolSize
	m.InitialPoolSize = s.Model.InitialPoolSize
	m.JwtProfileJwkUrl = s.Model.JwtProfileJwkUrl
	m.JwtProfileIssuer = s.Model.JwtProfileIssuer
	m.JwtProfileAudience = s.Model.JwtProfileAudience
	m.JwtProfileRoleClaimName = s.Model.JwtProfileRoleClaimName
	m.DatabaseActionsStatus = s.Model.DatabaseActionsStatus
	m.RestEnabledSqlStatus = s.Model.RestEnabledSqlStatus
	m.AdvancedProperties = s.Model.AdvancedProperties
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsdatabaseapigatewayconfigpooldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := CreateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetMaxPoolSize returns MaxPoolSize
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetMaxPoolSize() *int {
	return m.MaxPoolSize
}

// GetMinPoolSize returns MinPoolSize
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetMinPoolSize() *int {
	return m.MinPoolSize
}

// GetInitialPoolSize returns InitialPoolSize
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetInitialPoolSize() *int {
	return m.InitialPoolSize
}

// GetJwtProfileJwkUrl returns JwtProfileJwkUrl
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetJwtProfileJwkUrl() *string {
	return m.JwtProfileJwkUrl
}

// GetJwtProfileIssuer returns JwtProfileIssuer
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetJwtProfileIssuer() *string {
	return m.JwtProfileIssuer
}

// GetJwtProfileAudience returns JwtProfileAudience
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetJwtProfileAudience() *string {
	return m.JwtProfileAudience
}

// GetJwtProfileRoleClaimName returns JwtProfileRoleClaimName
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetJwtProfileRoleClaimName() *string {
	return m.JwtProfileRoleClaimName
}

// GetDatabaseActionsStatus returns DatabaseActionsStatus
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetDatabaseActionsStatus() CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum {
	return m.DatabaseActionsStatus
}

// GetRestEnabledSqlStatus returns RestEnabledSqlStatus
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetRestEnabledSqlStatus() CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum {
	return m.RestEnabledSqlStatus
}

// GetAdvancedProperties returns AdvancedProperties
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetPoolRouteValue returns PoolRouteValue
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetPoolRouteValue() *string {
	return m.PoolRouteValue
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolsdatabaseapigatewayconfigpooldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum(string(m.DatabaseActionsStatus)); !ok && m.DatabaseActionsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseActionsStatus: %s. Supported values are: %s.", m.DatabaseActionsStatus, strings.Join(GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum(string(m.RestEnabledSqlStatus)); !ok && m.RestEnabledSqlStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RestEnabledSqlStatus: %s. Supported values are: %s.", m.RestEnabledSqlStatus, strings.Join(GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum Enum with underlying type: string
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum string

// Set of constants representing the allowable values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum
const (
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnabled  CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum = "ENABLED"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusDisabled CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum = "DISABLED"
)

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum{
	"ENABLED":  CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnabled,
	"DISABLED": CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusDisabled,
}

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumLowerCase = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum{
	"enabled":  CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnabled,
	"disabled": CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusDisabled,
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumValues Enumerates the set of values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumValues() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum {
	values := make([]CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum, 0)
	for _, v := range mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumStringValues Enumerates the set of values in String for CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum(val string) (CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum, bool) {
	enum, ok := mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum Enum with underlying type: string
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum string

// Set of constants representing the allowable values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum
const (
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnabled  CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum = "ENABLED"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusDisabled CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum = "DISABLED"
)

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum{
	"ENABLED":  CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnabled,
	"DISABLED": CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusDisabled,
}

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumLowerCase = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum{
	"enabled":  CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnabled,
	"disabled": CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusDisabled,
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumValues Enumerates the set of values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumValues() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum {
	values := make([]CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum, 0)
	for _, v := range mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumStringValues Enumerates the set of values in String for CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum(val string) (CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum, bool) {
	enum, ok := mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
