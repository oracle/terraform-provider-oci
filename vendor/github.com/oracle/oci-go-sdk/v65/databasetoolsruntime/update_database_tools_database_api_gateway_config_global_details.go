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

// UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetails The content of the Database Tools database API gateway config global sub resource to be updated.
type UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetails interface {

	// How the target pool route value is determined for a HTTP request.
	GetPoolRoute() UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum

	// The request header name providing the pool route value.
	GetPoolRoutingHeader() *string

	// ORDS database API is a database management and monitoring REST API. Database Actions requires this feature.
	GetDatabaseApiStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum

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

type updatedatabasetoolsdatabaseapigatewayconfigglobaldetails struct {
	JsonData           []byte
	PoolRoute          UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum         `mandatory:"false" json:"poolRoute,omitempty"`
	PoolRoutingHeader  *string                                                                       `mandatory:"false" json:"poolRoutingHeader"`
	DatabaseApiStatus  UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum `mandatory:"false" json:"databaseApiStatus,omitempty"`
	HttpPort           *int                                                                          `mandatory:"false" json:"httpPort"`
	HttpsPort          *int                                                                          `mandatory:"false" json:"httpsPort"`
	CertificateBundle  databaseapigatewayconfigcertificatebundle                                     `mandatory:"false" json:"certificateBundle"`
	DocumentRoot       *string                                                                       `mandatory:"false" json:"documentRoot"`
	AdvancedProperties map[string]string                                                             `mandatory:"false" json:"advancedProperties"`
	Type               string                                                                        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedatabasetoolsdatabaseapigatewayconfigglobaldetails updatedatabasetoolsdatabaseapigatewayconfigglobaldetails
	s := struct {
		Model Unmarshalerupdatedatabasetoolsdatabaseapigatewayconfigglobaldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PoolRoute = s.Model.PoolRoute
	m.PoolRoutingHeader = s.Model.PoolRoutingHeader
	m.DatabaseApiStatus = s.Model.DatabaseApiStatus
	m.HttpPort = s.Model.HttpPort
	m.HttpsPort = s.Model.HttpsPort
	m.CertificateBundle = s.Model.CertificateBundle
	m.DocumentRoot = s.Model.DocumentRoot
	m.AdvancedProperties = s.Model.AdvancedProperties
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetPoolRoute returns PoolRoute
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) GetPoolRoute() UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum {
	return m.PoolRoute
}

// GetPoolRoutingHeader returns PoolRoutingHeader
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) GetPoolRoutingHeader() *string {
	return m.PoolRoutingHeader
}

// GetDatabaseApiStatus returns DatabaseApiStatus
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) GetDatabaseApiStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum {
	return m.DatabaseApiStatus
}

// GetHttpPort returns HttpPort
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) GetHttpPort() *int {
	return m.HttpPort
}

// GetHttpsPort returns HttpsPort
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) GetHttpsPort() *int {
	return m.HttpsPort
}

// GetCertificateBundle returns CertificateBundle
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) GetCertificateBundle() databaseapigatewayconfigcertificatebundle {
	return m.CertificateBundle
}

// GetDocumentRoot returns DocumentRoot
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) GetDocumentRoot() *string {
	return m.DocumentRoot
}

// GetAdvancedProperties returns AdvancedProperties
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedatabasetoolsdatabaseapigatewayconfigglobaldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum(string(m.PoolRoute)); !ok && m.PoolRoute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PoolRoute: %s. Supported values are: %s.", m.PoolRoute, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum(string(m.DatabaseApiStatus)); !ok && m.DatabaseApiStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseApiStatus: %s. Supported values are: %s.", m.DatabaseApiStatus, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum Enum with underlying type: string
type UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum string

// Set of constants representing the allowable values for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum
const (
	UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRoutePath   UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum = "PATH"
	UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteHeader UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum = "HEADER"
)

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum{
	"PATH":   UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRoutePath,
	"HEADER": UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteHeader,
}

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnumLowerCase = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum{
	"path":   UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRoutePath,
	"header": UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteHeader,
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnumValues Enumerates the set of values for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnumValues() []UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum {
	values := make([]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum, 0)
	for _, v := range mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnumStringValues Enumerates the set of values in String for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnumStringValues() []string {
	return []string{
		"PATH",
		"HEADER",
	}
}

// GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum(val string) (UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum, bool) {
	enum, ok := mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum Enum with underlying type: string
type UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum string

// Set of constants representing the allowable values for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum
const (
	UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnabled  UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum = "ENABLED"
	UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusDisabled UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum = "DISABLED"
)

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum{
	"ENABLED":  UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnabled,
	"DISABLED": UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusDisabled,
}

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnumLowerCase = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum{
	"enabled":  UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnabled,
	"disabled": UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusDisabled,
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnumValues Enumerates the set of values for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnumValues() []UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum {
	values := make([]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum, 0)
	for _, v := range mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnumStringValues Enumerates the set of values in String for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum(val string) (UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum, bool) {
	enum, ok := mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
