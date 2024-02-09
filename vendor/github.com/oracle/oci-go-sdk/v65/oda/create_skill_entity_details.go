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

// CreateSkillEntityDetails Properties that are required to create a skill entity.
type CreateSkillEntityDetails interface {

	// The entity name. This must be unique within the parent resource.
	GetName() *string
}

type createskillentitydetails struct {
	JsonData []byte
	Name     *string `mandatory:"true" json:"name"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createskillentitydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateskillentitydetails createskillentitydetails
	s := struct {
		Model Unmarshalercreateskillentitydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createskillentitydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ENUM_VALUES":
		mm := CreateSkillValueListEntityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPOSITE":
		mm := CreateSkillCompositeEntityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateSkillEntityDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetName returns Name
func (m createskillentitydetails) GetName() *string {
	return m.Name
}

func (m createskillentitydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createskillentitydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
