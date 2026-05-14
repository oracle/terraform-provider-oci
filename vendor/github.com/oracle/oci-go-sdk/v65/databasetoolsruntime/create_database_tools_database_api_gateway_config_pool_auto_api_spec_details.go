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

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails The content of a Database Tools database API gateway config auto API spec sub resource to be created.
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails interface {

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The name of the database object.
	GetDatabaseObjectName() *string

	// The type of the database object.
	GetDatabaseObjectType() CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum

	// Description of the autoApiSpec.
	GetDescription() *string

	// Used as the URI path element for this object. When not specified the objectName lowercase is the default value.
	GetAlias() *string

	// The operations to limit access to this resource. If not specified then the default is ["READ","WRITE"].
	GetOperations() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum

	// The security schemes that can access this resource. If not specified then the resource is public.
	GetSecuritySchemes() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum

	// The name of the database API gateway config privilege protecting the resource. Only valid for SCOPE JWT Profile pools and BEARER securitySchemes.
	GetScope() *string

	// The name of the database API gateway config roles protecting the resource. Only valid for RBAC JWT Profile pools and BEARER securitySchemes.
	GetRoles() []string
}

type createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails struct {
	JsonData           []byte
	Description        *string                                                                                 `mandatory:"false" json:"description"`
	Alias              *string                                                                                 `mandatory:"false" json:"alias"`
	Operations         []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum       `mandatory:"false" json:"operations,omitempty"`
	SecuritySchemes    []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum  `mandatory:"false" json:"securitySchemes,omitempty"`
	Scope              *string                                                                                 `mandatory:"false" json:"scope"`
	Roles              []string                                                                                `mandatory:"false" json:"roles"`
	DisplayName        *string                                                                                 `mandatory:"true" json:"displayName"`
	DatabaseObjectName *string                                                                                 `mandatory:"true" json:"databaseObjectName"`
	DatabaseObjectType CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum `mandatory:"true" json:"databaseObjectType"`
	Type               string                                                                                  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.DatabaseObjectName = s.Model.DatabaseObjectName
	m.DatabaseObjectType = s.Model.DatabaseObjectType
	m.Description = s.Model.Description
	m.Alias = s.Model.Alias
	m.Operations = s.Model.Operations
	m.SecuritySchemes = s.Model.SecuritySchemes
	m.Scope = s.Model.Scope
	m.Roles = s.Model.Roles
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetDescription() *string {
	return m.Description
}

// GetAlias returns Alias
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetAlias() *string {
	return m.Alias
}

// GetOperations returns Operations
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetOperations() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum {
	return m.Operations
}

// GetSecuritySchemes returns SecuritySchemes
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetSecuritySchemes() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum {
	return m.SecuritySchemes
}

// GetScope returns Scope
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetScope() *string {
	return m.Scope
}

// GetRoles returns Roles
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetRoles() []string {
	return m.Roles
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseObjectName returns DatabaseObjectName
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetDatabaseObjectName() *string {
	return m.DatabaseObjectName
}

// GetDatabaseObjectType returns DatabaseObjectType
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetDatabaseObjectType() CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum {
	return m.DatabaseObjectType
}

func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum(string(m.DatabaseObjectType)); !ok && m.DatabaseObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseObjectType: %s. Supported values are: %s.", m.DatabaseObjectType, strings.Join(GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumStringValues(), ",")))
	}

	for _, val := range m.Operations {
		if _, ok := GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum(string(val)); !ok && string(val) != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operations: %s. Supported values are: %s.", val, strings.Join(GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.SecuritySchemes {
		if _, ok := GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum(string(val)); !ok && string(val) != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecuritySchemes: %s. Supported values are: %s.", val, strings.Join(GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum Enum with underlying type: string
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum string

// Set of constants representing the allowable values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum
const (
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeFunction    CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "FUNCTION"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeMview       CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "MVIEW"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypePackage     CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "PACKAGE"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeProcedure   CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "PROCEDURE"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeTable       CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "TABLE"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeView        CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "VIEW"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeDualityview CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "DUALITYVIEW"
)

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum{
	"FUNCTION":    CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeFunction,
	"MVIEW":       CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeMview,
	"PACKAGE":     CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypePackage,
	"PROCEDURE":   CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeProcedure,
	"TABLE":       CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeTable,
	"VIEW":        CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeView,
	"DUALITYVIEW": CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeDualityview,
}

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumLowerCase = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum{
	"function":    CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeFunction,
	"mview":       CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeMview,
	"package":     CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypePackage,
	"procedure":   CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeProcedure,
	"table":       CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeTable,
	"view":        CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeView,
	"dualityview": CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeDualityview,
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumValues Enumerates the set of values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumValues() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum {
	values := make([]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum, 0)
	for _, v := range mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumStringValues Enumerates the set of values in String for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumStringValues() []string {
	return []string{
		"FUNCTION",
		"MVIEW",
		"PACKAGE",
		"PROCEDURE",
		"TABLE",
		"VIEW",
		"DUALITYVIEW",
	}
}

// GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum(val string) (CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum, bool) {
	enum, ok := mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum Enum with underlying type: string
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum string

// Set of constants representing the allowable values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum
const (
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsRead  CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum = "READ"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsWrite CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum = "WRITE"
)

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum{
	"READ":  CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsRead,
	"WRITE": CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsWrite,
}

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumLowerCase = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum{
	"read":  CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsRead,
	"write": CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsWrite,
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumValues Enumerates the set of values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumValues() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum {
	values := make([]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum, 0)
	for _, v := range mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumStringValues Enumerates the set of values in String for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumStringValues() []string {
	return []string{
		"READ",
		"WRITE",
	}
}

// GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum(val string) (CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum, bool) {
	enum, ok := mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum Enum with underlying type: string
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum string

// Set of constants representing the allowable values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum
const (
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBasic  CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum = "BASIC"
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBearer CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum = "BEARER"
)

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum{
	"BASIC":  CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBasic,
	"BEARER": CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBearer,
}

var mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumLowerCase = map[string]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum{
	"basic":  CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBasic,
	"bearer": CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBearer,
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumValues Enumerates the set of values for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumValues() []CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum {
	values := make([]CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum, 0)
	for _, v := range mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumStringValues Enumerates the set of values in String for CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum
func GetCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumStringValues() []string {
	return []string{
		"BASIC",
		"BEARER",
	}
}

// GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum(val string) (CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum, bool) {
	enum, ok := mappingCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
