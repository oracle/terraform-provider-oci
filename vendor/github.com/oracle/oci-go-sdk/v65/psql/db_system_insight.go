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

// DbSystemInsight Response envelope containing insight metadata and a polymorphic insight data payload.
type DbSystemInsight struct {

	// Echo of the requested insight type.
	InsightType DbSystemInsightInsightTypeEnum `mandatory:"true" json:"insightType"`

	// Echo of the requested insight data type.
	InsightDataType DbSystemInsightInsightDataTypeEnum `mandatory:"true" json:"insightDataType"`

	DateTimeRange *DateTimeRange `mandatory:"true" json:"dateTimeRange"`

	Data InsightData `mandatory:"true" json:"data"`
}

func (m DbSystemInsight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemInsight) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemInsightInsightTypeEnum(string(m.InsightType)); !ok && m.InsightType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InsightType: %s. Supported values are: %s.", m.InsightType, strings.Join(GetDbSystemInsightInsightTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemInsightInsightDataTypeEnum(string(m.InsightDataType)); !ok && m.InsightDataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InsightDataType: %s. Supported values are: %s.", m.InsightDataType, strings.Join(GetDbSystemInsightInsightDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DbSystemInsight) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InsightType     DbSystemInsightInsightTypeEnum     `json:"insightType"`
		InsightDataType DbSystemInsightInsightDataTypeEnum `json:"insightDataType"`
		DateTimeRange   *DateTimeRange                     `json:"dateTimeRange"`
		Data            insightdata                        `json:"data"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InsightType = model.InsightType

	m.InsightDataType = model.InsightDataType

	m.DateTimeRange = model.DateTimeRange

	nn, e = model.Data.UnmarshalPolymorphicJSON(model.Data.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Data = nn.(InsightData)
	} else {
		m.Data = nil
	}

	return
}

// DbSystemInsightInsightTypeEnum Enum with underlying type: string
type DbSystemInsightInsightTypeEnum string

// Set of constants representing the allowable values for DbSystemInsightInsightTypeEnum
const (
	DbSystemInsightInsightTypeQueryInsight DbSystemInsightInsightTypeEnum = "QUERY_INSIGHT"
)

var mappingDbSystemInsightInsightTypeEnum = map[string]DbSystemInsightInsightTypeEnum{
	"QUERY_INSIGHT": DbSystemInsightInsightTypeQueryInsight,
}

var mappingDbSystemInsightInsightTypeEnumLowerCase = map[string]DbSystemInsightInsightTypeEnum{
	"query_insight": DbSystemInsightInsightTypeQueryInsight,
}

// GetDbSystemInsightInsightTypeEnumValues Enumerates the set of values for DbSystemInsightInsightTypeEnum
func GetDbSystemInsightInsightTypeEnumValues() []DbSystemInsightInsightTypeEnum {
	values := make([]DbSystemInsightInsightTypeEnum, 0)
	for _, v := range mappingDbSystemInsightInsightTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemInsightInsightTypeEnumStringValues Enumerates the set of values in String for DbSystemInsightInsightTypeEnum
func GetDbSystemInsightInsightTypeEnumStringValues() []string {
	return []string{
		"QUERY_INSIGHT",
	}
}

// GetMappingDbSystemInsightInsightTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemInsightInsightTypeEnum(val string) (DbSystemInsightInsightTypeEnum, bool) {
	enum, ok := mappingDbSystemInsightInsightTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbSystemInsightInsightDataTypeEnum Enum with underlying type: string
type DbSystemInsightInsightDataTypeEnum string

// Set of constants representing the allowable values for DbSystemInsightInsightDataTypeEnum
const (
	DbSystemInsightInsightDataTypeAasTimeSeries DbSystemInsightInsightDataTypeEnum = "AAS_TIME_SERIES"
	DbSystemInsightInsightDataTypeTopQueries    DbSystemInsightInsightDataTypeEnum = "TOP_QUERIES"
)

var mappingDbSystemInsightInsightDataTypeEnum = map[string]DbSystemInsightInsightDataTypeEnum{
	"AAS_TIME_SERIES": DbSystemInsightInsightDataTypeAasTimeSeries,
	"TOP_QUERIES":     DbSystemInsightInsightDataTypeTopQueries,
}

var mappingDbSystemInsightInsightDataTypeEnumLowerCase = map[string]DbSystemInsightInsightDataTypeEnum{
	"aas_time_series": DbSystemInsightInsightDataTypeAasTimeSeries,
	"top_queries":     DbSystemInsightInsightDataTypeTopQueries,
}

// GetDbSystemInsightInsightDataTypeEnumValues Enumerates the set of values for DbSystemInsightInsightDataTypeEnum
func GetDbSystemInsightInsightDataTypeEnumValues() []DbSystemInsightInsightDataTypeEnum {
	values := make([]DbSystemInsightInsightDataTypeEnum, 0)
	for _, v := range mappingDbSystemInsightInsightDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemInsightInsightDataTypeEnumStringValues Enumerates the set of values in String for DbSystemInsightInsightDataTypeEnum
func GetDbSystemInsightInsightDataTypeEnumStringValues() []string {
	return []string{
		"AAS_TIME_SERIES",
		"TOP_QUERIES",
	}
}

// GetMappingDbSystemInsightInsightDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemInsightInsightDataTypeEnum(val string) (DbSystemInsightInsightDataTypeEnum, bool) {
	enum, ok := mappingDbSystemInsightInsightDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
