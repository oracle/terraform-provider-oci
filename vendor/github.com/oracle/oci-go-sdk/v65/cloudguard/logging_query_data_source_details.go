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

// LoggingQueryDataSourceDetails Information for a logging query for a data source.
type LoggingQueryDataSourceDetails struct {

	// List of logging query regions
	Regions []string `mandatory:"false" json:"regions"`

	// The continuous query expression that is run periodicall
	Query *string `mandatory:"false" json:"query"`

	// Interval in minutes that query is run periodically.
	IntervalInMinutes *int `mandatory:"false" json:"intervalInMinutes"`

	// The integer value that must be exceeded, fall below or equal to (depending on the operator), for the query result to trigger an event
	Threshold *int `mandatory:"false" json:"threshold"`

	QueryStartTime ContinuousQueryStartPolicy `mandatory:"false" json:"queryStartTime"`

	// The additional entities count used for data source query
	AdditionalEntitiesCount *int `mandatory:"false" json:"additionalEntitiesCount"`

	LoggingQueryDetails LoggingQueryDetails `mandatory:"false" json:"loggingQueryDetails"`

	// Operator used in data source
	Operator LoggingQueryOperatorTypeEnum `mandatory:"false" json:"operator,omitempty"`

	// Type of logging query for data source (Sighting/Insight)
	LoggingQueryType LoggingQueryTypeEnum `mandatory:"false" json:"loggingQueryType,omitempty"`
}

func (m LoggingQueryDataSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoggingQueryDataSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLoggingQueryOperatorTypeEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetLoggingQueryOperatorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLoggingQueryTypeEnum(string(m.LoggingQueryType)); !ok && m.LoggingQueryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoggingQueryType: %s. Supported values are: %s.", m.LoggingQueryType, strings.Join(GetLoggingQueryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LoggingQueryDataSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLoggingQueryDataSourceDetails LoggingQueryDataSourceDetails
	s := struct {
		DiscriminatorParam string `json:"dataSourceFeedProvider"`
		MarshalTypeLoggingQueryDataSourceDetails
	}{
		"LOGGINGQUERY",
		(MarshalTypeLoggingQueryDataSourceDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *LoggingQueryDataSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Regions                 []string                     `json:"regions"`
		Query                   *string                      `json:"query"`
		IntervalInMinutes       *int                         `json:"intervalInMinutes"`
		Threshold               *int                         `json:"threshold"`
		QueryStartTime          continuousquerystartpolicy   `json:"queryStartTime"`
		Operator                LoggingQueryOperatorTypeEnum `json:"operator"`
		LoggingQueryType        LoggingQueryTypeEnum         `json:"loggingQueryType"`
		AdditionalEntitiesCount *int                         `json:"additionalEntitiesCount"`
		LoggingQueryDetails     loggingquerydetails          `json:"loggingQueryDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Regions = make([]string, len(model.Regions))
	copy(m.Regions, model.Regions)
	m.Query = model.Query

	m.IntervalInMinutes = model.IntervalInMinutes

	m.Threshold = model.Threshold

	nn, e = model.QueryStartTime.UnmarshalPolymorphicJSON(model.QueryStartTime.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryStartTime = nn.(ContinuousQueryStartPolicy)
	} else {
		m.QueryStartTime = nil
	}

	m.Operator = model.Operator

	m.LoggingQueryType = model.LoggingQueryType

	m.AdditionalEntitiesCount = model.AdditionalEntitiesCount

	nn, e = model.LoggingQueryDetails.UnmarshalPolymorphicJSON(model.LoggingQueryDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LoggingQueryDetails = nn.(LoggingQueryDetails)
	} else {
		m.LoggingQueryDetails = nil
	}

	return
}
