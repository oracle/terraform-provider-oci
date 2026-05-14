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

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetails The content of a Database Tools database API gateway config pool sub resource to be updated.
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetails interface {

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
	GetDatabaseActionsStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum

	// Specifies whether the REST-Enabled SQL service is active.
	GetRestEnabledSqlStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum

	// Advanced pool properties.
	GetAdvancedProperties() map[string]string
}

type updatedatabasetoolsdatabaseapigatewayconfigpooldetails struct {
	JsonData                  []byte
	DisplayName               *string                                                                         `mandatory:"false" json:"displayName"`
	PoolRouteValue            *string                                                                         `mandatory:"false" json:"poolRouteValue"`
	DatabaseToolsConnectionId *string                                                                         `mandatory:"false" json:"databaseToolsConnectionId"`
	MaxPoolSize               *int                                                                            `mandatory:"false" json:"maxPoolSize"`
	MinPoolSize               *int                                                                            `mandatory:"false" json:"minPoolSize"`
	InitialPoolSize           *int                                                                            `mandatory:"false" json:"initialPoolSize"`
	JwtProfileJwkUrl          *string                                                                         `mandatory:"false" json:"jwtProfileJwkUrl"`
	JwtProfileIssuer          *string                                                                         `mandatory:"false" json:"jwtProfileIssuer"`
	JwtProfileAudience        *string                                                                         `mandatory:"false" json:"jwtProfileAudience"`
	JwtProfileRoleClaimName   *string                                                                         `mandatory:"false" json:"jwtProfileRoleClaimName"`
	DatabaseActionsStatus     UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum `mandatory:"false" json:"databaseActionsStatus,omitempty"`
	RestEnabledSqlStatus      UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum  `mandatory:"false" json:"restEnabledSqlStatus,omitempty"`
	AdvancedProperties        map[string]string                                                               `mandatory:"false" json:"advancedProperties"`
	Type                      string                                                                          `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatedatabasetoolsdatabaseapigatewayconfigpooldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedatabasetoolsdatabaseapigatewayconfigpooldetails updatedatabasetoolsdatabaseapigatewayconfigpooldetails
	s := struct {
		Model Unmarshalerupdatedatabasetoolsdatabaseapigatewayconfigpooldetails
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
func (m *updatedatabasetoolsdatabaseapigatewayconfigpooldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetPoolRouteValue returns PoolRouteValue
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetPoolRouteValue() *string {
	return m.PoolRouteValue
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetMaxPoolSize returns MaxPoolSize
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetMaxPoolSize() *int {
	return m.MaxPoolSize
}

// GetMinPoolSize returns MinPoolSize
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetMinPoolSize() *int {
	return m.MinPoolSize
}

// GetInitialPoolSize returns InitialPoolSize
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetInitialPoolSize() *int {
	return m.InitialPoolSize
}

// GetJwtProfileJwkUrl returns JwtProfileJwkUrl
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetJwtProfileJwkUrl() *string {
	return m.JwtProfileJwkUrl
}

// GetJwtProfileIssuer returns JwtProfileIssuer
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetJwtProfileIssuer() *string {
	return m.JwtProfileIssuer
}

// GetJwtProfileAudience returns JwtProfileAudience
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetJwtProfileAudience() *string {
	return m.JwtProfileAudience
}

// GetJwtProfileRoleClaimName returns JwtProfileRoleClaimName
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetJwtProfileRoleClaimName() *string {
	return m.JwtProfileRoleClaimName
}

// GetDatabaseActionsStatus returns DatabaseActionsStatus
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetDatabaseActionsStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum {
	return m.DatabaseActionsStatus
}

// GetRestEnabledSqlStatus returns RestEnabledSqlStatus
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetRestEnabledSqlStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum {
	return m.RestEnabledSqlStatus
}

// GetAdvancedProperties returns AdvancedProperties
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedatabasetoolsdatabaseapigatewayconfigpooldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum(string(m.DatabaseActionsStatus)); !ok && m.DatabaseActionsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseActionsStatus: %s. Supported values are: %s.", m.DatabaseActionsStatus, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum(string(m.RestEnabledSqlStatus)); !ok && m.RestEnabledSqlStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RestEnabledSqlStatus: %s. Supported values are: %s.", m.RestEnabledSqlStatus, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum Enum with underlying type: string
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum string

// Set of constants representing the allowable values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum
const (
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnabled  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum = "ENABLED"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusDisabled UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum = "DISABLED"
)

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum{
	"ENABLED":  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnabled,
	"DISABLED": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusDisabled,
}

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumLowerCase = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum{
	"enabled":  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnabled,
	"disabled": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusDisabled,
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumValues Enumerates the set of values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumValues() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum {
	values := make([]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum, 0)
	for _, v := range mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumStringValues Enumerates the set of values in String for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum(val string) (UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum, bool) {
	enum, ok := mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum Enum with underlying type: string
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum string

// Set of constants representing the allowable values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum
const (
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnabled  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum = "ENABLED"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusDisabled UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum = "DISABLED"
)

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum{
	"ENABLED":  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnabled,
	"DISABLED": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusDisabled,
}

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumLowerCase = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum{
	"enabled":  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnabled,
	"disabled": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusDisabled,
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumValues Enumerates the set of values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumValues() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum {
	values := make([]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum, 0)
	for _, v := range mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumStringValues Enumerates the set of values in String for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum(val string) (UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum, bool) {
	enum, ok := mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
