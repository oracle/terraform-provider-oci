// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// MetricExtension Detailed information of the Metric Extension resource
type MetricExtension struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of Metric Extension resource
	Id *string `mandatory:"true" json:"id"`

	// Metric Extension resource name
	Name *string `mandatory:"true" json:"name"`

	// Metric Extension resource display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Resource type to which Metric Extension applies
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Tenant Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	TenantId *string `mandatory:"true" json:"tenantId"`

	// Collection Method  Metric Extension applies
	CollectionMethod *string `mandatory:"true" json:"collectionMethod"`

	// The current status of the metric extension i.e. whether it is Draft or Published
	Status MetricExtensionLifeCycleDetailsEnum `mandatory:"true" json:"status"`

	// Schedule of metric extension should use RFC 5545 format -> recur-rule-part = "FREQ";"INTERVAL" where FREQ rule part identifies the type of recurrence rule. Valid values are "MINUTELY","HOURLY","DAILY" to specify repeating events based on an interval of a minute, an hour and a day or more. Example- FREQ=DAILY;INTERVAL=1
	CollectionRecurrences *string `mandatory:"true" json:"collectionRecurrences"`

	// List of metrics which are part of this metric extension
	MetricList []Metric `mandatory:"true" json:"metricList"`

	QueryProperties MetricExtensionQueryProperties `mandatory:"true" json:"queryProperties"`

	// Description of the metric extension.
	Description *string `mandatory:"false" json:"description"`

	// The current lifecycle state of the metric extension
	LifecycleState MetricExtensionLifeCycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Created by user
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Last updated by user
	LastUpdatedBy *string `mandatory:"false" json:"lastUpdatedBy"`

	// Metric Extension creation time. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Metric Extension update time. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// List of resource objects on which this metric extension is enabled.
	EnabledOnResources []EnabledResourceDetails `mandatory:"false" json:"enabledOnResources"`

	// Count of resources on which this metric extension is enabled.
	EnabledOnResourcesCount *int `mandatory:"false" json:"enabledOnResourcesCount"`

	// The URI path that the user can do a GET on to access the metric extension metadata
	ResourceUri *string `mandatory:"false" json:"resourceUri"`
}

func (m MetricExtension) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricExtension) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetricExtensionLifeCycleDetailsEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetMetricExtensionLifeCycleDetailsEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMetricExtensionLifeCycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMetricExtensionLifeCycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MetricExtension) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                             `json:"description"`
		LifecycleState          MetricExtensionLifeCycleStatesEnum  `json:"lifecycleState"`
		CreatedBy               *string                             `json:"createdBy"`
		LastUpdatedBy           *string                             `json:"lastUpdatedBy"`
		TimeCreated             *common.SDKTime                     `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                     `json:"timeUpdated"`
		EnabledOnResources      []EnabledResourceDetails            `json:"enabledOnResources"`
		EnabledOnResourcesCount *int                                `json:"enabledOnResourcesCount"`
		ResourceUri             *string                             `json:"resourceUri"`
		Id                      *string                             `json:"id"`
		Name                    *string                             `json:"name"`
		DisplayName             *string                             `json:"displayName"`
		ResourceType            *string                             `json:"resourceType"`
		CompartmentId           *string                             `json:"compartmentId"`
		TenantId                *string                             `json:"tenantId"`
		CollectionMethod        *string                             `json:"collectionMethod"`
		Status                  MetricExtensionLifeCycleDetailsEnum `json:"status"`
		CollectionRecurrences   *string                             `json:"collectionRecurrences"`
		MetricList              []Metric                            `json:"metricList"`
		QueryProperties         metricextensionqueryproperties      `json:"queryProperties"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecycleState = model.LifecycleState

	m.CreatedBy = model.CreatedBy

	m.LastUpdatedBy = model.LastUpdatedBy

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.EnabledOnResources = make([]EnabledResourceDetails, len(model.EnabledOnResources))
	copy(m.EnabledOnResources, model.EnabledOnResources)
	m.EnabledOnResourcesCount = model.EnabledOnResourcesCount

	m.ResourceUri = model.ResourceUri

	m.Id = model.Id

	m.Name = model.Name

	m.DisplayName = model.DisplayName

	m.ResourceType = model.ResourceType

	m.CompartmentId = model.CompartmentId

	m.TenantId = model.TenantId

	m.CollectionMethod = model.CollectionMethod

	m.Status = model.Status

	m.CollectionRecurrences = model.CollectionRecurrences

	m.MetricList = make([]Metric, len(model.MetricList))
	copy(m.MetricList, model.MetricList)
	nn, e = model.QueryProperties.UnmarshalPolymorphicJSON(model.QueryProperties.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryProperties = nn.(MetricExtensionQueryProperties)
	} else {
		m.QueryProperties = nil
	}

	return
}
