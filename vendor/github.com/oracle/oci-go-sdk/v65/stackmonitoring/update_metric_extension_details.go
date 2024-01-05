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

// UpdateMetricExtensionDetails The information about updating a metric extension resource
type UpdateMetricExtensionDetails struct {

	// Metric Extension resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the metric extension.
	Description *string `mandatory:"false" json:"description"`

	// Schedule of metric extension should use RFC 5545 format -> recur-rule-part = "FREQ";"INTERVAL" where FREQ rule part identifies the type of recurrence rule. Valid values are "MINUTELY","HOURLY","DAILY" to specify repeating events based on an interval of a minute, an hour and a day or more. Example- FREQ=DAILY;INTERVAL=1
	CollectionRecurrences *string `mandatory:"false" json:"collectionRecurrences"`

	// List of metrics which are part of this metric extension
	MetricList []Metric `mandatory:"false" json:"metricList"`

	QueryProperties MetricExtensionUpdateQueryProperties `mandatory:"false" json:"queryProperties"`
}

func (m UpdateMetricExtensionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMetricExtensionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMetricExtensionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                              `json:"displayName"`
		Description           *string                              `json:"description"`
		CollectionRecurrences *string                              `json:"collectionRecurrences"`
		MetricList            []Metric                             `json:"metricList"`
		QueryProperties       metricextensionupdatequeryproperties `json:"queryProperties"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.CollectionRecurrences = model.CollectionRecurrences

	m.MetricList = make([]Metric, len(model.MetricList))
	copy(m.MetricList, model.MetricList)
	nn, e = model.QueryProperties.UnmarshalPolymorphicJSON(model.QueryProperties.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryProperties = nn.(MetricExtensionUpdateQueryProperties)
	} else {
		m.QueryProperties = nil
	}

	return
}
