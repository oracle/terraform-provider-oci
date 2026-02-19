// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InsightData Forward-compatible payload. The schema is determined by kind.
// Servers may add additional fields without breaking older SDKs.
type InsightData interface {
}

type insightdata struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *insightdata) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinsightdata insightdata
	s := struct {
		Model Unmarshalerinsightdata
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *insightdata) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "SUMMARY":
		mm := SummaryData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE":
		mm := TableData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_SERIES":
		mm := TimeSeriesData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for InsightData: %s.", m.Kind)
		return *m, nil
	}
}

func (m insightdata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m insightdata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InsightDataKindEnum Enum with underlying type: string
type InsightDataKindEnum string

// Set of constants representing the allowable values for InsightDataKindEnum
const (
	InsightDataKindTimeSeries InsightDataKindEnum = "TIME_SERIES"
	InsightDataKindTable      InsightDataKindEnum = "TABLE"
	InsightDataKindSummary    InsightDataKindEnum = "SUMMARY"
)

var mappingInsightDataKindEnum = map[string]InsightDataKindEnum{
	"TIME_SERIES": InsightDataKindTimeSeries,
	"TABLE":       InsightDataKindTable,
	"SUMMARY":     InsightDataKindSummary,
}

var mappingInsightDataKindEnumLowerCase = map[string]InsightDataKindEnum{
	"time_series": InsightDataKindTimeSeries,
	"table":       InsightDataKindTable,
	"summary":     InsightDataKindSummary,
}

// GetInsightDataKindEnumValues Enumerates the set of values for InsightDataKindEnum
func GetInsightDataKindEnumValues() []InsightDataKindEnum {
	values := make([]InsightDataKindEnum, 0)
	for _, v := range mappingInsightDataKindEnum {
		values = append(values, v)
	}
	return values
}

// GetInsightDataKindEnumStringValues Enumerates the set of values in String for InsightDataKindEnum
func GetInsightDataKindEnumStringValues() []string {
	return []string{
		"TIME_SERIES",
		"TABLE",
		"SUMMARY",
	}
}

// GetMappingInsightDataKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInsightDataKindEnum(val string) (InsightDataKindEnum, bool) {
	enum, ok := mappingInsightDataKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
