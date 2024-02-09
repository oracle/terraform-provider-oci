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

// DataSourceEventSummary Summary information about a data source event.
type DataSourceEventSummary struct {

	// Data source event region
	Region *string `mandatory:"true" json:"region"`

	// Data source event date and time
	EventDate *common.SDKTime `mandatory:"true" json:"eventDate"`

	// Unique identifier of data source.
	DataSourceId *string `mandatory:"true" json:"dataSourceId"`

	// Data source event creation date and time
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	EventInfo DataSourceEventInfo `mandatory:"true" json:"eventInfo"`

	// Current data source event info status
	Status DataSourceEventInfoStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Data source event comments
	Comments *string `mandatory:"false" json:"comments"`
}

func (m DataSourceEventSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSourceEventSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataSourceEventInfoStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDataSourceEventInfoStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DataSourceEventSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Status       DataSourceEventInfoStatusEnum `json:"status"`
		Comments     *string                       `json:"comments"`
		Region       *string                       `json:"region"`
		EventDate    *common.SDKTime               `json:"eventDate"`
		DataSourceId *string                       `json:"dataSourceId"`
		TimeCreated  *common.SDKTime               `json:"timeCreated"`
		EventInfo    datasourceeventinfo           `json:"eventInfo"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Status = model.Status

	m.Comments = model.Comments

	m.Region = model.Region

	m.EventDate = model.EventDate

	m.DataSourceId = model.DataSourceId

	m.TimeCreated = model.TimeCreated

	nn, e = model.EventInfo.UnmarshalPolymorphicJSON(model.EventInfo.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EventInfo = nn.(DataSourceEventInfo)
	} else {
		m.EventInfo = nil
	}

	return
}
