// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DedicatedVantagePointSummary Information about dedicated vantage points.
type DedicatedVantagePointSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the dedicated vantage point.
	Id *string `mandatory:"true" json:"id"`

	// Unique dedicated vantage point name that cannot be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique permanent name of the vantage point.
	Name *string `mandatory:"true" json:"name"`

	// Status of the dedicated vantage point.
	Status DedicatedVantagePointStatusEnum `mandatory:"true" json:"status"`

	DvpStackDetails DvpStackDetails `mandatory:"true" json:"dvpStackDetails"`

	// Name of the region.
	Region *string `mandatory:"true" json:"region"`

	MonitorStatusCountMap *MonitorStatusCountMap `mandatory:"true" json:"monitorStatusCountMap"`

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DedicatedVantagePointSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedVantagePointSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedVantagePointStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDedicatedVantagePointStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DedicatedVantagePointSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated           *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated           *common.SDKTime                   `json:"timeUpdated"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		Id                    *string                           `json:"id"`
		DisplayName           *string                           `json:"displayName"`
		Name                  *string                           `json:"name"`
		Status                DedicatedVantagePointStatusEnum   `json:"status"`
		DvpStackDetails       dvpstackdetails                   `json:"dvpStackDetails"`
		Region                *string                           `json:"region"`
		MonitorStatusCountMap *MonitorStatusCountMap            `json:"monitorStatusCountMap"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.Name = model.Name

	m.Status = model.Status

	nn, e = model.DvpStackDetails.UnmarshalPolymorphicJSON(model.DvpStackDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DvpStackDetails = nn.(DvpStackDetails)
	} else {
		m.DvpStackDetails = nil
	}

	m.Region = model.Region

	m.MonitorStatusCountMap = model.MonitorStatusCountMap

	return
}
