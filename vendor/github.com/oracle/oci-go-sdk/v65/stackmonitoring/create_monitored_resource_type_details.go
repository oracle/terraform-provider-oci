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

// CreateMonitoredResourceTypeDetails The information about new monitored resource type. The resource type name should be unique across tenancy.
// A set of resource types are created by the service by default. These resource types are available
// for all tenancies. Service provided resource types can not be duplicated or overwritten in any tenancy.
type CreateMonitoredResourceTypeDetails struct {

	// A unique monitored resource type name. The name must be unique across tenancy.
	// Name can not be changed.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy containing the resource type.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Monitored resource type display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A friendly description.
	Description *string `mandatory:"false" json:"description"`

	// Metric namespace for resource type.
	MetricNamespace *string `mandatory:"false" json:"metricNamespace"`

	Metadata ResourceTypeMetadataDetails `mandatory:"false" json:"metadata"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMonitoredResourceTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMonitoredResourceTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateMonitoredResourceTypeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName     *string                           `json:"displayName"`
		Description     *string                           `json:"description"`
		MetricNamespace *string                           `json:"metricNamespace"`
		Metadata        resourcetypemetadatadetails       `json:"metadata"`
		FreeformTags    map[string]string                 `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{} `json:"definedTags"`
		Name            *string                           `json:"name"`
		CompartmentId   *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.MetricNamespace = model.MetricNamespace

	nn, e = model.Metadata.UnmarshalPolymorphicJSON(model.Metadata.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Metadata = nn.(ResourceTypeMetadataDetails)
	} else {
		m.Metadata = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	return
}
