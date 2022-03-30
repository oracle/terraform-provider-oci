// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// CreateScheduleDetails The saved schedule.
type CreateScheduleDetails struct {

	// The unique name of the schedule created by the user
	Name *string `mandatory:"true" json:"name"`

	// The tenancy of the customer
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ResultLocation ResultLocation `mandatory:"true" json:"resultLocation"`

	// In x-obmcs-recurring-time format shown here: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
	// Describes the frequency of when the schedule will be run
	ScheduleRecurrences *string `mandatory:"true" json:"scheduleRecurrences"`

	// The date and time of the first time job execution
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	QueryProperties *QueryProperties `mandatory:"true" json:"queryProperties"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateScheduleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags        map[string]string                 `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		Name                *string                           `json:"name"`
		CompartmentId       *string                           `json:"compartmentId"`
		ResultLocation      resultlocation                    `json:"resultLocation"`
		ScheduleRecurrences *string                           `json:"scheduleRecurrences"`
		TimeScheduled       *common.SDKTime                   `json:"timeScheduled"`
		QueryProperties     *QueryProperties                  `json:"queryProperties"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	nn, e = model.ResultLocation.UnmarshalPolymorphicJSON(model.ResultLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultLocation = nn.(ResultLocation)
	} else {
		m.ResultLocation = nil
	}

	m.ScheduleRecurrences = model.ScheduleRecurrences

	m.TimeScheduled = model.TimeScheduled

	m.QueryProperties = model.QueryProperties

	return
}
