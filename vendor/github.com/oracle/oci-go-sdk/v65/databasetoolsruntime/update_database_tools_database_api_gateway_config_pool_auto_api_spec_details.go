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

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails The content of a Database Tools database API gateway config auto API spec sub resource to be updated.
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails interface {

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The name of the database object.
	GetDatabaseObjectName() *string

	// The type of the database object.
	GetDatabaseObjectType() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum

	// Description of the autoApiSpec.
	GetDescription() *string

	// Used as the URI path element for this object. When not specified the objectName lowercase is the default value.
	GetAlias() *string

	// The operations to limit access to this resource. If not specified then the default is ["READ","WRITE"].
	GetOperations() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum

	// The security schemes that can access this resource. If not specified then the resource is public.
	GetSecuritySchemes() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum

	// The name of the database API gateway config privilege protecting the resource. Only valid for SCOPE JWT Profile pools and BEARER securitySchemes.
	GetScope() *string

	// The name of the database API gateway config roles protecting the resource. Only valid for RBAC JWT Profile pools and BEARER securitySchemes.
	GetRoles() []string
}

type updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails struct {
	JsonData           []byte
	DisplayName        *string                                                                                 `mandatory:"false" json:"displayName"`
	DatabaseObjectName *string                                                                                 `mandatory:"false" json:"databaseObjectName"`
	DatabaseObjectType UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum `mandatory:"false" json:"databaseObjectType,omitempty"`
	Description        *string                                                                                 `mandatory:"false" json:"description"`
	Alias              *string                                                                                 `mandatory:"false" json:"alias"`
	Operations         []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum       `mandatory:"false" json:"operations,omitempty"`
	SecuritySchemes    []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum  `mandatory:"false" json:"securitySchemes,omitempty"`
	Scope              *string                                                                                 `mandatory:"false" json:"scope"`
	Roles              []string                                                                                `mandatory:"false" json:"roles"`
	Type               string                                                                                  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails
	s := struct {
		Model Unmarshalerupdatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails
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
func (m *updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseObjectName returns DatabaseObjectName
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetDatabaseObjectName() *string {
	return m.DatabaseObjectName
}

// GetDatabaseObjectType returns DatabaseObjectType
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetDatabaseObjectType() UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum {
	return m.DatabaseObjectType
}

// GetDescription returns Description
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetDescription() *string {
	return m.Description
}

// GetAlias returns Alias
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetAlias() *string {
	return m.Alias
}

// GetOperations returns Operations
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetOperations() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum {
	return m.Operations
}

// GetSecuritySchemes returns SecuritySchemes
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetSecuritySchemes() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum {
	return m.SecuritySchemes
}

// GetScope returns Scope
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetScope() *string {
	return m.Scope
}

// GetRoles returns Roles
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) GetRoles() []string {
	return m.Roles
}

func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedatabasetoolsdatabaseapigatewayconfigpoolautoapispecdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum(string(m.DatabaseObjectType)); !ok && m.DatabaseObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseObjectType: %s. Supported values are: %s.", m.DatabaseObjectType, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumStringValues(), ",")))
	}
	for _, val := range m.Operations {
		if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum(string(val)); !ok && string(val) != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operations: %s. Supported values are: %s.", val, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.SecuritySchemes {
		if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum(string(val)); !ok && string(val) != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecuritySchemes: %s. Supported values are: %s.", val, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum Enum with underlying type: string
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum string

// Set of constants representing the allowable values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum
const (
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeFunction    UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "FUNCTION"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeMview       UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "MVIEW"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypePackage     UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "PACKAGE"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeProcedure   UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "PROCEDURE"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeTable       UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "TABLE"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeView        UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "VIEW"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeDualityview UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = "DUALITYVIEW"
)

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum{
	"FUNCTION":    UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeFunction,
	"MVIEW":       UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeMview,
	"PACKAGE":     UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypePackage,
	"PROCEDURE":   UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeProcedure,
	"TABLE":       UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeTable,
	"VIEW":        UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeView,
	"DUALITYVIEW": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeDualityview,
}

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumLowerCase = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum{
	"function":    UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeFunction,
	"mview":       UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeMview,
	"package":     UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypePackage,
	"procedure":   UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeProcedure,
	"table":       UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeTable,
	"view":        UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeView,
	"dualityview": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeDualityview,
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumValues Enumerates the set of values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumValues() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum {
	values := make([]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum, 0)
	for _, v := range mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumStringValues Enumerates the set of values in String for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumStringValues() []string {
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

// GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum(val string) (UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum, bool) {
	enum, ok := mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum Enum with underlying type: string
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum string

// Set of constants representing the allowable values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum
const (
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsRead  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum = "READ"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsWrite UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum = "WRITE"
)

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum{
	"READ":  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsRead,
	"WRITE": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsWrite,
}

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumLowerCase = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum{
	"read":  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsRead,
	"write": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsWrite,
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumValues Enumerates the set of values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumValues() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum {
	values := make([]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum, 0)
	for _, v := range mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumStringValues Enumerates the set of values in String for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumStringValues() []string {
	return []string{
		"READ",
		"WRITE",
	}
}

// GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum(val string) (UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum, bool) {
	enum, ok := mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum Enum with underlying type: string
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum string

// Set of constants representing the allowable values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum
const (
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBasic  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum = "BASIC"
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBearer UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum = "BEARER"
)

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum{
	"BASIC":  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBasic,
	"BEARER": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBearer,
}

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumLowerCase = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum{
	"basic":  UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBasic,
	"bearer": UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesBearer,
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumValues Enumerates the set of values for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumValues() []UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum {
	values := make([]UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum, 0)
	for _, v := range mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumStringValues Enumerates the set of values in String for UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumStringValues() []string {
	return []string{
		"BASIC",
		"BEARER",
	}
}

// GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum(val string) (UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum, bool) {
	enum, ok := mappingUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
