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

// TypedObjectWrapper A wrapper for a typed object.
type TypedObjectWrapper struct {
	TypedObject TypedObject `mandatory:"false" json:"typedObject"`
}

func (m TypedObjectWrapper) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TypedObjectWrapper) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TypedObjectWrapper) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TypedObject typedobject `json:"typedObject"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.TypedObject.UnmarshalPolymorphicJSON(model.TypedObject.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TypedObject = nn.(TypedObject)
	} else {
		m.TypedObject = nil
	}

	return
}
