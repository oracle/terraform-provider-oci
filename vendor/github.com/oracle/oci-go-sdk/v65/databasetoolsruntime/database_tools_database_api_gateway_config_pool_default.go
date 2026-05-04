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

// DatabaseToolsDatabaseApiGatewayConfigPoolDefault The content of a Database Tools database API gateway config pool sub resource.
type DatabaseToolsDatabaseApiGatewayConfigPoolDefault struct {

	// A system generated string that uniquely identifies a pool sub resource.
	Key *string `mandatory:"true" json:"key"`

	// The time the resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

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

	// Specifies to enable the Database Actions feature.
	DatabaseActionsStatus DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum `mandatory:"false" json:"databaseActionsStatus,omitempty"`

	// Specifies whether the REST-Enabled SQL service is active.
	RestEnabledSqlStatus DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum `mandatory:"false" json:"restEnabledSqlStatus,omitempty"`
}

// GetKey returns Key
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetKey() *string {
	return m.Key
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetDisplayName() *string {
	return m.DisplayName
}

// GetPoolRouteValue returns PoolRouteValue
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetPoolRouteValue() *string {
	return m.PoolRouteValue
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetMaxPoolSize returns MaxPoolSize
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetMaxPoolSize() *int {
	return m.MaxPoolSize
}

// GetMinPoolSize returns MinPoolSize
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetMinPoolSize() *int {
	return m.MinPoolSize
}

// GetInitialPoolSize returns InitialPoolSize
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetInitialPoolSize() *int {
	return m.InitialPoolSize
}

// GetJwtProfileJwkUrl returns JwtProfileJwkUrl
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetJwtProfileJwkUrl() *string {
	return m.JwtProfileJwkUrl
}

// GetJwtProfileIssuer returns JwtProfileIssuer
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetJwtProfileIssuer() *string {
	return m.JwtProfileIssuer
}

// GetJwtProfileAudience returns JwtProfileAudience
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetJwtProfileAudience() *string {
	return m.JwtProfileAudience
}

// GetJwtProfileRoleClaimName returns JwtProfileRoleClaimName
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetJwtProfileRoleClaimName() *string {
	return m.JwtProfileRoleClaimName
}

// GetDatabaseActionsStatus returns DatabaseActionsStatus
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetDatabaseActionsStatus() DatabaseToolsDatabaseApiGatewayConfigPoolDatabaseActionsStatusEnum {
	return m.DatabaseActionsStatus
}

// GetRestEnabledSqlStatus returns RestEnabledSqlStatus
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetRestEnabledSqlStatus() DatabaseToolsDatabaseApiGatewayConfigPoolRestEnabledSqlStatusEnum {
	return m.RestEnabledSqlStatus
}

// GetAdvancedProperties returns AdvancedProperties
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolDefault DatabaseToolsDatabaseApiGatewayConfigPoolDefault
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolDefault
	}{
		"DEFAULT",
		(MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolDefault)(m),
	}

	return json.Marshal(&s)
}
