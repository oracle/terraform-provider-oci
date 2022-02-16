// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ConfigParameterDefinition The configurable properties of an object type.
type ConfigParameterDefinition struct {
	ParameterType BaseType `mandatory:"false" json:"parameterType"`

	// This object represents the configurable properties for an object type.
	ParameterName *string `mandatory:"false" json:"parameterName"`

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The default value for the parameter.
	DefaultValue *interface{} `mandatory:"false" json:"defaultValue"`

	// The parameter class field name.
	ClassFieldName *string `mandatory:"false" json:"classFieldName"`

	// Specifies whether the parameter is static or not.
	IsStatic *bool `mandatory:"false" json:"isStatic"`

	// Specifies whether the parameter is a class field or not.
	IsClassFieldValue *bool `mandatory:"false" json:"isClassFieldValue"`
}

func (m ConfigParameterDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigParameterDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ConfigParameterDefinition) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ParameterType     basetype     `json:"parameterType"`
		ParameterName     *string      `json:"parameterName"`
		Description       *string      `json:"description"`
		DefaultValue      *interface{} `json:"defaultValue"`
		ClassFieldName    *string      `json:"classFieldName"`
		IsStatic          *bool        `json:"isStatic"`
		IsClassFieldValue *bool        `json:"isClassFieldValue"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ParameterType.UnmarshalPolymorphicJSON(model.ParameterType.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ParameterType = nn.(BaseType)
	} else {
		m.ParameterType = nil
	}

	m.ParameterName = model.ParameterName

	m.Description = model.Description

	m.DefaultValue = model.DefaultValue

	m.ClassFieldName = model.ClassFieldName

	m.IsStatic = model.IsStatic

	m.IsClassFieldValue = model.IsClassFieldValue

	return
}
