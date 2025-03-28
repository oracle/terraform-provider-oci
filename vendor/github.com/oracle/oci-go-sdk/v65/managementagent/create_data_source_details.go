// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDataSourceDetails A new data source.
type CreateDataSourceDetails interface {

	// Unique name of the DataSource.
	GetName() *string

	// Compartment owning this DataSource.
	GetCompartmentId() *string
}

type createdatasourcedetails struct {
	JsonData      []byte
	Name          *string `mandatory:"true" json:"name"`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
	Type          string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatasourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatasourcedetails createdatasourcedetails
	s := struct {
		Model Unmarshalercreatedatasourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.CompartmentId = s.Model.CompartmentId
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatasourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "PROMETHEUS_EMITTER":
		mm := CreatePrometheusEmitterDataSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDataSourceDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetName returns Name
func (m createdatasourcedetails) GetName() *string {
	return m.Name
}

// GetCompartmentId returns CompartmentId
func (m createdatasourcedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createdatasourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatasourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
