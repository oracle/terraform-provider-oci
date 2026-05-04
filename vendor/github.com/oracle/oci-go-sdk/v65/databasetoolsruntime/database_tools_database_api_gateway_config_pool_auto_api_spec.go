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

// DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec The content of a Database Tools database API gateway config auto API spec sub resource defined within a pool.
type DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec interface {

	// A system generated string that uniquely identifies an auto API spec sub resource within a given pool.
	GetKey() *string

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The time the resource was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the resource was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// The name of the database object.
	GetDatabaseObjectName() *string

	// The type of the database object.
	GetDatabaseObjectType() DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum

	// Description of the autoApiSpec.
	GetDescription() *string

	// Used as the URI path element for this object. When not specified the objectName lowercase is the default value.
	GetAlias() *string

	// The operations to limit access to this resource. If not specified then the default is ["READ","WRITE"].
	GetOperations() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum

	// The security schemes that can access this resource. If not specified then the resource is public.
	GetSecuritySchemes() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum

	// The name of the database API gateway config privilege protecting the resource. Only valid for SCOPE JWT Profile pools and BEARER securitySchemes.
	GetScope() *string

	// The name of the database API gateway config roles protecting the resource. Only valid for RBAC JWT Profile pools and BEARER securitySchemes.
	GetRoles() []string
}

type databasetoolsdatabaseapigatewayconfigpoolautoapispec struct {
	JsonData           []byte
	DatabaseObjectName *string                                                                    `mandatory:"false" json:"databaseObjectName"`
	DatabaseObjectType DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum `mandatory:"false" json:"databaseObjectType,omitempty"`
	Description        *string                                                                    `mandatory:"false" json:"description"`
	Alias              *string                                                                    `mandatory:"false" json:"alias"`
	Operations         []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum       `mandatory:"false" json:"operations,omitempty"`
	SecuritySchemes    []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum  `mandatory:"false" json:"securitySchemes,omitempty"`
	Scope              *string                                                                    `mandatory:"false" json:"scope"`
	Roles              []string                                                                   `mandatory:"false" json:"roles"`
	Key                *string                                                                    `mandatory:"true" json:"key"`
	DisplayName        *string                                                                    `mandatory:"true" json:"displayName"`
	TimeCreated        *common.SDKTime                                                            `mandatory:"true" json:"timeCreated"`
	TimeUpdated        *common.SDKTime                                                            `mandatory:"true" json:"timeUpdated"`
	Type               string                                                                     `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsdatabaseapigatewayconfigpoolautoapispec) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsdatabaseapigatewayconfigpoolautoapispec databasetoolsdatabaseapigatewayconfigpoolautoapispec
	s := struct {
		Model Unmarshalerdatabasetoolsdatabaseapigatewayconfigpoolautoapispec
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
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
func (m *databasetoolsdatabaseapigatewayconfigpoolautoapispec) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefault{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec: %s.", m.Type)
		return *m, nil
	}
}

// GetDatabaseObjectName returns DatabaseObjectName
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetDatabaseObjectName() *string {
	return m.DatabaseObjectName
}

// GetDatabaseObjectType returns DatabaseObjectType
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetDatabaseObjectType() DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum {
	return m.DatabaseObjectType
}

// GetDescription returns Description
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetDescription() *string {
	return m.Description
}

// GetAlias returns Alias
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetAlias() *string {
	return m.Alias
}

// GetOperations returns Operations
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetOperations() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum {
	return m.Operations
}

// GetSecuritySchemes returns SecuritySchemes
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetSecuritySchemes() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum {
	return m.SecuritySchemes
}

// GetScope returns Scope
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetScope() *string {
	return m.Scope
}

// GetRoles returns Roles
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetRoles() []string {
	return m.Roles
}

// GetKey returns Key
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetKey() *string {
	return m.Key
}

// GetDisplayName returns DisplayName
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsdatabaseapigatewayconfigpoolautoapispec) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum(string(m.DatabaseObjectType)); !ok && m.DatabaseObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseObjectType: %s. Supported values are: %s.", m.DatabaseObjectType, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnumStringValues(), ",")))
	}
	for _, val := range m.Operations {
		if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum(string(val)); !ok && string(val) != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operations: %s. Supported values are: %s.", val, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.SecuritySchemes {
		if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum(string(val)); !ok && string(val) != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecuritySchemes: %s. Supported values are: %s.", val, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeFunction    DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum = "FUNCTION"
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeMview       DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum = "MVIEW"
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypePackage     DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum = "PACKAGE"
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeProcedure   DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum = "PROCEDURE"
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeTable       DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum = "TABLE"
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeView        DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum = "VIEW"
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeDualityview DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum = "DUALITYVIEW"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum{
	"FUNCTION":    DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeFunction,
	"MVIEW":       DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeMview,
	"PACKAGE":     DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypePackage,
	"PROCEDURE":   DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeProcedure,
	"TABLE":       DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeTable,
	"VIEW":        DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeView,
	"DUALITYVIEW": DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeDualityview,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum{
	"function":    DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeFunction,
	"mview":       DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeMview,
	"package":     DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypePackage,
	"procedure":   DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeProcedure,
	"table":       DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeTable,
	"view":        DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeView,
	"dualityview": DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeDualityview,
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnumValues() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnumStringValues() []string {
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

// GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDatabaseObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsRead  DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum = "READ"
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsWrite DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum = "WRITE"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum{
	"READ":  DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsRead,
	"WRITE": DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsWrite,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum{
	"read":  DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsRead,
	"write": DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsWrite,
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnumValues() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnumStringValues() []string {
	return []string{
		"READ",
		"WRITE",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesBasic  DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum = "BASIC"
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesBearer DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum = "BEARER"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum{
	"BASIC":  DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesBasic,
	"BEARER": DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesBearer,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum{
	"basic":  DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesBasic,
	"bearer": DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesBearer,
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnumValues() []DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum
func GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnumStringValues() []string {
	return []string{
		"BASIC",
		"BEARER",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
