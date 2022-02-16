// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// QueryAggregation Query results.
type QueryAggregation struct {

	// Percentage progress completion of the query.
	PercentComplete *int `mandatory:"true" json:"percentComplete"`

	// Number of rows query retrieved. Up to maxTotalCount limit.
	TotalCount *int `mandatory:"false" json:"totalCount"`

	// Number of rows matched by query.
	TotalMatchedCount *int64 `mandatory:"false" json:"totalMatchedCount"`

	// True if query did not complete processing all data.
	ArePartialResults *bool `mandatory:"false" json:"arePartialResults"`

	// Explanation of why results may be partial. Only set if arePartialResults is true.
	PartialResultReason *string `mandatory:"false" json:"partialResultReason"`

	// Query result columns
	Columns []AbstractColumn `mandatory:"false" json:"columns"`

	// Query result fields
	Fields []AbstractColumn `mandatory:"false" json:"fields"`

	// Query result data
	Items []map[string]interface{} `mandatory:"false" json:"items"`

	// Time ellapsed executing query in milli-seconds.
	QueryExecutionTimeInMs *int64 `mandatory:"false" json:"queryExecutionTimeInMs"`
}

func (m QueryAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *QueryAggregation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TotalCount             *int                     `json:"totalCount"`
		TotalMatchedCount      *int64                   `json:"totalMatchedCount"`
		ArePartialResults      *bool                    `json:"arePartialResults"`
		PartialResultReason    *string                  `json:"partialResultReason"`
		Columns                []abstractcolumn         `json:"columns"`
		Fields                 []abstractcolumn         `json:"fields"`
		Items                  []map[string]interface{} `json:"items"`
		QueryExecutionTimeInMs *int64                   `json:"queryExecutionTimeInMs"`
		PercentComplete        *int                     `json:"percentComplete"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TotalCount = model.TotalCount

	m.TotalMatchedCount = model.TotalMatchedCount

	m.ArePartialResults = model.ArePartialResults

	m.PartialResultReason = model.PartialResultReason

	m.Columns = make([]AbstractColumn, len(model.Columns))
	for i, n := range model.Columns {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Columns[i] = nn.(AbstractColumn)
		} else {
			m.Columns[i] = nil
		}
	}

	m.Fields = make([]AbstractColumn, len(model.Fields))
	for i, n := range model.Fields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Fields[i] = nn.(AbstractColumn)
		} else {
			m.Fields[i] = nil
		}
	}

	m.Items = make([]map[string]interface{}, len(model.Items))
	for i, n := range model.Items {
		m.Items[i] = n
	}

	m.QueryExecutionTimeInMs = model.QueryExecutionTimeInMs

	m.PercentComplete = model.PercentComplete

	return
}
