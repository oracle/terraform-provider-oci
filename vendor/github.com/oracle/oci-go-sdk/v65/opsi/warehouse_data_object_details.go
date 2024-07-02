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

// WarehouseDataObjectDetails Warehouse data object details.
type WarehouseDataObjectDetails interface {
}

type warehousedataobjectdetails struct {
	JsonData       []byte
	DataObjectType string `json:"dataObjectType"`
}

// UnmarshalJSON unmarshals json
func (m *warehousedataobjectdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerwarehousedataobjectdetails warehousedataobjectdetails
	s := struct {
		Model Unmarshalerwarehousedataobjectdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DataObjectType = s.Model.DataObjectType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *warehousedataobjectdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DataObjectType {
	case "VIEW":
		mm := WarehouseViewDataObjectDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE":
		mm := WarehouseTableDataObjectDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for WarehouseDataObjectDetails: %s.", m.DataObjectType)
		return *m, nil
	}
}

func (m warehousedataobjectdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m warehousedataobjectdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
