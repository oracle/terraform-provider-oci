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

// MonitoredResource The response object for create monitored resource and get monitored resource operations.
// This contains information about the monitored resource. Credentials and credential aliases attributes
// will be returned as null due to security reasons.
type MonitoredResource struct {

	// Monitored resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// Monitored resource name.
	Name *string `mandatory:"true" json:"name"`

	// Monitored Resource Type.
	Type *string `mandatory:"true" json:"type"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Tenancy Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	TenantId *string `mandatory:"true" json:"tenantId"`

	// Monitored resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Monitored resource host name.
	HostName *string `mandatory:"false" json:"hostName"`

	// The external resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// External resource is any OCI resource which is not a Stack Monitoring service resource.
	// Currently supports only following resource types - Container database, non-container database,
	// pluggable database and OCI compute instance.
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Management Agent Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// Time zone in the form of tz database canonical zone ID.
	ResourceTimeZone *string `mandatory:"false" json:"resourceTimeZone"`

	// The date and time when the monitored resource was created, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the monitored resource was last updated, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Lifecycle state of the monitored resource.
	LifecycleState ResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// License edition of the monitored resource.
	License LicenseTypeEnum `mandatory:"false" json:"license,omitempty"`

	// Source type to indicate if the resource is stack monitoring discovered, OCI native resource, etc.
	SourceType SourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// Resource Category to indicate the kind of resource type.
	ResourceCategory ResourceCategoryEnum `mandatory:"false" json:"resourceCategory,omitempty"`

	// List of monitored resource properties.
	Properties []MonitoredResourceProperty `mandatory:"false" json:"properties"`

	DatabaseConnectionDetails *ConnectionDetails `mandatory:"false" json:"databaseConnectionDetails"`

	Credentials MonitoredResourceCredential `mandatory:"false" json:"credentials"`

	Aliases *MonitoredResourceAliasCredential `mandatory:"false" json:"aliases"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MonitoredResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseTypeEnum(string(m.License)); !ok && m.License != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for License: %s. Supported values are: %s.", m.License, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceCategoryEnum(string(m.ResourceCategory)); !ok && m.ResourceCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceCategory: %s. Supported values are: %s.", m.ResourceCategory, strings.Join(GetResourceCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MonitoredResource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName               *string                           `json:"displayName"`
		HostName                  *string                           `json:"hostName"`
		ExternalId                *string                           `json:"externalId"`
		ManagementAgentId         *string                           `json:"managementAgentId"`
		ResourceTimeZone          *string                           `json:"resourceTimeZone"`
		TimeCreated               *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated               *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState            ResourceLifecycleStateEnum        `json:"lifecycleState"`
		License                   LicenseTypeEnum                   `json:"license"`
		SourceType                SourceTypeEnum                    `json:"sourceType"`
		ResourceCategory          ResourceCategoryEnum              `json:"resourceCategory"`
		Properties                []MonitoredResourceProperty       `json:"properties"`
		DatabaseConnectionDetails *ConnectionDetails                `json:"databaseConnectionDetails"`
		Credentials               monitoredresourcecredential       `json:"credentials"`
		Aliases                   *MonitoredResourceAliasCredential `json:"aliases"`
		FreeformTags              map[string]string                 `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                map[string]map[string]interface{} `json:"systemTags"`
		Id                        *string                           `json:"id"`
		Name                      *string                           `json:"name"`
		Type                      *string                           `json:"type"`
		CompartmentId             *string                           `json:"compartmentId"`
		TenantId                  *string                           `json:"tenantId"`
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

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.License = model.License

	m.SourceType = model.SourceType

	m.ResourceCategory = model.ResourceCategory

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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.Type = model.Type

	m.CompartmentId = model.CompartmentId

	m.TenantId = model.TenantId

	return
}
