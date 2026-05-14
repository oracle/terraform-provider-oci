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

// DatabaseToolsDatabaseApiGatewayConfigGlobalDefault The content of a Database Tools database API gateway config global resource.
type DatabaseToolsDatabaseApiGatewayConfigGlobalDefault struct {

	// A string that uniquely identifies a Database Tools database API gateway config global settings resource.
	Key *string `mandatory:"true" json:"key"`

	// The time the resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The request header name providing the pool route value.
	PoolRoutingHeader *string `mandatory:"false" json:"poolRoutingHeader"`

	// Specifies the HTTP listen port. 0 disables HTTP. Use of ports below 1024 requires elevated (root) privileges and is generally discouraged; deployment on non-privileged ports (1024–65535) is recommended.
	HttpPort *int `mandatory:"false" json:"httpPort"`

	// Specifies the HTTPS listen port. 0 disables HTTPS. Use of ports below 1024 requires elevated (root) privileges and is generally discouraged; deployment on non-privileged ports (1024–65535) is recommended. ORDS will use a self-signed certificate if a certificate bundle is not provided.
	HttpsPort *int `mandatory:"false" json:"httpsPort"`

	CertificateBundle DatabaseApiGatewayConfigCertificateBundle `mandatory:"false" json:"certificateBundle"`

	// The location of the static resources to be served under the / root server path.
	DocumentRoot *string `mandatory:"false" json:"documentRoot"`

	// Advanced global properties.
	AdvancedProperties map[string]string `mandatory:"false" json:"advancedProperties"`

	// How the target pool route value is determined for a HTTP request.
	PoolRoute DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum `mandatory:"false" json:"poolRoute,omitempty"`

	// ORDS database API is a database management and monitoring REST API. Database Actions requires this feature.
	DatabaseApiStatus DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum `mandatory:"false" json:"databaseApiStatus,omitempty"`

	// The RESTful service definition location.
	MetadataSource DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum `mandatory:"false" json:"metadataSource,omitempty"`
}

// GetKey returns Key
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetKey() *string {
	return m.Key
}

// GetPoolRoute returns PoolRoute
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetPoolRoute() DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum {
	return m.PoolRoute
}

// GetPoolRoutingHeader returns PoolRoutingHeader
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetPoolRoutingHeader() *string {
	return m.PoolRoutingHeader
}

// GetDatabaseApiStatus returns DatabaseApiStatus
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetDatabaseApiStatus() DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum {
	return m.DatabaseApiStatus
}

// GetMetadataSource returns MetadataSource
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetMetadataSource() DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum {
	return m.MetadataSource
}

// GetHttpPort returns HttpPort
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetHttpPort() *int {
	return m.HttpPort
}

// GetHttpsPort returns HttpsPort
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetHttpsPort() *int {
	return m.HttpsPort
}

// GetCertificateBundle returns CertificateBundle
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetCertificateBundle() DatabaseApiGatewayConfigCertificateBundle {
	return m.CertificateBundle
}

// GetDocumentRoot returns DocumentRoot
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetDocumentRoot() *string {
	return m.DocumentRoot
}

// GetAdvancedProperties returns AdvancedProperties
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsDatabaseApiGatewayConfigGlobalDefault DatabaseToolsDatabaseApiGatewayConfigGlobalDefault
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsDatabaseApiGatewayConfigGlobalDefault
	}{
		"DEFAULT",
		(MarshalTypeDatabaseToolsDatabaseApiGatewayConfigGlobalDefault)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsDatabaseApiGatewayConfigGlobalDefault) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PoolRoute          DatabaseToolsDatabaseApiGatewayConfigGlobalPoolRouteEnum         `json:"poolRoute"`
		PoolRoutingHeader  *string                                                          `json:"poolRoutingHeader"`
		DatabaseApiStatus  DatabaseToolsDatabaseApiGatewayConfigGlobalDatabaseApiStatusEnum `json:"databaseApiStatus"`
		MetadataSource     DatabaseToolsDatabaseApiGatewayConfigGlobalMetadataSourceEnum    `json:"metadataSource"`
		HttpPort           *int                                                             `json:"httpPort"`
		HttpsPort          *int                                                             `json:"httpsPort"`
		CertificateBundle  databaseapigatewayconfigcertificatebundle                        `json:"certificateBundle"`
		DocumentRoot       *string                                                          `json:"documentRoot"`
		AdvancedProperties map[string]string                                                `json:"advancedProperties"`
		Key                *string                                                          `json:"key"`
		TimeCreated        *common.SDKTime                                                  `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                                                  `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.PoolRoute = model.PoolRoute

	m.PoolRoutingHeader = model.PoolRoutingHeader

	m.DatabaseApiStatus = model.DatabaseApiStatus

	m.MetadataSource = model.MetadataSource

	m.HttpPort = model.HttpPort

	m.HttpsPort = model.HttpsPort

	nn, e = model.CertificateBundle.UnmarshalPolymorphicJSON(model.CertificateBundle.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CertificateBundle = nn.(DatabaseApiGatewayConfigCertificateBundle)
	} else {
		m.CertificateBundle = nil
	}

	m.DocumentRoot = model.DocumentRoot

	m.AdvancedProperties = model.AdvancedProperties

	m.Key = model.Key

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}
