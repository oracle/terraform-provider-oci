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

// WarehouseDataObjectSummary Summary of a Warehouse data object.
type WarehouseDataObjectSummary struct {

	// Type of the data object.
	DataObjectType DataObjectTypeEnum `mandatory:"true" json:"dataObjectType"`

	// Name of the data object, which can be used in data object queries just like how view names are used in a query.
	Name *string `mandatory:"false" json:"name"`

	// Owner of the data object, which can be used in data object queries in front of data object names just like SCHEMA.VIEW notation in queries.
	Owner *string `mandatory:"false" json:"owner"`

	Details WarehouseDataObjectDetails `mandatory:"false" json:"details"`
}

func (m WarehouseDataObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WarehouseDataObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataObjectTypeEnum(string(m.DataObjectType)); !ok && m.DataObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataObjectType: %s. Supported values are: %s.", m.DataObjectType, strings.Join(GetDataObjectTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *WarehouseDataObjectSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name           *string                    `json:"name"`
		Owner          *string                    `json:"owner"`
		Details        warehousedataobjectdetails `json:"details"`
		DataObjectType DataObjectTypeEnum         `json:"dataObjectType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Owner = model.Owner

	nn, e = model.Details.UnmarshalPolymorphicJSON(model.Details.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Details = nn.(WarehouseDataObjectDetails)
	} else {
		m.Details = nil
	}

	m.DataObjectType = model.DataObjectType

	return
}
