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

// UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails The content of the Database Tools database API gateway config global sub resource to be updated.
type UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails struct {

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
	PoolRoute UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum `mandatory:"false" json:"poolRoute,omitempty"`

	// ORDS database API is a database management and monitoring REST API. Database Actions requires this feature.
	DatabaseApiStatus UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum `mandatory:"false" json:"databaseApiStatus,omitempty"`
}

// GetPoolRoute returns PoolRoute
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) GetPoolRoute() UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum {
	return m.PoolRoute
}

// GetPoolRoutingHeader returns PoolRoutingHeader
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) GetPoolRoutingHeader() *string {
	return m.PoolRoutingHeader
}

// GetDatabaseApiStatus returns DatabaseApiStatus
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) GetDatabaseApiStatus() UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum {
	return m.DatabaseApiStatus
}

// GetHttpPort returns HttpPort
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) GetHttpPort() *int {
	return m.HttpPort
}

// GetHttpsPort returns HttpsPort
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) GetHttpsPort() *int {
	return m.HttpsPort
}

// GetCertificateBundle returns CertificateBundle
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) GetCertificateBundle() DatabaseApiGatewayConfigCertificateBundle {
	return m.CertificateBundle
}

// GetDocumentRoot returns DocumentRoot
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) GetDocumentRoot() *string {
	return m.DocumentRoot
}

// GetAdvancedProperties returns AdvancedProperties
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) GetAdvancedProperties() map[string]string {
	return m.AdvancedProperties
}

func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails
	}{
		"DEFAULT",
		(MarshalTypeUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PoolRoute          UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum         `json:"poolRoute"`
		PoolRoutingHeader  *string                                                                       `json:"poolRoutingHeader"`
		DatabaseApiStatus  UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum `json:"databaseApiStatus"`
		HttpPort           *int                                                                          `json:"httpPort"`
		HttpsPort          *int                                                                          `json:"httpsPort"`
		CertificateBundle  databaseapigatewayconfigcertificatebundle                                     `json:"certificateBundle"`
		DocumentRoot       *string                                                                       `json:"documentRoot"`
		AdvancedProperties map[string]string                                                             `json:"advancedProperties"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.PoolRoute = model.PoolRoute

	m.PoolRoutingHeader = model.PoolRoutingHeader

	m.DatabaseApiStatus = model.DatabaseApiStatus

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

	return
}
