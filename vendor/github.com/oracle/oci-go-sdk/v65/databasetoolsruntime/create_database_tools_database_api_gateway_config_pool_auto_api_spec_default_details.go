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

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails The content of a Database Tools database API gateway config auto API spec sub resource to be created.
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails struct {

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the database object.
	DatabaseObjectName *string `mandatory:"true" json:"databaseObjectName"`

	// Description of the autoApiSpec.
	Description *string `mandatory:"false" json:"description"`

	// Used as the URI path element for this object. When not specified the objectName lowercase is the default value.
	Alias *string `mandatory:"false" json:"alias"`

	// The name of the database API gateway config privilege protecting the resource. Only valid for SCOPE JWT Profile pools and BEARER securitySchemes.
	Scope *string `mandatory:"false" json:"scope"`

	// The name of the database API gateway config roles protecting the resource. Only valid for RBAC JWT Profile pools and BEARER securitySchemes.
	Roles []string `mandatory:"false" json:"roles"`

	// The type of the database object.
	DatabaseObjectType CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum `mandatory:"true" json:"databaseObjectType"`

	// The operations to limit access to this resource. If not specified then the default is ["READ","WRITE"].
	Operations []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum `mandatory:"false" json:"operations,omitempty"`

	// The security schemes that can access this resource. If not specified then the resource is public.
	SecuritySchemes []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum `mandatory:"false" json:"securitySchemes,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseObjectName returns DatabaseObjectName
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetDatabaseObjectName() *string {
	return m.DatabaseObjectName
}

// GetDatabaseObjectType returns DatabaseObjectType
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetDatabaseObjectType() CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum {
	return m.DatabaseObjectType
}

// GetDescription returns Description
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetDescription() *string {
	return m.Description
}

// GetAlias returns Alias
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetAlias() *string {
	return m.Alias
}

// GetOperations returns Operations
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetOperations() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum {
	return m.Operations
}

// GetSecuritySchemes returns SecuritySchemes
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetSecuritySchemes() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum {
	return m.SecuritySchemes
}

// GetScope returns Scope
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetScope() *string {
	return m.Scope
}

// GetRoles returns Roles
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetRoles() []string {
	return m.Roles
}

func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum(string(m.DatabaseObjectType)); !ok && m.DatabaseObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseObjectType: %s. Supported values are: %s.", m.DatabaseObjectType, strings.Join(GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumStringValues(), ",")))
	}
	for _, val := range m.Operations {
		if _, ok := GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operations: %s. Supported values are: %s.", val, strings.Join(GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.SecuritySchemes {
		if _, ok := GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecuritySchemes: %s. Supported values are: %s.", val, strings.Join(GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails
	}{
		"DEFAULT",
		(MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails)(m),
	}

	return json.Marshal(&s)
}
