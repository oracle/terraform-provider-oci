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

// DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault The content of a Database Tools database API gateway config auto API spec sub resource defined within a pool.
type DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault struct {

	// A system generated string that uniquely identifies an auto API spec sub resource within a given pool.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the database object.
	DatabaseObjectName *string `mandatory:"false" json:"databaseObjectName"`

	// Description of the autoApiSpec.
	Description *string `mandatory:"false" json:"description"`

	// Used as the URI path element for this object. When not specified the objectName lowercase is the default value.
	Alias *string `mandatory:"false" json:"alias"`

	// The name of the database API gateway config privilege protecting the resource. Only valid for SCOPE JWT Profile pools and BEARER securitySchemes.
	Scope *string `mandatory:"false" json:"scope"`

	// The name of the database API gateway config roles protecting the resource. Only valid for RBAC JWT Profile pools and BEARER securitySchemes.
	Roles []string `mandatory:"false" json:"roles"`

	// The type of the database object.
	DatabaseObjectType DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum `mandatory:"false" json:"databaseObjectType,omitempty"`

	// The operations to limit access to this resource. If not specified then the default is ["READ","WRITE"].
	Operations []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum `mandatory:"false" json:"operations,omitempty"`

	// The security schemes that can access this resource. If not specified then the resource is public.
	SecuritySchemes []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum `mandatory:"false" json:"securitySchemes,omitempty"`
}

// GetKey returns Key
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetKey() *string {
	return m.Key
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseObjectName returns DatabaseObjectName
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetDatabaseObjectName() *string {
	return m.DatabaseObjectName
}

// GetDatabaseObjectType returns DatabaseObjectType
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetDatabaseObjectType() DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum {
	return m.DatabaseObjectType
}

// GetDescription returns Description
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetDescription() *string {
	return m.Description
}

// GetAlias returns Alias
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetAlias() *string {
	return m.Alias
}

// GetOperations returns Operations
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetOperations() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum {
	return m.Operations
}

// GetSecuritySchemes returns SecuritySchemes
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetSecuritySchemes() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum {
	return m.SecuritySchemes
}

// GetScope returns Scope
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetScope() *string {
	return m.Scope
}

// GetRoles returns Roles
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetRoles() []string {
	return m.Roles
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum(string(m.DatabaseObjectType)); !ok && m.DatabaseObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseObjectType: %s. Supported values are: %s.", m.DatabaseObjectType, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnumStringValues(), ",")))
	}
	for _, val := range m.Operations {
		if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operations: %s. Supported values are: %s.", val, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.SecuritySchemes {
		if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecuritySchemes: %s. Supported values are: %s.", val, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault
	}{
		"DEFAULT",
		(MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault)(m),
	}

	return json.Marshal(&s)
}
