// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMonitoredResourceDetails The information about new monitored resource to be created.
// The combination of monitored resource name and type should be unique across tenancy.
type CreateMonitoredResourceDetails struct {

	// Monitored Resource Name.
	Name *string `mandatory:"true" json:"name"`

	// Monitored Resource Type.
	Type *string `mandatory:"true" json:"type"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Monitored resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Host name of the monitored resource.
	HostName *string `mandatory:"false" json:"hostName"`

	// External resource is any OCI resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	// which is not a Stack Monitoring service resource.
	// Currently supports only OCI compute instance.
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Management Agent Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// Time zone in the form of tz database canonical zone ID. Specifies the preference with
	// a value that uses the IANA Time Zone Database format (x-obmcs-time-zone).
	// For example - America/Los_Angeles
	ResourceTimeZone *string `mandatory:"false" json:"resourceTimeZone"`

	// License edition of the monitored resource. If not provided
	// the default license type for the compartment will be used.
	License LicenseTypeEnum `mandatory:"false" json:"license,omitempty"`

	// List of monitored resource properties.
	Properties []MonitoredResourceProperty `mandatory:"false" json:"properties"`

	DatabaseConnectionDetails *ConnectionDetails `mandatory:"false" json:"databaseConnectionDetails"`

	Credentials MonitoredResourceCredential `mandatory:"false" json:"credentials"`

	Aliases *MonitoredResourceAliasCredential `mandatory:"false" json:"aliases"`

	// List of MonitoredResourceCredentials. This property complements the existing
	// "credentials" property by allowing user to specify more than one credential.
	// If both "credential" and "additionalCredentials" are specified, union of the
	// values is used as list of credentials applicable for this resource.
	// If any duplicate found in the combined list of "credentials" and "additionalCredentials",
	// an error will be thrown.
	AdditionalCredentials []MonitoredResourceCredential `mandatory:"false" json:"additionalCredentials"`

	// List of MonitoredResourceAliasCredentials. This property complements the existing
	// "aliases" property by allowing user to specify more than one credential alias.
	// If both "aliases" and "additionalAliases" are specified, union of the
	// values is used as list of aliases applicable for this resource.
	// If any duplicate found in the combined list of "alias" and "additionalAliases",
	// an error will be thrown.
	AdditionalAliases []MonitoredResourceAliasCredential `mandatory:"false" json:"additionalAliases"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMonitoredResourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMonitoredResourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLicenseTypeEnum(string(m.License)); !ok && m.License != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for License: %s. Supported values are: %s.", m.License, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateMonitoredResourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName               *string                            `json:"displayName"`
		HostName                  *string                            `json:"hostName"`
		ExternalId                *string                            `json:"externalId"`
		ManagementAgentId         *string                            `json:"managementAgentId"`
		ResourceTimeZone          *string                            `json:"resourceTimeZone"`
		License                   LicenseTypeEnum                    `json:"license"`
		Properties                []MonitoredResourceProperty        `json:"properties"`
		DatabaseConnectionDetails *ConnectionDetails                 `json:"databaseConnectionDetails"`
		Credentials               monitoredresourcecredential        `json:"credentials"`
		Aliases                   *MonitoredResourceAliasCredential  `json:"aliases"`
		AdditionalCredentials     []monitoredresourcecredential      `json:"additionalCredentials"`
		AdditionalAliases         []MonitoredResourceAliasCredential `json:"additionalAliases"`
		FreeformTags              map[string]string                  `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{}  `json:"definedTags"`
		Name                      *string                            `json:"name"`
		Type                      *string                            `json:"type"`
		CompartmentId             *string                            `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.HostName = model.HostName

	m.ExternalId = model.ExternalId

	m.ManagementAgentId = model.ManagementAgentId

	m.ResourceTimeZone = model.ResourceTimeZone

	m.License = model.License

	m.Properties = make([]MonitoredResourceProperty, len(model.Properties))
	copy(m.Properties, model.Properties)
	m.DatabaseConnectionDetails = model.DatabaseConnectionDetails

	nn, e = model.Credentials.UnmarshalPolymorphicJSON(model.Credentials.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Credentials = nn.(MonitoredResourceCredential)
	} else {
		m.Credentials = nil
	}

	m.Aliases = model.Aliases

	m.AdditionalCredentials = make([]MonitoredResourceCredential, len(model.AdditionalCredentials))
	for i, n := range model.AdditionalCredentials {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.AdditionalCredentials[i] = nn.(MonitoredResourceCredential)
		} else {
			m.AdditionalCredentials[i] = nil
		}
	}
	m.AdditionalAliases = make([]MonitoredResourceAliasCredential, len(model.AdditionalAliases))
	copy(m.AdditionalAliases, model.AdditionalAliases)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Name = model.Name

	m.Type = model.Type

	m.CompartmentId = model.CompartmentId

	return
}
