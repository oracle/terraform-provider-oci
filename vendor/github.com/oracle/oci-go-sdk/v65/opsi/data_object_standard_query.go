// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataObjectStandardQuery Information required to execute query on data objects. Query is given in standard SQL syntax providing flexibility
// to form complex queries such as queries with joins and nested queries.
type DataObjectStandardQuery struct {

	// List of bind parameters to be applied in the query.
	BindParams []DataObjectBindParameter `mandatory:"false" json:"bindParams"`

	// Timeout (in seconds) to be set for the data object query execution.
	QueryExecutionTimeoutInSeconds *float64 `mandatory:"false" json:"queryExecutionTimeoutInSeconds"`

	// SQL query statement with standard Oracle supported SQL syntax.
	// - When Warehouse (e.g: Awr hub) data objects are queried, use the actual names of underlying data objects (e.g: tables, views) in the query.
	// The same query that works through JDBC connection with the OperationsInsightsWarehouseUsers credentials will work here and vice-versa.
	// SCHEMA.VIEW syntax can also be used here.
	// - When OPSI data objects are queried, use name of the respective OPSI data object, just like how views are used in a query.
	// Identifier of the OPSI data object cannot be used in the query.
	Statement *string `mandatory:"false" json:"statement"`

	TimeFilters *DataObjectQueryTimeFilters `mandatory:"false" json:"timeFilters"`
}

// GetBindParams returns BindParams
func (m DataObjectStandardQuery) GetBindParams() []DataObjectBindParameter {
	return m.BindParams
}

// GetQueryExecutionTimeoutInSeconds returns QueryExecutionTimeoutInSeconds
func (m DataObjectStandardQuery) GetQueryExecutionTimeoutInSeconds() *float64 {
	return m.QueryExecutionTimeoutInSeconds
}

func (m DataObjectStandardQuery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectStandardQuery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectStandardQuery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectStandardQuery DataObjectStandardQuery
	s := struct {
		DiscriminatorParam string `json:"queryType"`
		MarshalTypeDataObjectStandardQuery
	}{
		"STANDARD_QUERY",
		(MarshalTypeDataObjectStandardQuery)(m),
	}

	return json.Marshal(&s)
}
