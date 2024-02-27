// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Entity Metadata for an entity.
type Entity interface {

	// Unique immutable identifier that was assigned when the resource was created.
	GetId() *string

	// The entity name. This must be unique within the parent resource.
	GetName() *string
}

type entity struct {
	JsonData []byte
	Id       *string `mandatory:"true" json:"id"`
	Name     *string `mandatory:"true" json:"name"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *entity) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerentity entity
	s := struct {
		Model Unmarshalerentity
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.Name = s.Model.Name
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *entity) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "COMPOSITE":
		mm := CompositeEntity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENUM_VALUES":
		mm := ValueListEntity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Entity: %s.", m.Type)
		return *m, nil
	}
}

// GetId returns Id
func (m entity) GetId() *string {
	return m.Id
}

// GetName returns Name
func (m entity) GetName() *string {
	return m.Name
}

func (m entity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m entity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
