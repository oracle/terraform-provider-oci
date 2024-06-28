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

// QueryDataObjectResultSetRowsCollection Collection of result set rows from the data object query.
type QueryDataObjectResultSetRowsCollection interface {
}

type querydataobjectresultsetrowscollection struct {
	JsonData []byte
	Format   string `json:"format"`
}

// UnmarshalJSON unmarshals json
func (m *querydataobjectresultsetrowscollection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerquerydataobjectresultsetrowscollection querydataobjectresultsetrowscollection
	s := struct {
		Model Unmarshalerquerydataobjectresultsetrowscollection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Format = s.Model.Format

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *querydataobjectresultsetrowscollection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Format {
	case "JSON":
		mm := QueryDataObjectJsonResultSetRowsCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for QueryDataObjectResultSetRowsCollection: %s.", m.Format)
		return *m, nil
	}
}

func (m querydataobjectresultsetrowscollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m querydataobjectresultsetrowscollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryDataObjectResultSetRowsCollectionFormatEnum Enum with underlying type: string
type QueryDataObjectResultSetRowsCollectionFormatEnum string

// Set of constants representing the allowable values for QueryDataObjectResultSetRowsCollectionFormatEnum
const (
	QueryDataObjectResultSetRowsCollectionFormatJson QueryDataObjectResultSetRowsCollectionFormatEnum = "JSON"
)

var mappingQueryDataObjectResultSetRowsCollectionFormatEnum = map[string]QueryDataObjectResultSetRowsCollectionFormatEnum{
	"JSON": QueryDataObjectResultSetRowsCollectionFormatJson,
}

var mappingQueryDataObjectResultSetRowsCollectionFormatEnumLowerCase = map[string]QueryDataObjectResultSetRowsCollectionFormatEnum{
	"json": QueryDataObjectResultSetRowsCollectionFormatJson,
}

// GetQueryDataObjectResultSetRowsCollectionFormatEnumValues Enumerates the set of values for QueryDataObjectResultSetRowsCollectionFormatEnum
func GetQueryDataObjectResultSetRowsCollectionFormatEnumValues() []QueryDataObjectResultSetRowsCollectionFormatEnum {
	values := make([]QueryDataObjectResultSetRowsCollectionFormatEnum, 0)
	for _, v := range mappingQueryDataObjectResultSetRowsCollectionFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryDataObjectResultSetRowsCollectionFormatEnumStringValues Enumerates the set of values in String for QueryDataObjectResultSetRowsCollectionFormatEnum
func GetQueryDataObjectResultSetRowsCollectionFormatEnumStringValues() []string {
	return []string{
		"JSON",
	}
}

// GetMappingQueryDataObjectResultSetRowsCollectionFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryDataObjectResultSetRowsCollectionFormatEnum(val string) (QueryDataObjectResultSetRowsCollectionFormatEnum, bool) {
	enum, ok := mappingQueryDataObjectResultSetRowsCollectionFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
