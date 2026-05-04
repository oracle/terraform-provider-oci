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

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails The content of a Database Tools database API gateway config pool sub resource to be updated.
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails struct {

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
	DatabaseActionsStatus UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum `mandatory:"false" json:"databaseActionsStatus,omitempty"`

	// Specifies whether the REST-Enabled SQL service is active.
	RestEnabledSqlStatus UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum `mandatory:"false" json:"restEnabledSqlStatus,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetPoolRouteValue returns PoolRouteValue
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetPoolRouteValue() *string {
	return m.PoolRouteValue
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetMaxPoolSize returns MaxPoolSize
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetMaxPoolSize() *int {
	return m.MaxPoolSize
}

// GetMinPoolSize returns MinPoolSize
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetMinPoolSize() *int {
	return m.MinPoolSize
}

// GetInitialPoolSize returns InitialPoolSize
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetInitialPoolSize() *int {
	return m.InitialPoolSize
}

// GetJwtProfileJwkUrl returns JwtProfileJwkUrl
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetJwtProfileJwkUrl() *string {
	return m.JwtProfileJwkUrl
}

// GetJwtProfileIssuer returns JwtProfileIssuer
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetJwtProfileIssuer() *string {
	return m.JwtProfileIssuer
}

// GetJwtProfileAudience returns JwtProfileAudience
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetJwtProfileAudience() *string {
	return m.JwtProfileAudience
}

// GetJwtProfileRoleClaimName returns JwtProfileRoleClaimName
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetJwtProfileRoleClaimName() *string {
	return m.JwtProfileRoleClaimName
}

// GetDatabaseActionsStatus returns DatabaseActionsStatus
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetDatabaseActionsStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum {
	return m.DatabaseActionsStatus
}

// GetRestEnabledSqlStatus returns RestEnabledSqlStatus
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetRestEnabledSqlStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum {
	return m.RestEnabledSqlStatus
}

// GetAdvancedProperties returns AdvancedProperties
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails
	}{
		"DEFAULT",
		(MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails)(m),
	}

	return json.Marshal(&s)
}
