// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataObjectQuery Information required to form and execute query on a data object.
type DataObjectQuery interface {

	// List of bind parameters to be applied in the query.
	GetBindParams() []DataObjectBindParameter

	// Timeout (in seconds) to be set for the data object query execution.
	GetQueryExecutionTimeoutInSeconds() *float64
}

type dataobjectquery struct {
	JsonData                       []byte
	BindParams                     []DataObjectBindParameter `mandatory:"false" json:"bindParams"`
	QueryExecutionTimeoutInSeconds *float64                  `mandatory:"false" json:"queryExecutionTimeoutInSeconds"`
	QueryType                      string                    `json:"queryType"`
}

// UnmarshalJSON unmarshals json
func (m *dataobjectquery) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataobjectquery dataobjectquery
	s := struct {
		Model Unmarshalerdataobjectquery
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.BindParams = s.Model.BindParams
	m.QueryExecutionTimeoutInSeconds = s.Model.QueryExecutionTimeoutInSeconds
	m.QueryType = s.Model.QueryType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataobjectquery) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.QueryType {
	case "STANDARD_QUERY":
		mm := DataObjectStandardQuery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEMPLATIZED_QUERY":
		mm := DataObjectTemplatizedQuery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DataObjectQuery: %s.", m.QueryType)
		return *m, nil
	}
}

// GetBindParams returns BindParams
func (m dataobjectquery) GetBindParams() []DataObjectBindParameter {
	return m.BindParams
}

// GetQueryExecutionTimeoutInSeconds returns QueryExecutionTimeoutInSeconds
func (m dataobjectquery) GetQueryExecutionTimeoutInSeconds() *float64 {
	return m.QueryExecutionTimeoutInSeconds
}

func (m dataobjectquery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataobjectquery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataObjectQueryQueryTypeEnum Enum with underlying type: string
type DataObjectQueryQueryTypeEnum string

// Set of constants representing the allowable values for DataObjectQueryQueryTypeEnum
const (
	DataObjectQueryQueryTypeTemplatizedQuery DataObjectQueryQueryTypeEnum = "TEMPLATIZED_QUERY"
	DataObjectQueryQueryTypeStandardQuery    DataObjectQueryQueryTypeEnum = "STANDARD_QUERY"
)

var mappingDataObjectQueryQueryTypeEnum = map[string]DataObjectQueryQueryTypeEnum{
	"TEMPLATIZED_QUERY": DataObjectQueryQueryTypeTemplatizedQuery,
	"STANDARD_QUERY":    DataObjectQueryQueryTypeStandardQuery,
}

var mappingDataObjectQueryQueryTypeEnumLowerCase = map[string]DataObjectQueryQueryTypeEnum{
	"templatized_query": DataObjectQueryQueryTypeTemplatizedQuery,
	"standard_query":    DataObjectQueryQueryTypeStandardQuery,
}

// GetDataObjectQueryQueryTypeEnumValues Enumerates the set of values for DataObjectQueryQueryTypeEnum
func GetDataObjectQueryQueryTypeEnumValues() []DataObjectQueryQueryTypeEnum {
	values := make([]DataObjectQueryQueryTypeEnum, 0)
	for _, v := range mappingDataObjectQueryQueryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectQueryQueryTypeEnumStringValues Enumerates the set of values in String for DataObjectQueryQueryTypeEnum
func GetDataObjectQueryQueryTypeEnumStringValues() []string {
	return []string{
		"TEMPLATIZED_QUERY",
		"STANDARD_QUERY",
	}
}

// GetMappingDataObjectQueryQueryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectQueryQueryTypeEnum(val string) (DataObjectQueryQueryTypeEnum, bool) {
	enum, ok := mappingDataObjectQueryQueryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
