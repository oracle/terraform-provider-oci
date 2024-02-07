// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSource A representation of a source configuration setup in the Management Agent.
type DataSource interface {

	// Identifier for DataSource. This represents the type and name for the data source associated with the Management Agent.
	GetKey() *string

	// Unique name of the DataSource.
	GetName() *string

	// Compartment owning this DataSource.
	GetCompartmentId() *string

	// State of the DataSource.
	GetState() LifecycleStatesEnum

	// The time the DataSource was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the DataSource data was last received. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime
}

type datasource struct {
	JsonData      []byte
	Key           *string             `mandatory:"true" json:"key"`
	Name          *string             `mandatory:"true" json:"name"`
	CompartmentId *string             `mandatory:"true" json:"compartmentId"`
	State         LifecycleStatesEnum `mandatory:"true" json:"state"`
	TimeCreated   *common.SDKTime     `mandatory:"true" json:"timeCreated"`
	TimeUpdated   *common.SDKTime     `mandatory:"true" json:"timeUpdated"`
	Type          string              `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *datasource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatasource datasource
	s := struct {
		Model Unmarshalerdatasource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.Name = s.Model.Name
	m.CompartmentId = s.Model.CompartmentId
	m.State = s.Model.State
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datasource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "KUBERNETES_CLUSTER":
		mm := KubernetesClusterDataSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PROMETHEUS_EMITTER":
		mm := PrometheusEmitterDataSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DataSource: %s.", m.Type)
		return *m, nil
	}
}

// GetKey returns Key
func (m datasource) GetKey() *string {
	return m.Key
}

// GetName returns Name
func (m datasource) GetName() *string {
	return m.Name
}

// GetCompartmentId returns CompartmentId
func (m datasource) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetState returns State
func (m datasource) GetState() LifecycleStatesEnum {
	return m.State
}

// GetTimeCreated returns TimeCreated
func (m datasource) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m datasource) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m datasource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m datasource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
