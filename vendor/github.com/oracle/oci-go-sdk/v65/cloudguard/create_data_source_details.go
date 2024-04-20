// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDataSourceDetails Parameters for creating a data source (DataSource resource).
type CreateDataSourceDetails struct {

	// Data source display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID of the data source
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of data source feed provider (LoggingQuery)
	DataSourceFeedProvider DataSourceFeedProviderEnum `mandatory:"true" json:"dataSourceFeedProvider"`

	// Enablement status of data source.
	Status DataSourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	DataSourceDetails DataSourceDetails `mandatory:"false" json:"dataSourceDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDataSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataSourceFeedProviderEnum(string(m.DataSourceFeedProvider)); !ok && m.DataSourceFeedProvider != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSourceFeedProvider: %s. Supported values are: %s.", m.DataSourceFeedProvider, strings.Join(GetDataSourceFeedProviderEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDataSourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDataSourceStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDataSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Status                 DataSourceStatusEnum              `json:"status"`
		DataSourceDetails      datasourcedetails                 `json:"dataSourceDetails"`
		FreeformTags           map[string]string                 `json:"freeformTags"`
		DefinedTags            map[string]map[string]interface{} `json:"definedTags"`
		DisplayName            *string                           `json:"displayName"`
		CompartmentId          *string                           `json:"compartmentId"`
		DataSourceFeedProvider DataSourceFeedProviderEnum        `json:"dataSourceFeedProvider"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Status = model.Status

	nn, e = model.DataSourceDetails.UnmarshalPolymorphicJSON(model.DataSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataSourceDetails = nn.(DataSourceDetails)
	} else {
		m.DataSourceDetails = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.DataSourceFeedProvider = model.DataSourceFeedProvider

	return
}
