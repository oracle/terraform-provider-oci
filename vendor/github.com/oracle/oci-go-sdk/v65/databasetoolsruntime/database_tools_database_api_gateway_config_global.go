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

// DatabaseToolsDatabaseApiGatewayConfigGlobal The content of a Database Tools database API gateway config global resource.
type DatabaseToolsDatabaseApiGatewayConfigGlobal interface {

	// A string that uniquely identifies a Database Tools database API gateway config global settings resource.
	GetKey() *string

	// The time the resource was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the resource was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// How the target pool route value is determined for a HTTP request.
	GetPoolRoute() DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum

	// The request header name providing the pool route value.
	GetPoolRoutingHeader() *string

	// ORDS database API is a database management and monitoring REST API. Database Actions requires this feature.
	GetDatabaseApiStatus() DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum

	// The RESTful service definition location.
	GetMetadataSource() DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum

	// Specifies the HTTP listen port. 0 disables HTTP. Use of ports below 1024 requires elevated (root) privileges and is generally discouraged; deployment on non-privileged ports (1024–65535) is recommended.
	GetHttpPort() *int

	// Specifies the HTTPS listen port. 0 disables HTTPS. Use of ports below 1024 requires elevated (root) privileges and is generally discouraged; deployment on non-privileged ports (1024–65535) is recommended. ORDS will use a self-signed certificate if a certificate bundle is not provided.
	GetHttpsPort() *int

	GetCertificateBundle() DatabaseApiGatewayConfigCertificateBundle

	// The location of the static resources to be served under the / root server path.
	GetDocumentRoot() *string

	// Advanced global properties.
	GetAdvancedProperties() map[string]string
}

type databasetoolsdatabaseapigatewayconfigglobal struct {
	JsonData           []byte
	PoolRoute          DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum         `mandatory:"false" json:"poolRoute,omitempty"`
	PoolRoutingHeader  *string                                                          `mandatory:"false" json:"poolRoutingHeader"`
	DatabaseApiStatus  DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum `mandatory:"false" json:"databaseApiStatus,omitempty"`
	MetadataSource     DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum    `mandatory:"false" json:"metadataSource,omitempty"`
	HttpPort           *int                                                             `mandatory:"false" json:"httpPort"`
	HttpsPort          *int                                                             `mandatory:"false" json:"httpsPort"`
	CertificateBundle  databaseapigatewayconfigcertificatebundle                        `mandatory:"false" json:"certificateBundle"`
	DocumentRoot       *string                                                          `mandatory:"false" json:"documentRoot"`
	AdvancedProperties map[string]string                                                `mandatory:"false" json:"advancedProperties"`
	Key                *string                                                          `mandatory:"true" json:"key"`
	TimeCreated        *common.SDKTime                                                  `mandatory:"true" json:"timeCreated"`
	TimeUpdated        *common.SDKTime                                                  `mandatory:"true" json:"timeUpdated"`
	Type               string                                                           `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsdatabaseapigatewayconfigglobal) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsdatabaseapigatewayconfigglobal databasetoolsdatabaseapigatewayconfigglobal
	s := struct {
		Model Unmarshalerdatabasetoolsdatabaseapigatewayconfigglobal
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.PoolRoute = s.Model.PoolRoute
	m.PoolRoutingHeader = s.Model.PoolRoutingHeader
	m.DatabaseApiStatus = s.Model.DatabaseApiStatus
	m.MetadataSource = s.Model.MetadataSource
	m.HttpPort = s.Model.HttpPort
	m.HttpsPort = s.Model.HttpsPort
	m.CertificateBundle = s.Model.CertificateBundle
	m.DocumentRoot = s.Model.DocumentRoot
	m.AdvancedProperties = s.Model.AdvancedProperties
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsdatabaseapigatewayconfigglobal) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DatabaseToolsDatabaseApiGatewayConfigGlobalDefault{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsDatabaseApiGatewayConfigGlobal: %s.", m.Type)
		return *m, nil
	}
}

// GetPoolRoute returns PoolRoute
func (m databasetoolsdatabaseapigatewayconfigglobal) GetPoolRoute() DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum {
	return m.PoolRoute
}

// GetPoolRoutingHeader returns PoolRoutingHeader
func (m databasetoolsdatabaseapigatewayconfigglobal) GetPoolRoutingHeader() *string {
	return m.PoolRoutingHeader
}

// GetDatabaseApiStatus returns DatabaseApiStatus
func (m databasetoolsdatabaseapigatewayconfigglobal) GetDatabaseApiStatus() DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum {
	return m.DatabaseApiStatus
}

// GetMetadataSource returns MetadataSource
func (m databasetoolsdatabaseapigatewayconfigglobal) GetMetadataSource() DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum {
	return m.MetadataSource
}

// GetHttpPort returns HttpPort
func (m databasetoolsdatabaseapigatewayconfigglobal) GetHttpPort() *int {
	return m.HttpPort
}

// GetHttpsPort returns HttpsPort
func (m databasetoolsdatabaseapigatewayconfigglobal) GetHttpsPort() *int {
	return m.HttpsPort
}

// GetCertificateBundle returns CertificateBundle
func (m databasetoolsdatabaseapigatewayconfigglobal) GetCertificateBundle() databaseapigatewayconfigcertificatebundle {
	return m.CertificateBundle
}

// GetDocumentRoot returns DocumentRoot
func (m databasetoolsdatabaseapigatewayconfigglobal) GetDocumentRoot() *string {
	return m.DocumentRoot
}

// GetAdvancedProperties returns AdvancedProperties
func (m databasetoolsdatabaseapigatewayconfigglobal) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

// GetKey returns Key
func (m databasetoolsdatabaseapigatewayconfigglobal) GetKey() *string {
	return m.Key
}

// GetTimeCreated returns TimeCreated
func (m databasetoolsdatabaseapigatewayconfigglobal) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolsdatabaseapigatewayconfigglobal) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m databasetoolsdatabaseapigatewayconfigglobal) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsdatabaseapigatewayconfigglobal) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum(string(m.PoolRoute)); !ok && m.PoolRoute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PoolRoute: %s. Supported values are: %s.", m.PoolRoute, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum(string(m.DatabaseApiStatus)); !ok && m.DatabaseApiStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseApiStatus: %s. Supported values are: %s.", m.DatabaseApiStatus, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum(string(m.MetadataSource)); !ok && m.MetadataSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetadataSource: %s. Supported values are: %s.", m.MetadataSource, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRoutePath   DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum = "PATH"
	DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteHeader DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum = "HEADER"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum{
	"PATH":   DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRoutePath,
	"HEADER": DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteHeader,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum{
	"path":   DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRoutePath,
	"header": DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteHeader,
}

// GetDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum
func GetDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnumValues() []DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum
func GetDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnumStringValues() []string {
	return []string{
		"PATH",
		"HEADER",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnabled  DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum = "ENABLED"
	DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusDisabled DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum = "DISABLED"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum{
	"ENABLED":  DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnabled,
	"DISABLED": DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusDisabled,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum{
	"enabled":  DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnabled,
	"disabled": DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusDisabled,
}

// GetDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum
func GetDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnumValues() []DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum
func GetDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceDatabase DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum = "DATABASE"
	DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceCloud    DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum = "CLOUD"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum{
	"DATABASE": DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceDatabase,
	"CLOUD":    DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceCloud,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum{
	"database": DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceDatabase,
	"cloud":    DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceCloud,
}

// GetDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum
func GetDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnumValues() []DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum
func GetDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnumStringValues() []string {
	return []string{
		"DATABASE",
		"CLOUD",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
