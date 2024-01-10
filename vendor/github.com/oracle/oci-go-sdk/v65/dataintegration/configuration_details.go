// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigurationDetails A key map. If provided, key is replaced with generated key.
type ConfigurationDetails struct {
	DataAsset DataAsset `mandatory:"false" json:"dataAsset"`

	Connection Connection `mandatory:"false" json:"connection"`

	// The compartment ID of the object store.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	Schema *Schema `mandatory:"false" json:"schema"`
}

func (m ConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataAsset     dataasset  `json:"dataAsset"`
		Connection    connection `json:"connection"`
		CompartmentId *string    `json:"compartmentId"`
		Schema        *Schema    `json:"schema"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.DataAsset.UnmarshalPolymorphicJSON(model.DataAsset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataAsset = nn.(DataAsset)
	} else {
		m.DataAsset = nil
	}

	nn, e = model.Connection.UnmarshalPolymorphicJSON(model.Connection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Connection = nn.(Connection)
	} else {
		m.Connection = nil
	}

	m.CompartmentId = model.CompartmentId

	m.Schema = model.Schema

	return
}
