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

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails The content of a Database Tools database API gateway config auto API spec sub resource to be updated.
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails struct {

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

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
	DatabaseObjectType UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum `mandatory:"false" json:"databaseObjectType,omitempty"`

	// The operations to limit access to this resource. If not specified then the default is ["READ","WRITE"].
	Operations []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum `mandatory:"false" json:"operations,omitempty"`

	// The security schemes that can access this resource. If not specified then the resource is public.
	SecuritySchemes []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum `mandatory:"false" json:"securitySchemes,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseObjectName returns DatabaseObjectName
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetDatabaseObjectName() *string {
	return m.DatabaseObjectName
}

// GetDatabaseObjectType returns DatabaseObjectType
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetDatabaseObjectType() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum {
	return m.DatabaseObjectType
}

// GetDescription returns Description
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetDescription() *string {
	return m.Description
}

// GetAlias returns Alias
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetAlias() *string {
	return m.Alias
}

// GetOperations returns Operations
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetOperations() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum {
	return m.Operations
}

// GetSecuritySchemes returns SecuritySchemes
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetSecuritySchemes() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum {
	return m.SecuritySchemes
}

// GetScope returns Scope
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetScope() *string {
	return m.Scope
}

// GetRoles returns Roles
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) GetRoles() []string {
	return m.Roles
}

func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum(string(m.DatabaseObjectType)); !ok && m.DatabaseObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseObjectType: %s. Supported values are: %s.", m.DatabaseObjectType, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumStringValues(), ",")))
	}
	for _, val := range m.Operations {
		if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operations: %s. Supported values are: %s.", val, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.SecuritySchemes {
		if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecuritySchemes: %s. Supported values are: %s.", val, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails
	}{
		"DEFAULT",
		(MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails)(m),
	}

	return json.Marshal(&s)
}
