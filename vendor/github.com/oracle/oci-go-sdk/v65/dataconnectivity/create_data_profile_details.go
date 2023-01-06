// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDataProfileDetails The data profile payload.
type CreateDataProfileDetails struct {
	ReadOperationConfig *ReadOperationConfig `mandatory:"false" json:"readOperationConfig"`

	DataAsset *DataAsset `mandatory:"false" json:"dataAsset"`

	Connection *Connection `mandatory:"false" json:"connection"`

	Schema *Schema `mandatory:"false" json:"schema"`

	DataEntity DataEntity `mandatory:"false" json:"dataEntity"`

	ProfileConfig *ProfileConfig `mandatory:"false" json:"profileConfig"`
}

func (m CreateDataProfileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataProfileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDataProfileDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ReadOperationConfig *ReadOperationConfig `json:"readOperationConfig"`
		DataAsset           *DataAsset           `json:"dataAsset"`
		Connection          *Connection          `json:"connection"`
		Schema              *Schema              `json:"schema"`
		DataEntity          dataentity           `json:"dataEntity"`
		ProfileConfig       *ProfileConfig       `json:"profileConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ReadOperationConfig = model.ReadOperationConfig

	m.DataAsset = model.DataAsset

	m.Connection = model.Connection

	m.Schema = model.Schema

	nn, e = model.DataEntity.UnmarshalPolymorphicJSON(model.DataEntity.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataEntity = nn.(DataEntity)
	} else {
		m.DataEntity = nil
	}

	m.ProfileConfig = model.ProfileConfig

	return
}
