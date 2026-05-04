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

// DatabaseToolsDatabaseApiGatewayConfigPool The content of a Database Tools database API gateway config pool sub resource.
type DatabaseToolsDatabaseApiGatewayConfigPool interface {

	// A system generated string that uniquely identifies a pool sub resource.
	GetKey() *string

	// The time the resource was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the resource was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

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
	GetDatabaseActionsStatus() DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum

	// Specifies whether the REST-Enabled SQL service is active.
	GetRestEnabledSqlStatus() DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum

	// Advanced pool properties.
	GetAdvancedProperties() map[string]string
}

type databasetoolsdatabaseapigatewayconfigpool struct {
	JsonData                  []byte
	DisplayName               *string                                                            `mandatory:"false" json:"displayName"`
	PoolRouteValue            *string                                                            `mandatory:"false" json:"poolRouteValue"`
	DatabaseToolsConnectionId *string                                                            `mandatory:"false" json:"databaseToolsConnectionId"`
	MaxPoolSize               *int                                                               `mandatory:"false" json:"maxPoolSize"`
	MinPoolSize               *int                                                               `mandatory:"false" json:"minPoolSize"`
	InitialPoolSize           *int                                                               `mandatory:"false" json:"initialPoolSize"`
	JwtProfileJwkUrl          *string                                                            `mandatory:"false" json:"jwtProfileJwkUrl"`
	JwtProfileIssuer          *string                                                            `mandatory:"false" json:"jwtProfileIssuer"`
	JwtProfileAudience        *string                                                            `mandatory:"false" json:"jwtProfileAudience"`
	JwtProfileRoleClaimName   *string                                                            `mandatory:"false" json:"jwtProfileRoleClaimName"`
	DatabaseActionsStatus     DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum `mandatory:"false" json:"databaseActionsStatus,omitempty"`
	RestEnabledSqlStatus      DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum  `mandatory:"false" json:"restEnabledSqlStatus,omitempty"`
	AdvancedProperties        map[string]string                                                  `mandatory:"false" json:"advancedProperties"`
	Key                       *string                                                            `mandatory:"true" json:"key"`
	TimeCreated               *common.SDKTime                                                    `mandatory:"true" json:"timeCreated"`
	TimeUpdated               *common.SDKTime                                                    `mandatory:"true" json:"timeUpdated"`
	Type                      string                                                             `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsdatabaseapigatewayconfigpool) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsdatabaseapigatewayconfigpool databasetoolsdatabaseapigatewayconfigpool
	s := struct {
		Model Unmarshalerdatabasetoolsdatabaseapigatewayconfigpool
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
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
func (m *databasetoolsdatabaseapigatewayconfigpool) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DatabaseToolsDatabaseApiGatewayConfigPoolDefault{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsDatabaseApiGatewayConfigPool: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m databasetoolsdatabaseapigatewayconfigpool) GetDisplayName() *string {
	return m.DisplayName
}

// GetPoolRouteValue returns PoolRouteValue
func (m databasetoolsdatabaseapigatewayconfigpool) GetPoolRouteValue() *string {
	return m.PoolRouteValue
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m databasetoolsdatabaseapigatewayconfigpool) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetMaxPoolSize returns MaxPoolSize
func (m databasetoolsdatabaseapigatewayconfigpool) GetMaxPoolSize() *int {
	return m.MaxPoolSize
}

// GetMinPoolSize returns MinPoolSize
func (m databasetoolsdatabaseapigatewayconfigpool) GetMinPoolSize() *int {
	return m.MinPoolSize
}

// GetInitialPoolSize returns InitialPoolSize
func (m databasetoolsdatabaseapigatewayconfigpool) GetInitialPoolSize() *int {
	return m.InitialPoolSize
}

// GetJwtProfileJwkUrl returns JwtProfileJwkUrl
func (m databasetoolsdatabaseapigatewayconfigpool) GetJwtProfileJwkUrl() *string {
	return m.JwtProfileJwkUrl
}

// GetJwtProfileIssuer returns JwtProfileIssuer
func (m databasetoolsdatabaseapigatewayconfigpool) GetJwtProfileIssuer() *string {
	return m.JwtProfileIssuer
}

// GetJwtProfileAudience returns JwtProfileAudience
func (m databasetoolsdatabaseapigatewayconfigpool) GetJwtProfileAudience() *string {
	return m.JwtProfileAudience
}

// GetJwtProfileRoleClaimName returns JwtProfileRoleClaimName
func (m databasetoolsdatabaseapigatewayconfigpool) GetJwtProfileRoleClaimName() *string {
	return m.JwtProfileRoleClaimName
}

// GetDatabaseActionsStatus returns DatabaseActionsStatus
func (m databasetoolsdatabaseapigatewayconfigpool) GetDatabaseActionsStatus() DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum {
	return m.DatabaseActionsStatus
}

// GetRestEnabledSqlStatus returns RestEnabledSqlStatus
func (m databasetoolsdatabaseapigatewayconfigpool) GetRestEnabledSqlStatus() DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum {
	return m.RestEnabledSqlStatus
}

// GetAdvancedProperties returns AdvancedProperties
func (m databasetoolsdatabaseapigatewayconfigpool) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

// GetKey returns Key
func (m databasetoolsdatabaseapigatewayconfigpool) GetKey() *string {
	return m.Key
}

// GetTimeCreated returns TimeCreated
func (m databasetoolsdatabaseapigatewayconfigpool) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolsdatabaseapigatewayconfigpool) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m databasetoolsdatabaseapigatewayconfigpool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsdatabaseapigatewayconfigpool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum(string(m.DatabaseActionsStatus)); !ok && m.DatabaseActionsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseActionsStatus: %s. Supported values are: %s.", m.DatabaseActionsStatus, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum(string(m.RestEnabledSqlStatus)); !ok && m.RestEnabledSqlStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RestEnabledSqlStatus: %s. Supported values are: %s.", m.RestEnabledSqlStatus, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnabled  DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum = "ENABLED"
	DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusDisabled DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum = "DISABLED"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum{
	"ENABLED":  DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnabled,
	"DISABLED": DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusDisabled,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum{
	"enabled":  DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnabled,
	"disabled": DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusDisabled,
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnumValues() []DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnabled  DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum = "ENABLED"
	DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusDisabled DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum = "DISABLED"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum{
	"ENABLED":  DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnabled,
	"DISABLED": DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusDisabled,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum{
	"enabled":  DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnabled,
	"disabled": DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusDisabled,
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnumValues() []DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
